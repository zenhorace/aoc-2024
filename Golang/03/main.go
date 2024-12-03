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
	OperRE   = regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)|don\'t\(\)|do\(\)`)
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
	found := OperRE.FindAll(b, -1)

	do := true
	sum := 0
	for _, v := range found {
		if "don't()" == string(v) {
			do = false
		} else if "do()" == string(v) {
			do = true
		} else if do {
			sum += mul(v)
		}
	}
	println("pt2:", sum)
}

func part1(b []byte) {
	found := validMul.FindAll(b, -1)
	sum := 0
	for _, v := range found {
		sum += mul(v)
	}

	println("pt1:", sum)
}

func mul(v []byte) int {
	inside := strings.Trim(string(v), "mul()")
	parts := strings.Split(inside, ",")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	return x * y
}
