package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"sort"

	"github.com/andrewstuart/aoc2021/pkg/ezaoc"
)

func main() {
	f, err := os.OpenFile("../input", os.O_RDONLY, 0400)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	br := bufio.NewReader(f)

	out, err := aoc(br)
	if err != nil {
		log.Fatal("error in aoc method: ", err)
	}
	log.Println(out)
}

var scores = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

func aoc(r io.Reader) (int, error) {
	inputs, err := ezaoc.ReadAOC(r, func(st string) (string, error) {
		if st == "" {
			return st, io.EOF
		}
		return st, nil
	})
	if err != nil {
		return 0, err
	}
	var autocompletes []int
ins:
	for _, in := range inputs {
		s := ezaoc.Stack[rune]{}
		for _, ch := range in {
			switch ch {
			case '<':
				s.Push('>')
			case '[':
				s.Push(']')
			case '(':
				s.Push(')')
			case '{':
				s.Push('}')
			default:
				expected := s.Pop()
				if ch != expected {
					continue ins
				}
			}
		}
		score := 0
		for len(s) > 0 {
			score = score*5 + scores[s.Pop()]
		}
		if score > 0 {
			autocompletes = append(autocompletes, score)
		}
	}

	sort.Ints(autocompletes)

	return Median(autocompletes), nil
}

func Median(ts []int) int {
	d2 := ts[len(ts)/2]
	if len(ts)&1 == 1 {
		return d2
	}
	return (d2 + ts[len(ts)/2+1]) / 2
}
