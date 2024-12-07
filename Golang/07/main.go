package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	b, err := os.ReadFile("input1.txt")
	if err != nil {
		log.Fatalf("couldn't read file, %v", err)
	}
	lines := strings.Split(string(b), "\n")
	var tests []int
	var operands [][]int
	for _, c := range lines {
		t, o, _ := strings.Cut(c, ": ")
		test, _ := strconv.Atoi(t)
		var l []int
		for _, v := range strings.Split(o, " ") {
			n, _ := strconv.Atoi(v)
			l = append(l, n)
		}
		tests = append(tests, test)
		operands = append(operands, l)
	}

	part1(tests, operands)
	part2(tests, operands)

}

func part1(tests []int, operands [][]int) {
	sum := 0
	for i, t := range tests {
		if helper1(t, operands[i]) {
			sum += t
		}
	}
	println(sum)
}

func helper1(test int, operands []int) bool {
	if len(operands) == 1 {
		return test-operands[0] == 0
	}
	pc := []int{operands[0] + operands[1]}
	mc := []int{operands[0] * operands[1]}
	if len(operands) > 2 {
		pc = append(pc, operands[2:]...)
		mc = append(mc, operands[2:]...)
	}
	return helper1(test, pc) || helper1(test, mc)
}

func part2(tests []int, operands [][]int) {
	sum := 0
	for i, t := range tests {
		if helper2(t, operands[i]) {
			sum += t
		}
	}
	println(sum)
}

func helper2(test int, operands []int) bool {
	if len(operands) == 1 {
		return test-operands[0] == 0
	}
	pc := []int{operands[0] + operands[1]}
	mc := []int{operands[0] * operands[1]}
	cv, _ := strconv.Atoi(fmt.Sprintf("%d%d", operands[0], operands[1]))
	cc := []int{cv}
	if len(operands) > 2 {
		pc = append(pc, operands[2:]...)
		mc = append(mc, operands[2:]...)
		cc = append(cc, operands[2:]...)
	}
	return helper2(test, pc) || helper2(test, mc) || helper2(test, cc)
}
