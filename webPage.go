package main

type webPage struct {
	Uri      string
	HttpBody string
}

func (this *webPage) NewWebPage(uri string, httpBody string) {
	this.Uri = uri
	this.HttpBody = httpBody
}
