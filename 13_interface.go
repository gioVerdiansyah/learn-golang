package main

import (
	"fmt"
	"math"
)

type calc interface {
	area() float64
	perimeter() float64
}

type square[T int | float64] struct {
	side T
}

type circle[T int | float64] struct {
	r T
}

//[T int | float64]
// is a multiple data type

func (s square[T]) area() float64 {
	return math.Pow(float64(s.side), 2)
}

func (s square[T]) perimeter() float64 {
	return float64(s.side * 4)
}

func (c circle[T]) area() float64 {
	return math.Pi * math.Pow(float64(c.r), 2)
}

func (c circle[T]) perimeter() float64 {
	return 2 * math.Pi * float64(c.r)
}

func (c circle[T]) diameter() float64 {
	return float64(c.r * 2)
}

// Embedded Interface
type cube interface {
	calc
	volume() float64
}

type cubeStruct struct {
	side float32
}

func (c cubeStruct) area() float64 {
	return 6 * math.Pow(float64(c.side), 2)
}

func (c cubeStruct) perimeter() float64 {
	return 12 * float64(c.side)
}

func (c cubeStruct) volume() float64 {
	return math.Pow(float64(c.side), 3)
}

func main() {
	var c calc

	c = square[int]{side: 12}
	fmt.Println("--------Square--------")
	fmt.Println("Side		:", c.(square[int]).side)
	fmt.Println("Area	 	:", c.area())
	fmt.Println("Perimeter	:", c.perimeter())

	c = circle[float64]{r: 28.8}
	fmt.Println("-------Triangle-------")
	fmt.Println("Radius		:", c.(circle[float64]).r)
	fmt.Println("Diameter 	:", c.(circle[float64]).diameter())
	fmt.Println("Area 		:", c.area())
	fmt.Println("Perimeter 	:", c.perimeter())

	var cu cube
	cu = cubeStruct{side: 20}
	fmt.Println("---------Cube---------")
	fmt.Println("Side		:", cu.(cubeStruct).side)
	fmt.Println("Area		:", cu.area())
	fmt.Println("Perimeter	:", cu.perimeter())
	fmt.Println("Volume		:", cu.volume())

	var freeType interface{}
	freeType = "Verdi"
	fmt.Println("FreeType	:", freeType)

	var bag = []interface{}{1, 2.5, "3+", []int{4, 5}}
	for _, v := range bag {
		fmt.Printf("%v\n", v)
	}
}
