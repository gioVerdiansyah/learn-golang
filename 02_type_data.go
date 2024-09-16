package main

import (
	"fmt"
)

//uint8 0 ↔ 255
//uint16 0 ↔ 65535
//uint32 0 ↔ 4294967295
//uint64 0 ↔ 18446744073709551615
//uint sama dengan uint32 atau uint64 (tergantung nilai)
//byte sama dengan uint8
//int8
//-128 ↔ 127
//int16
//-32768 ↔ 32767
//int32
//-2147483648 ↔ 2147483647
//int64
//-9223372036854775808 ↔ 9223372036854775807
//int sama dengan int32 atau int64 (tergantung nilai)
//rune sama dengan int32

func num_n_dec() {
	var positiveNum uint32 = 34
	negativeNum := -1234566

	fmt.Printf("positiveNum:%d negativeNum:%d\n", positiveNum, negativeNum)
}

func num_dec() {
	decimalNum := 2.66
	fmt.Printf("decimalNum:%f\n", decimalNum)
	fmt.Printf("decimalNum:%.4f\n", decimalNum)
}

func boolean() {
	var open bool = false

	fmt.Printf("is open? %t", open)
}

func str() {
	var msg string = "Hello World"
	fmt.Println(msg)

	msg = `Hallo saya verdi.
Salam kenal.
mari kita sama sama belajar Golang.`

	fmt.Println(msg)
}

func constant() {
	const firstName string = "Sofyan"
	const lastName = "Verdi"

	fmt.Println(firstName, lastName)
}

func main() {
	//num_n_dec()
	//num_dec()
	//boolean()
	//str()
	constant()
}
