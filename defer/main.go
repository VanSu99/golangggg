package main

import "fmt"

func main() {
	// Defer

	/*
		-	defer là trì hoãn việc thực thi hàm cho tới khi hàm bao ngoài nó return.
		- Mỗi lời gọi defer được push vào stack và thực thi theo thứ tự ngược lại khi hàm bao ngoài nó kết thúc.
		- Sử dụng defer cho việc đóng hoặc giải phóng tài nguyên.

		- Tham khảo: https://zalopay-oss.github.io/go-advanced/ch1-basic/ch1-04-func-method-interface.html
	*/

	fmt.Println("line 1")
	fmt.Println("line 2")
	defer fmt.Println("line 3")
	fmt.Println("line 4")

	/*
		Output: line 1
						line 2
						line 4
						line 3

		-> Sau khi func main kết thúc sẽ thực hiện lệnh phía sau defer tức là nó sẽ pop ra khỏi stack.
	*/

}
