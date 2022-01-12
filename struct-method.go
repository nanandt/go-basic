package main

import "fmt"

type Person struct {
	Name, Country string
	Age           int
}

//func sayHi(person Person, name string) {
//	fmt.Println("Hello ", name, "My name is ", person.Name)
//}
func (person Person)sayHi(name string) {
	fmt.Println("Hello", name, "My name is ", person.Name)
}

func (p Person) sayHuu(){
	fmt.Println("Huu from ", p.Name)
}

func main() {
	var rizky Person
	rizky.Name = "rizky"
	rizky.Country = "Amerika"
	rizky.Age = 23

	//sayHi(rizky,"tubagus")
	rizky.sayHi("tubagus")
	rizky.sayHuu()
}
