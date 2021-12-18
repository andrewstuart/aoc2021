package main

import (
	"bufio"
	"encoding/hex"
	"io"
	"log"
	"math"
	"os"

	"astuart.co/advent2021/pkg/ezaoc"
	"github.com/oyi812/bitpack"
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

type packet struct {
	Version uint64 `bits:"3"`
	TypeID  uint64 `bits:"4"`
	Value   uint64
	Sub     []packet
}

const (
	typeAdd  = 0
	typeMult = 1
	typeMin  = 2
	typeMax  = 3
	typeVal  = 4
	typeGT   = 5
	typeLT   = 6
	typeEQ   = 7
)

func (p packet) eval() uint64 {
	if p.TypeID == typeVal {
		return p.Value
	}

	var out uint64
	switch p.TypeID {
	case typeMin:
		out = math.MaxUint64
	case typeMult:
		out = 1
	case typeGT:
		if p.Sub[0].eval() > p.Sub[1].eval() {
			return 1
		}
		return 0
	case typeLT:
		if p.Sub[0].eval() < p.Sub[1].eval() {
			return 1
		}
		return 0
	case typeEQ:
		if p.Sub[0].eval() == p.Sub[1].eval() {
			return 1
		}
		return 0
	}

	for _, sub := range p.Sub {
		switch p.TypeID {
		case typeAdd:
			out += sub.eval()
		case typeMult:
			out *= sub.eval()
		case typeMin:
			if e := sub.eval(); e < out {
				out = e
			}
		case typeMax:
			if e := sub.eval(); e > out {
				out = e
			}
		}
	}
	return out
}

func aoc(r io.Reader) (int, error) {
	inputs, err := ezaoc.ReadAOC(r, func(st string) ([]byte, error) {
		if st == "" {
			return nil, io.EOF
		}
		bs := []byte(st)
		n, err := hex.Decode(bs, bs)
		if err != nil {
			return nil, err
		}
		return bs[:n], nil
	})
	if err != nil {
		return 0, err
	}

	var sum uint64
	for _, in := range inputs {
		p, err := parsePacket(in)
		if err != nil {
			log.Fatal(err)
		}
		sum += p.eval()
	}

	return int(sum), nil
}

func parsePacket(bs []byte) (*packet, error) {
	packet, _, err := parse(bs, 0)
	if err != io.EOF {
		return nil, err
	}
	return packet, nil
}

func parse(bs []byte, o int) (*packet, int, error) {
	var v uint64
	v, o = bitpack.Unpack(bs, o, 3)
	var t uint64
	t, o = bitpack.Unpack(bs, o, 3)
	p := packet{
		Version: v,
		TypeID:  t,
	}

	switch p.TypeID {
	case 4:
		var out uint64
		for {
			var contBit uint64
			contBit, o = bitpack.Unpack(bs, o, 1)

			var next uint64
			next, o = bitpack.Unpack(bs, o, 4)
			out = out << 4
			out |= next
			if contBit == 0 {
				break
			}
		}
		p.Value = out
	default:
		var lID uint64
		lID, o = bitpack.Unpack(bs, o, 1)
		switch lID {
		case 0:
			var bitLen uint64
			bitLen, o = bitpack.Unpack(bs, o, 15)
			start := o
			for uint64(o-start) < bitLen {
				sub, off, err := parse(bs, o)
				o = off
				if err != nil && err != io.EOF {
					return nil, o, err
				}
				p.Sub = append(p.Sub, *sub)
			}
		case 1:
			var count uint64
			count, o = bitpack.Unpack(bs, o, 11)
			for len(p.Sub) < int(count) {
				sub, off, err := parse(bs, o)
				o = off
				if err != nil && err != io.EOF {
					return nil, o, err
				}
				p.Sub = append(p.Sub, *sub)
			}
		}
	}
	return &p, o, io.EOF
}
