package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	f, err := os.OpenFile("../input", os.O_RDONLY, 0400)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	br := bufio.NewReader(f)

	out, err := aoc(br, days)
	if err != nil {
		log.Fatal("error in aoc method: ", err)
	}
	log.Println(out)
}

const days = 256

func aoc(r io.Reader, n int) (int, error) {
	br := bufio.NewReader(r)

	var fish [9]int
	for {
		st, err := br.ReadString('\n')

		if st != "" {
			istrs := strings.Split(strings.TrimSpace(st), ",")
			for _, str := range istrs {
				if i, err := strconv.Atoi(strings.TrimSpace(str)); err == nil {
					fish[i]++
				}
			}
		}
		if err != nil {
			log.Println(err)
			break
		}
	}

	spew.Dump(fish)

	for i := 0; i < n; i++ {
		spawns := fish[0]
		for j := 1; j < len(fish); j++ {
			fish[j-1] = fish[j]
		}
		fish[8] = spawns
		fish[6] += spawns
	}

	count := 0
	for _, n := range fish {
		count += n
	}

	return count, nil
}
