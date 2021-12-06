package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("{{ .Config.Output }}", os.O_RDONLY, 0400)
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

	for {
		st, err := br.ReadString('\n')
		if err != nil {
			log.Println(err)
			break
		}

		// TODO: write string
		log.Println(st)
	}

	return 0, fmt.Errorf("not implemented")
}
