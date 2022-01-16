package main

import "fmt"

func main() {
	name := "rizky"
	counter := 0
	increment := func() {
		name = "hasan"
		fmt.Println("increment")
		counter++
	}

	increment()
	increment()
	fmt.Println(counter)
	fmt.Println(name)

	
}
