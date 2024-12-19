package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("couldn't read file, %v", err)
	}
	parts := strings.Split(string(b), "\n\n")
	towels := strings.Split(parts[0], ", ")
	patterns := strings.Split(parts[1], "\n")

	println(len(towels))

	count := 0
	var sum int64
	memo := make(map[string]int64)
	for _, p := range patterns {
		if possible(p, towels) {
			count++
		}
		sum += allPossible(p, towels, memo)
	}
	println(count, sum)
}

func possible(pattern string, towels []string) bool {
	if len(pattern) == 0 {
		return true
	}
	for _, t := range towels {
		if strings.HasPrefix(pattern, t) {
			if possible(strings.TrimPrefix(pattern, t), towels) {
				return true
			}
		}
	}
	return false
}

func allPossible(pattern string, towels []string, memo map[string]int64) int64 {
	if ans, ok := memo[pattern]; ok {
		return ans
	}
	if len(pattern) == 0 {
		return 1
	}
	var count int64
	for _, t := range towels {
		if strings.HasPrefix(pattern, t) {
			count += allPossible(strings.TrimPrefix(pattern, t), towels, memo)
		}
	}
	memo[pattern] = count
	return count
}
