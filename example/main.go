package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	"time"
)

var HomeHandler = func(w http.ResponseWriter,r *http.Request){
	fmt.Println("test")
	w.Write([]byte("test"))
}

var HomeHandler2 = func(w http.ResponseWriter,r *http.Request){
	fmt.Println("test2")
	w.Write([]byte("test2"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/home", HomeHandler)
	http.Handle("/", r)
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	srv.ListenAndServe()
}
