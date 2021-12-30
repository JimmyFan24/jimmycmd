package main

import "fmt"

func main() {
	b := M2Test()
	fmt.Printf("%d", b)
}

type M1 interface {
	M1Eat() (i int)
}
