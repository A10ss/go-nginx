package main

import (
	"log"
	"net/http"
)

var (
	err  error
	path = "WebApp/my-app"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(path)))

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Println(err)
	}

}
