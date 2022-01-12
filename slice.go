package main

import "fmt"

func main() {

	// deklarasi array 3 (tidak tahu kapasitas arraynya)
	var months = [...]string{
		"januari",
		"februari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}
	fmt.Println(months)

	fmt.Println("===========================")
	slice1 := months[3:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
	fmt.Println("===========================")
	//months[4] = "bukan mei"
	//fmt.Println(slice1)

	//slice1[0] = "bukan april"
	//fmt.Println(slice1)

	days := [...]string{
		"senin",
		"selasa",
		"rabu",
		"kamis",
		"jum'at",
		"sabtu",
		"minggu",
	}
	daySlice1 := days[5:]
	daySlice1[0] = "sabtu baru"
	daySlice1[1] = "minggu baru"
	fmt.Println(days)

	daySlice2 := append(daySlice1,"libur baru")
	fmt.Println(daySlice2)
	fmt.Println("=========")
	daySlice2[0] = "ups"
	fmt.Println(daySlice2)
	fmt.Println(days)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "gunawan"
	newSlice[1] = "dwi"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string,len(newSlice),cap(newSlice))
	copy(copySlice,newSlice)
	fmt.Println(copySlice)

	iniArray := [5]int{
		1,2,3,4,5,
	}

	iniSlice := []int{1,2,3,4,5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
