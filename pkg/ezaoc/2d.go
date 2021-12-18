package ezaoc

import "fmt"

func Print2dGrid[T any](ts [][]T) {
	for _, row := range ts {
		for _, cell := range row {
			fmt.Print(cell)
		}
		fmt.Println()
	}
}

// Make2DSlice creates a 2d slice of type T and sets i,j of the 2d array to the
// result of f(i,j)
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
	gtZero := i >= 0 && j >= 0
	inBounds := i < len(ts) && len(ts) > 0 && j < len(ts[0])
	return gtZero && inBounds
}

type Cell[T any] struct {
	I, J  int
	Value T
}

func (c Cell[T]) Set(ts [][]T, to T) {
	ts[c.I][c.J] = to
}

func (c Cell[T]) Point() [2]int {
	return [2]int{c.I, c.J}
}

// SliceNeighbors is a utility function to get the elements surrounding a particular 2d index.
func SliceNeighbors[T any](ts [][]T, n, m int) []Cell[T] {
	var out []Cell[T]
	for i := n - 1; i < n+2; i++ {
		for j := m - 1; j < m+2; j++ {
			if IsSafe(ts, i, j) && !(i == n && j == m) { // You are not your own neighbor
				out = append(out, Cell[T]{I: i, J: j, Value: ts[i][j]})
			}
		}
	}
	return out
}

// NonDiagSliceNeighbors is a utility function to get the elements surrounding a
// particular 2d index
func NonDiagSliceNeighbors[T any](ts [][]T, n, m int) []Cell[T] {
	var out []Cell[T]
	for i := n - 1; i < n+2; i++ {
		for j := m - 1; j < m+2; j++ {
			if IsSafe(ts, i, j) && !(i == n && j == m) && !(i != n && j != m) { // You are not your own neighbor, ignore diags
				out = append(out, Cell[T]{I: i, J: j, Value: ts[i][j]})
			}
		}
	}
	return out
}

// VisitCells calls a function for a Cell of each value.
func VisitCells[T any](ts [][]T, f func(Cell[T]) error) {
	var c Cell[T]
	for i, row := range ts {
		for j := range row {
			c.I, c.J, c.Value = i, j, ts[i][j]
			if f(c) != nil {
				return
			}
		}
	}
}

// VisitNeighbors iterates over a 2d array, calling a func with each index and
// a list of neighbors.
func VisitNeighbors[T any](ts [][]T, f func(Cell[T], []Cell[T]) error) {
	var c Cell[T]
	for i, row := range ts {
		for j := range row {
			c.I, c.J, c.Value = i, j, ts[i][j]
			if f(c, SliceNeighbors(ts, i, j)) != nil {
				return
			}
		}
	}
}

// VisitNeighbors iterates over a 2d array, calling a func with each index and
// a list of neighbors.
func VisitNonDiagNeighbors[T any](ts [][]T, f func(Cell[T], []Cell[T]) error) {
	var c Cell[T]
	for i, row := range ts {
		for j := range row {
			c.I, c.J, c.Value = i, j, ts[i][j]
			if f(c, NonDiagSliceNeighbors(ts, i, j)) != nil {
				return
			}
		}
	}
}
