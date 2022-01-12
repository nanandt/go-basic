package main

import "fmt"

func main() {
	// deklarasi array 1
	var names [3]string
	names[0] = "Muhamad"
	names[1] = "fatih"
	names[2] = "rizky" // mengubah isi index array
	//names[3] = "bagas" tidak bisa di compile karna data arraynya sudah didefinisasikan 3

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// deklarasi array 2
	var values = [3]int{
		20,
		30,
		49,
	}
	fmt.Println(values)

	fmt.Println(len(names))
	fmt.Println(len(values))

}
