package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func basicHandlingError() {
	var input string
	fmt.Print("Enter a number: ")
	fmt.Scanln(&input)

	convVal, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println("Your input isn't number")
		fmt.Println(err.Error())
	} else {
		fmt.Println("Your input is:", convVal)
	}
}

func validate(val string) (bool, error) {
	if strings.TrimSpace(val) == "" {
		return false, errors.New("value can't be empty")
	}

	return true, nil
}
func customError() {
	var input string
	fmt.Print("Enter something: ")
	fmt.Scanln(&input)

	val, err := validate(input)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Your input is:", val)
	}
}

func usePanic() {
	var input string
	fmt.Print("Enter a number: ")
	fmt.Scanln(&input)

	if val, err := strconv.Atoi(input); err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Your input is:", val)
	}
}

func catchStrConv() {
	if e := recover(); e != nil {
		fmt.Println("Error Occurred!")
	} else {
		fmt.Println("Application run normally!")
	}
}
func useRecover() {
	defer catchStrConv()

	var input string
	fmt.Print("Enter a number: ")
	fmt.Scanln(&input)

	if val, err := strconv.Atoi(input); err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Your input is:", val)
	}
}

func errorIIFE() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error Occurred:", err)
		}
	}()

	panic("Something bad happened!")
}

func errorIIFELoop() {
	for _, val := range []string{"A", "B", "C"} {
		fmt.Println(val)
		func() {
			defer func() {
				if err := recover(); err != nil {
					fmt.Println("Error Occurred:", err)
				} else {
					fmt.Println("Application run normally!")
				}
			}()
			panic("Something bad happened!")
		}()
		//The loop is not stopped because it has been wrapped by the IIFE function
	}
}

func main() {
	//basicHandlingError()
	//customError()
	//usePanic()
	//useRecover()
	//errorIIFE()
	errorIIFELoop()
}
