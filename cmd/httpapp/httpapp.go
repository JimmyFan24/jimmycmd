package main

import (
	"jimmycmd/internal/httpapp"
)

func main() {

	httpapp.NewApp("httpapp").Run()
	//fmt.Println("main app is running ")
}
