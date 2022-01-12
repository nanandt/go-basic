package main

import "fmt"

func main() {
	ujian := 80
	absen := 70

	//lulusUjian := ujian >= 80 // true
	//lulusAbsen := absen > 70 // false
	//fmt.Println(lulusUjian)
	//fmt.Println(lulusAbsen)
	//
	//var naikKelas bool = lulusUjian && lulusAbsen
	//fmt.Println(naikKelas)

	fmt.Println(ujian >= 80 && absen > 70)

	
}
