package main

import "fmt"

func main() {
	// i := 0
	// for i <= 5 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// fmt.Println("lasr value of i : ", i)

	// for {
	// 	fmt.Println("This will print forever...")
	// }

	// for i := 0; i <= 5; i++ {
	// 	if i == 3 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }

	// breakpoint:
	// 	for x := 0; x < 3; x++ {
	// 		for y := 0; y < 3; y++ {
	// 			if x == y && y == 1 {
	// 				break breakpoint
	// 			}
	// 			fmt.Println("X:", x, "Y:", y)
	// 		}
	// 	}

	class_details := map[string]int{"class1": 20, "class2": 40, "class3": 30}
	for _, v := range class_details {
		// fmt.Println("where key is :", k)
		fmt.Println("where value is :", v)
	}
}
