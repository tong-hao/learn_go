package main

import "fmt"

func max (i1 int, i2 int) int {
	if i1 > i2 {
		return i1
	} else {
		return i2
	}
}

func main() {
	m := max(1,50)
	fmt.Println(m)

}
