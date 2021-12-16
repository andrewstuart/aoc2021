package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"astuart.co/advent2020/pkg/myaoc"
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
	inputs, err := myaoc.ReadAOC(r, func(st string) string {
		return st
	})
	if err != nil {
		return 0, err
	}

	spew.Dump(inputs)

	return 0, fmt.Errorf("not implemented")
}
