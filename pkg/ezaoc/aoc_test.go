package ezaoc

import (
	"testing"

	"github.com/alecthomas/assert"
)

func TestQueue(t *testing.T) {
	asrt := assert.New(t)
	q := Queue[int]{}
	q.Enqueue(3)
	asrt.Equal(1, len(q))
	asrt.Equal(len(q), q.Len())

	asrt.Equal(3, q.Dequeue())
	asrt.Zero(q.Len())
}
