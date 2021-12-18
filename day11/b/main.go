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

func aoc(r io.Reader) (int, error) {
	inputs, err := ezaoc.ReadAOC(r, ezaoc.IntSlicer(""))
	if err != nil {
		return 0, err
	}

	ct := len(inputs) * len(inputs[0])
	flashes := 0
	i := 0
	for flashes < ct {
		flashes = 0
		i++
		flashed := ezaoc.Set[[2]int]{}
		toFlash := ezaoc.Queue[ezaoc.Cell[int]]{}

		ezaoc.VisitCells(inputs, func(c ezaoc.Cell[int]) error {
			inputs[c.I][c.J]++
			if inputs[c.I][c.J] > 9 {
				toFlash.Enqueue(c)
			}
			return nil
		})

		for len(toFlash) > 0 {
			c := toFlash.Dequeue()
			if flashed.Contains(c.Point()) {
				continue
			}
			flashed.Add(c.Point())

			inputs[c.I][c.J]++
			for _, neigh := range ezaoc.SliceNeighbors(inputs, c.I, c.J) {
				inputs[neigh.I][neigh.J]++
				if inputs[neigh.I][neigh.J] > 9 && !flashed.Contains(neigh.Point()) {
					toFlash.Enqueue(neigh)
				}
			}
		}

		ezaoc.VisitCells(inputs, func(c ezaoc.Cell[int]) error {
			if c.Value > 9 {
				flashes++
				inputs[c.I][c.J] = 0
			}
			return nil
		})
	}

	return i, nil
}
