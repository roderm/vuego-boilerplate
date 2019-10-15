package main

import (
	"github.com/GeertJohan/go.rice"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(rice.MustFindBox("static").HTTPBox()))
	http.ListenAndServe(":3000", nil)
}
