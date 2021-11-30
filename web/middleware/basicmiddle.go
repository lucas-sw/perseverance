package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Logger struct {
	hanlder http.Handler
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.hanlder.ServeHTTP(w, r)
	log.Printf("%s %s %v", r.Method, r.URL, time.Since(start))
}

func NewLogger(hanlerToWrap http.Handler) *Logger {
	return &Logger{hanlerToWrap}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func CurTimeHandler(w http.ResponseWriter, r *http.Request) {
	curTime := time.Now().Format(time.Kitchen)
	w.Write([]byte(fmt.Sprintf("Current time is %v", curTime)))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/hello", HelloHandler)
	mux.HandleFunc("/v1/time", CurTimeHandler)
	wrappedMux := NewLogger(mux)
	addr := "localhost:7773"
	log.Printf("listen at:%s", addr)
	log.Fatal(http.ListenAndServe(addr, wrappedMux))

}
