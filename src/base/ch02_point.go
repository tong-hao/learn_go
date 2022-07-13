package main

import "fmt"

func main() {

	// 1.
	var i1 int = 1
	var pi1 *int

	pi1 = &i1
	fmt.Println(i1, pi1, *pi1)

}
