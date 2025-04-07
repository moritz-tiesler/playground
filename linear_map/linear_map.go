package lmap

import (
	"bytes"
	"encoding/gob"
	"hash/fnv"
)

// TODO impl resizing

type LMap[T comparable, U any] struct {
	table []Entry[T, U]
	capa  uint64
	size  int
	enc   *gob.Encoder
	buff  *bytes.Buffer
	colls int
}

func New[T comparable, U any](capa uint64) *LMap[T, U] {
	var b bytes.Buffer
	return &LMap[T, U]{
		table: make([]Entry[T, U], capa, capa),
		capa:  capa,
		enc:   gob.NewEncoder(&b),
		buff:  &b,
	}
}

type Entry[T comparable, U any] struct {
	key   T
	value U
	used  bool
}

func (m *LMap[T, U]) Hash(k T) uint64 {
	var h = fnv.New64a()
	m.encode(k)
	h.Write(m.buff.Bytes())
	m.buff.Reset()
	hash := h.Sum64()
	truc := hash & (m.capa - 1)
	return truc
}

func (m *LMap[T, U]) Put(k T, v U) {
	index := m.Hash(k)
	for m.table[index].used {
		m.colls++
		index = (index + 1) & (m.capa - 1)
	}
	m.table[index] = Entry[T, U]{k, v, true}
	m.size++
	// Todo: if size ~ 1/4 cap then resize underlying slice
}

func (m *LMap[T, U]) Zero() U {
	var z U
	return z
}

func (m *LMap[T, U]) Clear() {
	clear(m.table)
}

func (m *LMap[T, U]) Get(k T) (U, bool) {
	index := m.Hash(k)

	// Linear probing to find the key
	originalIndex := index
	for m.table[index].used {
		if m.table[index].key == k {
			return m.table[index].value, true
		}
		index = (index + 1) & (m.capa - 1) // Wrap around if necessary

		if index == originalIndex {
			// If we've looped back to the start, the key is not present
			return m.Zero(), false
		}
	}

	return m.Zero(), false // Key not found
}

// TODO impl Delete()

func (m *LMap[T, U]) encode(v T) error {

	err := m.enc.Encode(v)
	if err != nil {
		return err
	}
	return nil
}
