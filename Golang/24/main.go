package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

var (
	wires = make(map[string]chan bool)
	wmu   sync.RWMutex

	zValues []bool
	zmap    = make(map[string]chan bool)
	// zmu     sync.RWMutex

	operMap = map[string]oper{"AND": and, "XOR": xor, "OR": or}
)

type oper func(bool, bool) bool

func xor(i1, i2 bool) bool { return i1 != i2 }
func and(i1, i2 bool) bool { return i1 && i2 }
func or(i1, i2 bool) bool  { return i1 || i2 }

type gate struct {
	in1, in2, out string
	f             oper
}

func wait(wg *sync.WaitGroup, c chan bool) {
	defer wg.Done()
	<-c
}

func (g gate) perform() {
	wmu.RLock()
	cin1 := wires[g.in1]
	cin2 := wires[g.in2]
	cout := wires[g.out]
	wmu.RUnlock()
	var wg sync.WaitGroup
	wg.Add(2)
	wait(&wg, cin1)
	wait(&wg, cin2)
	wg.Wait()
	i1 := <-cin1
	i2 := <-cin2
	r := g.f(i1, i2)
	go broadcast(cout, r)
}

func broadcast(c chan bool, v bool) {
	for {
		c <- v
	}
}

func zWrite(wg *sync.WaitGroup, z string, out chan bool) {
	defer wg.Done()
	r := <-out
	zValues[atoi(strings.TrimPrefix(z, "z"))] = r
}

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("couldn't read file, %v", err)
	}
	parts := strings.Split(string(b), "\n\n")

	for _, line := range strings.Split(parts[0], "\n") {
		p := strings.Split(line, ": ")
		c := make(chan bool, 1)
		wires[p[0]] = c
		go broadcast(c, p[1] == "1")
	}

	// first get all the gates and other wires
	gates := []gate{}
	for _, line := range strings.Split(parts[1], "\n") {
		left, res, _ := strings.Cut(line, " -> ")
		cout, ok := wires[res]
		if !ok {
			cout = make(chan bool, 1)
			if strings.HasPrefix(res, "z") {
				zmap[res] = cout
			}
			wires[res] = cout

		}

		p := strings.Fields(left)
		cin1, ok := wires[p[0]]
		if !ok {
			cin1 = make(chan bool, 1)
			wires[p[0]] = cin1
		}
		cin2, ok := wires[p[2]]
		if !ok {
			cin2 = make(chan bool, 1)
			wires[p[2]] = cin2
		}

		gates = append(gates, gate{in1: p[0], in2: p[2], out: res, f: operMap[p[1]]})
	}

	for _, g := range gates {
		go g.perform()
	}

	zValues = make([]bool, len(zmap))
	var wg sync.WaitGroup
	for k, v := range zmap {
		wg.Add(1)
		go zWrite(&wg, k, v)
	}
	wg.Wait() // wait for all Z's to be written

	res := ""
	for _, b := range zValues {
		if b {
			res = "1" + res
		} else {
			res = "0" + res
		}
	}
	val, _ := strconv.ParseInt(res, 2, 64)
	println(val)
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
