package main

import "fmt"

type Name struct {
	Name string
	From string
}

func (n Name) getMethod(message string) {
	fmt.Println(message, n.Name, "from", n.From)
}

// Method Pointer
func (n Name) changeName(name string) {
	n.Name = name
}

func (n *Name) changeNamePointer(name string) {
	n.Name = name
}

func main() {
	person := Name{Name: "John", From: "America"}
	//person.getMethod("Hello")
	person.changeName("Verdi")
	fmt.Println(person.Name)
	person.changeNamePointer("Sofyan")
	fmt.Println(person.Name)
}
