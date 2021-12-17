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
	asrt.Equal(45, out)

	out, err = aoc(strings.NewReader("target area: x=217..240, y=-126..-69"))
	asrt.NoError(err)
	// TODO replace assert expected value here
	asrt.Equal(5050, out)
}
