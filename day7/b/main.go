package main

import (
	"bufio"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
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
			if err != io.EOF {
				logrus.WithError(err).Error("error reading")
			}
			break
		}
	}

	max := 0
	for _, n := range pos {
		if n > max {
			max = n
		}
	}

	cost := math.MaxInt
	for p := 0; p < max; p++ {
		thisCost := 0
		for _, pp := range pos {
			n := pp - p
			if n < 0 {
				n = -n
			}
			cost := n * (n + 1) / 2
			thisCost += cost
		}
		if thisCost < cost {
			cost = thisCost
		}
	}

	return cost, nil
}
