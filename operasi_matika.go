package main

import "fmt"

func main() {
	var a = 10
	var b = 10
	var d = 12
	var e = 5
	var c = a - b + d*e
	fmt.Println(c)

	var i = 10
	i += 5 // menambah  nilai i => i= i +5
	fmt.Println(i)

	//unary operator
	var k = 1
	k++ // k = k +1
	fmt.Println(k)
	k++ // k = k + 1
	fmt.Println(k)

	k-- // k = k -1
	fmt.Println(k)
	k-- // k = k - 1
	fmt.Println(k)

}
