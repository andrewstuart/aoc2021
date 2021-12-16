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
		{3, "C200B40A82"},
		{54, "04005AC33890"},
		{7, "880086C3E88112"},
		{9, "CE00C43D881120"},
		{1, "D8005AC2A8F0"},
		{0, "F600BC2D8F"},
		{0, "9C005AC2F8F0"},
		{1, "9C0141080250320F1802104A08"},
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
