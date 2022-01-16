package main

import "fmt"

func main() {
	name1 := "wahyu"
	name2 := "wahyu"
	var result bool = name1 == name2
	fmt.Println(result)

	value1 := 500
	value2 := 300
	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}
