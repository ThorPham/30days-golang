package main

import "fmt"

type Student struct {
	name string
	id   string
}

// create method for struct
func (s *Student) updateName(newName string) {
	s.name = newName
}

// interface
type Animal interface {
	walk() string
	message(s string) string
}

type duck struct {
	name string
	foot string
}

// implement interface
func (d *duck) walk() string {
	return d.name + " walk " + d.foot
}

func (d *duck) message(s string) string {
	return "duck say : " + s
}
func main() {
	s1 := Student{name: "thorpham", id: "123"}
	s2 := &Student{name: "thorpham1", id: "1231"}
	fmt.Println(s1)
	fmt.Println(*s2)
	// access field
	fmt.Println("name ", s1.name, "id ", s1.id)
	fmt.Println("name ", (*s2).name, "id ", (*s2).id)
	s1.updateName("dungpham")
	fmt.Println(s1)
	// anonymous struct
	Elemts := struct {
		name string
		age  int
	}{name: "dung", age: 10}
	fmt.Println(Elemts)

	// anonymous field
	type anonymousField struct {
		int
		string
	}
	a1 := anonymousField{1, "one"}
	fmt.Println(a1.int, a1.string)

	// use interface
	var dInterface Animal
	dInterface = &duck{name: "donal", foot: "two feet"}
	fmt.Printf(dInterface.message("hello everyone"))

}
