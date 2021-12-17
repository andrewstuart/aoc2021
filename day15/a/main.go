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

	// graph := ezaoc.Make2DSlice(len(input), len(input[0]), func(i, j int) int {
	// 	return int(input[i][j] - 48)
	// })

	// var i, j, cost int
	// for {
	// 	cost += graph[i][j]
	// 	fmt.Printf("%d,%d: %d\n", i, j, graph[i][j])
	// 	if i+1 >= len(graph) && j+1 >= len(graph[0]) {
	// 		break
	// 	}
	// 	if i+1 >= len(graph) {
	// 		j++
	// 		continue
	// 	}
	// 	if j+1 >= len(graph[0]) {
	// 		i++
	// 		continue
	// 	}
	// 	j1 := graph[i][j+1]
	// 	fmt.Printf("j1 = %+v\n", j1)
	// 	i1 := graph[i+1][j]
	// 	fmt.Printf("i1 = %+v\n", i1)
	// 	if j1 < i1 {
	// 		j++
	// 		continue
	// 	}
	// 	i++
	// }
	// return cost, nil

	h := &Cells{}

	graph := ezaoc.Make2DSlice(len(input), len(input[0]), func(i, j int) *Cell {
		cell := &Cell{
			i:       i,
			j:       j,
			heapIdx: h.Len(),
			n:       int(input[i][j] - 48),
			cost:    math.Inf(1),
		}
		heap.Push(h, cell)
		return cell
	})

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
			if !seen.Contains([2]int{n.I, n.J}) && next.cost+float64(n.Value.n) <= n.Value.cost {
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
