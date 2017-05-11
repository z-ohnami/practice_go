package main

import "fmt"

func main() {
	var greet [2]string
	greet[0] = "Hello"
	greet[1] = "World"
	fmt.Println(greet[0], greet[1])
	fmt.Println(greet)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	var s []int = primes[1:4]
	fmt.Println(s)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX" // !!
	fmt.Println(a, b)
	fmt.Println(names)
}
