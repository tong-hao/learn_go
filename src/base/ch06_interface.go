package main

import "fmt"

type Phone interface {
	call() int
}

type NokiaPhone struct {
}

func (n NokiaPhone) call() int  {
	fmt.Println("this is a nokia phone.", n)
	return 1
}

type Iphone struct {
}

func (p Iphone) call()  int{
	fmt.Println("this is a iphone. ", p)
	return 2
}

func main() {
	var p Phone
	p = new (NokiaPhone)
	p.call()

	p = new (Iphone)
	p.call()

	fmt.Println(&p)
}