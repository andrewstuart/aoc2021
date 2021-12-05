package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/sirupsen/logrus"
)

type Point struct {
	X, Y int
}

type Line struct {
	A, B Point
}

func main() {
	f, err := os.OpenFile("../input", os.O_RDONLY, 0400)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	br := bufio.NewReader(f)

	log.Println(aoc(br))
}

func aoc(r io.Reader) (int, error) {
	br := bufio.NewReader(r)

	var maxx, maxy int
	var lines []Line
	for {
		st, err := br.ReadString('\n')
		if err != nil {
			log.Println(err)
			break
		}
		var l Line
		_, err = fmt.Sscanf(st, "%d,%d -> %d,%d", &l.A.X, &l.A.Y, &l.B.X, &l.B.Y)
		if err != nil {
			logrus.WithError(err).Fatal("error scanning")
		}

		if l.A.X > maxx {
			maxx = l.A.X
		}
		if l.B.X > maxx {
			maxx = l.B.X
		}
		if l.A.Y > maxy {
			maxy = l.A.Y
		}
		if l.B.Y > maxy {
			maxy = l.B.Y
		}

		lines = append(lines, l)
	}

	grid := make([][]int, maxx+1)
	for i := range grid {
		grid[i] = make([]int, maxy+1)
	}

	cols := 0
	for _, l := range lines {
		if l.A.X == l.B.X {
			x := l.A.X
			a, b := order(l.A.Y, l.B.Y)
			for i := a; i <= b; i++ {
				if grid[x][i] == 1 {
					cols++
				}
				grid[x][i]++
			}
		}

		if l.A.Y == l.B.Y {
			y := l.A.Y
			a, b := order(l.A.X, l.B.X)
			for i := a; i <= b; i++ {
				if grid[i][y] == 1 {
					cols++
				}
				grid[i][y]++
			}
		}
	}

	return cols, nil
}

func order(a, b int) (int, int) {
	if a > b {
		return b, a
	}
	return a, b
}
