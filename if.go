package main

import "fmt"

func main() {
	name := "hendr"

	if name == "rizky"{
		fmt.Println("Hello", name)
	} else if name == "rizzky" {
		fmt.Println("Hello", name)
	} else {
		fmt.Println("kamu Siapaaa?", name)
	}

	if length := len(name); length > 5 {
		fmt.Println("nama terlalu panjang")
	} else {
		fmt.Println("nama sudah cukup")
	}

	
}
