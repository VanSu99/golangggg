package main

import (
	"fmt"
)

func main() {
	// Array

	/*
		- Chúng ta khai báo cho array có 4 items nhưng ta chỉ truyền vào có 2 items thì các item còn lại sẽ có default value = 0
		=> [100, 200, 0, 0]
		- truy cập phần tử trong array thông qua index
		- array là value type không phải là reference type -> sự khác biệt so vói các language lập trình khác.
	*/

	// Khai báo c1 : [tên_arr] := [số_lượng_item] type {value_items}
	myArr := [4]int{1, 2, 3, 4}
	fmt.Println(myArr) // [1,2,3,4]

	myArr2 := [3]int{100}
	fmt.Println(myArr2) // [100, 0, 0]

	// Size của array
	fmt.Println(len(myArr)) // 4

	// khai báo array không cần setup size
	// dấu [...] thì complier sẽ auto đếm số lượng phần tử có trong array
	myArr3 := [...]string{"BMW", "Honda", "Dream"}
	fmt.Println(myArr3)
	fmt.Println(myArr3[1])

	// Value type != Reference type
	countries := [...]string{"VietNam", "China", "UK", "France"}

	/*
		- khi ta gán các giá trị của biến 'countries' cho 'copyCountries' thì biến 'copyCountries' sẽ tự tạo ra 1 vùng nhớ riêng biệt với biến 'countries'.
		Cho nên dù 'copyContries' có thay đổi gì trong array cũng không làm ảnh hưởng đến biến 'countries'.
	*/
	copyCountries := countries

	copyCountries[0] = "USA"
	fmt.Println(countries)               // [VietNam, China, UK, France]
	fmt.Println("Copy: ", copyCountries) // [USA, China, UK, France]

	// Loop qua từng phần tử trong array
	// Cách 1:
	for i := 0; i < len(countries); i++ {
		fmt.Println(countries[i])
	}

	// cách 2:
	for index, value := range countries {
		fmt.Printf("i=%d value=%s", index, value)
		fmt.Println()
	}

	// mảng 2 chiều [row][column]
	matrix := [4][2]int{
		{1, 2},
		{1, 2},
		{3, 5},
		{6, 7},
	}
	fmt.Println(matrix)

	// Loop mảng 2 chiều
	for i := 0; i < 4; i++ {
		for j := 0; j < 2; j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}

}
