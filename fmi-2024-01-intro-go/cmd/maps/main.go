package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Distance(o Vertex) float64 {
	return math.Sqrt((v.X-o.X)*(v.X-o.X) + (v.Y-o.Y)*(v.Y-o.Y))
}

func CalcualteDistance(places map[string]Vertex, from string, to string) (float64, error) {
	var (
		v1, v2 Vertex
		ok     bool
	)
	if v1, ok = places[from]; !ok {
		return 0, fmt.Errorf("Cannot find %s in places %v", from, places)
	}
	if v2, ok = places[to]; !ok {
		return 0, fmt.Errorf("Cannot find %s in places %v", to, places)
	}
	return v1.Distance(v2), nil
}

func main() {
	places := make(map[string]Vertex, 10)
	places["Bell Labs"] = Vertex{40.68433, -74.39967}
	places["Microsoft"] = Vertex{60.68433, -84.39967}
	places["Vitosha Soft"] = Vertex{Y: 23.32415, X: 42.69751}

}
