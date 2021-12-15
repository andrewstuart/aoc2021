package myaoc

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
	return i > 0 && i < len(ts) && len(ts) > 0 && j > 0 && j < len(ts[0])
}

func SliceNeighbors[T any](ts [][]T, n, m int) []T {
	var out []T
	for i := n - 1; i < n+2; i++ {
		for j := m - 1; j < m+2; j++ {
			if IsSafe(ts, i, j) {
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
