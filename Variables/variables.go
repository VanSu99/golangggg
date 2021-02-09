package main

import (
	"fmt"
	"math"
)

func variables() {
	// Boolean -> true / false
	var myBool = true
	fmt.Println("Boolean: ", myBool)

	// String -> buộc nằm trong dấu ""
	var myString = "Anh la vo dich"
	fmt.Println("String: ", myString)

	// int -> số nguyên
	var myInt = 18
	fmt.Println("Int: ", myInt)

	// int 8, 16, 32, 64 -> số lượng bit
	// 1. Range (từ ? đến ?)
	fmt.Println(math.MinInt8) // -128
	fmt.Println(math.MaxInt8) // 127

	// uint -> số nguyên dương
	var myUInt = 10
	fmt.Println(myUInt)

	// float -> số thực
	var myFloat float64 = 12.5
	fmt.Println(myFloat)

	// complex -> số phức
	var myComplex complex64 = 10 + 1i
	fmt.Println(myComplex)

	// Const -> hằng số
	/*
		- không thể thay đổi value của biến
		- auto khởi tạo giá trị cho Enum bằng iota
		- dùng dấu _ để setup biến không cần thiết
	*/
	const myConst = 12
	const (
		_ = iota
		red
		yellow
		pink
		green
	)
	fmt.Println(myConst) // 12
	fmt.Println(red)     // 1
	fmt.Println(yellow)  // 2

	// uintptr -> dùng để luu trữ địa chỉ pointer

}

func main() {
	variables()
}
