package main

import "fmt"

type book struct {
	id int
	name string
}

func main() {
	var book1 book
	book1.id = 100
	book1.name = "道德经"
	fmt.Println(book1)

	// 指针
	var book2 *book
	book2 = &book1
	fmt.Println(book2)
	fmt.Println(book2.name)

}