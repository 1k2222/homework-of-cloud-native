package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func HandlerGreet(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		w.Header().Add(k, strings.Join(v, ","))
	}
	w.Header().Add(VERSION, os.Getenv(VERSION))
	msg := fmt.Sprintf("Hello, %s\n", r.Host)
	if _, err := w.Write([]byte(msg)); err != nil {
		log.Printf("write response failed, err: %s", err.Error())
	}
	w.WriteHeader(http.StatusOK)
	log.Printf("access log, IP: %s, Path: %s, HTTP Code: %d", r.Host, r.URL.Path, http.StatusOK)
}

func HandlerHealthz(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("OK\n")); err != nil {
		log.Printf("write response failed, err: %s", err.Error())
	}
	w.WriteHeader(http.StatusOK)
	log.Printf("access log, IP: %s, Path: %s, HTTP Code: %d", r.Host, r.URL.Path, http.StatusOK)
}

func RunServer() {
	http.HandleFunc("/greet", HandlerGreet)
	http.HandleFunc("/healthz", HandlerHealthz)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func main() {
	RunServer()
}
