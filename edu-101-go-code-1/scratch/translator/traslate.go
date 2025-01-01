package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/translate", translate)

	err := http.ListenAndServe("0.0.0.0:3333", nil)
	if err != nil {
		log.Fatalln("failed to start server")
	}

}

func translate(rw http.ResponseWriter, r *http.Request) {

	// time.Sleep(time.Second * 6)
	name := r.URL.Query().Get("name")
	io.WriteString(rw, fmt.Sprintf("Hola! %s", name))
}
