package main

import "fmt"

func main() {
	//const firstName string = "slamet"
	//const age = 23
	//fmt.Println(firstName)

	const (
	firstName string = "slamet"
	age = 23
	)
	fmt.Println(firstName)
	fmt.Println(age)

	// jika menggunakan const maka tipe datanya tidak bisa berubah
}
