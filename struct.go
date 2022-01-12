package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
	Merried       bool
}

func main() {
	// iterasi data struct ke1
	var rizky Customer
	rizky.Name = "Rizky"
	rizky.Address = "Jakarta"
	rizky.Age = 23

	fmt.Println(rizky.Name)
	fmt.Println(rizky.Address)
	fmt.Println(rizky.Age)

	// iterasi data struct ke2
	wahyu := Customer{
		Name:    "wahyu",
		Address: "Bandung",
		Age:     40,
	}
	fmt.Println(wahyu)
	// iterasi data struct ke3
	/**
	cara iterasi data struct ke 3, datanya harus berurutan jika tdk dan structnya diubah maka akan error
	 */
	hasan := Customer{"hasan", "Surabaya", 20, true}
	fmt.Println(hasan)
}
