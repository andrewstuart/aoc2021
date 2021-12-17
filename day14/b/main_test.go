package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAOC(t *testing.T) {
	asrt, rq := assert.New(t), require.New(t)

	f, err := os.OpenFile("../test", os.O_RDONLY, 0400)
	rq.NoError(err)

	out, err := aoc(bufio.NewReader(f), 40)
	asrt.NoError(err)
	// TODO replace assert expected value here
	asrt.Equal(2188189693529, out)
}

func BenchmarkAOC(b *testing.B) {
	bs, err := ioutil.ReadFile("../test")
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		aoc(bytes.NewReader(bs), 40)
	}
}
