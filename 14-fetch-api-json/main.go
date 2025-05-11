package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Astronaut struct {
	Name  string `json:"name"`
	Craft string `json:"craft"`
}
type ApiResponse struct {
	People  []Astronaut `json:"people"`
	Message string      `json:"message"`
	Number  int         `json:"number"`
}

func main() {
	url := "http://api.open-notify.org/astros.json"
	spaceClient := http.Client{Timeout: time.Second * 2}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "spacecount-tutorial")
	res, err := spaceClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if res.Body != nil {
		defer res.Body.Close()
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	p := ApiResponse{}
	err = json.Unmarshal(body, &p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(p)

}
