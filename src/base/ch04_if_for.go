package main

import "fmt"

func main() {

	// if
	b := true
	if b {
		fmt.Println("1", b)
	} else if b {
		fmt.Println("2")
	}


	// switch
	{
		i :=1
		switch i {

		case 1:
			fmt.Println(1)
		case 2:
			fmt.Println(2)
		default:
			fmt.Println("other")
		}
	}



	// for
	{
		for i:=1; i<10;i++ {
			fmt.Print(i,",")
		}
		fmt.Println()

		i2:=5
		for i2<10{
			fmt.Print(i2,",")
			i2++
		}
		fmt.Println()

		i3:=6
		for{
			i3++
			fmt.Print(i3)
			if i3 > 10 {
				break
			}
		}
		fmt.Println()

	}

	// for range
	 arr1 := []int  {1,2,3}

	for i,v := range arr1 {
		fmt.Print("arr1[",i,"]=", v, ", ")
	}
	fmt.Println()


}