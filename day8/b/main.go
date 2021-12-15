package main

import (
	"bufio"
	"io"
	"log"
	"math"
	"os"
	"regexp"
	"sort"
	"strings"
)

// var lookup = map[string]int{
// 	"abcefg":  0,
// 	"cg":      1,
// 	"acdeg":   2,
// 	"acdfg":   3,
// 	"bcdf":    4,
// 	"abdfg":   5,
// 	"abdefg":  6,
// 	"acf":     7,
// 	"abcdefg": 8,
// 	"abcdfg":  9,
// }

var randLookup = map[string]string{}

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
			log.Println(st, err)
			break
		}
		strs = append(strs, st)
	}

	out := 0
	for _, st := range strs {
		sts := strings.Fields(st)
		sortStrs(sts)
		rhs := strings.Fields(strings.Split(st, "|")[1])
		sortStrs(rhs)
		m := map[string]int{}
		for _, st := range sts {
			if st == "|" {
				continue
			}

			switch len(st) {
			case 2:
				m[st] = 1
			case 3:
				m[st] = 7
			case 4:
				m[st] = 4
			case 7:
				m[st] = 8
			}
		}
		mInv := invert(m)

		for _, st := range sts {
			if st == "|" {
				continue
			}

			if len(st) == 6 {
				if len(regexp.MustCompile("["+mInv[1]+"]").ReplaceAllString(st, "")) == 5 {
					m[st] = 6
				} else if len(regexp.MustCompile("["+mInv[4]+"]").ReplaceAllString(st, "")) == 2 {
					m[st] = 9
				} else {
					m[st] = 0
				}
			}
			if len(st) == 5 {
				if len(regexp.MustCompile("["+mInv[1]+"]").ReplaceAllString(st, "")) == 3 {
					m[st] = 3
				} else if len(regexp.MustCompile("["+mInv[4]+"]").ReplaceAllString(st, "")) == 3 {
					m[st] = 2
				} else {
					m[st] = 5
				}
			}
		}

		// TODO: convert rhs
		line := 0.0
		for i, st := range rhs {
			n := m[st]
			place := math.Pow(10, 3-float64(i)) * float64(n)
			line += place
		}
		out += int(line)
	}
	return out, nil
}

func sortStrs(strs []string) {
	for i, str := range strs {
		bs := []byte(str)
		sort.Slice(bs, func(i, j int) bool {
			return bs[i] < bs[j]
		})
		strs[i] = string(bs)

	}
}

func invert[T comparable, U comparable](m map[T]U) map[U]T {
	out := map[U]T{}
	for k, v := range m {
		out[v] = k
	}
	return out
}
