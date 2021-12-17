package main

import (
	"bufio"
	"io"
	"log"
	"os"

	"astuart.co/advent2020/pkg/ezaoc"
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
	ezaoc.VisitNonDiagNeighbors(inputs, func(i, j int, ns []int) error {
		val := inputs[i][j]
		for _, v := range ns {
			if v <= val {
				return nil
			}
		}
		out += 1 + val
		return nil
	})

	return out, nil
}
