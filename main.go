package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func errorHTML(err error, rw http.ResponseWriter) {
	if err != nil {
		error, _ := ioutil.ReadFile("show/err.html")
		fmt.Fprint(rw, strings.ReplaceAll(string(error), "{error}", err.Error()))
	}
}

func redirect(rw http.ResponseWriter, r *http.Request) {
	req := 302
	http.Redirect(rw, r, "/site", req)
}

func main() {
	port := ":8080"
	src := http.FileServer(http.Dir("source"))
	img := http.FileServer(http.Dir("images"))
	show := http.FileServer(http.Dir("show"))
	download := http.FileServer(http.Dir("download"))

	go fmt.Printf("Listening on port %v", port[1:])

	http.Handle("/source/", http.StripPrefix("/source", src))
	http.Handle("/images/", http.StripPrefix("/images", img))
	http.Handle("/site/", http.StripPrefix("/site", show))
	http.Handle("/download/", http.StripPrefix("/download", download))
	http.HandleFunc("/", redirect)

	err := http.ListenAndServe(port, nil)

	if err != nil {
		fmt.Println(err)
	}
}
