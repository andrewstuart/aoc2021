package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/andrewstuart/advent2021/pkg/ezaoc"
)

func main() {
	f, err := os.OpenFile("../input", os.O_RDONLY, 0400)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	br := bufio.NewReader(f)

	t := time.Now()
	out, err := aoc(br, 40)
	if err != nil {
		log.Fatal("error in aoc method: ", err)
	}
	log.Println(out)
	fmt.Println(time.Since(t))
}

type Seq struct {
	From, To string
}

func aoc(r io.Reader, n int) (int, error) {
	t := time.Now()
	br := bufio.NewReader(r)
	head, err := br.ReadString('\n')
	if err != nil {
		return 0, fmt.Errorf("error reading header: %w", err)
	}
	head = strings.TrimSpace(head)
	br.ReadString('\n')

	inputs, err := ezaoc.ReadAOC(br, func(st string) (Seq, error) {
		if st == "" {
			return Seq{}, io.EOF
		}
		strs := strings.Split(st, " -> ")
		return Seq{
			From: strs[0],
			To:   strs[1],
		}, nil
	})
	if err != nil {
		return 0, err
	}
	fmt.Println(time.Since(t))

	m := map[string]string{}
	for _, in := range inputs {
		m[in.From] = in.To
	}

	counts := map[string]int{}
	syllables := map[string]int{}
	for j := 0; j < len(head)-1; j++ {
		syllables[head[j:j+2]]++
	}
	for _, ch := range head {
		counts[string([]rune{ch})]++
	}

	t = time.Now()
	// Main loop
	for i := 0; i < n; i++ {
		inc := map[string]int{}

		for syl, ct := range syllables {
			ch := m[syl]
			counts[ch] += ct
			a, b := syl[:1]+ch, ch+syl[1:]
			inc[a] += ct // Two new ones appear
			inc[b] += ct
		}
		syllables = inc
	}
	fmt.Printf("time.Since(t): %v\n", time.Since(t))

	type charct struct {
		ch string
		ct int
	}
	var cts []charct
	for ch, ct := range counts {
		cts = append(cts, charct{ch, ct})
	}

	sort.Slice(cts, func(i, j int) bool {
		return cts[i].ct < cts[j].ct
	})

	return cts[len(cts)-1].ct - cts[0].ct, nil
}
