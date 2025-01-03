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
	lang := r.URL.Query().Get("lang")

	if lang == "" {
		lang = "spanish"
	}
	log.Println("Am I coming here")

	log.Println(lang)

	switch lang {
	case "french":
		io.WriteString(rw, fmt.Sprintf("Bon jour! %s", name))

	default:
		io.WriteString(rw, fmt.Sprintf("Hola! %s", name))
	}
}
