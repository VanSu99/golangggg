package function

import "fmt"

func main() {
	sum := findSum(1, 2, 3, 4, 5)
	fmt.Println("Tổng: ", sum)

	// err là xử lý case khi có lỗi -> case: 5/0
	res, err := calDivide(5, 10)
	fmt.Println("Chia: ", res, err)
}

// Cách khai báo 1 function
// func name_func(arg1 data_type, arg2 data_type) data_type_return { // do something }

/*
- Nếu tham số truyền vào cùng 1 data type thì khai báo như sau:
func name_func(arg1, arg2 data_type) data_type_return {}
*/

/*
- Đối với các data types không phải là Slice và Map, khi ta truyền tham số vào 1 function thì
biến tham số truyền vào sẽ tạo ra 1 ô nhớ mới do đó dù có thay đổi value của biến truyền vào làm
tham số cũng không bị ảnh hưởng.

*/

// dùng toán tử ... để truyền Slice vào func - giống như Rest Parameters của JS
func findSum(a ...int) (sum int) {
	for _, value := range a {
		sum += value
	}
	return
}

func calDivide(a, b int) (int, error) {
	result := 0
	if b == 0 {
		return 0, fmt.Errorf("Cannot divide by zero")
	}
	result = a / b
	return result, nil
}

// Anonymous function - không có tên hàm, khai báo trong 1 func khác

func parent() {
	func() {
		fmt.Println("Hello")
	}()
}
