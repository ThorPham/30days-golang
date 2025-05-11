package main

import (
	"container/list"
	"fmt"
)

func main() {
	ls := list.List{}
	ls.PushBack(1)
	ls.PushBack(2)
	ls.PushBack(3)
	for c := ls.Front(); c != nil; c = c.Next() {
		fmt.Println(c.Value)
	}
}
