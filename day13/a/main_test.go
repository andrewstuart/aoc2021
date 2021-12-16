package main

import (
	"bufio"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAOC(t *testing.T) {
	asrt, rq := assert.New(t), require.New(t)

	f, err := os.OpenFile("../test", os.O_RDONLY, 0400)
	rq.NoError(err)

	out, err := aoc(bufio.NewReader(f))
	asrt.NoError(err)
	// TODO replace assert expected value here
	asrt.Equal(17, out)
}
