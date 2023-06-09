package main

import (
	"fmt"
	"github.com/nitishfy/go-server/pkg/handler"
	"log"
	"net/http"
)

const portNum = ":8082"

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/hello", handler.Hello)
	http.HandleFunc("/form", handler.FormHandler)
	fmt.Println("Listening on port ", portNum)
	if err := http.ListenAndServe(portNum, nil); err != nil {
		log.Fatal(err)
	}
}
