package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type tempTitle struct {
	Title string
	User  string
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	view, _ := template.ParseFiles("./views/index.html")
	data := tempTitle{Title: "ICT - Password Reset", User: "None"}
	view.Execute(w, data)
}

func invitationHandler(w http.ResponseWriter, req *http.Request) {
	view, _ := template.ParseFiles("./views/invitation.html")
	data := tempTitle{Title: "Invitation!", User: "None"}
	view.Execute(w, data)
}

func prankedHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	// logic part of log in
	fmt.Println("username:", req.Form["txtUserName"][0])
	//usr := req.Form["txtUserName"]
	view, _ := template.ParseFiles("./views/pranked.html")
	data := tempTitle{Title: "Get Pranked", User: req.Form["txtUserName"][0]}
	view.Execute(w, data)
}

func main() {
	// handler
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/phished", prankedHandler)
	http.HandleFunc("/invitation", invitationHandler)

	// styling
	styles := http.FileServer(http.Dir("./views/styles"))
	http.Handle("/styles/", http.StripPrefix("/styles/", styles))
	// additional assets
	assets := http.FileServer(http.Dir("./views/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", assets))

	log.Println("Listing for requests at http://localhost:80/")
	log.Fatal(http.ListenAndServe(":80", nil))
}
