package main

import (
	"log"
	"os"
	"strings"
)

type node struct {
	p pos
	r rune
	n []pos
}

type pos struct {
	x, y int
}

// panics if pos isn't in grid
func (p pos) node() node {
	return nodeList[rowLen*p.y+p.x]
}

type dir pos

var (
	nodeList []node
	regions  = make(map[rune][]map[pos]bool)
	grid     = make(map[pos]rune)
	left     = dir{-1, 0}
	right    = dir{1, 0}
	up       = dir{0, -1}
	down     = dir{0, 1}
	dirs     = []dir{left, right, up, down}

	rowLen int
)

func main() {
	b, err := os.ReadFile("input1.txt")
	if err != nil {
		log.Fatalf("couldn't read file, %v", err)
	}

	rows := strings.Split(string(b), "\n")
	rowLen = len(rows[0])
	for i, row := range rows {
		for j, plant := range row {
			p := pos{x: j, y: i}
			nodeList = append(nodeList, node{p: p, r: plant})
			grid[p] = plant
		}
	}

	for i := 0; i < len(nodeList); i++ {
		nodeList[i].n = regionalNeighbours(nodeList[i].r, nodeList[i].p)
	}

	for _, node := range nodeList {
		if inRegions(node.p, node.r) {
			continue
		}
		m := make(map[pos]bool)
		dfs(node.p, m)
		regions[node.r] = append(regions[node.r], m)
	}

	sum := 0
	sump2 := 0
	for _, sub := range regions {
		sigh := 0
		umm := 0
		for _, s := range sub {
			area := len(s)
			side := sides(s)
			pm := 0
			for l := range s {
				pm += (4 - len(l.node().n))
			}
			sigh += (area * pm)
			umm += (area * side)
		}
		// fmt.Printf("%v, %v, %+v\n", string(k), sigh, sub)
		sum += sigh
		sump2 += umm
	}

	println(sum, sump2)

}

// Count the number of verticies (since that must equal the number of sides).
func sides(region map[pos]bool) int {
	count := 0
	for loc := range region {
		if len(loc.node().n) >= 4 {
			continue // inside piece. skip
		}
		for _, d := range dirs {
			next := pos{loc.x + d.x, loc.y + d.y}
			if grid[next] == loc.node().r {
				continue
			}
			if d.x == 0 && region[pos{loc.x - d.y, loc.y}] && !region[pos{loc.x - d.y, loc.y + d.y}] {
				continue
			}
			if d.y == 0 && region[pos{loc.x, loc.y - d.x}] && !region[pos{loc.x + d.x, loc.y - d.x}] {
				continue
			}
			count++
		}
	}
	return count
}

func inRegions(p pos, r rune) bool {
	for _, m := range regions[r] {
		if m[p] {
			return true
		}
	}
	return false
}

func dfs(p pos, seen map[pos]bool) {
	if seen[p] { // deleteable
		return
	}
	seen[p] = true
	for _, l := range p.node().n {
		if !seen[l] {
			dfs(l, seen)
		}
	}
}

func regionalNeighbours(r rune, p pos) (res []pos) {
	for _, d := range dirs {
		x := p.x + d.x
		y := p.y + d.y
		p := pos{x, y}
		if grid[p] == r {
			res = append(res, p)
		}
	}
	return res
}
