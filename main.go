package main

import (
	"log"
	"net/http"
	"text/template"
)

type tempTitle struct {
	Title string
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	view, _ := template.ParseFiles("./views/index.html")
	data := tempTitle{Title: "ICT - Password Reset"}
	view.Execute(w, data)
}

func prankedHandler(w http.ResponseWriter, req *http.Request) {
	view, _ := template.ParseFiles("./views/pranked.html")
	data := tempTitle{Title: "Get Pranked"}
	view.Execute(w, data)
}

func main() {
	// handler
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/phished", prankedHandler)

	// styling
	styles := http.FileServer(http.Dir("./views/styles"))
	http.Handle("/styles/", http.StripPrefix("/styles/", styles))
	// additional assets
	assets := http.FileServer(http.Dir("./views/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", assets))

	log.Println("Listing for requests at http://localhost:80/")
	log.Fatal(http.ListenAndServe(":80", nil))
}
