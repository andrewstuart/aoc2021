package ezaoc

import (
	"bufio"
	"io"
	"strings"
)

// ReadAOC takes any bufio.Reader (to prevent loss of bytes in io.EOF cases)
// and calls the provided func on every space-trimmed input line, returning a
// slice of that item and any errors encountered. The callee should return
// io.EOF to cease use of the reader, e.g. in the case of header or footer.
func ReadAOC[T any](br bufio.Reader, f func(string) (T, error)) ([]T, error) {
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
