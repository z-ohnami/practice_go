package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])


	m = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967, },
		"Google": {37.42202, -122.08408 },
	}

	fmt.Println(m)

	q := make(map[string]int)

	q["Answer"] = 42
	fmt.Println("The value:", q["Answer"])

	q["Answer"] = 48
	fmt.Println("The value:", q["Answer"])

	delete(q, "Answer")
	fmt.Println("The value:", q["Answer"])

	v, ok := q["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
