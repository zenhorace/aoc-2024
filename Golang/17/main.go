package main

import (
	"fmt"
	"log"
	"math"
	"strings"
)

type computer struct {
	regA, regB, regC, pgmPtr int

	output []string
}

func (c *computer) combo(in int) int {
	switch in {
	case 4:
		return c.regA
	case 5:
		return c.regB
	case 6:
		return c.regC
	default:
		return in
	}
}

func (c *computer) adv(op int) { c.regA = c.regA / exp(c.combo(op)) }

func exp(p int) int { return int(math.Pow(2, float64(p))) }

func (c *computer) bxl(op int) { c.regB = c.regB ^ op }

func (c *computer) bst(op int) { c.regB = c.combo(op) % 8 }

func (c *computer) bxc(_ int) { c.regB = c.regB ^ c.regC }

func (c *computer) bdv(op int) { c.regB = c.regA / exp(c.combo(op)) }

func (c *computer) cdv(op int) { c.regC = c.regA / exp(c.combo(op)) }

func (c *computer) jnz() bool { return c.regA != 0 }

func (c *computer) out(op int) { c.output = append(c.output, fmt.Sprintf("%d", c.combo(op)%8)) }

func (c *computer) printout() string { return strings.Join(c.output, ",") }

func (c *computer) run(pgm []int) {
	i := 0
	for i < len(pgm) {
		jump := false
		switch pgm[i] {
		case 0:
			c.adv(pgm[i+1])
		case 1:
			c.bxl(pgm[i+1])
		case 2:
			c.bst(pgm[i+1])
		case 3:
			jump = c.jnz()
		case 4:
			c.bxc(pgm[i+1])
		case 5:
			c.out(pgm[i+1])
		case 6:
			c.bdv(pgm[i+1])
		case 7:
			c.cdv(pgm[i+1])
		default:
			log.Fatalf("unknown opcode %v", pgm[i])
		}
		if jump {
			i = pgm[i+1]
			println("jump")
		} else {
			i += 2
		}
		fmt.Printf("a: %d, b: %d, c: %d\n", c.regA, c.regB, c.regC)
	}
	println(c.printout())
}

func main() {
	c := computer{regA: 30878003}
	pgm := []int{2, 4, 1, 2, 7, 5, 0, 3, 4, 7, 1, 7, 5, 5, 3, 0}
	c.run(pgm)

}
