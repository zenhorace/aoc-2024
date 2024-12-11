package main

import (
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
	grid  [][]int
	left  = dir{-1, 0}
	right = dir{1, 0}
	up    = dir{0, -1}
	down  = dir{0, 1}
	dirs  = []dir{left, right, up, down}
)

func main() {
	b, err := os.ReadFile("input1.txt")
	if err != nil {
		log.Fatalf("couldn't read file, %v", err)
	}
	for _, rowStr := range strings.Split(string(b), "\n") {
		row := make([]int, len(rowStr))
		for i, v := range rowStr {
			row[i], _ = strconv.Atoi(string(v))
		}
		grid = append(grid, row)
	}

	sumP1 := 0
	sumP2 := 0
	for y, row := range grid {
		for x, v := range row {
			if v != 0 {
				continue
			}
			track := make(map[pos]bool)
			explore(pos{x: x, y: y}, 0, track)
			sumP1 += len(track)
			sumP2 += explorePt2(pos{x: x, y: y}, 0)
		}
	}
	println("Part1", sumP1)
	println("Part2:", sumP2)
}

func explorePt2(p pos, start int) int {
	if start == 9 {
		return 1
	}
	branches := nextOpts(p, start)
	if len(branches) == 0 {
		return 0
	}
	sum := 0
	for _, b := range branches {
		sum += explorePt2(b, start+1)
	}
	return sum
}

func explore(p pos, start int, ends map[pos]bool) {
	if start == 9 {
		ends[p] = true
		return
	}
	branches := nextOpts(p, start)
	if len(branches) == 0 {
		return
	}
	for _, b := range branches {
		explore(b, start+1, ends)
	}
}

func nextOpts(p pos, val int) (res []pos) {
	for _, d := range dirs {
		x := p.x + d.x
		y := p.y + d.y
		if x >= len(grid[0]) || x < 0 || y < 0 || y >= len(grid) {
			continue //out of bounds
		}
		if grid[y][x] == (val + 1) {
			res = append(res, pos{x: x, y: y})
		}
	}
	return res
}
