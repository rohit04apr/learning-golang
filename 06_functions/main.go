package main

import "fmt"

func new() {
	fmt.Println("new function")
}

func sum(num1 int, num2 int) {
	sum := num1 + num2

	fmt.Println(sum)
}

func sum1(num1 int, num2 int) (int, string) {
	sum := num1 + num2
	msg := ""

	if sum < 10 {
		msg = "less than 10"
	} else {
		msg = "greater than 10"
	}

	return sum, msg
}

func sum2(num1 int, num2 int) (sum int, msg string) {
	sum = num1 + num2
	msg = ""

	if sum < 10 {
		msg = "less than 10"
	} else {
		msg = "greater than 10"
	}

	return
}

func main() {
	s, m := sum2(2, 11)
	fmt.Println(s)
	fmt.Println(m)
}
