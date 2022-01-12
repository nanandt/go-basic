package main

import "fmt"

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Terjadi Error dengan :", message)
	}
	fmt.Println("App Selesai")
}


// panic mengentikan progam tetapi defer akan tetap diexecute
// recover tetap menjalankn app meski teerjadi panic dan menangkap pesan panic tsb
// dan menyimpan nya di defer func


func runApp(error bool) {
	defer endApp()
	if error {
		panic("Aplikasi Error")
	}
	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApp(true)
	fmt.Println("rizky")
}
