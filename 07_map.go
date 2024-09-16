package main

import (
	"fmt"
	"strconv"
)

func basic_map() {
	var month map[string]int
	month = map[string]int{}

	//Others method to initialization
	var _ = map[string]int{}
	var _ = make(map[string]int)
	var _ = *new(map[string]int)

	month["january"] = 1
	month["february"] = 2
	month["march"] = 3

	fmt.Println("January num: ", month["january"])
	fmt.Println("February num: ", month["february"])
	fmt.Println("March num: ", month["march"])

	var person = map[string]string{
		"firstName": "Sofyan",
		"lastName":  "Verdi",
		"age":       strconv.Itoa(18),
	}

	fmt.Println(person["firstName"])
	fmt.Println(person["lastName"])
}

func iteration_with_range() {
	// interface{} is an empty data type in Go that can store values of any data type
	my_obj := map[string]interface{}{
		"firstName": "Sofyan",
		"lastName":  "Verdiansyah",
		"age":       strconv.Itoa(18),
		"married":   false,
	}

	for key, value := range my_obj {
		fmt.Println("Object Key:", key, " \t | value:", value)
	}
}

func delete_item() {
	my_obj := map[string]string{
		"firstName": "Sofyan",
		"lastName":  "Verdiansyah",
		"age":       strconv.Itoa(18),
	}

	fmt.Println(len(my_obj)) // 3
	delete(my_obj, "age")
	fmt.Println(len(my_obj)) // 2
}

func check_key() {
	my_obj := map[string]string{
		"firstName": "Sofyan",
		"lastName":  "Verdiansyah",
		"age":       strconv.Itoa(18),
	}

	value, isMarried := my_obj["married"]

	if isMarried {
		fmt.Println(my_obj["firstName"], my_obj["lastName"], "has been married: ", value)
	} else {
		fmt.Println("Item married doesn't exist")
	}
}

func map_and_slice() {
	genders := []map[string]string{
		{"name": "Verdi", "gender": "male"},
		{"name": "Adi", "gender": "male"},
		{"name": "Beta", "gender": "female"},
	}

	for _, value := range genders {
		fmt.Println(value["name"], "gender is", value["gender"])
	}
}

// From mas Radar
type Person struct {
	name        string
	description string
	age         int
}

func data_struct() {
	persons := Person{
		name:        "radar",
		description: "hallo dek",
		age:         1,
	}

	var person = map[string]string{
		"firstName": persons.name,
		"lastName":  persons.description,
		"age":       strconv.Itoa(persons.age),
	}

	fmt.Println(person["firstName"])
	fmt.Println(person["lastName"])
	fmt.Println(person["age"])
}

func main() {
	//basic_map()
	//iteration_with_range()
	//delete_item()
	//check_key()
	map_and_slice()
}
