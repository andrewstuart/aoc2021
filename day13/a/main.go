package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"astuart.co/advent2020/pkg/ezaoc"
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

type Fold struct {
	Direction string
	Index     int
}

func aoc(r io.Reader) (int, error) {
	maxx, maxy := 0, 0
	br := bufio.NewReader(r)
	inputs, err := ezaoc.ReadAOC(br, func(st string) ([2]int, error) {
		if st == "" {
			return [2]int{}, io.EOF
		}
		var a [2]int
		fmt.Sscanf(st, "%d,%d", &a[0], &a[1])
		if a[0] > maxx {
			maxx = a[0]
		}
		if a[1] > maxy {
			maxy = a[1]
		}
		return a, nil
	})
	if err != nil {
		return 0, fmt.Errorf("error reading points: %w", err)
	}

	folds, err := ezaoc.ReadAOC(br, func(st string) (Fold, error) {
		var f Fold
		if st == "" {
			return f, io.EOF
		}

		fs := strings.Fields(st)
		last := fs[len(fs)-1]

		strs := strings.Split(last, "=")
		f.Direction = strs[0]
		f.Index, _ = strconv.Atoi(strs[1])
		return f, nil
	})
	if err != nil {
		return 0, fmt.Errorf("error reading folds: %w", err)
	}

	s := ezaoc.SetFrom(inputs)

	paper := ezaoc.Make2DSlice(maxx+1, maxy+1, func(i, j int) bool {
		return s.Contains([2]int{i, j})
	})

	i := 0
	printPaper := func() {
		i = 0
		for x := range paper[0] {
			for y := range paper {
				if paper[y][x] {
					i++
					fmt.Print("X")
					continue
				}
				fmt.Print(".")
			}
			fmt.Println()
		}
	}

	for i, fold := range folds {
		fmt.Printf("\nFold %d on %s at %d.\n\n", i+1, fold.Direction, fold.Index)
		if fold.Direction == "y" {
			for y := range paper {
				for x := range paper[y] {
					xDest := x
					if x > fold.Index {
						xDest = 2*fold.Index - x
						if xDest < 0 {
							continue
						}
					}
					paper[y][xDest] = paper[y][xDest] || paper[y][x]
				}
				paper[y] = paper[y][:fold.Index]
			}
			continue
		}
		for y, row := range paper {
			for x := range row {
				yDest := y
				if y > fold.Index {
					yDest = 2*fold.Index - y
					if yDest < 0 {
						continue
					}
				}
				paper[yDest][x] = paper[yDest][x] || paper[y][x]
			}
		}
		paper = paper[:fold.Index]
	}
	printPaper()

	return i, nil
}
