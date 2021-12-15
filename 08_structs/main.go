package main

import "fmt"

type person struct {
	first string
	last  string
}

type details struct {
	person
	company string
}

func main() {

	// p1 := person{
	// 	first: "Steve",
	// 	last:  "Jobs",
	// }

	// p2 := person{
	// 	first: "Bill",
	// 	last:  "Gates",
	// }

	det1 := details{
		person: person{
			first: "Steve",
			last:  "Jobs",
		},
		company: "Apple",
	}

	det2 := details{
		person: person{
			first: "Bill",
			last:  "Gates",
		},
		company: "Microsoft",
	}

	fmt.Printf("FIRST NAME : %s, LAST NAME : %s, COMPANY : %s\n", det1.first, det1.last, det1.company)
	fmt.Printf("FIRST NAME : %s, LAST NAME : %s, COMPANY : %s\n", det2.first, det2.last, det2.company)

}
