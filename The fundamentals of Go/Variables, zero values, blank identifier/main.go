package main

import "fmt"

func main() {
	z, _ := 42, 43

	b, c, d, a, f := 0, 1, 2, 3, "happiness"
	fmt.Println(z, a, b, c, d, f)

	// this would not work
	/*
		b, c, d, e := 0, 1, 2, 3
		fmt.Println(b, c, d)
	*/

	// this does work
	/*
		var g int
		fmt.Println(g)
		g = 43
		fmt.Println(g)

		// declare a variable to hold a VALUE of a certain TYPE
		// initializing a variable
		var h int = 44
		fmt.Println(h)
	*/
}
