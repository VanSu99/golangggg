package main

import "fmt"

type People struct {
	name string
	age  int
}

// Thừa kế từ 1 Struct
type Student struct {
	// define attribute
	People
	mssv     int
	gender   bool
	subjects []string
}

func structTypes() {
	// Struct

	/*
		- Là 1 nơi chứa nhiều data types khác nhau.
		-	return về 1 Object.
		- Cách đặt tên Struct:
			+ viết hoa tên struct
		- không tham chiếu tương tự như Array.
	*/

	// Cách khai báo 1 struct và object nhanh
	studentYub := struct{ name string }{name: "Yub"}
	fmt.Println(studentYub)

	// khời tạo Object dựa trên Struct
	studentA := Student{}
	studentA.name = "Evan"
	studentA.age = 12
	studentA.mssv = 1
	studentA.gender = true

	fmt.Println(studentA) // {3 Evan true [Toan Vat Ly Sinh hoc]}

	// truy xuất value trong Object (giống cách trong JS)
	fmt.Println(studentA.name) // Evan

	// khi ta khởi tạo 1 Object rỗng -> cần thêm thuộc tính
	//studentA.name = "EvanTran"

	// Copy
	// biến copy từ 1 struct sẽ tự tạo ra 1 vùng nhớ mới, không ảnh hưởng struct được copy.

}
