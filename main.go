package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
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

func index(rw http.ResponseWriter, r *http.Request) {
	file, err := ioutil.ReadFile("show/index.html")

	errorHTML(err, rw)

	fmt.Fprint(rw, string(file))
}


func main() {
	port := ":8080"
	src := http.FileServer(http.Dir("source"))
	img := http.FileServer(http.Dir("images"))
  show := http.FileServer(http.Dir("show"))

	go fmt.Printf("Listening on port %v", port[1:])

	http.Handle("/source/", http.StripPrefix("/source", src))
	http.Handle("/images/", http.StripPrefix("/images", img))
  http.Handle("/site/", http.StripPrefix("/site", show)) 
	// http.HandleFunc("/index", index)
	http.HandleFunc("/", redirect) 

	err := http.ListenAndServe(port, nil)

	if err != nil {
		fmt.Println(err)
	}
}

//rozumiesz wszystko?
//rozumiem costam

//no powiem tyle ciezko mi zrozumiec, tak jakbym mial cos zrobic w c++ nie wiedzac nic o nim poza tym jak sie nazywa
//chyba nie bede mogl ci pomoc bo tylko bede przeszkadzal. zwlaszcza ze nic nie wiem o golangu
//oh ok
//sory serio, bym pomogl ale to bedzie bol dla mnie w rozumieniu i dla ciebie w tlumaczeniu mi tego/pracowaniu ze mna :/
// :/