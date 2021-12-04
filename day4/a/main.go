package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Cell struct {
	N      int
	Marked bool
}

type Board struct {
	Cells [5][5]Cell
}

func (b *Board) sum() int {
	sum := 0
	for _, r := range b.Cells {
		for _, cell := range r {
			if !cell.Marked {
				sum += cell.N
			}
		}
	}
	return sum
}

func (b *Board) checkRows() bool {
top:
	for _, col := range b.Cells {
		for _, c := range col {
			if !c.Marked {
				continue top // skip rows
			}
		}
		// End of row, nothing unmarked
		return true
	}
	// end of cols, haven't returned true yet
	return false
}

func (b *Board) checkCols() bool {
top:
	for i := 0; i < len(b.Cells); i++ {
		for j := 0; j < len(b.Cells[0]); j++ {
			if !b.Cells[j][i].Marked {
				continue top
			}
		}
		// End of col, nothing unmarked
		return true
	}
	// end of rows, haven't returned true yet
	return false
}

func (b *Board) mark(n int) bool {
	for i, r := range b.Cells {
		for j, cell := range r {
			if cell.N == n {
				b.Cells[i][j].Marked = true
			}
		}
	}

	return b.checkRows() || b.checkCols()
}

func main() {
	f, err := os.OpenFile("../input", os.O_RDONLY, 0400)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	br := bufio.NewReader(f)

	var ints []int
	st, err := br.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	for _, ist := range strings.Split(strings.TrimSpace(st), ",") {
		i, _ := strconv.Atoi(strings.TrimSpace(ist))
		ints = append(ints, i)
	}

	var boards []*Board
	board := &Board{}
	var i int
	for {
		st, err := br.ReadString('\n')
		if err != nil {
			boards = append(boards, board)
			break
		}
		fs := strings.Fields(st)
		if len(fs) == 0 {
			if *board != (Board{}) {
				boards = append(boards, board)
			}
			board = &Board{}
			i = 0
			continue
		}

		for j, f := range fs {
			k, _ := strconv.Atoi(f)
			board.Cells[i][j].N = k
		}
		i++
	}

	for _, n := range ints {
		for _, b := range boards {
			if b.mark(n) {
				log.Println(boards[0].sum() * n)
				os.Exit(0)
			}
		}
	}

	os.Exit(1)
}
