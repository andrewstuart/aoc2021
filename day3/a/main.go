package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	var strs []string
	for {
		st, err := br.ReadString('\n')
		if err != nil {
			if st != "" {
				fmt.Printf("st = %+v\n", st)
				strs = append(strs, st)
			}
			fmt.Printf("err = %+v\n", err)
			break
		}

		strs = append(strs, strings.TrimSpace(st))
	}

	fmt.Printf("len(strs[0]) = %+v\n", len(strs[0]))
	fmt.Printf("len(strs) = %+v\n", len(strs))

	var gamma, eps string
	for i := 0; i < len(strs[0]); i++ {
		var zeros, ones int
		for _, st := range strs {
			if st[i] == '0' {
				zeros++
				continue
			}
			ones++
		}
		fmt.Printf("%d.\t%d\t%d\n", i, zeros, ones)
		if zeros > ones {
			gamma += "0"
			eps += "1"
			continue
		}
		gamma += "1"
		eps += "0"
	}

	fmt.Printf("a = %+v\n", gamma)
	fmt.Printf("b = %+v\n", eps)

	ai, _ := strconv.ParseInt(gamma, 2, 64)
	bi, _ := strconv.ParseInt(eps, 2, 64)
	fmt.Printf("ai = %+v\n", ai)
	fmt.Printf("bi = %+v\n", bi)

	fmt.Printf("ai*bi = %+v\n", ai*bi)
}
