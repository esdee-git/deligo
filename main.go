package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Page struct {
	Title string
}

var templates = template.Must(template.ParseFiles("header.html"))

func RootHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		if error := request.ParseForm(); error == nil {
			url := request.PostForm.Get("url")
			tags := request.PostForm.Get("tags")
			description := request.PostForm.Get("description")
			addLinkResult := deliGo(url, tags, description)
			buffer := bytes.NewBufferString(addLinkResult)
			response.Write(buffer.Bytes())
			log.Println(addLinkResult)
		}
	} else {
		response.Header().Set("Content-type", "text/html")
		err := request.ParseForm()
		if err != nil {
			http.Error(response, fmt.Sprintf("error parsing url %v", err), 500)
		}
		templates.ExecuteTemplate(response, "header.html", Page{Title: "Home"})
	}
}

func main() {
	var host = flag.String("host", "0.0.0.0", "IP of host to run webserver on")//must be at least 0.0.0.0 because 127.0.0.1 doesn't allow remote connections! http://stackoverflow.com/a/32404451
	var port = flag.Int("port", 8080, "Port to run webserver on")
	var staticPath = flag.String("staticPath", "dist/", "Path to static files")

	flag.Parse()
	readDeliciousData()

	router := mux.NewRouter()
	router.HandleFunc("/", RootHandler)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(*staticPath))))

	addr := fmt.Sprintf("%s:%d", *host, *port)
	log.Printf("Listening on %s", addr)

	err := http.ListenAndServe(addr, router)
	if err != nil {
		log.Fatal("ListenAndServe error: ", err)
	}
}
