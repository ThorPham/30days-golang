package main

import "fmt"

// the * and &
// New keywork
type Employee struct {
	name   string
	empid  int
	salary int
}

func main() {
	p := 10
	ptr := &p
	fmt.Println("p value ", p)
	fmt.Println("p value by pointer ", *ptr)
	fmt.Println("p address ", &p)
	fmt.Println("ptr value by pointer ", &ptr)
	fmt.Println("ptr hold address of p ", ptr)
	emp1 := Employee{name: "thorpham", empid: 10001, salary: 10000}
	fmt.Println(emp1)
	emp2 := new(Employee)
	emp2.empid = 200
	emp2.name = "dungpham"
	emp2.salary = 100000
	fmt.Println(*emp2)
}
