package ezaoc

import (
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
