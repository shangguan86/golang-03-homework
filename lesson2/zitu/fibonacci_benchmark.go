package main

import (
	"flag"
	"log"
	"math"
	"time"
)

type funcDef func(int) int

type benchResult struct {
	result int
	cost   int64
	desc   string
}

func fibRecu(n int) int {
	// 0, 1, 1, 2, 3, 5...
	if n <= 1 {
		return n
	}
	return fibRecu(n-1) + fibRecu(n-2)
}

func fibIter(n int) int {
	// best
	// 0, 1, 1, 2, 3, 5...
	// a b n
	// b a+b n-1
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return a
}

func fibCalc(n int) int {
	// precision and overflow could cause wrong result
	return int(math.Pow(math.Phi, float64(n))/math.Sqrt(5) + 0.5)
}

func bench(f funcDef, n int, ch chan benchResult, desc string) {
	s := time.Now().UnixNano()
	r := f(n)
	delta := time.Now().UnixNano() - s
	ch <- benchResult{r, delta, desc}
}

func main() {
	// use current cpu count as max proc
	// runtime.GOMAXPROCS(runtime.NumCPU())

	// unbuffered channel
	// both sides should ready
	ch := make(chan benchResult)
	pn := flag.Int("n", 0, "n count")
	flag.Parse()
	*pn = *pn + 1
	// use a timer to prevent long time waiting
	go bench(fibRecu, *pn, ch, "Recursion")
	go bench(fibIter, *pn, ch, "Iteration")
	go bench(fibCalc, *pn, ch, "Calculation")

	for i := 0; i < 3; i++ {
		br, _ := <-ch
		log.Printf("result is: %v, use %v approach, cost %vns\n", br.result, br.desc, br.cost)
	}
}
