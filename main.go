package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	err  error
	path = "D:/GoWorkSpace/go-nginx/WebApp/demo-app"
)

func FilePathScanner(localpath string, routeMap map[string]string) {
	dir, err := ioutil.ReadDir(localpath)
	if err != nil {
		fmt.Println(err)
	}
	for _, ele := range dir {
		if ele.IsDir() {
			fmt.Println(localpath + "/" + ele.Name())
			FilePathScanner(localpath+"/"+ele.Name(), routeMap)
		} else {
			routepath := []rune(localpath)[len([]rune(path)):]
			routeMap[localpath+"/"+ele.Name()] = string(routepath) + "/" + ele.Name()
		}
	}
}

func main() {
	routeMap := make(map[string]string)
	FilePathScanner(path, routeMap)

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("D:\\GoWorkSpace\\go-nginx\\WebApp\\my-app")))

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Println(err)
	}

}
