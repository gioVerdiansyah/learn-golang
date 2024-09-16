package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Print(rand.Int()%100 + 1)
}
