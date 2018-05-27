package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

type Path []Point

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)

}

func pathDistance(p Path) {
	lenc := 0.0
	for i := 0; i < len(p)-1; i++ {
		lenc += p[i].Distance(p[i+1])
	}
	fmt.Println(lenc)
}

func main() {
	p := Path{{-1, 0}, {0, 0}, {3, 4}}
	pathDistance(p)
}
