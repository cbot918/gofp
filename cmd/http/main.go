package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	port = ":5678"
)

func main() {
	fmt.Println("listening ", port)
	http.HandleFunc("/", handle)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}

func handle(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("read body failed")
		return
	}
	fmt.Println(string(body))
	_, err = io.WriteString(w, string(body))
	if err != nil {
		fmt.Println("write failed")
		return
	}
}
