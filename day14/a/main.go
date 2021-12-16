package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"

	"astuart.co/advent2020/pkg/myaoc"
)

func main() {
	f, err := os.OpenFile("../input", os.O_RDONLY, 0400)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	br := bufio.NewReader(f)

	out, err := aoc(br, 10)
	if err != nil {
		log.Fatal("error in aoc method: ", err)
	}
	log.Println(out)
}

type Seq struct {
	From, To string
}

func aoc(r io.Reader, n int) (int, error) {
	br := bufio.NewReader(r)
	head, err := br.ReadString('\n')
	if err != nil {
		return 0, fmt.Errorf("error reading header: %w", err)
	}
	head = strings.TrimSpace(head)
	fmt.Printf("strings.TrimSpace(head) = %+v\n", strings.TrimSpace(head))

	inputs, err := myaoc.ReadAOC(br, func(st string) Seq {
		strs := strings.Split(st, " -> ")
		return Seq{
			From: strs[0],
			To:   strs[1],
		}
	})
	if err != nil {
		return 0, err
	}

	m := map[string]string{}
	for _, in := range inputs {
		m[in.From] = in.To
	}

	for i := 0; i < n; i++ {
		tpl := ""
		if i < 5 {
			fmt.Printf("head = %+v\n", head)
		}
		for j := 0; j < len(head)-1; j++ {
			tpl += string([]byte{head[j]})
			lo := head[j : j+2]
			tpl += m[lo]
		}
		tpl += string([]byte{head[len(head)-1]})

		head = tpl
	}

	mm := map[rune]int{}
	for _, ch := range head {
		mm[ch]++
	}
	type charct struct {
		ch rune
		ct int
	}
	var cts []charct
	for ch, ct := range mm {
		cts = append(cts, charct{ch, ct})
	}

	sort.Slice(cts, func(i, j int) bool {
		return cts[i].ct < cts[j].ct
	})

	return cts[len(cts)-1].ct - cts[0].ct, nil
}
