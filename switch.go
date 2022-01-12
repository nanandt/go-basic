package main

import "fmt"

func main() {
	name := "rizky"

	switch name {
	case "rizky":
		fmt.Println("Hello", name)
	case "Gaga":
		fmt.Println("Hello", name)
	default:
		fmt.Println("Kamu siapaaa?")
	}
	// short statement
	//switch length := len(name); length > 5 {
	//case true:
	//	fmt.Println("Nama terlalu panjang")
	//case false:
	//	fmt.Println("Nama sudah benar")
	//}

	// switch tanpa kondisi
	length := len(name)
	switch {
	case length > 5:
		fmt.Println("kepanjangan boss")
	case length <= 5:
		fmt.Println("nama sudah cukup")
	default:
		fmt.Println("nama benar")
	}


}
