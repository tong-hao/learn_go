package main

import (
	"errors"
	"fmt"
)

func MoreThan1(f float64) (float64, error) {
	if f < 1 {
		return 0, errors.New("this is a error.")
	}

	return 1, nil
}

func main() {

	 r, err := MoreThan1(-1.0)
	 if err != nil {
		 fmt.Println(err)
	 } else {
		 fmt.Println(r)
	 }
}