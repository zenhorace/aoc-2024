package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type pos struct {
	x, y int
}

type dir pos

var (
	grid  [][]string
	ends  = make(map[pos]bool)
	left  = dir{-1, 0}
	right = dir{1, 0}
	up    = dir{0, -1}
	down  = dir{0, 1}
	moves = map[dir][]dir{
		left:  {left, down, up},
		right: {right, down, up},
		up:    {right, left, up},
		down:  {right, down, left},
	}
	neighbors = []dir{left, right, up, down}
)

func main() {
	b, err := os.ReadFile("sample.txt")
	if err != nil {
		log.Fatalf("couldn't read file, %v", err)
	}
	for _, row := range strings.Split(string(b), "\n") {
		grid = append(grid, strings.Split(row, ""))
	}
	start := pos{1, 3}
	end := pos{5, 7}
	// start := pos{81, 51}
	// end := pos{73, 31}
	grid[end.y][end.x] = "." // simplifies
	grid[start.y][start.x] = "."

	part2(start, end)
}

func part2(start, end pos) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == "." {
				continue
			}
			if canStartOrEndPath(x, y) {
				ends[pos{x, y}] = true
			}
		}
	}

	// fmt.Printf("%+v\n", ends)

	baseline := bfs(start, end)
	println(baseline)
	cheats := make(map[int]int)
	for st := range ends {
		opts := endOpts(st)
		for _, e := range opts {
			// fmt.Printf("st: %+v, e: %+v\n", st, e)
			res := bfsWall(st, e)
			if len(res) < 1 || len(res) > 19 {
				continue
			}
			delta := baseline - eval(start, end, res)
			if delta >= 50 {
				cheats[delta]++
			}
			if delta == 62 {
				fmt.Printf("s: %+v, e: %+v\n", st, e)
			}
		}
	}
	fmt.Printf("%+v\n", cheats)
}

func eval(start, end pos, replace []pos) int {
	for _, p := range replace {
		grid[p.y][p.x] = "."
	}
	defer func() {
		for _, p := range replace {
			grid[p.y][p.x] = "#"
		}
	}()
	return bfs(start, end)
}

func part1(start, end pos) {
	st := time.Now()
	baseline := bfs(start, end)
	println(baseline)
	ch := 0
	for y := 1; y < len(grid)-1; y++ {
		for x := 1; x < len(grid[0])-1; x++ {
			if grid[y][x] != "#" {
				continue
			}
			if canRemove(x, y) {
				grid[y][x] = "."
			}
			if baseline-bfs(start, end) >= 100 {
				ch++
			}
			grid[y][x] = "#"
		}
	}
	fmt.Printf("Took: %v\n", time.Now().Sub(st))
	fmt.Printf("%+v\n", ch)
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

		for _, m := range neighbors {
			p := pos{curr.x + m.x, curr.y + m.y}
			if _, ok := visited[p]; !ok && validPos(p) && grid[p.y][p.x] == "." {
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

func canRemove(x, y int) bool {
	count := 0
	for _, n := range neighbors {
		if grid[n.y+y][n.x+x] == "." {
			count++
		}
	}
	return count >= 2
}

func canStartOrEndPath(x, y int) bool {
	count := 0
	for _, n := range neighbors {
		p := pos{n.x + x, n.y + y}
		if validPos(p) && grid[p.y][p.x] == "." {
			count++
		}
	}
	return count >= 1
}

func bfsWall(start, end pos) []pos {
	q := []pos{start}
	visited := make(map[pos]pos)
	visited[start] = pos{-1, -1}

	found := false
outer:
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		for _, m := range neighbors {
			p := pos{curr.x + m.x, curr.y + m.y}
			if _, ok := visited[p]; !ok && validPos(p) && grid[p.y][p.x] == "#" {
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
		return nil
	}
	curr := end
	res := []pos{curr}
	for curr != start {
		curr = visited[curr]
		res = append(res, curr)
	}
	return res
}

func endOpts(start pos) (res []pos) {
	for y := max(start.y-19, 0); y <= min(start.y+19, len(grid)-1); y++ {
		delta := abs(start.y - y)
		for x := max(start.x-delta, 0); x <= min(start.x+delta, len(grid[0])-1); x++ {
			if (x == start.x && y == start.y) || grid[y][x] != "#" {
				continue
			}
			if ends[pos{x, y}] {
				res = append(res, pos{x, y})
			}
		}
	}
	return res
}

func validPos(p pos) bool {
	return p.x >= 0 && p.x < len(grid[0]) && p.y >= 0 && p.y < len(grid)
}

func abs(v int) int {
	if v < 0 {
		return v * -1
	}
	return v
}

func printGrid() {
	for _, r := range grid {
		println(strings.Join(r, ""))
	}
}
