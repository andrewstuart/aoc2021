package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	t := time.Now()
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
	log.Println(time.Since(t))
}

type Area struct {
	X0, XN, Y0, YN int
}

func aoc(r io.Reader) (int, error) {
	var a Area

	_, err := fmt.Fscanf(bufio.NewReader(r), "target area: x=%d..%d, y=%d..%d", &a.X0, &a.XN, &a.Y0, &a.YN)
	if err != nil {
		return 0, err
	}

	var ct int
	for i := 0; i < a.XN+1; i++ {
	l:
		for j := a.Y0 - 1; j < 300; j++ {
			vx := i
			vy := j

			var x, y int
			for {
				if y < a.Y0 || x > a.XN {
					break
				}
				y += vy
				x += vx
				vy--
				if vx > 0 {
					vx--
				} else if vx < 0 {
					vx++
				}
				if a.Y0 <= y && y <= a.YN && x >= a.X0 && x <= a.XN {
					ct++
					continue l
				}
			}
		}
	}

	return ct, nil
}
