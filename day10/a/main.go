package main

import (
	"bufio"
	"io"
	"log"
	"os"

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
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
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
	errs := 0
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
					errs += scores[ch]
					log.Printf("error, expected %c found %c\n", expected, ch)
					continue ins
				}
			}
		}
	}

	return errs, nil
}
