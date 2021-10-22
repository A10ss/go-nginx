package main

import "io/ioutil"

type webPage struct {
	Uri      string
	HttpBody string
}

func NewWebPage(uri string, httpBody string) *webPage {
	return &webPage{
		Uri:      uri,
		HttpBody: httpBody,
	}
}

func (this *webPage) ScanUrl() (files []string) {
	dir, err := ioutil.ReadDir("./WebApp")
	if err != nil {
		return nil
	}
	for _, ele := range dir {
		files = append(files, ele.Name())
	}
	return files
}

func main() {

}
