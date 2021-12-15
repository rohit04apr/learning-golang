package main

import "fmt"

func main() {
	// string - anything in quotes
	var name1 string
	name1 = "Rohit1"

	// var name2 string = "Rohit2"

	// name3 := "Rohit3"

	fmt.Println(name1)
	// fmt.Println(name2)
	// fmt.Println(name3)

	// int - number
	var number int
	number = 3

	fmt.Println(number)

	// float - decimal

	var pie float32
	pie = 3.14

	fmt.Println(pie)

	// bool - true/false
	var flag bool
	flag = true

	fmt.Println(flag)

	var newVar string = "old var"
	fmt.Println(newVar)
	newVar = "new var"

	fmt.Println(newVar)

	// const newVar = "old var"
	// fmt.Println(newVar)

	fmt.Println(name1)
	fmt.Printf("%s\n", name1)
	fmt.Printf("%d\n", number)
	fmt.Printf("%f\n", pie)
	fmt.Printf("%t\n", flag)

}
