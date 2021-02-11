package main

import "fmt"

func mapType() {
	// Map

	/*
		- Map là 1 data type ở dạng Hash Table.
		- Phần tử ở dạng KEY-VALUE.
		- Truy xuất value của item thông qua KEY.
		- KEY phải cùng data type, và ở dạng so sánh. Trong trường hợp vẫn muốn dùng kiểu không so sánh được làm khóa
		thì ta dùng hàm fmt.Sprintf() để chuyển giá trị kiểu này thành chuỗi trước, sau đó dùng chuỗi đó làm khóa.
		- VALUE cũng phải cùng data type.
		- Slice không được dùng làm KEY trong map.
	*/

	/*
		- Khai báo:
			C1: var name_map map[data_types_key]data_types_value
			C2: name_map := map[data_types_key]data_types_value
	*/
	student := map[string]int{
		"John":   19,
		"Evan":   21,
		"Selena": 25,
	}
	fmt.Println(student)

	/*
		- Trong trường hợp, nếu ta chưa biết KEY-VALUE nó như thế nào thì dùng make() để xử lý
	*/

	teacher := make(map[string]int)
	fmt.Println("With make: ", teacher) // map[]

	// Lấy ra các giá trị của KEY và VALUE trong map
	fmt.Println("Truy xuất vào key: ", student["Evan"]) //21

	// Muốn thêm 1 student nữa
	student["Luffy"] = 300

	// Xóa 1 key trong map -> delete(name_map, key_delete)
	delete(student, "Evan")
	fmt.Println("Xóa 1 key: ", student)

	// Check KEY có exist trong map hay không?
	_, isExist := student["Evan"]
	fmt.Println("Check exist of 1 key: ", isExist) // false

	// Check length
	fmt.Println("length: ", len(student))

	// Reference
	// việc copy 1 map ra 1 biến mới là dạng tham chiếu, trỏ cùng đến 1 vùng nhớ

	copyMap := student
	copyMap["Obama"] = 45

	fmt.Println("original map: ", student)
	fmt.Println("copy map: ", copyMap)


}
