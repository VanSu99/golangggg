package main

import "fmt"

func slice() {
	// Slice

	/*
		- Có size động -> dùng nhiều hơn Array.
		- Không cần thêm size -> nếu thêm size vào thì coi như là 1 array.
		- Slice là 1 reference type.
	*/

	// Khai báo Slice - cách 1
	var mySlice []int
	fmt.Println(mySlice)

	// khai báo - cách 2
	mySlice1 := []int{1, 2, 3, 4}
	fmt.Println(mySlice1)

	// khai báo - cách 3 từ 1 array cho trước
	// syntax: sliceName := arrName[start:end-1]
	/*
		- start: vị trí bắt đầu từ
		- end: vị trí kết thúc
	*/

	var myArr = [5]int{1, 2, 3, 4, 5}
	mySlice2 := myArr[1:3] // giá trị của mySlice2 này bắt đầu từ myArr[1] đến myArr[3-1=2]
	fmt.Println(mySlice2)

	mySlice3 := myArr[:] // [:] : lấy tất cả value có trong array
	fmt.Println(mySlice3)

	mySlice4 := myArr[2:] // [2:] : lấy từ vị trí thứ 2 cho đến hết
	fmt.Println(mySlice4) // [3,4,5]

	// tạo 1 Slice từ 1 Slice khác
	var mySlice5 = []int{1, 2, 3, 4, 5, 6}
	mySlice6 := mySlice5
	fmt.Println(mySlice6)

	mySlice7 := mySlice5[1:]
	fmt.Println(mySlice7) // [2,3,4,5,6]

	// Reference type
	myArr2 := [5]int{1, 2, 3, 4, 5}
	mySlice10 := myArr2[:]
	fmt.Println(mySlice10)

	// thay đổi giá trị của phần tử có index là 0
	mySlice10[0] = 777
	fmt.Println(mySlice10, "-", myArr2)

	// length - capacity Slice
	/*
		- length: số lượng phần tử trong Slice
		- capacity (sức chứa): số lượng phần tử của underlying array bắt đầu từ vị trí start khi Slice được init
	*/
	countries := [...]string{"VietNam", "Myamar", "UK", "CANADA"}
	mySlice11 := countries[2:3]

	fmt.Println(len(mySlice11)) // 1 <== 3 - 2 = 1
	fmt.Println(cap(mySlice11)) // 2 <== 4 - 2 = 2

	// tạo slice bằng keyword: make()
	mySlice12 := make([]int, 10)
	// length của mySlice12 là 10 items
	/*
		make(array, length, capacity)
	*/
	fmt.Println("Slice with make: ", mySlice12)

	// Append - thêm phần tử vào Slice
	/*
		- return về 1 Slice mới (có size mới)
		- syntax : append(name_slice, value_add)
	*/
	mySlice8 := []string{"BMW", "Mercedes", "Honda Civic"}
	mySlice9 := append(mySlice8, "Pagani")
	fmt.Println("Cars slice: ", mySlice9)

	// dùng append để thêm 1 array vào Slice -> dùng kèm với dấu ...
	mySlice14 := append(mySlice8, []string{"A", "B", "C"}...)
	fmt.Println("Cars slice with array: ", mySlice14)

	/*
		- Có thể dùng Slice để thao tác các CTDL: Stack và Queue
		- Thực hiện Pop, Push thông qua syntax: [start:end-1]
		- Yêu cầu hiểu về Stack và Queue
	*/
	// Stack - LIFO
	stackA := []int{1, 2, 3, 4, 5, 6}
	stackB := stackA[:5] // bóc item 6 ra trước
	fmt.Println("Stack LIFO: ", stackB)

	// Queue - FIFO
	stackC := stackA[1:] // bóc item 1 ra trước
	fmt.Println("Queue FIFO: ", stackC)

	// Thao tác với item ở giữa
	stackD := []int{1, 2, 3, 4, 5, 6}
	stackE := append(stackD[:2], stackD[3:]...)
	fmt.Println(stackE) // [1,2,4,5,6]

}
