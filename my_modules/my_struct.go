package my_modules

import f "fmt"

type Person struct {
	Name    string
	Age     int
	Married bool
}

type person struct {
	Name    string
	Age     int
	Married bool
}

func PrintIdentityPerson(p Person) {
	f.Printf("Name: %s\n Age: %d\n Married: %t\n", p.Name, p.Age, p.Married)
}

func printIdentityPerson(p person) {
	f.Printf("Name: %s\nAge: %d\nMarried: %t\n", p.Name, p.Age, p.Married)
}
