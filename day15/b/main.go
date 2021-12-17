package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strings"

	"astuart.co/advent2020/pkg/ezaoc"
)

type Cell struct {
	i, j, heapIdx, n int
	cost             float64
	prev             *Cell
}

type Cells []*Cell

// Len is the number of elements in the collection.
func (c Cells) Len() int {
	return len(c)
}

// Less reports whether the element with index i
// must sort before the element with index j.
//
// If both Less(i, j) and Less(j, i) are false,
// then the elements at index i and j are considered equal.
// Sort may place equal elements in any order in the final result,
// while Stable preserves the original input order of equal elements.
//
// Less must describe a transitive ordering:
//  - if both Less(i, j) and Less(j, k) are true, then Less(i, k) must be true as well.
//  - if both Less(i, j) and Less(j, k) are false, then Less(i, k) must be false as well.
//
// Note that floating-point comparison (the < operator on float32 or float64 values)
// is not a transitive ordering when not-a-number (NaN) values are involved.
// See Float64Slice.Less for a correct implementation for floating-point values.
func (c Cells) Less(i int, j int) bool {
	return c[i].cost < c[j].cost
}

// Swap swaps the elements with indexes i and j.
func (c Cells) Swap(i int, j int) {
	c[i], c[j] = c[j], c[i]
	c[j].heapIdx = j
	c[i].heapIdx = i
}

func (c *Cells) Push(x any) {
	cell := x.(*Cell)
	cell.heapIdx = len(*c)
	*c = append(*c, cell)
}

func (c *Cells) Pop() any {
	last := c.Len() - 1
	out := (*c)[last]
	*c = (*c)[:last]
	return out
}

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
	br := bufio.NewReader(r)

	var input []string
	for {
		st, err := br.ReadString('\n')
		if err != nil {
			log.Println(err)
			break
		}

		input = append(input, strings.TrimSpace(st))
	}

	h := &Cells{}

	x, y := len(input), len(input[0])
	graph := ezaoc.Make2DSlice(5*x, 5*y, func(i, j int) *Cell {
		in := int(input[i%x][j%y] - 48)
		in += i / x
		in += j / y
		if in > 9 {
			in = 1 + (in % 10)
		}

		cell := &Cell{
			i:       i,
			j:       j,
			heapIdx: h.Len(),
			n:       in,
			cost:    math.Inf(1),
		}
		heap.Push(h, cell)
		return cell
	})

	// for _, g := range graph {
	// 	for _, n := range g {
	// 		fmt.Print(n.n)
	// 	}
	// 	fmt.Println()
	// }

	graph[0][0].cost = float64(graph[0][0].n)
	heap.Init(h)

	seen := ezaoc.Set[[2]int]{}

	for h.Len() > 0 {
		next := heap.Pop(h).(*Cell)
		if next.i == len(graph) && next.j == len(graph[0]) {
			break
		}
		seen.Add([2]int{next.i, next.j})
		for _, n := range ezaoc.SliceNeighbors(graph, next.i, next.j) {
			if !(next.i == n.I || next.j == n.J) { // constrain to sideways
				continue
			}
			if !seen.Contains(n.Point()) && next.cost+float64(n.Value.n) <= n.Value.cost {
				n.Value.prev = next
				n.Value.cost = float64(n.Value.n) + next.cost
				heap.Fix(h, n.Value.heapIdx)
			}
		}
	}

	last := graph[len(graph)-1][len(graph[0])-1]
	cost := last.n
	for last.prev != nil {
		fmt.Printf("last.prev = %+v\n", last.prev)
		cost += last.prev.n
		last = last.prev
	}

	return int(graph[len(graph)-1][len(graph[0])-1].cost) - graph[0][0].n, nil
}
