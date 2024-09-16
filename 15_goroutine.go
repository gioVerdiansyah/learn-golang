package main

import (
	"fmt"
	"runtime"
)

func printLoop(length int, message string) {
	for i := 0; i < length; i++ {
		fmt.Println(i+1, message)
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // auto adjust

	go printLoop(20, "Hello")
	printLoop(10, "Hai")

	//	Blocking
	var input string
	fmt.Scanln(&input)
}
