package main

import "fmt"

func var_basic() {
	var firstName string = "Verdi"
	fmt.Println("Hello", "world")
	fmt.Printf("Nama saya adalah %s", firstName)
}

func var_notype() {
	freeVar := "Iam free"
	fmt.Print(freeVar)
}

var test_var = "tess"

//test_var := "tess" //error

func multi_var() {
	var one, two, three string
	one, two, three = "satu", "dua", "tiga"

	var four, five, six string = "empat", "lima", "enam"

	seven, eight, nine := "tujuh", 8, "sembilan"

	println(one, two, three, four, five, six, seven, eight, nine)
}

func underscore_var() {
	//	Underscore ( _ ) adalah reserved variable yang bisa dimanfaatkan untuk menampung nilai yang tidak dipakai.
	_ = "Iam zero"
	name, _ := "Verdi", "black hole"

	println(name)
}

func keywoard_new() {
	//	Keyword new digunakan untuk membuat variabel pointer dengan tipe data tertentu. Nilai data default-nya akan menyesuaikan tipe
	//datanya.

	name := new(string)
	println(name)  //0xc000055f30
	println(*name) // ""
}

func main() {
	//var_basic()
	//var_notype()
	//multi_var()
	//underscore_var()
	keywoard_new()
}
