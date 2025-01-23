package main

import (
	"strconv"
	"strings"
)

var (

	// +---+---+---+
	// | 7 | 8 | 9 |
	// +---+---+---+
	// | 4 | 5 | 6 |
	// +---+---+---+
	// | 1 | 2 | 3 |
	// +---+---+---+
	//     | 0 | A |
	//     +---+---+

	kpadPaths = map[string][]string{
		"A": {"<", "^<<", "<^", "^", "^^<<", "<^^", "^^", "^^^<<", "^^^<", "^^^", ""},
		"0": {"", "^<", "^", "^>", "^^<", "^^", "^^>", "^^^<", "^^^", "^^^>", ">"},
		"1": {">v", "", ">", ">>", "^", "^>", "^>>", "^^", "^^>", "^^>>", ">>v"},
		"2": {"v", "<", "", ">", "<^", "^", "^>", "^^<", "^^", "^^>", ">v"},
		"3": {"v<", "<<", "<", "", "<<^", "<^", "^", "<<^^", "<^^", "^^", "v"},
		"4": {">vv", "v", "v>", "v>>", "", ">", ">>", "^", "^>", "^>>", ">>vv"},
		"5": {"vv", "v<", "v", "v>", "<", "", ">", "^<", "^", "^>", "vv>"},
		"6": {"<vv", "<<v", "<v", "v", "<<", "<", "", "^<<", "^<", "^", "vv"},
		"7": {">vvv", "vv", ">vv", ">>vv", "v", ">v", ">>v", "", ">", ">>", ">>vvv"},
		"8": {"vvv", "vv<", "vv", "vv>", "v<", "v", "v>", "<", "", ">", "vvv>"},
		"9": {"vvv<", "vv<<", "vv<", "vv", "v<<", "v<", "v", "<<", "<", "", "vvv"},
	}
	kpadIdx = map[string]int{"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "A": 10}

	//		 +---+---+
	// 	   | ^ | A |
	// +---+---+---+
	// | < | v | > |
	// +---+---+---+

	dpadPaths = map[string][]string{
		"A": {"", "<", "v<<", "v<", "v"},
		"^": {">", "", "v<", "v", "v>"},
		"<": {">>^", ">^", "", ">", ">>"},
		"v": {"^>", "^", "<", "", ">"},
		">": {"^", "^<", "<<", "<", ""},
	}
	dpadIdx = map[string]int{"A": 0, "^": 1, "<": 2, "v": 3, ">": 4}
)

func expandKpad(orig string) string {
	start := "A"
	var res []string
	for _, r := range orig {
		k := string(r)
		// fmt.Printf("%+v\n", kpadIdx[k])
		res = append(res, kpadPaths[start][kpadIdx[k]])
		start = k
	}
	return strings.Join(res, "A") + "A"
}

func expandDpad(orig string) string {
	start := "A"
	var res []string
	for _, r := range orig {
		k := string(r)
		res = append(res, dpadPaths[start][dpadIdx[k]])
		start = k
	}
	return strings.Join(res, "A") + "A"
}

func main() {
	input := []string{"593A", "508A", "386A", "459A", "246A"}
	// input := []string{"029A", "980A", "179A", "456A", "379A"}
	// input := []string{"179A"}
	sum := 0
	for _, code := range input {
		res := expandDpad(expandDpad(expandKpad(code)))
		v, _ := strconv.Atoi(strings.TrimSuffix(code, "A"))
		sum += (v * len(res))
		println(v, len(res))
	}
	println("part 1:", sum)
}
