package main

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	input  = "814 1183689 0 1 766231 4091 93836 46"
	sample = "125 17"
)

func main() {
	initL := strings.Split(input, " ")
	curr := make(map[string]int)
	for _, v := range initL {
		curr[v]++
	}
	for i := 0; i < 75; i++ {
		mod := make(map[string]int)
		for stone, count := range curr {
			if stone == "0" {
				mod["1"] += count
			} else if len(stone)%2 == 0 {
				f := stone[:len(stone)/2]
				s := stone[len(stone)/2:]
				mod[f] += count
				mod[removeLeadingZeroes(s)] += count
			} else {
				val, _ := strconv.Atoi(stone)
				val *= 2024
				mod[fmt.Sprintf("%d", val)] += count
			}
		}
		curr = mod
	}
	sum := 0
	for _, v := range curr {
		sum += v
	}
	println(sum)
}

func removeLeadingZeroes(s string) string {
	res, _ := strconv.Atoi(s)
	return fmt.Sprintf("%d", res)
}
