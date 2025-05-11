package main

import (
	"fmt"
	"os"
)

func main2() {
	// file, err := os.Open("input.txt")
	// if err != nil {
	// 	return
	// }
	// defer file.Close()
	// stat, err := file.Stat()
	// if err != nil {
	// 	return
	// }
	// bs := make([]byte, stat.Size())
	// _, err = file.Read(bs)
	// if err != nil {
	// 	return
	// }
	// str := string(bs)
	// fmt.Printf(str)
	file, err := os.ReadFile("input.txt")
	if err != nil {
		return
	}
	str := string(file)
	fmt.Println(str)
	//write file
	text := []byte("hello\ngo\n")
	writeFile := os.WriteFile("output.txt", text, 677)
	if writeFile != nil {
		return
	}

}
