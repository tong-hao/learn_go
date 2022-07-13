package main

import "fmt"

func main() {

	// 1. 变量
	{
		// 声明
		var i1 int
		var i2 int = 2
		i1 = 1
		i3 := 3
		fmt.Println("i1=", i1, ", i2=", i2, i3)

		// 声明 多个变量
		var s1, s2, s3 string
		s1, s2, s3 = "a", "b", "c"
		fmt.Println(s1, s2, s3)

		// 声明
		var (
			d1 float64
			u1 uint16
		)
		fmt.Println(d1, u1)
	}

	// 2.常量
	{
		const c1 int = 10
		const (
			a = iota
			b
			c
		)
		fmt.Println(c1, a, b, c)
	}


	// 3.类型转换
	var i1 int = 1
	var f1 float32 = 2.1
	var f2 float32 = float32(i1)
	var i2 int = int(f1)
	fmt.Println(i1,i2,f1,f2)

}
