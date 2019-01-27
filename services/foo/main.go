package main

import (
	"log"
	"net/http"

        "common/router/muxrouter"
)

func main() {
        log.Println(routes)
	router := muxrouter.NewRouter(routes)

	log.Fatal(http.ListenAndServe(":8080", router))
}
