package main

import "fmt"

func basic() {
	var strVal [4]string

	strVal[0] = "Golang"
	strVal[1] = "is"
	strVal[2] = "fun"
	strVal[3] = "!"
	fmt.Println(strVal[0], strVal[1], strVal[2], strVal[3])
}

func array_init() {
	var ranArr = [4]string{"My", "name", "is", "Verdi"}
	var fruits = [3]string{
		"Apel",
		"Banana",
		"Lemon",
	}

	fmt.Println(ranArr)
	fmt.Println("And I likes ", len(fruits), " fruits:")
	fmt.Println(fruits)
}

func dynamic_array() {
	var progLang = [...]string{
		"PHP",
		"JS",
		"Python",
		"Golang",
	}

	fmt.Println("I likes ", len(progLang), " programming language:")
	fmt.Println(progLang)
}

func array_multidimension() {
	var myArr = [...][3]int{[3]int{1, 2, 3}, [3]int{4, 5, 6}}
	// when the 1-dimensional array is dynamic and the 2-dimensional array is not dynamic, the contents of the 2-dimensional array must follow
	// the rules of the second dimension array [3]. [3]int grayed out because the 2-dimensional array rule type
	// already represents
	var myArr2 = [...][]int{{1, 2, 3, 4}, {5, 6}}
	// when the parent array is already dynamic and if you want a dynamic 2d array again there is no need
	//to make it dynamic again like this [...][...]

	fmt.Println(myArr)
	fmt.Println(myArr2)
}

func basic_arr_loop() {
	var fruits = [5]string{"Apple", "Orange", "Banana", "Pear", "Grape"}

	for i := 0; i < len(fruits); i++ {
		fmt.Println(fruits[i])
	}
}

func arr_range() {
	var fruits = [5]string{"Apple", "Orange", "Banana", "Pear", "Grape"}

	for i, v := range fruits {
		fmt.Println("Index ", i, " contain ", v)
	}

	//	If i (arr index) doesn't used, you can use _ to ignore that
}

func arr_with_make() {
	var my_arr = make([]string, 2)
	my_arr[0] = "Apple"
	my_arr[1] = "Orange"
	fmt.Println(my_arr)
}

func arr_slice() {
	my_arr := []int{1, 2, 3}

	fmt.Print(my_arr[1])

	//	different array and slice
	_ = []int{1, 2, 3}    // slice
	_ = [3]int{1, 2, 3}   // array
	_ = [...]int{1, 2, 3} // array
}

func arr_and_slice() {
	my_arr := []int{1, 2, 3, 4, 5}
	new_arr := my_arr[1:3]

	fmt.Println(new_arr) // [2 3]

	new_arr[0] = 9
	fmt.Println(new_arr[0])
}

func len_and_cap() {
	my_arr := []int{1, 2, 3, 4, 5}

	//The len() function is used to count the number of slice elements present.
	//The cap() function is used to calculate the maximum width or capacity of the slice.

	fmt.Println(len(my_arr[0:2])) // 2
	fmt.Println(cap(my_arr[0:2])) // 5
}

func arr_append() {
	my_arr := []int{1, 2, 3, 4, 5}
	my_arr2 := append(my_arr, 6)

	fmt.Println(my_arr)  // [1 2 3 4 5]
	fmt.Println(my_arr2) // [1 2 3 4 5 6]
}

func arr_copy() {
	my_fruits := make([]string, 4)
	fruits := []string{"Apple", "Orange", "Banana", "Pear", "Grape"}
	x := copy(my_fruits, fruits)

	fmt.Println(my_fruits) // [Apple Orange Banana Pear]
	fmt.Println(fruits)    // [Apple Orange Banana Pear Grape]
	fmt.Println(x)         // 4

	my_fruits = []string{"Apple", "Orange", "Banana"}
	fruits = []string{"Pear", "Grape"}
	x = copy(my_fruits, fruits)

	fmt.Println(my_fruits) // [Pear Grape Banana]
	fmt.Println(fruits)    // [Pear Grape]
	fmt.Println(x)         // 2
}

func slice_with_3_index() {
	fruits := []string{"apple", "banana", "cherry", "date", "watermelon", "pear", "grape"}
	aFruits := fruits[1:3]
	bFruits := fruits[2:4:4]

	fmt.Println(fruits)
	fmt.Println(len(fruits))
	fmt.Println(cap(fruits))

	fmt.Println(aFruits)
	fmt.Println(len(aFruits))
	fmt.Println(cap(aFruits))

	fmt.Println(bFruits)
	fmt.Println(len(bFruits))
	fmt.Println(cap(bFruits))
}

func main() {
	//basic()
	//array_init()
	//dynamic_array()
	//array_multidimension()
	//basic_arr_loop()
	//arr_range()
	//arr_with_make()
	//arr_slice()
	//arr_and_slice()
	//len_and_cap()
	//arr_append()
	//arr_copy()
	slice_with_3_index()
}
