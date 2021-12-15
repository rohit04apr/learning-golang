package main

import "fmt"

func main() {
	var name = "rohit"
	// copy(&name)
	// fmt.Println(name)

	fmt.Println(name)
	fmt.Println(&name)
	copy(&name)
}

func copy(name *string) {
	fmt.Println("copy func :", name)
	*name = "rohit2"
}
