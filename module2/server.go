package module2

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type Handler func(http.ResponseWriter, *http.Request)

type Middleware func(*http.Request)

func NewHandler(h Handler, middlewares ...Middleware) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, m := range middlewares {
			m(r)
		}
		h(w, r)
	}
}

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
	w.WriteHeader(http.StatusOK)
	log.Printf("access log, IP: %s, Path: %s, HTTP Code: %d", r.Host, r.URL.Path, http.StatusOK)
}

func RunServer() {
	http.HandleFunc("/greet", HandlerGreet)
	http.HandleFunc("/healthz", HandlerHealthz)
	if err := http.ListenAndServe(":80", nil); err != nil {
		panic(err)
	}
}
