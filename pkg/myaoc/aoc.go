package myaoc

import (
	"bufio"
	"constraints"
	"io"
	"strings"
)

// Make2DSlice creates a 2d slice of type T and sets i,j of the 2d array to the result of f(i,j)
func Make2DSlice[T any](i, j int, f func(i, j int) T) [][]T {
	m := make([][]T, i)
	for i := range m {
		m[i] = make([]T, j)
		for j := range m[i] {
			m[i][j] = f(i, j)
		}
	}
	return m
}

// IsSafe returns for any 2d slice whether the given ints are in bounds
func IsSafe[T any](ts [][]T, i, j int) bool {
	return i >= 0 && j >= 0 && i < len(ts) && len(ts) > 0 && j < len(ts[0])
}

// SliceNeighbors is a utility function to get the elements surrounding a particular 2d index.
func SliceNeighbors[T any](ts [][]T, n, m int) []T {
	var out []T
	for i := n - 1; i < n+2; i++ {
		for j := m - 1; j < m+2; j++ {
			if IsSafe(ts, i, j) && !(i == n && j == m) { // You are not your own neighbor
				out = append(out, ts[i][j])
			}
		}
	}
	return out
}

// type Coster interface {
// 	Cost() int
// }

// type Dijkstra[T Coster] struct {
// 	Cells [][]T
// }

// func (receiver type) name(params) type {

// }

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

// ReadAOC takes a reader and calls the func on every space-trimmed input line.
func ReadAOC[T any](r io.Reader, f func(string) (T, error)) ([]T, error) {
	br := bufio.NewReader(r)

	var ts []T
	for {
		st, err := br.ReadString('\n')
		if err != nil && err != io.EOF {
			return nil, err
		}
		if st == "" && err == io.EOF {
			return ts, nil
		}

		next, err := f(strings.TrimSpace(st))
		if err != nil {
			if err == io.EOF { // Callees may return io.EOF to end our use of this reader.
				return ts, nil
			}
			return ts, err
		}
		ts = append(ts, next)
	}
}

func MaxOf[T any, U constraints.Ordered](ts []T, f func(T) U) U {
	if len(ts) == 0 {
		panic("nothin there")
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
