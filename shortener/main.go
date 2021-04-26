package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/zpnk/go-bitly"
)

var tpl *template.Template

func shorturl(res http.ResponseWriter, req *http.Request) {
	fmt.Println("method: ", req.Method)
	if req.Method == "GET" {
		tpl, err := template.ParseFiles("index.html")
		if err != nil {
			log.Fatalln(err)
		}
		tpl.Execute(res, nil)
	} else {
		req.ParseForm()
		fmt.Println("URL je: ", strings.Join(req.Form["longurl"], " "))
		io.WriteString(res, strings.Join(req.Form["longurl"], " "))
		url := strings.Join(req.Form["longurl"], " ")

		token := bitly.New("e7b5b209b33cbe08bf83af5436f40d37d491c734") //there is my generated token from bitly
		short, err := token.Links.Shorten(url)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(short)

	}
}

func main() {
	http.HandleFunc("/index", shorturl)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
