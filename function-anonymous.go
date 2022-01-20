package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, b Blacklist) {
	if b(name) {
		fmt.Println("You are Blocked ", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "admin"
	}

	registerUser("admin", blacklist)
	registerUser("rizky", blacklist)

	registerUser("root", func(name string) bool {
		return name == "root"
	})
}
