package pointer

import "fmt"

func main() {
	// Pointer

	/*
		- khai báo pointer: đặt * trước data type
	*/

	var a int = 12
	var b *int = &a // pointer b trỏ về địa chỉ ô nhớ của a
	fmt.Println(a, b) // 12 12

	// xuất ra value của 1 pointer
	fmt.Println(*b)

	a = 100
	// lúc này a và b cùng ở chung 1 địa chỉ nên a hoặc b thay đổi value thì biến còn lại cũng bị thay đổi theo
	fmt.Println(a, *b) // 24 24

	// pointer với array
	myArr := [4]int{1,2,3,4}
	myPointer1 := myArr
	fmt.Println(myArr, myPointer1)

	myArr[1] = 99
	// do array không tham chiếu
	fmt.Println(myArr, myPointer1) // [1,99,3,4] [1,2,3,4]

	// pointer với Struct
	var structNe *myStruct
	structNe = new(myStruct)

	fmt.Println(structNe)

	// truy xuất vào object đang được trỏ đến
	// Cách 1
	(*structNe).number = 20
	fmt.Println((*structNe).number)

	// Cách 2
	structNe.number = 20
	fmt.Println(structNe.number)

	// pointer với Map, có reference
	myMapNe := map[string]int{"Evan": 12, "Selena":23, "John Doe": 34}
	fmt.Println(myMapNe)
}

type myStruct struct {
	number int
}
