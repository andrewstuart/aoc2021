package main

import (
	"bufio"
	"os"
	"strings"
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
	asrt.Equal(61229, out)
}

func TestAOC2(t *testing.T) {
	asrt := assert.New(t)

	out, err := aoc(strings.NewReader("acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf\n"))
	asrt.NoError(err)
	// TODO replace assert expected value here
	asrt.Equal(5353, out)
}
