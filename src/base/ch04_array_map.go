package main

import "fmt"

func printArray(arr []int) {
	for e := range arr {
		fmt.Print(e, ",")
	}
	fmt.Println()
}

func main() {

	// 1.array
	{
		var arr1 [10]int
		for i := 0; i < 10; i++ {
			arr1[i] = i
		}

		for e := range arr1 {
			fmt.Print(e, ",")
		}
		fmt.Println()
		fmt.Println(arr1[0:3])
	}

	// 2.slice
	// todo:切片和数组的区别是什么？
	{
		var num = make([]int, 10)
		printArray(num)


		var num2 = make([]int, 3, 10)
		printArray(num2)
	}

	// 3.map
	var map1 map[string]string
	map1 = make(map[string]string)
	map1["a"] = "apple"
	map1["b"] = "basketball"

	for k := range map1 {
		fmt.Println(k, "=", map1[k])
	}

	v,ok := map1["c"]
	r := ok == true
	if r {
		fmt.Println(v)
	} else {
		fmt.Println("不存在c")
	}


}