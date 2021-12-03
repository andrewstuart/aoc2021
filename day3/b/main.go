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
			break
		}

		strs = append(strs, strings.TrimSpace(st))
	}

	strsa := make([]string, len(strs))
	copy(strsa, strs)

	var a string
	for i := 0; i < len(strs[0]); i++ {
		var zeros, ones int
		for _, st := range strsa {
			if st[i] == '0' {
				zeros++
				continue
			}
			ones++
		}
		var tmp []string
		for _, str := range strsa {
			if zeros > ones && str[i] == '0' {
				tmp = append(tmp, str)
				continue
			}
			if ones >= zeros && str[i] == '1' {
				tmp = append(tmp, str)
			}
		}
		if len(tmp) == 1 {
			a = tmp[0]
			break
		}
		strsa = tmp
	}

	strsb := make([]string, len(strs))
	copy(strsb, strs)

	var b string
	for i := 0; i < len(strs[0]); i++ {
		var zeros, ones int
		for _, st := range strsb {
			if st[i] == '0' {
				zeros++
				continue
			}
			ones++
		}
		var tmp []string
		for _, str := range strsb {
			if zeros > ones && str[i] == '1' {
				tmp = append(tmp, str)
				continue
			}
			if ones >= zeros && str[i] == '0' {
				tmp = append(tmp, str)
			}
		}
		if len(tmp) == 1 {
			b = tmp[0]
			break
		}
		strsb = tmp
	}

	fmt.Printf("a = %+v\n", a)
	fmt.Printf("b = %+v\n", b)

	ai, _ := strconv.ParseInt(a, 2, 64)
	bi, _ := strconv.ParseInt(b, 2, 64)

	fmt.Printf("ai*bi = %+v\n", ai*bi)
}
