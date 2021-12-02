package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func maina() {
	f, err := os.OpenFile("input", os.O_RDONLY, 0400)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	br := bufio.NewReader(f)

	last := 0
	n := 0
	for i := 0; ; i++ {
		st, err := br.ReadString('\n')
		if err != nil {
			log.Println(n)
			return
		}
		num, err := strconv.Atoi(strings.TrimSpace(st))
		if err != nil {
			log.Fatal(err)
		}
		if i > 0 && num > last {
			n++
		}
		last = num
	}
}
