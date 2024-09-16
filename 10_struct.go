package main

import "fmt"

type PersonStruct struct {
	FullName string
	Name     string
	Age      int
	Married  bool
}

func useStruct() {
	person := PersonStruct{
		FullName: "Verdiansyah",
		Name:     "Verdi",
		Age:      21,
		Married:  false,
	}

	fmt.Println(person)
	fmt.Println("Full Name		: ", person.FullName)
	fmt.Println("Name			: ", person.Name)
	fmt.Println("Age			: ", person.Age)
	fmt.Println("Married			: ", person.Married)
}

type SquareStruct struct {
	width, height int
}

func initStruct() {
	var square1 SquareStruct
	square1.width = 12
	square1.height = 12

	var square2 = SquareStruct{10, 10}

	var square3 = SquareStruct{height: 5, width: 5}

	fmt.Printf("Squeare %d x %d = %d \n", square1.width, square1.height, square1.width*square1.height)
	fmt.Printf("Squeare %d x %d = %d \n", square2.width, square2.height, square2.width*square2.height)
	fmt.Printf("Squeare %d x %d = %d \n", square3.width, square3.height, square3.width*square3.height)
}

// Embedded Struct
type person struct {
	name    string
	age     int
	married bool
}

type student struct {
	grade rune
	class string
	person
}

func embeddedStruct() {
	var _ = student{
		person: person{
			name:    "Verdi",
			age:     21,
			married: false,
		},
		class: "XI RPL 2",
		grade: 'A',
	}

	var s2 = student{}
	s2.name = "Adi"
	s2.age = 20
	s2.married = true
	s2.class = "X RPL 1"
	s2.grade = 'B'

	//	if person and student have the same property name then it can be like this
	//	s2.person.class
}

func anonymousStruct() {
	var s1 = struct {
		person
		class string
		grade rune
	}{}

	s1.person = person{name: "Adi", age: 21, married: false}
	s1.class = "X RPL 1"
	s1.grade = 'B'

	//	Or

	var _ = struct {
		person
		class string
		grade rune
	}{
		person: person{name: "Adi", age: 21, married: false},
		class:  "X RPL 1",
		grade:  'B',
	}
}

func sliceStruct() {
	var persons = []person{
		{name: "Verdi", age: 21, married: false},
		{name: "Adi", age: 20, married: true},
		{name: "Niam", age: 18, married: false},
	}

	for _, v := range persons {
		fmt.Printf("Name: %s \nAge: %d\nMarried: %t\n\n", v.name, v.age, v.married)
	}
}

func sliceAnonymousStruct() {
	var persons = []struct {
		person
		grade rune
	}{
		{person: person{name: "Verdi", age: 21, married: true}, grade: 'A'},
		{person: person{name: "Adi", age: 20, married: false}, grade: 'B'},
		{person: person{name: "Niam", age: 19, married: true}, grade: 'C'},
	}
	for _, v := range persons {
		fmt.Printf("Name: %s \nAge: %d\nMarried: %t\n\n", v.name, v.age, v.married)
	}
}

// Nested struct
type car struct {
	specification struct {
		speed        int
		acceleration int
		engine       string
	}
	merk       string
	color      string
	production string
}

// Tag property
type ReportValue struct {
	grade rune   `A`
	time  string `2024-06-12 13:00`
}

// People Alias
type People = Person

func main() {
	//useStruct()
	//initStruct()
	//sliceStruct()
	sliceAnonymousStruct()
}
