package set

import (
	"fmt"
	"strings"
)

type Set[T comparable] map[T]struct{}

func New[T comparable]() Set[T] {
	return make(Set[T])
}

func (s Set[T]) String() string {
	items := make([]string, 0, len(s))

	for elem := range s {
		items = append(items, fmt.Sprintf("%v", elem))
	}
	return fmt.Sprintf("Set{%s}", strings.Join(items, ", "))
}

func (s Set[T]) Clone() Set[T] {
	clonedSet := New[T]()
	for elem := range s {
		clonedSet.Add(elem)
	}
	return clonedSet
}

func (s Set[T]) Equal(other Set[T]) bool {
	if len(s) != len(other) {
		return false
	}
	for elem := range s {
		if !other.Contains(elem) {
			return false
		}
	}
	return true
}

func (s Set[T]) Add(v T) {
	s[v] = struct{}{}
}

func (s Set[T]) Pop() (v T, ok bool) {
	for item := range s {
		delete(s, item)
		return item, true
	}
	return v, false
}

func (s Set[T]) Contains(v T) (ok bool) {
	_, ok = s[v]
	return ok
}

func (s Set[T]) Difference(other Set[T]) Set[T] {
	diff := New[T]()
	for elem := range s {
		if !other.Contains(elem) {
			diff.Add(elem)
		}
	}
	return diff
}

func (s Set[T]) SymmetricDifference(other Set[T]) Set[T] {
	sd := New[T]()
	for elem := range s {
		if !other.Contains(elem) {
			sd.Add(elem)
		}
	}
	for elem := range other {
		if !s.Contains(elem) {
			sd.Add(elem)
		}
	}
	return sd
}

func (s Set[T]) Union(other Set[T]) Set[T] {
	n := len(s)
	if len(other) > n {
		n = len(other)
	}
	unionedSet := make(Set[T], n)

	for elem := range s {
		unionedSet.Add(elem)
	}
	for elem := range other {
		unionedSet.Add(elem)
	}
	return unionedSet
}

func (s Set[T]) Intersection(other Set[T]) Set[T] {
	intersection := New[T]()
	// loop over smaller set
	if len(s) < len(other) {
		for elem := range s {
			if other.Contains(elem) {
				intersection.Add(elem)
			}
		}
	} else {
		for elem := range other {
			if s.Contains(elem) {
				intersection.Add(elem)
			}
		}
	}
	return intersection
}
