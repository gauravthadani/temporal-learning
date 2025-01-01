package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	fmt.Println(get())
}

func get() string {
	r, err := http.Get("http://localhost:3333/translate?name=Gaurav")
	if err != nil {
		log.Fatalln("Failed to get a response")
	}
	defer r.Body.Close()
	result, _ := io.ReadAll(r.Body)
	return string(result)
}
