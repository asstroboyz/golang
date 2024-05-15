package main

import "fmt"

func main() {
	names := [...]string{"Eko", "kurniawan", "khannedy", "joko", "budi", "nugraha"}

	// index 4 hingga 6
	slice1 := names[4:6]
	fmt.Println(slice1)

	// index pertama hingga ke 3
	slice2 := names[:3]
	fmt.Println(slice2)

	// index dari ke 3 hingga terakhir
	slice3 := names[3:]
	fmt.Println(slice3)

	// mengambil semua data
	slice4 := names[:]
	fmt.Println(slice4)

	days := [...]string{"senin", "selasa", "rabu", "kamis", "Jumat", "sabtu", "minggu"}

	daySlice1 := days[5:] //sabtu minggu
	fmt.Println(daySlice1)

	daySlice1[0] = "Sabtu Baru" // data 5,6 yang sabtu minggu diganti nama
	daySlice1[1] = "Minggu Baru"
	fmt.Println(daySlice1)
	fmt.Println(days)

	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[0] = "Sabtu Lama"
	// days := [...]string{"senin", "selasa", "rabu", "kamis", "Jumat", "sabtu", "minggu"}
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)

	var newSlice []string = make([]string, 2, 5) // panjnag 2 kapasitas 5
	newSlice[0] = "eko"
	newSlice[1] = "eko"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Eko")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "Budi"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	iniArray := [...]int{1, 2, 3, 4, 5}
	inislice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(inislice)

}
