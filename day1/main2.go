package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.OpenFile("input", os.O_RDONLY, 0400)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	br := bufio.NewReader(f)

	ms := []int{}
	for {
		st, err := br.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		num, err := strconv.Atoi(strings.TrimSpace(st))
		if err != nil {
			log.Fatal(err)
		}
		ms = append(ms, num)
	}

	last := 0
	n := 0
	for i := 0; i < len(ms)-2; i++ {
		sum := ms[i] + ms[i+1] + ms[i+2]
		if i > 0 {
			if sum > last {
				n++
			}
		}
		last = sum
	}
	fmt.Printf("n = %+v\n", n)
}
