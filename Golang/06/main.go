package main

import (
	"log"
	"os"
	"strings"
)

type pos struct {
	x, y int
}

type dir pos

var (
	left  = dir{-1, 0}
	right = dir{1, 0}
	up    = dir{0, -1}
	down  = dir{0, 1}

	turns = make(map[dir]dir)
	grid  []string
)

func main() {
	b, err := os.ReadFile("input1.txt")
	if err != nil {
		log.Fatalf("couldn't read file, %v", err)
	}
	grid = strings.Split(string(b), "\n")

	turns[left] = up
	turns[up] = right
	turns[right] = down
	turns[down] = left

	var start pos
	for y, row := range grid {
		if x := strings.Index(row, "^"); x >= 0 {
			start = pos{x: x, y: y}
		}
	}
	seen := make(map[pos]bool)
	seen[start] = true
	part1(start, up, seen)
	println("part1:", len(seen))
	part2(start, up)
}

func part1(p pos, d dir, seen map[pos]bool) {
	if exit(p) {
		return
	}
	if grid[p.y][p.x] == '#' {
		nd := turns[d]
		np := pos{x: p.x - d.x + nd.x, y: p.y - d.y + nd.y}
		part1(np, nd, seen)
		return
	}
	seen[p] = true
	part1(pos{x: p.x + d.x, y: p.y + d.y}, d, seen)
}

func part2(start pos, d dir) {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] != '.' {
				continue
			}
			seen := make(map[pos]bool)
			seen[start] = true
			obs := pos{x: j, y: i}
			// fmt.Printf(" %+v ", obs)
			if loop(start, obs, d, seen) {
				count++
			}
		}
	}
	println("part2:", count)
}

func loop(p, obs pos, d dir, seen map[pos]bool) bool {
	// cycle detection heuristic: the grid is 130x130...
	//If we've traveled to 520 spots without seeing a new spot,
	// (the length of the perimeter) we're probably in a cycle.
	sinceLastAdd := 0
	for !exit(p) && sinceLastAdd <= 520 {
		if grid[p.y][p.x] == '#' || p == obs {
			nd := turns[d]
			p = pos{x: p.x - d.x + nd.x, y: p.y - d.y + nd.y}
			d = nd
			continue
		}
		if _, ok := seen[p]; ok {
			sinceLastAdd++
		} else {
			seen[p] = true
			sinceLastAdd = 0
		}
		p = pos{x: p.x + d.x, y: p.y + d.y}
	}
	return sinceLastAdd >= 520
}

func exit(p pos) bool {
	return p.x < 0 || p.x == len(grid[0]) || p.y < 0 || p.y == len(grid)
}
