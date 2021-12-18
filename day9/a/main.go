package main

import (
	"bufio"
	"io"
	"log"
	"os"

	"github.com/andrewstuart/aoc2021/pkg/ezaoc"
	"github.com/davecgh/go-spew/spew"
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

func aoc(r io.Reader) (int, error) {
	inputs, err := ezaoc.ReadAOC(r, ezaoc.IntSlicer(""))
	if err != nil {
		return 0, err
	}

	spew.Dump(inputs)

	out := 0
	ezaoc.VisitNonDiagNeighbors(inputs, func(c ezaoc.Cell[int], ns []ezaoc.Cell[int]) error {
		for _, v := range ns {
			if v.Value <= c.Value {
				return nil
			}
		}
		out += 1 + c.Value
		return nil
	})

	return out, nil
}
