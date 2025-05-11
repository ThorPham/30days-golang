package main

import (
	"fmt"
	"strings"
)

func main1() {
	// check string contain sub-string
	fmt.Println(strings.Contains("golang", "go"))
	// count sub-string in string
	fmt.Println(strings.Count("golang", "g"))
	// check prefix and suffix
	fmt.Println(strings.HasPrefix("golang", "go"))
	fmt.Println(strings.HasSuffix("golang", "g"))
	// get index
	fmt.Println(strings.Index("golang", "go0"))
	//join string
	fmt.Println(strings.Join([]string{"go", "lang"}, "--->"))
	fmt.Println(strings.Replace("golang", "go", "cc", 1))

	arr := []byte("acb")
	fmt.Println(arr)
	str := string([]byte{'a', 'c', 'm', 'n'})
	fmt.Println(str)

}
