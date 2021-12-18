package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/andrewstuart/aoc2021/pkg/ezaoc"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	f, err := os.OpenFile("../input", os.O_RDONLY, 0400)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	br := bufio.NewReader(f)

	log.Println(aoc(br))
}

func aoc(r io.Reader) (int) {
	inputs, err := ezaoc.ReadAOC(r, func(st string) (string, error) {
    if st == "" {
      return st, io.EOF
    }
		return st, nil
	})
	if err != nil {
		return 0, err
	}


  // Add challenge logic here probably
  count := 0
	spew.Dump(inputs)
  count = len(inputs)

	return count
}
