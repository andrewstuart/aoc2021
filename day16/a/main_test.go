package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAOC(t *testing.T) {
	var tests = []struct {
		expected int
		given    string
	}{
		{-1, "D2FE28"},
		{-1, "38006F45291200"},
		{16, "8A004A801A8002F478"},
		{12, "620080001611562C8802118E34"},
		{23, "C0015000016115A2E0802F182340"},
		{31, "A0016C880162017C3686B18A3D4780"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.given, func(t *testing.T) {
			asrt := assert.New(t)
			actual, err := aoc(strings.NewReader(tt.given))
			asrt.NoError(err)
			if tt.expected >= 0 {
				asrt.Equal(tt.expected, actual)
			}
		})
	}
}
