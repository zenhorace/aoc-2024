package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	validMul = regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	dontRE   = regexp.MustCompile(`don\'t\(\)`)
	doRE     = regexp.MustCompile(`do\(\)`)
)

func main() {
	b, err := os.ReadFile("input1.txt")
	if err != nil {
		log.Fatalf("couldn't read file, %v", err)
	}
	part1(b)
	part2(b)
}

func part2(b []byte) {
	mulIdx := validMul.FindAllIndex(b, -1)
	found := validMul.FindAll(b, -1)
	doIdx := doRE.FindAllIndex(b, -1)
	dontIdx := dontRE.FindAllIndex(b, -1)
	doPtr, dontPtr := 0, 0

	sum := 0
	for i, v := range mulIdx {
		doPtr = advance(doIdx, doPtr, v[0])
		dontPtr = advance(dontIdx, dontPtr, v[0])

		if v[0] < dontIdx[dontPtr][0] {
			sum += mul(found[i])
			continue
		}

		if dontIdx[dontPtr][0] < doIdx[doPtr][0] && doIdx[doPtr][0] < v[0] {
			sum += mul(found[i])
		}
	}
	println("pt2:", sum)
}

func advance(indexes [][]int, curr, opIdx int) int {
	for i := curr; i < len(indexes)-1; i++ {
		if indexes[i][0] > opIdx || indexes[i+1][0] > opIdx {
			return i
		}
	}
	return len(indexes) - 1
}

func part1(b []byte) {
	found := validMul.FindAll(b, -1)
	sum := 0
	for _, v := range found {
		sum += mul(v)
	}

	println(sum)
}

func mul(v []byte) int {
	inside := strings.Trim(string(v), "mul()")
	parts := strings.Split(inside, ",")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	return x * y
}
