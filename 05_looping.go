package main

import "fmt"

func basic_loop() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

func loop_with_condition() {
	i := 1

	for i <= 10 {
		fmt.Println(i)
		i++
	}
}

func loop_with_break() {
	i := 1

	for {
		fmt.Println(i)

		if i == 5 {
			break
		}
		i++
	}
}

func loop_with_continue() {
	for i := 1; i <= 10; i++ {
		if i%2 == 2 {
			continue
		}

		fmt.Println(i)
	}
}

func loop_recursive() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j, " ")
		}

		fmt.Println()
	}
}

func label_loop() {
outer_loop:
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 5; j++ {
			if i == 3 {
				break outer_loop
			}
			fmt.Println("Matrix ", "[", i, "]", "[", j, "]")
		}
	}
}

func main() {
	//basic_loop()
	//loop_with_condition()
	//loop_with_break()
	//loop_with_continue()
	//loop_recursive()
	label_loop()
}
