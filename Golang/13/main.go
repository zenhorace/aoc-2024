package main

import (
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const deltaP2 = float64(10000000000000)

type equation struct {
	a, b, p float64
}

func main() {
	b, err := os.ReadFile("input1.txt")
	if err != nil {
		log.Fatalf("couldn't read file, %v", err)
	}
	batches := strings.Split(string(b), "\n\n")

	sumP1, sumP2 := 0, 0
	for _, batch := range batches {
		lines := strings.Split(batch, "\n")
		if len(lines) != 3 {
			log.Fatal("unexpected number of lines for batch", len(lines))
		}
		eq1, eq2 := getEqs(lines)
		if a, b, v := valid(solve(eq1, eq2)); v {
			sumP1 += (a*3 + b)
		}
		eq1.p += deltaP2
		eq2.p += deltaP2
		if a, b, v := validP2(solve(eq1, eq2)); v {
			sumP2 += (a*3 + b)
		}
	}
	println("part1:", sumP1, "part2:", sumP2)
}

func getEqs(lines []string) (eq1, eq2 equation) {
	ba := strings.TrimPrefix(lines[0], "Button A: X+")
	ms, ns, _ := strings.Cut(ba, ", Y+")
	m, _ := strconv.ParseFloat(ms, 64)
	n, _ := strconv.ParseFloat(ns, 64)
	eq1.a, eq2.a = m, n

	bb := strings.TrimPrefix(lines[1], "Button B: X+")
	ms, ns, _ = strings.Cut(bb, ", Y+")
	m, _ = strconv.ParseFloat(ms, 64)
	n, _ = strconv.ParseFloat(ns, 64)
	eq1.b, eq2.b = m, n

	p := strings.TrimPrefix(lines[2], "Prize: X=")
	ms, ns, _ = strings.Cut(p, ", Y=")
	m, _ = strconv.ParseFloat(ms, 64)
	n, _ = strconv.ParseFloat(ns, 64)
	eq1.p, eq2.p = m, n
	return eq1, eq2
}

// Intuition: We have two equations of the form mA + nB = P,
// and two unknowns so we can solve by substitution.
func solve(eq1, eq2 equation) (a, b float64) {
	b = math.Abs((eq2.p - (eq2.a * eq1.p / eq1.a)) / (eq2.b - (eq1.b * eq2.a / eq1.a)))
	a = (eq1.p - b*eq1.b) / eq1.a
	return a, b
}

func valid(a, b float64) (ai, bi int, v bool) {
	if a < 0 || a > 100 || b > 100 {
		return ai, bi, false
	}
	ar, br := math.Round(a), math.Round(b)
	if math.Abs(a-ar) > 0.05 || math.Abs(b-br) > 0.05 {
		return ai, bi, false
	}
	ai, bi = int(ar), int(br)
	return ai, bi, true
}

func validP2(a, b float64) (ai, bi int, v bool) {
	if a < 0 {
		return ai, bi, false
	}
	ar, br := math.Round(a), math.Round(b)
	if math.Abs(a-ar) > 0.0001 || math.Abs(b-br) > 0.0001 {
		return ai, bi, false
	}
	ai, bi = int(ar), int(br)
	return ai, bi, true
}
