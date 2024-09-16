package main

// This file only shows how to access module files that can be used publicly or privately.

import (
	example "first_golang/my_modules"
	"fmt"
)

// example is alias. Original is my_module
//Use alias . if you want to import anonymously

func main() {
	example.SayHello("Verdi")
	//example.sayHello("Verdi") // Error because func module is unexported

	//example.PrintIdentityPerson(example.Person{Name: "Verdi", Age: 18, Married: false})

	//	If the same module just called the module name
	//sayToUser() // from partial.go

	//	But to run is different, like that
	// go run 12_public_private_property.go partial.go

	// Init function from module
	fmt.Println("Name: ", example.User.Name)
	fmt.Println("Email: ", example.User.Email)
}
