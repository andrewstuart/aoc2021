package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("../input", os.O_RDONLY, 0400)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	br := bufio.NewReader(f)

	var x, y int
	for i := 0; ; i++ {
		st, err := br.ReadString('\n')
		if err != nil {
			break
		}

		dir, n := "", 0
		fmt.Sscanf(st, "%s %d", &dir, &n)

		fmt.Printf("dir = %+v\n", dir)
		fmt.Printf("n = %+v\n", n)

		switch dir {
		case "forward":
			x += n
		case "down":
			y += n
		case "up":
			y -= n
		}
	}

	fmt.Printf("x = %+v\n", x)
	fmt.Printf("y = %+v\n", y)
	fmt.Printf("x*y = %+v\n", x*y)
}
