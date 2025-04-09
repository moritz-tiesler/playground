package set

import (
	"iter"
	"maps"
	"slices"
)

type Set[E comparable] struct {
	items map[E]struct{}
}

func (s *Set[E]) Put(v E) {
	s.items[v] = struct{}{}
}

func (s *Set[E]) Del(v E) {
	delete(s.items, v)
}

func (s *Set[E]) Clear() {
	clear(s.items)
}

func (s *Set[E]) Len() int {
	return len(s.items)
}

func (s *Set[E]) Items() iter.Seq[E] {
	return maps.Keys(s.items)
}

func (s *Set[E]) Has(v E) bool {
	_, ok := s.items[v]
	return ok
}

func (s *Set[E]) Intersection(other *Set[E]) *Set[E] {
	empty := slices.Values([]E{})
	intersected := New(empty)
	for item := range s.items {
		if other.Has(item) {
			intersected.Put(item)
		}
	}
	return intersected
}

func New[E comparable](rangable iter.Seq[E]) *Set[E] {
	s := Set[E]{items: make(map[E]struct{})}
	for v := range rangable {
		s.Put(v)
	}
	return &s
}
