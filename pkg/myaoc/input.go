package myaoc

import (
	"bufio"
	"io"
	"strings"
)

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
