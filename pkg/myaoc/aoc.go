package myaoc

import (
	"constraints"
)

type Queue[T any] []T

func (q Queue[T]) Len() int {
	return len(q)
}

func (q Queue[T]) Enqueue(t T) {
	q = append(q, t)
}

func (q Queue[T]) Dequeue() T {
	o := q[0]
	q = q[1:]
	return o
}

type Set[T comparable] map[T]struct{}

func (s Set[T]) Add(t T) {
	s[t] = struct{}{}
}

func (s Set[T]) Contains(t T) bool {
	_, ok := s[t]
	return ok
}

func SetFrom[T comparable](ts []T) Set[T] {
	s := Set[T]{}
	for _, t := range ts {
		s.Add(t)
	}
	return s
}

func SetFromFunc[U any, T comparable](ts []U, f func(U) T) Set[T] {
	s := Set[T]{}
	for _, t := range ts {
		s.Add(f(t))
	}
	return s
}

func MaxOf[T any, U constraints.Ordered](ts []T, f func(T) U) U {
	if len(ts) == 0 {
		var u U
		return u
	}
	u := f(ts[0])
	for _, t := range ts {
		ft := f(t)
		if ft < u {
			u = ft
		}
	}
	return u
}

func Last[T any](ts []T) T {
	return ts[len(ts)-1]
}
