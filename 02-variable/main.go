package main

import "fmt"

func main() {
	// khai báo biến trong go
	// Go là một statically typed language nên khai báo biến cần biết kiểu
	// khi khai báo kiểu được đặt sau tên biến
	var x int = 10
	// khai báo biến không gán giá trị mặc định là 0
	var y int
	// khai báo tự ngầm định kiểu
	z := 30.0 // tự định kiểu dựa vào kiểu được gán
	// khai báo nhiều biến một lúc : kiểu dữ liệu chỉ cần để 1 cái cuối cùng
	var m, n, l string = "m", "n", "l"
	fmt.Println(x, y, z, m, n, l)
}
