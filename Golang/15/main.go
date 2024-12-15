package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type pos struct {
	x, y int
}

type dir pos

type interval struct {
	start, stop int
}

var (
	left   = dir{-1, 0}
	right  = dir{1, 0}
	up     = dir{0, -1}
	down   = dir{0, 1}
	dirMap = map[rune]dir{
		'>': right,
		'<': left,
		'^': up,
		'v': down,
	}
	nGrid [][]rune
	wGrid [][]rune
)

func main() {
	b, err := os.ReadFile("sample.txt")
	if err != nil {
		log.Fatalf("couldn't read file, %v", err)
	}
	parts := strings.Split(string(b), "\n\n")
	var start pos
	for j, row := range strings.Split(parts[0], "\n") {
		rw := make([]rune, len(row))
		for i, r := range row {
			rw[i] = r
			if r == '@' {
				start = pos{i, j}
			}
		}
		nGrid = append(nGrid, rw)
	}
	fillWideGrid()
	wStart := pos{start.x * 2, start.y}

	moves := strings.ReplaceAll(parts[1], "\n", "")
	for _, m := range moves {
		start = next(m, start, nGrid)
		if m == '<' || m == '>' {
			wStart = horizNext(m, wStart, wGrid)
		} else {
			wStart = vertNext(m, wStart, wGrid)
		}
	}
	println("Part 1:", gpsSum('O', nGrid))
	println("Part 2:", gpsSum('[', wGrid))
}

func next(move rune, start pos, grid [][]rune) pos {
	d := dirMap[move]
	cx := d.x + start.x
	cy := d.y + start.y
	cr := grid[cy][cx]
	for cr == 'O' {
		cx = d.x + cx
		cy = d.y + cy
		cr = grid[cy][cx]
	}
	if cr == '#' {
		return start
	}
	// we move
	for cx != start.x || cy != start.y {
		// swap
		px := cx - d.x
		py := cy - d.y
		grid[cy][cx], grid[py][px] = grid[py][px], grid[cy][cx]
		cx, cy = px, py
	}
	return pos{start.x + d.x, start.y + d.y}
}

func horizNext(move rune, start pos, grid [][]rune) pos {
	d := dirMap[move]
	cx := d.x + start.x
	cy := d.y + start.y
	cr := grid[cy][cx]
	for cr == '[' || cr == ']' {
		cx = d.x + cx
		cy = d.y + cy
		cr = grid[cy][cx]
	}
	if cr == '#' {
		return start
	}
	// we move
	for cx != start.x || cy != start.y {
		// swap
		px := cx - d.x
		py := cy - d.y
		grid[cy][cx], grid[py][px] = grid[py][px], grid[cy][cx]
		cx, cy = px, py
	}
	return pos{start.x + d.x, start.y + d.y}
}

func vertNext(move rune, start pos, grid [][]rune) pos {
	d := dirMap[move]
	c := pos{d.x + start.x, d.y + start.y}
	cr := grid[c.y][c.x]
	if cr == '#' {
		return start
	}
	if cr == '.' {
		grid[c.y][c.x], grid[start.y][start.x] = grid[start.y][start.x], grid[c.y][c.x]
		return c
	}

	// hard work needed
	ranges := map[int][]interval{c.x: {{start.y, c.y}}}

	clear := false
	for !clear {
		clear = true
		// evaluate
		for px, list := range ranges {
			py := list[len(list)-1].stop
			if grid[py][px] == '.' {
				continue
			}
			clear = false
			if grid[py][px] == '[' {
				if _, ok := ranges[px+1]; ok && grid[ranges[px+1][len(ranges[px+1])-1].stop][px+1] != '.' {
					continue
				}
				ranges[px+1] = append(ranges[px+1], interval{py, py})
			} else if grid[py][px] == ']' {
				if _, ok := ranges[px-1]; ok && grid[ranges[px-1][len(ranges[px-1])-1].stop][px-1] != '.' {
					continue
				}
				ranges[px-1] = append(ranges[px-1], interval{py, py})
			} else { // #
				return start
			}
		}
		// advance
		for x, list := range ranges {
			vIdx := len(list) - 1
			if grid[list[vIdx].stop][x] == '.' {
				continue
			}
			list[vIdx].stop += d.y
			ranges[x] = list
		}
	}
	// we move
	for r, list := range ranges {
		for _, v := range list {
			cx, cy := r, v.stop
			for v.start != cy {
				px, py := cx-d.x, cy-d.y
				grid[cy][cx], grid[py][px] = grid[py][px], grid[cy][cx]
				cx, cy = px, py
			}
		}
	}
	return c
}

func gpsSum(obs rune, g [][]rune) (sum int) {
	for j, r := range g {
		for i, l := range r {
			if l == obs {
				sum += (100*j + i)
			}
		}
	}
	return sum
}

func fillWideGrid() {
	for _, row := range nGrid {
		nr := make([]rune, len(row)*2)
		for j, r := range row {
			switch r {
			case '@':
				nr[j*2] = '@'
				nr[j*2+1] = '.'
			case '#':
				nr[j*2] = '#'
				nr[j*2+1] = '#'
			case 'O':
				nr[j*2] = '['
				nr[j*2+1] = ']'
			case '.':
				nr[j*2] = '.'
				nr[j*2+1] = '.'
			}
		}
		wGrid = append(wGrid, nr)
	}
}

// These functions were only for debgging/validating setps.

func printGrid(g [][]rune) {
	for _, r := range g {
		for _, i := range r {
			fmt.Printf("%v", string(i))
		}
		fmt.Printf("\n")
	}
}

func clone(g [][]rune) (c [][]rune) {
	for _, v := range g {
		s := slices.Clone(v)
		c = append(c, s)
	}
	return c
}

func validate() bool {
	for y, r := range wGrid {
		for x, obs := range r {
			if obs == '[' && wGrid[y][x+1] != ']' {
				return false
			}
			if obs == ']' && wGrid[y][x-1] != '[' {
				return false
			}
		}
	}
	return true
}
