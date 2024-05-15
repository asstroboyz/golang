package main

import "fmt"

func main() {
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke", counter)
	// 	counter++ //menaikkan counter satu persatu hingga maksimal 10 karena kurang dari atau sama dengan 10
	// }

	//for statement
	for counter :=1; counter <=10; counter++{
		fmt.Println("Perulangan ke", counter)
	}

	fmt.Println("selesai")

	// init statment yaitu sebelum for di eksekusi
	//post statement yaitu statement yang akan selalu di eksekusi di akhir tiap perulangan

	//for range
	names :=[]string{"Ganda","Gunawan","Risdandi"}
	for i :=0; i< len(names); i++ {
		fmt.Println(names[i])
	}
	// cara manual diatas

	for index, name := range names{
		fmt.Println("Index", index, "=",name)
	} // cetak bersama index ke

	for _, name := range names{
		fmt.Println(name)
	} // cetak tanpa index
}