package main

import (
	"log"
	"net/http"
	"os"
)

var err error

func main() {

	file, err := os.ReadFile("./WebApp/HelloWorld.html")
	if err != nil {
		log.Println(err)
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, err := writer.Write(file)
		if err != nil {
			log.Println(err)
		}
	})

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println(err)
	}

}
