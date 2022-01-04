package main

import "fmt"

func main() {
	type NoKTP string
	type Merried bool

	var NoKtpRizky NoKTP = "982329832983788"
	fmt.Println(NoKtpRizky)

	var merriedStatus Merried = false
	fmt.Println(merriedStatus)
}
