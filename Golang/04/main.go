package main

import (
	"log"
	"os"
	"strings"
)

var countPt1 = 0

func main() {
	b, err := os.ReadFile("input1.txt")
	if err != nil {
		log.Fatalf("couldn't read file, %v", err)
	}
	grid := strings.Split(string(b), "\n")
	part2(grid)
}

func part2(grid []string) {
	count := 0
	for y, r := range grid {
		for x, l := range r {
			if l != 'A' {
				continue
			}
			if cross(x, y, grid) {
				count++
			}
		}
	}
	println(count)
}

func cross(x, y int, grid []string) bool {
	if x == 0 || y == 0 || x == (len(grid[y])-1) || y == (len(grid)-1) {
		return false
	}
	lb := grid[y-1][x-1]
	la := grid[y+1][x+1]
	rb := grid[y-1][x+1]
	ra := grid[y+1][x-1]

	return ((lb == 'M' && la == 'S') || (lb == 'S' && la == 'M')) &&
		((rb == 'M' && ra == 'S') || (rb == 'S' && ra == 'M'))
}

func part1(grid []string) {
	for y, r := range grid {
		for x, l := range r {
			if l != 'X' {
				continue
			}
			right(x, y, grid)
			left(x, y, grid)
			up(x, y, grid)
			down(x, y, grid)
			ul(x, y, grid)
			ur(x, y, grid)
			dl(x, y, grid)
			dr(x, y, grid)
		}
	}
	println(countPt1)
}

func right(x, y int, grid []string) {
	for i, v := range "XMAS" {
		if x+i >= len(grid[y]) {
			return
		}
		if rune(grid[y][x+i]) != v {
			return
		}
	}
	countPt1++
}

func left(x, y int, grid []string) {
	for i, v := range "XMAS" {
		if x-i < 0 {
			return
		}
		if rune(grid[y][x-i]) != v {
			return
		}
	}
	countPt1++
}

func down(x, y int, grid []string) {
	for i, v := range "XMAS" {
		if y+i >= len(grid) {
			return
		}
		if rune(grid[y+i][x]) != v {
			return
		}
	}
	countPt1++
}

func up(x, y int, grid []string) {
	for i, v := range "XMAS" {
		if y-i < 0 {
			return
		}
		if rune(grid[y-i][x]) != v {
			return
		}
	}
	countPt1++
}

func dr(x, y int, grid []string) {
	for i, v := range "XMAS" {
		if x+i >= len(grid[y]) || y+i >= len(grid) {
			return
		}
		if rune(grid[y+i][x+i]) != v {
			return
		}
	}
	countPt1++
}

func ur(x, y int, grid []string) {
	for i, v := range "XMAS" {
		if x+i >= len(grid[y]) || y-i < 0 {
			return
		}
		if rune(grid[y-i][x+i]) != v {
			return
		}
	}
	countPt1++
}

func dl(x, y int, grid []string) {
	for i, v := range "XMAS" {
		if x-i < 0 || y+i >= len(grid) {
			return
		}
		if rune(grid[y+i][x-i]) != v {
			return
		}
	}
	countPt1++
}

func ul(x, y int, grid []string) {
	for i, v := range "XMAS" {
		if x-i < 0 || y-i < 0 {
			return
		}
		if rune(grid[y-i][x-i]) != v {
			return
		}
	}
	countPt1++
}
