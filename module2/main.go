package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	Handler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, fmt.Sprintf("VERSION : %s \n", os.Getenv("VERSION")))
		for k, v := range req.Header {
			io.WriteString(w, fmt.Sprintf("%s : %s \n", k, v))
		}
		fmt.Println(req.RemoteAddr)
	}

	healthzHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "200")
	}

	http.HandleFunc("/", Handler)
	http.HandleFunc("/healthz", healthzHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
