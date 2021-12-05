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
		if l.A.X != l.B.X && l.A.Y != l.B.Y {
			xdir, ydir := 1, 1
			if l.A.X > l.B.X {
				xdir = -1
			}
			if l.A.Y > l.B.Y {
				ydir = -1
			}

			for x, y := l.A.X, l.A.Y; ; x, y = x+xdir, y+ydir {
				if grid[x][y] == 1 {
					cols++
				}
				grid[x][y]++
				if x == l.B.X && y == l.B.Y {
					break
				}
			}
			continue
		}

		if l.A.X == l.B.X {
			x := l.A.X
			a, b := order(l.A.Y, l.B.Y)
			for y := a; y <= b; y++ {
				if grid[x][y] == 1 {
					cols++
				}
				grid[x][y]++
			}
			continue
		}

		if l.A.Y == l.B.Y {
			y := l.A.Y
			a, b := order(l.A.X, l.B.X)
			for x := a; x <= b; x++ {
				if grid[x][y] == 1 {
					cols++
				}
				grid[x][y]++
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
