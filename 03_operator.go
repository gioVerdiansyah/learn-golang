package main

import "fmt"

func aritmatich_operator() {
	val := (((144 * 3) / 5) - 10) / 2

	fmt.Print(val)
}

func comparison_operator() {
	val := (((144 * 3) / 5) - 10) / 2
	equal := val % 2
	fmt.Printf("Value %d %%2 is equal 1 ? %t", val, equal == 1)
}

func logic_operator() {
	stockReady := true
	machineReady := false

	stockAndmachine := stockReady && machineReady
	stockOrMachine := stockReady || machineReady
	stockNotReady := !stockReady
	fmt.Printf("Stock And Machine: %t\n", stockAndmachine)
	fmt.Printf("Stock Or Machine: %t\n", stockOrMachine)
	fmt.Printf("Stock is Ready: %t\n", stockNotReady)
}

func main() {
	//aritmatich_operator()
	//comparison_operator()
	logic_operator()
}
