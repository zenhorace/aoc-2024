package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	roomWidth  = 101
	roomHeight = 103
)

var (
	bots []bot
)

type pos struct {
	x, y int
}

type bot struct {
	start  pos
	dX, dY int
}

func main() {
	d, err := os.ReadFile("input1.txt")
	if err != nil {
		log.Fatalf("couldn't read file, %v", err)
	}
	lines := strings.Split(string(d), "\n")
	bots = make([]bot, len(lines))
	for i, line := range lines {
		ss, vs, _ := strings.Cut(line, " ")
		b := bot{}
		sx, sy, _ := strings.Cut(strings.TrimPrefix(ss, "p="), ",")
		vx, vy, _ := strings.Cut(strings.TrimPrefix(vs, "v="), ",")
		b.start.x, _ = strconv.Atoi(sx)
		b.start.y, _ = strconv.Atoi(sy)
		b.dX, _ = strconv.Atoi(vx)
		b.dY, _ = strconv.Atoi(vy)
		bots[i] = b
	}
	part2()
}

// 1078, 1179, 1280
// 101
func part2() {
	for i := 1280; i < 10000; i += 101 {
		println(i)
		var grid [][]bool
		for i := 0; i < roomHeight; i++ {
			grid = append(grid, make([]bool, roomWidth))
		}
		for _, b := range bots {
			p := calculatePos(b, i)
			grid[p.y][p.x] = true
		}
		printGrid(grid)
	}
}

func part1() {
	tiles := make(map[pos]int)
	for _, b := range bots {
		tiles[calculatePos(b, 100)] += 1
	}
	midX := roomWidth / 2
	midY := roomHeight / 2

	var q1, q2, q3, q4 int
	for loc, count := range tiles {
		if loc.x == midX || loc.y == midY {
			continue
		}
		switch {
		case loc.x < roomWidth/2 && loc.y < roomHeight/2:
			q1 += count
		case loc.x < roomWidth/2:
			q2 += count
		case loc.y < roomHeight/2:
			q3 += count
		default:
			q4 += count
		}
	}
	println(q1 * q2 * q3 * q4)
}

func calculatePos(b bot, iters int) pos {
	x := (b.start.x + iters*b.dX) % roomWidth
	if x < 0 {
		x = roomWidth + x
	}
	y := (b.start.y + iters*b.dY) % roomHeight
	if y < 0 {
		y = roomHeight + y
	}
	return pos{x, y}
}

func printGrid(grid [][]bool) {
	out := strings.Builder{}
	for _, row := range grid {
		for _, b := range row {
			if b {
				out.WriteRune('@')
			} else {
				out.WriteRune('.')
			}
		}
		out.WriteRune('\n')
	}
	println(out.String())
}
