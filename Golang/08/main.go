package main

import (
	"log"
	"os"
	"strings"
)

type pos struct {
	x, y int
}

var (
	grid []string
)

func (p pos) antinode(other pos) pos {
	dX := p.x - other.x
	dY := p.y - other.y
	return pos{x: p.x + dX, y: p.y + dY}
}

func (p pos) allAntinodes(other pos) (res []pos) {
	dX := p.x - other.x
	dY := p.y - other.y
	curr := pos{x: p.x + dX, y: p.y + dY}
	for inbounds(curr) {
		res = append(res, curr)
		curr = pos{x: curr.x + dX, y: curr.y + dY}
	}
	return res
}

func main() {
	b, err := os.ReadFile("input1.txt")
	if err != nil {
		log.Fatalf("couldn't read file, %v", err)
	}
	grid = strings.Split(string(b), "\n")
	group := make(map[rune][]pos)

	for y, row := range grid {
		for x, v := range row {
			if v == '.' {
				continue
			}
			group[v] = append(group[v], pos{x: x, y: y})
		}
	}
	part1(group)
	part2(group)
}

func part2(group map[rune][]pos) {
	antinodeSet := make(map[pos]bool)
	for _, g := range group {
		if len(g) == 1 {
			continue
		}
		antinodesForGroupPt2(g, antinodeSet)
	}
	println("part2:", len(antinodeSet))
}

func antinodesForGroupPt2(g []pos, set map[pos]bool) {
	for i, left := range g {
		set[left] = true
		for j := 0; j < len(g); j++ {
			if i == j {
				continue
			}
			for _, item := range left.allAntinodes(g[j]) {
				set[item] = true
			}
		}
	}
}

func part1(group map[rune][]pos) {
	antinodeSet := make(map[pos]bool)
	for _, g := range group {
		if len(g) == 1 {
			continue
		}
		antinodeForGroup(g, antinodeSet)
	}
	println("part1", len(antinodeSet))
}

func antinodeForGroup(g []pos, set map[pos]bool) {
	for i, left := range g {
		for j := 0; j < len(g); j++ {
			if i == j {
				continue
			}
			p := left.antinode(g[j])
			if inbounds(p) {
				set[p] = true
			}
		}
	}
}

func inbounds(p pos) bool {
	return p.x >= 0 && p.x < len(grid[0]) && p.y >= 0 && p.y < len(grid)
}
