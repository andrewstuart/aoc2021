package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"

	"astuart.co/advent2021/pkg/ezaoc"
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

	var points []ezaoc.Cell[int]
	ezaoc.VisitNonDiagNeighbors(inputs, func(c ezaoc.Cell[int], ns []ezaoc.Cell[int]) error {
		for _, v := range ns {
			if v.Value <= c.Value {
				return nil
			}
		}
		points = append(points, c)
		return nil
	})

	spew.Dump(points)

	var basins []basin

	for _, pt := range points {
		b := basin{pt}
		visited := ezaoc.Set[[2]int]{}
		root := ezaoc.Queue[ezaoc.Cell[int]]{pt}
		for root.Len() > 0 {
			pt := root.Dequeue()
			ns := ezaoc.NonDiagSliceNeighbors(inputs, pt.I, pt.J)
			for _, cell := range ns {
				if cell.Value < 9 && !visited.Contains(cell.Point()) {
					visited.Add(cell.Point())
					b = append(b, cell)
					root.Enqueue(cell)
				}
			}
		}
		basins = append(basins, b)
	}

	spew.Dump(basins[2])
	for _, b := range basins {
		fmt.Printf("b.size() = %+v\n", b.size())
	}

	sort.Slice(basins, func(i, j int) bool {
		return basins[i].size() < basins[j].size()
	})

	out := 1
	for i := len(basins) - 3; i < len(basins); i++ {
		out *= basins[i].size()
	}
	return out, nil
}

type basin []ezaoc.Cell[int]

func (b basin) size() int {
	return len(b) - 1
}
