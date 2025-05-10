package main

import "fmt"

func main() {
	// primative datatype trong golang
	// int : int8 , int16, int32, uint8, uint16 ..
	// float : float32, float64
	// string,
	// bool : true hoặc false
	// Kiểu Byte và Rune
	// byte : Là bí danh của uint8, đại diện cho một ký tự ASCII.
	// rune : Là bí danh của int32, đại diện cho một điểm mã Unicode
	// Kiểu mảng
	var array [3]string = [3]string{"one", "two", "three"}
	fmt.Println(array)
	// Kiểu slice : Slice là một cấu trúc dữ liệu động, biểu diễn một phần của mảng. Slice có thể thay đổi kích thước.
	var slice []string = []string{"one", "two", "three"}
	fmt.Println(slice)
	//Slice có ba thành phần: con trỏ (pointer), độ dài (length), và dung lượng (capacity).
	// Map là một cấu trúc từ điển (key-value store). Nó lưu trữ các cặp khóa-giá trị.
	var mapType map[string]int = map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println(mapType)

	// Kiểu struct
	// khai báo struct
	type Car struct {
		name  string
		age   int
		color string
	}
	// sử dụng struct
	var car = Car{name: "toyota", age: 20, color: "red"}
	fmt.Println(car)
	// Truy xuất từng phần tử
	fmt.Println("name ", car.name, "age ", car.age, "color ", car.color)
	// Truy xuất từng phần tử dùng pointer
	ptr := &car
	fmt.Println("name ", (*ptr).name, "age ", (*ptr).age, "color ", (*ptr).color)

	// Thêm method cho struct
	// 	func (c *Car) printString() string {
	// 		return fmt.Sprintf("Name: %s, Age: %d, Color: %s", c.name, c.age, c.color)
	//   }
	// 	fmt.Println(car.printString())

}
