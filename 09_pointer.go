package main

import "fmt"

func pinterBasic() {
	//Pointer type variables are characterized by an asterisk sign ( * )
	var num int = 5
	var numPtr *int = &num

	fmt.Println("Number (value)		: ", num)
	fmt.Println("Number (address)	: ", &num)

	fmt.Println("Pointer (address)	 	: ", numPtr)
	fmt.Println("Pointer (value)	: ", *numPtr)
}

func valueChanges() {
	var numA int = 5
	var numB *int = &numA
	fmt.Println("NumberA (value)		: ", numA)
	fmt.Println("NumberA (address)	: ", &numA)
	fmt.Println("NumberB (address)	: ", numB)
	fmt.Println("NumberB (value)		: ", *numB)

	fmt.Println("")

	numA = 10

	fmt.Println("NumberA (value)		: ", numA)
	fmt.Println("NumberA (address)	: ", &numA)
	fmt.Println("NumberB (address)	: ", numB)
	fmt.Println("NumberB (value)		: ", *numB)
}

func changeWithoutPointer(v int) {
	v = 5
}
func changeWithPointer(v *int) {
	*v = 5
}

func main() {
	//pinterBasic()
	//valueChanges()

	val := 10
	fmt.Printf("Original: %d\n", val)
	changeWithoutPointer(val)
	fmt.Printf("Change without pointer: %d\n", val)
	changeWithPointer(&val)
	fmt.Printf("Change with pointer: %d\n", val)
}
