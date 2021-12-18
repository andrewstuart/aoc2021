package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"astuart.co/advent2021/pkg/ezaoc"
	"github.com/davecgh/go-spew/spew"
)

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
	inputs, err := ezaoc.ReadAOC(r, func(st string) (string, error) {
    if st == "" {
      return st, io.EOF
    }
		return st, nil
	})
	if err != nil {
		return 0, err
	}

	spew.Dump(inputs)

	return 0, fmt.Errorf("not implemented")
}
