package main

import (
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

var rulemap = make(map[int][]int)

type rule struct {
	before, after int
}

func main() {
	b, err := os.ReadFile("input1.txt")
	if err != nil {
		log.Fatalf("couldn't read file, %v", err)
	}
	parts := strings.Split(string(b), "\n\n")
	rulesLines := strings.Split(parts[0], "\n")
	updateLines := strings.Split(parts[1], "\n")

	rules := make([]rule, len(rulesLines))
	for i, v := range rulesLines {
		bs, as, _ := strings.Cut(v, "|")
		b, _ := strconv.Atoi(bs)
		a, _ := strconv.Atoi(as)
		rules[i] = rule{before: b, after: a}
	}

	updates := make([][]int, len(updateLines))
	for i, v := range updateLines {
		u := strings.Split(v, ",")
		update := make([]int, len(u))
		for j, page := range u {
			num, _ := strconv.Atoi(page)
			update[j] = num
		}
		updates[i] = update
	}

	for _, v := range rules {
		rulemap[v.before] = append(rulemap[v.before], v.after)
	}

	partA(updates)
	partB(updates)

}

func cmp(a, b int) int {
	if slices.Contains(rulemap[a], b) {
		return -1
	}
	return 0
}

func partA(updates [][]int) {
	sum := 0
	for _, update := range updates {
		cp := slices.Clone(update)
		slices.SortFunc(cp, cmp)
		if slices.Equal(update, cp) {
			midIdx := (len(update) / 2)
			sum += update[midIdx]
		}
	}
	println("partA:", sum)
}

func partB(updates [][]int) {
	sum := 0
	for _, update := range updates {
		cp := slices.Clone(update)
		slices.SortFunc(cp, cmp)
		if slices.Equal(update, cp) {
			continue
		}
		midIdx := (len(cp) / 2)
		sum += cp[midIdx]
	}
	println("partB:", sum)
}
