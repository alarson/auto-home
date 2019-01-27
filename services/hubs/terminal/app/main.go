package main

import (
	"net/http"
	"log"
	"io/ioutil"
)

func main() {
	MakeRequest()
}

func MakeRequest() {
	resp, err := http.Get("http://localhost:8080/")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
}
