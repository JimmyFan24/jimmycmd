package main

import "fmt"

type M2 struct {
	name string
}

func M2Test() (i int) {
	a := 1
	i = i + a
	fmt.Println(i)
	return i
}
