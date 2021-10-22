package main

import (
	"fmt"
	"io/ioutil"
)

var path = "D:/GoWorkSpace/go-nginx/WebApp/my-app"

func FilePathScanner(localpath string, indexMap map[string]string) {
	dir, err := ioutil.ReadDir(localpath)
	if err != nil {
		fmt.Println(err)
	}
	for _, ele := range dir {
		if ele.IsDir() {
			fmt.Println(localpath + "/" + ele.Name())
			FilePathScanner(localpath+"/"+ele.Name(), indexMap)
		} else {
			routepath := []rune(localpath)[len([]rune(path)):]
			indexMap[localpath+"/"+ele.Name()] = string(routepath) + "/" + ele.Name()
		}
	}
}

func main() {
	routeMap := make(map[string]string)
	FilePathScanner(path, routeMap)
	fmt.Println(routeMap)
}
