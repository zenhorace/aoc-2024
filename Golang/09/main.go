package main

import (
	"log"
	"os"
	"slices"
	"strconv"
)

func main() {
	b, err := os.ReadFile("input1.txt")
	if err != nil {
		log.Fatalf("couldn't read file, %v", err)
	}
	// convert to digits and dots (but -1 for dots)
	var layout []int
	for i, v := range b {
		digit, _ := strconv.Atoi(string(v))
		item := -1
		if i%2 == 0 {
			item = i / 2
		}
		for j := 0; j < digit; j++ {
			layout = append(layout, item)
		}
	}
	part1(slices.Clone(layout))
	part2(slices.Clone(layout))
}

func part2(layout []int) {
	i := len(layout)
	for i >= 0 {
		i = nextToMove(layout, i)
		if i < 0 {
			break
		}
		mvSize := spanBack(layout, i)
		to := spaceOfSize(layout, mvSize, i-mvSize)
		if to == -1 {
			i -= (mvSize - 1)
			continue
		}
		digit := layout[i]
		for j := 0; j < mvSize; j++ {
			layout[to+j] = digit
			layout[i-j] = -1
		}
		i -= (mvSize - 1)
	}

	sum := 0
	for i, v := range layout {
		if v == -1 {
			continue
		}
		sum += (i * v)
	}

	println("part2:", sum)
}

func spaceOfSize(layout []int, size, stop int) int {
	for i, v := range layout {
		if i >= stop {
			return -1
		}
		if v != -1 {
			continue
		}
		if spanFwd(layout, i) >= size {
			return i
		}
	}
	return -1
}

func spanFwd(layout []int, idx int) int {
	count := 0
	val := layout[idx]
	for i := idx; i < len(layout); i++ {
		if layout[i] != val {
			return count
		}
		count++
	}
	return count
}

func spanBack(layout []int, idx int) int {
	count := 0
	val := layout[idx]
	for i := idx; i >= 0; i-- {
		if layout[i] != val {
			return count
		}
		count++
	}
	return count
}

func part1(layout []int) {
	mvIdx := nextToMove(layout, len(layout))
	for i, v := range layout {
		if i >= mvIdx {
			break
		}
		if v != -1 {
			continue
		}
		layout[i], layout[mvIdx] = layout[mvIdx], layout[i]
		mvIdx = nextToMove(layout, mvIdx)
	}

	sum := 0
	for i, v := range layout {
		if v == -1 {
			break
		}
		sum += (i * v)
	}

	println("part1:", sum)
}

func nextToMove(layout []int, idx int) int {
	idx--
	for idx >= 0 {
		if layout[idx] != -1 {
			return idx
		}
		idx--
	}
	return idx
}
