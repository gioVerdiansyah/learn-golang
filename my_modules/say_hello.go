package my_modules

import "fmt"

// Exported / Public
func SayHello(name string) {
	fmt.Print("Hello ", name)
}

// Unexported / Private
func sayHello(name string) {
	fmt.Print("Hello ", name)
}

// The difference is whether or not a module begins with an uppercase letter.
