package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Cell struct {
	i, j, hi, n int
	cost        float64
	prev        *Cell
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
	c[j].hi = i
	c[i].hi = j
}

func (c *Cells) Push(x any) {
	cell := x.(*Cell)
	cell.hi = len(*c)
	*c = append(*c, cell)

}

func (c *Cells) Pop() any {
	last := len(*c) - 1
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

	// graph := myaoc.Make2DSlice(len(input), len(input[0]), func(i, j int) int {
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

	// h := &Cells{}

	// graph := myaoc.Make2DSlice(len(input), len(input[0]), func(i, j int) *Cell {
	// 	cell := &Cell{
	// 		i:    i,
	// 		j:    j,
	// 		hi:   h.Len(),
	// 		n:    int(input[i][j] - 48),
	// 		cost: math.Inf(1),
	// 	}
	// 	heap.Push(h, cell)
	// 	return cell
	// })

	// var i, j int

	// graph[0][0].cost = float64(graph[0][0].n)
	// heap.Init(h)

	// var path []*Cell
	// seen := myaoc.Set[[2]int]{}

	// for {
	// 	next := heap.Pop(h).(*Cell)
	// 	seen.Add([2]int{next.i, next.j})
	// 	for _, n := range myaoc.SliceNeighbors(graph, next.i, next.j) {
	// 		if !seen.Contains([2]int{n.i, n.j}) {
	// 			n.prev = next
	// 			n.cost = float64(n.n)
	// 			heap.Fix(h, n.hi)
	// 		}
	// 	}
	// }

	return 0, fmt.Errorf("not implemented")
}
