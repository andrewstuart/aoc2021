package ezaoc

import (
	"constraints"
	"strconv"
	"strings"
)

// IntSlicer takes a string and returns a function to slice input strings by
// that, and convert the results to integers.
func IntSlicer(delimiter string) func(string) ([]int, error) {
	return func(st string) ([]int, error) {
		sts := strings.Split(st, delimiter)
		out := make([]int, len(sts))
		var err error
		for i := range sts {
			out[i], err = strconv.Atoi(sts[i])
			if err != nil {
				return nil, err
			}
		}
		return out, nil
	}
}

// MaxOf returns the index and highest valued of the ordered input items based
// on the given func. If the given slice is zero length or nil, the zero values
// of types T and U will be returned.
func MaxOf[T any, U constraints.Ordered](ts []T, value func(T) U) (int, U) {
	if len(ts) == 0 {
		var u U
		return -1, u
	}
	idx := -1
	u := value(ts[0])
	for i, t := range ts {
		ft := value(t)
		if ft < u {
			idx = i
			u = ft
		}
	}
	return idx, u
}
