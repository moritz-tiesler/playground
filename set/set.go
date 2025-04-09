package set

import (
	"iter"
	"maps"
	"reflect"
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
	for item := range s.Iter() {
		if other.Has(item) {
			intersected.Put(item)
		}
	}
	return intersected
}

func (s *Set[E]) Difference(other *Set[E]) *Set[E] {
	empty := slices.Values([]E{})
	difference := New(empty)
	for item := range s.Iter() {
		if !other.Has(item) {
			difference.Put(item)
		}
	}
	return difference
}

func (s *Set[E]) Equals(other *Set[E]) bool {
	if s.Len() != other.Len() {
		return false
	}
	return reflect.DeepEqual(s.items, other.items)
}

func (s *Set[E]) Union(other *Set[E]) *Set[E] {
	unioned := New(s.Items())
	for item := range other.Iter() {
		unioned.Put(item)
	}
	return unioned
}

func (s *Set[E]) SymmetricDifference(other *Set[E]) *Set[E] {
	empty := slices.Values([]E{})
	diff1 := New(empty)
	diff2 := New(empty)
	for item := range s.Iter() {
		if !other.Has(item) {
			diff1.Put(item)
		}
	}
	for item := range other.Iter() {
		if !s.Has(item) {
			diff2.Put(item)
		}
	}
	return diff1.Union(diff2)
}

func (s *Set[E]) SubsetOf(other *Set[E]) bool {
	for item := range s.Iter() {
		if !other.Has(item) {
			return false
		}
	}
	return true
}

func (s *Set[E]) SuperSetOf(other *Set[E]) bool {
	return other.SubsetOf(s)
}

func (s *Set[E]) Iter() iter.Seq[E] {
	return func(yield func(E) bool) {
		for e := range s.items {
			if !yield(e) {
				return
			}
		}
	}
}

func New[E comparable](rangable iter.Seq[E]) *Set[E] {
	s := Set[E]{items: make(map[E]struct{})}
	for v := range rangable {
		s.Put(v)
	}
	return &s
}
