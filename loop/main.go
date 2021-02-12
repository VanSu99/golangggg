package main

import (
	"fmt"
)


func main() {
	// For Loop
	/*
		- dùng để lặp qua các phần tử trong 1 array
		- vòng for sẽ bao gồm 3 yếu tố sau:
			+ biến chạy từ
			+ điểm dừng
			+ bước nhảy
		- keyword: break -> thoát khỏi loop
	*/

	for i := 1; i <= 10; i++ {
		for j := 1; i < 5; i++ {
			fmt.Println(i * j)
		}
	}

	/*
		- mỗi lần lặp i sẽ thực hiện tất cả các lần lặp j ->
		tức là: sẽ lấy 1 của lần loop 1 của i * loop 1 của j, tiếp tục: 1*2, 1*3, 1*4, 1*5 -> kết thúc vòng loop của j.
		- tiếp tục loop lần 2 của i -> cứ thế mà xoay vòng
	*/

	// dùng keyword: range để duyệt các phần tử trong array
	myArr := [3]int{1,2,3}
	for index, value := range myArr {
		fmt.Println(index, value)
	}

	mySlice := []int{1,2,3,4} // Slice
	for index, value := range mySlice {
		fmt.Println(index, value)
	}

	// Map
	myMap := map[string]int {
		"Evan": 18,
		"Tom": 21,
		"Selena": 28,
	}
	for key, value := range myMap {
		fmt.Println(key, value)
	}

	// String
	myStr := "Hello world"
	for index, value := range myStr {
		// value sẽ là mã unicode
		// muốn xuất ra ký tụ thì ép kiểu string
		fmt.Println(index, string(value))
	}

	/*
	- dùng 'range' để duyệt các phần tử giá trị arg đầu tiên luôn trả về index.
	*/

	// muốn lấy ra value mà thôi -> dùng dấu _
	for _, value := range myStr {
		fmt.Println(string(value))
	}

}
