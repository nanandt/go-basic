package main

import "fmt"

func main() {
	var name string // cara deklarasi variable ke 1

	name = "wahyu aji leksono"
	fmt.Println(name)

	name = "jaka harun"
	fmt.Println(name)

	var country = "indonesia" // cara deklarasi variable ke 2
	fmt.Println(country)

	var age = 70
	fmt.Println(age)

	firstName := "muhamad" // cara deklarasi variable ke 3, dan hanya untuk deklarasi awal saja
	fmt.Println(firstName)

	var ( // multiple variable
		lastName = "fatikhurrizky"
		middleName = "kechaw"
	)

	fmt.Println(lastName)
	fmt.Println(middleName)
}
