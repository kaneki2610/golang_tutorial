package main

import "fmt"

func main() {
	//1.khai báo biến
	var number int
	number = 10
	fmt.Println(number)

	var number2 int = 100
	fmt.Println(number2)

	//type inference
	var number3 = 101
	fmt.Println(number3)

	//2.khai báo nhiều biến
	//2.1 Khai báo nhiều biến cùng kiểu dữ liệu
	var a,b int
	a = 102
	b = 103
	fmt.Println(a)
	fmt.Println(b)

	var c,d int = 104,105
	fmt.Println(c + d)

	var e,f = 105,106
	fmt.Println(e+f)

	//2.2 Khai báo nhiều biến khác kiểu dữ liệu
	var a1,a2,a3 = "a1", "a2", 3
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)

	//khai báo short hand
	tes2:= 100
	fmt.Println(tes2)
}
