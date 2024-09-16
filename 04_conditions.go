package main

import "fmt"

func ifelse() {
	point := 88

	if point == 100 {
		print("Perfect")
	} else if point >= 80 {
		print("Good")
	} else {
		print("enough")
	}
}

func temporary_var() {
	if point := 85; point >= 90 {
		print("A")
	} else if point >= 80 {
		print("B")
	} else if point >= 70 {
		print("C")
	} else {
		print("D")
	}
}

func switchcase() {
	grade := "A"

	switch grade {
	case "A":
		print("Perfect")
	case "B":
		print("Good")
	case "C":
		print("Enough")
	default:
		print("Bad")
	}
}

func switchcase_next() {
	point := 7

	switch point {
	case 10:
		fmt.Print("Perfect")
	case 7, 8, 9:
		fmt.Print("Good")
	default:
		{
			print("Very bad \n")
			print("Try learn more")
		}
	}
}

func switchcase_next_2() {
	point := 6
	switch {
	case point == 10:
		print("Perfect")
	case (point <= 7) && (point >= 9):
		print("Good")
	default:
		{
			print("Enough")
			print("You must learn more")
		}
	}
}

func switchcase_fallthrough() {
	point := 8

	switch {
	case (point >= 9):
		fmt.Println("Perfect")
	case (point <= 8) && (point >= 6):
		fmt.Println("Awesome")
		fallthrough
	case (point <= 5):
		fmt.Println("Keep going")
	default:
		fmt.Println("Bad")
	}
}

func main() {
	//ifelse()
	//temporary_var()
	//switchcase()
	//switchcase_next()
	//switchcase_next_2()
	switchcase_fallthrough()
}
