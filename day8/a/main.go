package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

var lookup = map[string]int{
	"abcefg":  0,
	"cf":      1,
	"acdeg":   2,
	"acdfg":   3,
	"bcdf":    4,
	"abdfg":   5,
	"abdefg":  6,
	"acf":     7,
	"abcdefg": 8,
	"abcdfg":  9,
}

var randLookup = make(map[string]int)

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

	var strs []string
	for {
		st, err := br.ReadString('\n')
		if err != nil {
			log.Println(err)
			break
		}
		strs = append(strs, st)
	}

	n := 0
	for _, st := range strs {
		for _, st := range strings.Fields(strings.Split(st, "|")[1]) {
			if len(st) <= 4 || len(st) == 7 {
				n++
			}
			bs := []byte(st)
			sort.Slice(bs, func(i, j int) bool {
				return bs[i] < bs[j]
			})

			fmt.Printf("%s: %d\n", string(bs), lookup[string(bs)])
		}
	}

	return n, nil
}
