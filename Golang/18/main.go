package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type pos struct {
	x, y int
}

type dir pos

var (
	grid  [][]bool
	left  = dir{-1, 0}
	right = dir{1, 0}
	up    = dir{0, -1}
	down  = dir{0, 1}
	moves = []dir{left, right, up, down}
)

func main() {
	b, err := os.ReadFile("input1.txt")
	if err != nil {
		log.Fatalf("couldn't read file, %v", err)
	}

	for i := 0; i < 71; i++ {
		grid = append(grid, make([]bool, 71))
	}

	lines := strings.Split(string(b), "\n")
	for i := 0; i < 1024; i++ {
		xs, ys, _ := strings.Cut(lines[i], ",")
		grid[atoi(ys)][atoi(xs)] = true
	}
	println("part1:", bfs(pos{0, 0}, pos{70, 70}))

	for i := 1024; i < len(lines); i++ {
		xs, ys, _ := strings.Cut(lines[i], ",")
		grid[atoi(ys)][atoi(xs)] = true
		if bfs(pos{0, 0}, pos{70, 70}) < 0 {
			fmt.Printf("Part 2: %v,%v'n", xs, ys)
			break
		}
	}
}

func bfs(start, end pos) int {
	q := []pos{start}
	visited := make(map[pos]pos)
	visited[start] = pos{-1, -1}

	found := false
outer:
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		for _, m := range moves {
			p := pos{curr.x + m.x, curr.y + m.y}
			if _, ok := visited[p]; !ok && validPos(p) {
				q = append(q, p)
				visited[p] = curr
			}
			if p == end {
				found = true
				break outer
			}
		}
	}
	if !found {
		return -1
	}
	curr := end
	count := 0
	for curr != start {
		curr = visited[curr]
		count++
	}
	return count
}

func validPos(p pos) bool {
	return p.x >= 0 && p.x < len(grid[0]) && p.y >= 0 && p.y < len(grid) && !grid[p.y][p.x]
}

func atoi(s string) int {
	r, _ := strconv.Atoi(s)
	return r
}

func printGrid() {
	for _, row := range grid {
		for _, c := range row {
			if c {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Printf("\n")
	}
}
