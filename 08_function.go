package main

import (
	"fmt"
	"math"
	"math/rand/v2"
	"strings"
)

func printMessage(message string, data []string) {
	dataJoin := strings.Join(data, " ")
	fmt.Println(message, dataJoin)
}

func randNumRange(min int, max int) int {
	return rand.Int()%(max-min+1) - min
}

func divide(m int, n int) {
	if n == 0 {
		fmt.Printf("Invalid divider, %d cannot divide by %d", m, n)
		return //break program
	}

	fmt.Printf("Divide by %d:%d = %d", m, n, m/n)
}

func multipleReturn(a, h float64) (float64, float64) {
	A := 0.5 * a * h
	c := math.Sqrt(math.Pow(a, 2) + math.Pow(h, 2))
	P := a + h + c

	return A, P
}

func predefinedReturnVal(a, h float64) (area float64, perimeter float64) {
	area = 0.5 * a * h
	c := math.Sqrt(math.Pow(a, 2) + math.Pow(h, 2))
	perimeter = a + h + c

	return
}

func variadic(x ...int) float64 {
	var total int
	for val := range x {
		total += val
	}

	return float64(total / len(x))
}

func dynamicVariadic(n ...interface{}) {
	for _, val := range n {
		fmt.Printf("Type: %T Value: %v\n", val, val)
	}
}

func basicClosure() {
	avg := func(n ...int) float64 {
		var total int
		for _, val := range n {
			total += val
		}

		return float64(total / len(n))
	}

	nums := []int{9, 9, 4, 5, 2, 3, 5, 3}
	fmt.Printf("from avg %v is %.2f", nums, avg(nums...))
	// %v is used to display data of any type
}

func IIFEClosure() {
	//	Closure will be executed immediately upon declaration.

	randNums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	searchEvenNums := func(n ...int) []int {
		var evenNums []int
		for _, val := range n {
			if val%2 == 0 {
				evenNums = append(evenNums, val)
			}
		}

		return evenNums
	}(randNums...)

	fmt.Printf("Original: %v\n", randNums)
	fmt.Printf("Even nums: %v", searchEvenNums)
}

func findMax(numbers []int, max int) (int, func() []int) {
	var res []int
	for _, e := range numbers {
		if e <= max {
			res = append(res, e)
		}
	}
	return len(res), func() []int {
		return res
	}
}

// Alias Scheme Structure
type callbackStruct func(e string) bool

// func callback(data []string, callback func(each string) bool) {
func callback(data []string, callback callbackStruct) {
	var tmp []string
	for _, v := range data {
		if filtered := callback(v); filtered {
			tmp = append(tmp, v)
		}
	}

	fmt.Print(tmp)
}

func main() {
	//printMessage("Hello", []string{"Ryo", "Yamada"})
	//fmt.Print(randNumRange(1, 5))
	//divide(144, 5)
	//a, p := multipleReturn(12, 4)
	//a, p := predefinedReturnVal(125.08, 40.12)
	//fmt.Printf("Area: %.2f Perimeter: %.2f", a, p)
	//dynamicVariadic(1, 2, 3, "Hello", "World", []int{1, 2, 3, 4, 5})
	//basicClosure()
	//IIFEClosure()
	//length, getVal := findMax([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5)
	//fmt.Printf("Length: %d Value: %v\n", length, getVal())

	filterCallback := callbackStruct(func(e string) bool {
		return strings.Contains(strings.ToLower(e), "c")
	})
	callback([]string{"Apple", "Cocomelon", "Banana", "Coconut"}, filterCallback)
}
