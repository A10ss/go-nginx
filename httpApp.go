package main

import (
	"log"
	"net/http"
)

var (
	thiserr  error
	httpPath = "WebApp/my-app"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(httpPath)))

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	thiserr = server.ListenAndServe()
	if thiserr != nil {
		log.Println(thiserr)
	}

}
