package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func getCurrentIP(r *http.Request) string {
	ip := r.Header.Get("X-Real-IP")
	if ip == "" {
		// 当请求头不存在即不存在代理时直接获取ip
		ip = strings.Split(r.RemoteAddr, ":")[0]
	}
	return ip

}

func main() {

	Handler := func(w http.ResponseWriter, req *http.Request) {
		os.Setenv("VERSION", "v0.0.1")
		io.WriteString(w, fmt.Sprintf("VERSION : %s \n", os.Getenv("VERSION")))
		for k, v := range req.Header {
			io.WriteString(w, fmt.Sprintf("%s : %s \n", k, v))
		}
		fmt.Println(getCurrentIP)
	}

	healthzHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "200")
	}

	http.HandleFunc("/", Handler)
	http.HandleFunc("/healthz", healthzHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
