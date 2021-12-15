package main

import "fmt"

func main() {
	// var flag bool = false

	// if flag {
	// 	fmt.Println("condition matches")
	// } else {
	// 	fmt.Println("condition doesn't match")
	// }

	var age int = 128

	// 0-2 infant
	// 2-5 toddlers
	// 5-12 child
	// 12-18 teen
	// 18+ adult

	// if age > 0 && age <= 2 {
	// 	fmt.Println("Infant")
	// } else if age > 2 && age <= 5 {
	// 	fmt.Println("Toddlers")
	// } else if age > 5 && age <= 12 {
	// 	fmt.Println("Child")
	// } else if age > 12 && age <= 18 {
	// 	fmt.Println("teen")
	// } else {
	// 	fmt.Println("Adult")
	// }

	switch {
	case age > 0 && age <= 2:
		fmt.Println("Infant")
	case age > 2 && age <= 5:
		fmt.Println("Toddlers")
	case age > 5 && age <= 12:
		fmt.Println("Child")
	case age > 12 && age <= 18:
		fmt.Println("teen")
	default:
		fmt.Println("Adult")
	}
}
