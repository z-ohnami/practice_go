package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	// add
	fmt.Println(add(42, 13))
	// swap
	fmt.Println(swap("orange", "apple"))
	// split
	fmt.Println(split(20))
}
