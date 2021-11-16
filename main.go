package main

import (
	"github.com/valyala/fasthttp"
	"log"
)

var (
	err  error
	path = "WebApp/my-app"
)

func main() {

	handler := fasthttp.FSHandler(path, 0)

	err := fasthttp.ListenAndServe(":8081", handler)
	if err != nil {
		log.Println(err)
	}

}
