package main

import (
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	b, err := os.ReadFile("input1.txt")
	if err != nil {
		log.Fatalf("couldn't read file, %v", err)
	}
	data := string(b)
	lines := strings.Split(data, "\n")
	l1 := make([]int, len(lines))
	l2 := make([]int, len(lines))
	for i, v := range lines {
		parts := strings.Fields(v)
		l1[i], _ = strconv.Atoi(parts[0])
		l2[i], _ = strconv.Atoi(parts[1])
	}
	// part1(l1, l2)
	part2(l1, l2)
}

func part2(l1, l2 []int) {
	freq := make(map[int]int)
	for _, v := range l2 {
		freq[v] += 1
	}

	sum := 0
	for _, v := range l1 {
		sum += (freq[v] * v)
	}
	println(sum)
}

func part1(l1, l2 []int) {
	slices.Sort(l1)
	slices.Sort(l2)
	sum := 0.0

	for i := 0; i < len(l1); i++ {
		sum += math.Abs(float64(l1[i] - l2[i]))
	}
	println(int(sum))
}
