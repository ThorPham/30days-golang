package main

import (
	"fmt"
	"log"
	"net/http"
)

// func handler(w http.ResponseWriter, r *http.Request) {
// 	log.Println("Hello world receive request")
// 	defer log.Println("End of hello request")
// 	fmt.Fprint(w, "hello world")
// }

type HandleViaStruct struct{}

func (*HandleViaStruct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello world receive request")
	defer log.Println("End of hello request")
	fmt.Fprint(w, "hello world")
}
func main() {
	log.Print("Hello world sample started.")

	http.Handle("/", &HandleViaStruct{})
	http.ListenAndServe(":8081", nil)

}
