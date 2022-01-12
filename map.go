package main

import "fmt"

func main() {

	person := map[string]string{
		"name":    "Rizky",
		"address": "tegal",
	}
	person["title"] = "progammer"
	person["title"] = "progammer anyar"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])

	book := make(map[string]string)
	book["title"] = "Buku Cara Ngoding"
	book["author"] = "Anak Singkong"
	book["age"] = "22"

	delete(book, "age")

	fmt.Println(book["title"])
	fmt.Println(book["author"])
	fmt.Println(book["age"])
}
