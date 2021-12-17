package ezaoc

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

// NonDiagSliceNeighbors is a utility function to get the elements surrounding a
// particular 2d index
func NonDiagSliceNeighbors[T any](ts [][]T, n, m int) []T {
	var out []T
	for i := n - 1; i < n+2; i++ {
		for j := m - 1; j < m+2; j++ {
			if IsSafe(ts, i, j) && !(i == n && j == m) && !(i != n && j != m) { // You are not your own neighbor, ignore diags
				out = append(out, ts[i][j])
			}
		}
	}
	return out
}

// VisitNeighbors iterates over a 2d array, calling a func with each index and
// a list of neighbors.
func VisitNeighbors[T any](ts [][]T, f func(i, j int, ns []T) error) {
	for i, row := range ts {
		for j := range row {
			if f(i, j, SliceNeighbors(ts, i, j)) != nil {
				return
			}
		}
	}
}

// VisitNeighbors iterates over a 2d array, calling a func with each index and
// a list of neighbors.
func VisitNonDiagNeighbors[T any](ts [][]T, f func(i, j int, ns []T) error) {
	for i, row := range ts {
		for j := range row {
			if f(i, j, NonDiagSliceNeighbors(ts, i, j)) != nil {
				return
			}
		}
	}
}
