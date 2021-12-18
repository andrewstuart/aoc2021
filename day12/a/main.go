package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"

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

	out, err := aoc(br)
	if err != nil {
		log.Fatal("error in aoc method: ", err)
	}
	log.Println(out)
}

type Vtx struct {
	Name string
	ct   int
}

func (v Vtx) small() bool {
	for _, ch := range v.Name {
		if 'A' <= ch && ch <= 'Z' {
			return false
		}
	}
	return true
}

type Edge struct {
	From, To *Vtx
}

func aoc(r io.Reader) (int, error) {
	inputs, err := ezaoc.ReadAOC(r, func(st string) (Edge, error) {
		if st == "" {
			return Edge{}, io.EOF
		}
		sts := strings.Split(st, "-")
		return Edge{
			From: &Vtx{Name: sts[0]},
			To:   &Vtx{Name: sts[1]},
		}, nil
	})
	if err != nil {
		return 0, err
	}

	m := ezaoc.GroupByFunc(inputs, func(e Edge) (string, *Vtx) {
		return e.From.Name, e.To
	})
	mb := ezaoc.GroupByFunc(inputs, func(e Edge) (string, *Vtx) {
		return e.To.Name, e.From
	})
	for k, v := range mb {
		m[k] = append(m[k], v...)
	}
	spew.Dump(m)

	paths := dfs(&Vtx{Name: "start"}, "end", m)

	ezaoc.Print2dGrid(paths[:5])

	return len(paths), nil
}

func dfs(from *Vtx, to string, edges map[string][]*Vtx) [][]*Vtx {
	if from.Name == "start" && from.ct > 0 {
		return nil
	}
	if from.small() && from.ct >= 2 {
		return nil
	}
	from.ct++
	defer func() {
		from.ct--
	}()
	// Exit cond
	if from.Name == to {
		return [][]*Vtx{{from}}
	}

	var paths [][]*Vtx
	for _, v := range edges[from.Name] {
		paths2 := dfs(v, to, edges)
		for i := range paths {
			paths[i] = append([]*Vtx{from}, paths[i]...)
		}
		paths = append(paths, paths2...)
	}
	return paths
}
