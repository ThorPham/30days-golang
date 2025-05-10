package main

import "fmt"

var message = make(chan string)

func main() {
	go createPing()
	msg := <-message
	fmt.Println(msg)
}
func createPing() {
	message <- "PING"
}
