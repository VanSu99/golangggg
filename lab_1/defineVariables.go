package main

import "fmt"

func main() {
	// function main là hàm đặc biệt được tự động gọi khi chạy chương trình
	var number int = 10
	fmt.Println(number)

	// Type inference - không cần thêm data types mà Go tự hiểu biến đó là thuộc data types nào
	// var [tên biến] = [giá trị]
	var email = "evan@gmail.com"
	fmt.Println(email)

	// Khai báo nhiếu biến
	// case 1: khai báo nhiều biến cùng data type

	var a, b = 1, 2
	fmt.Println(a)
	fmt.Println(b)

	// case 2: khai báo nhiều biến khác data type

	var (
		name    string = "Anh la vo dich"
		age     int    = 18
		address string = "Viet Nam"
	)

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(address)

	// case 2 áp dụng Type inference

	var name2, address2, age2 = "Evan", "Quan 2", 19
	fmt.Println(name2)
	fmt.Println(age2)
	fmt.Println(address2)

	// Short hand : [tên biến] := [giá trị] -> dùng nhiều
	language := "Golang"
	fmt.Println(language)
}
