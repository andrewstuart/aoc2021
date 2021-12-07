package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
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
	br := bufio.NewReader(r)

	var pos []int
	for {
		st, err := br.ReadString('\n')

		if st != "" {
			for _, ist := range strings.Split(strings.TrimSpace(st), ",") {
				if i, err := strconv.Atoi(ist); err == nil {
					pos = append(pos, i)
				}
			}
		}

		if err != nil {
			log.Println(err)
			break
		}
	}

	sort.Ints(pos)
	n := pos[len(pos)/2] // lazy median
	if len(pos)%2 == 1 {
		n += pos[len(pos)/2+1]
		n /= 2
	}
	fmt.Printf("n = %+v\n", n)

	diff := 0.0
	for _, p := range pos {
		diff += math.Abs(float64(p) - float64(n))
	}

	return int(diff), nil
}
