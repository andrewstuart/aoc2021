package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

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

type Area struct {
	X0, XN, Y0, YN int
}

func aoc(r io.Reader) (int, error) {
	var a Area

	_, err := fmt.Fscanf(bufio.NewReader(r), "target area: x=%d..%d, y=%d..%d", &a.X0, &a.XN, &a.Y0, &a.YN)
	if err != nil {
		return 0, err
	}
	spew.Dump(a)

	var vy int
	var okMax, okvy int
	for {
		vy++

		var y int
		maxH := 0
		vy0 := vy
		for {
			y += vy0
			vy0--
			if y > maxH {
				maxH = y
			}
			if y >= a.Y0 && y <= a.YN {
				if maxH > okMax {
					okMax = maxH
					okvy = vy
				}
			}
			if y < a.Y0 {
				break
			}
		}

		if vy > 500 {
			break
		}
	}
	fmt.Printf("okMax = %+v\n", okMax)

	var vx int
l:
	for {
		vx++

		vx0 := vx
		for x := 0; vx0 > 0 && x <= a.XN; x += vx0 {
			if vx0 > 0 {
				vx0--
			} else if vx < 0 {
				vx0++
			}
			if x >= a.X0 && x <= a.XN {
				break l
				// enough = true
				// ok = true
			}
		}
		if vx > 500 {
			break
		}
	}
	fmt.Printf("okvx = %+v\n", vx-1)
	fmt.Printf("okvy = %+v\n", okvy)
	// fmt.Printf("okvx = %+v\n", vx-2)

	return okMax, nil
}
