package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func HandlerGetConfigs(w http.ResponseWriter, r *http.Request) {
	c, err := os.ReadFile("/etc/httpserver/config.properties")
	if err != nil {
		log.Println("read config failed, err: %w", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if _, err := w.Write(c); err != nil {
		log.Println("write config to client failed, err: %w", err)
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

func HandlerDelayedGreet(w http.ResponseWriter, r *http.Request) {
	log.Printf("waiting for 10 seconds...")
	time.Sleep(time.Second * 10)
	HandlerGreet(w, r)
}

func HandlerHealthz(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("OK\n")); err != nil {
		log.Printf("write response failed, err: %s", err.Error())
	}
	w.WriteHeader(http.StatusOK)
	log.Printf("access log, IP: %s, Path: %s, HTTP Code: %d", r.Host, r.URL.Path, http.StatusOK)
}

func RunServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/get_configs", HandlerGetConfigs)
	mux.HandleFunc("/greet", HandlerGreet)
	mux.HandleFunc("/delayed_greet", HandlerDelayedGreet)
	mux.HandleFunc("/healthz", HandlerHealthz)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	go server.ListenAndServe()
	listenSignal(context.Background(), server)
}

func listenSignal(ctx context.Context, httpSrv *http.Server) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	var s os.Signal
	select {
	case s = <-sigs:
		fmt.Printf("notify sigs %d\n", s)
		if err := httpSrv.Shutdown(ctx); err != nil {
			log.Printf("shutdown failed, err: %s", err.Error())
		}
		fmt.Println("http shutdown")
	}
}

func main() {
	RunServer()
}
