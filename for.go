package main

import "fmt"

func main() {
	//counter := 1
	//for counter <= 10{
	//	fmt.Println("perulangan ke ",counter)
	//	counter++
	//}

	// for dengan statement
	for number := 1;number <= 10; number++{
		fmt.Println("Nomer ke ", number)
	}

	slice := []string{
		"rizky","wawan","agung",
	}
	for i := 0; i<len(slice);i++{
		fmt.Println(slice[i])
	}

	for _, value := range slice{
		//fmt.Println("index ",index, "=", value)
		fmt.Println(value)
	}

	person := make(map[string]string)
	person["nama"] = "ayu"
	person["country"] = "indonesia"

	for _, value := range person{
		fmt.Println(value)
	}
}
