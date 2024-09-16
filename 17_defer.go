package main

import (
	"fmt"
	"os"
	"strings"
)

func useDefer() {
	defer fmt.Println("Hello Defer") // Always in end line
	fmt.Println("Hello World")
}

func deferReturn(menu string, orderType string) {
	defer fmt.Println("Thank You for order, please wait your orders")

	if strings.ToLower(menu) == "pizza" && strings.ToLower(orderType) == "big" {
		fmt.Println("Ordering Big Pizza got 15% discount!")
	}

	fmt.Println("Your order is", menu, "type is", orderType)
}

func deferBlock() {
	var num int = 3

	fmt.Println("begin")
	if num == 3 {
		fmt.Println("main 1")
		defer fmt.Println("defer 1")
	}

	fmt.Println("end")

	fmt.Println("")

	fmt.Println("begin")
	if num == 3 {
		fmt.Println("main 2")
		func() {
			defer fmt.Println("defer 2")
			// will not in the end
		}()
	}

	fmt.Println("end")
}

func exit() {
	fmt.Println("Welcome")
	os.Exit(1)

	defer fmt.Println("End")
	fmt.Println("Hello User")
}

func main() {
	//useDefer()
	//deferReturn("Pizza", "Big")
	//deferReturn("Burger", "Small")
	//deferBlock()
	exit()
}
