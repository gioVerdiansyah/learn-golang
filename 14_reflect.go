package main

import (
	"fmt"
	"reflect"
)

func basicUse(val interface{}) {
	refVal := reflect.ValueOf(val)
	fmt.Println("Type		:", refVal.Type())       // Specific data type
	fmt.Println("Kind		:", refVal.Kind())       // General data type
	fmt.Println("Value		:", refVal.Interface()) // Interface return anything value
	//fmt.Println("Value		:", refVal.Interface().([]string)) // Casting
	if refVal.Kind() == reflect.Slice {
		for i := 0; i < refVal.Len(); i++ {
			fmt.Println(refVal.Index(i))
		}
	}
}

type name struct {
	firstName string
	lastName  string
	length    int
}

func (n *name) getPropertyInfo() {
	refVal := reflect.ValueOf(n)

	if refVal.Kind() == reflect.Ptr {
		refVal = refVal.Elem()
	}

	refType := refVal.Type()

	for i := 0; i < refVal.NumField(); i++ {
		fmt.Println("Name	:", refType.Field(i).Name)
		fmt.Println("Type	:", refType.Field(i).Type)
		fmt.Println("Value	:", refVal.Field(i))
		fmt.Println("")
	}
}

func (n *name) SetFirstName(name string) {
	n.firstName = name
}
func (n *name) SetLastName(name string) {
	n.lastName = name
}

func main() {
	//basicUse([]string{"a", "b", "c", "d"})

	name := &name{
		firstName: "Sofyan",
		lastName:  "Verdi",
		length:    24,
	}
	//name.getPropertyInfo()

	var refVal = reflect.ValueOf(name)
	// List Method
	fmt.Println(refVal.NumMethod())

	//var method = refVal.MethodByName("SetFirstName")
	var method = refVal.Method(0)
	method.Call([]reflect.Value{
		reflect.ValueOf("Adi"),
	})

	fmt.Println("FirstName	:", name.firstName)
	fmt.Println("LastName	:", name.lastName)
}
