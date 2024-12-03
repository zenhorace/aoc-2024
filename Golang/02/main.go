package main

import (
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

// Part 2 is a DP problem (coin change) but n is small enough that brute force works.

func main() {
	b, err := os.ReadFile("input1.txt")
	if err != nil {
		log.Fatalf("couldn't read file, %v", err)
	}
	data := string(b)
	reports := strings.Split(data, "\n")
	countP1, countP2 := 0, 0
	for _, r := range reports {
		l := strings.Split(r, " ")
		levels := make([]int, len(l))
		for i, v := range l {
			levels[i], _ = strconv.Atoi(v)
		}
		if validDecrease(levels) || validIncrease(levels) {
			countP1++
			continue
		}

		if tolerant(levels) {
			countP2++
		}

	}
	println("Part1", countP1)
	println("Part2", countP1+countP2)
}

func tolerant(l []int) bool {
	for i := 0; i < len(l); i++ {

		// weird things happen in Go when you try to repeatedly slice :/
		n := slices.Clone(l)
		m := slices.Clone(l)
		a := append(n[:i], m[i+1:]...)
		if validIncrease(a) || validDecrease(a) {
			return true
		}
	}
	return false
}

func validIncrease(l []int) bool {
	// assume no single level reports
	for i := 1; i < len(l); i++ {
		diff := l[i] - l[i-1]
		if diff > 3 || diff < 1 {
			return false
		}
	}
	return true
}

func validDecrease(l []int) bool {
	// assume no single level reports
	for i := 1; i < len(l); i++ {
		diff := l[i-1] - l[i]
		if diff > 3 || diff < 1 {
			return false
		}
	}
	return true
}
