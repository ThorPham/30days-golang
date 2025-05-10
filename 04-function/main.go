package main

import "fmt"

// khai bao function : func + ten function + tham so function + return type
// function trong go ko co default parameter
func addby2(a int, b int) int {
	return a + b

}

// function ... is array
func addmore(a ...string) {
	fmt.Println(a)
}

// function using pointer
func add2(a *int, b *int) {
	*a = *a + 2
	*b = *b + 2
}

// function return 2 value
func swap(a int, b int) (int, int) {
	return b, a
}
func main() {

	fmt.Println(addby2(1, 2))
	addmore("1", "2", "3", "4")
	a := 2
	b := 2
	patr := &a
	pbtr := &b
	fmt.Println("a ", a, "b ", b)
	add2(patr, pbtr)
	fmt.Println("a ", a, "b ", b)
	x := 15
	y := 20
	m, l := swap(x, y)
	fmt.Println("m ", m, "l ", l)

}
