package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Rsvp struct{
	Name, Email, Phone string 
	WillAttend bool
}

var responses = make([]*Rsvp, 0, 10)
var templates = make(map[string]*template.Template, 3)

func loadTemplates(){
	templateNames := [5]string {"welcome", "form", "thanks", "sorry", "list"}
	for index, name := range templateNames {
		t, err := template.ParseFiles("layout.html", name + ".html")
		if (err == nil) {
			templates[name] = t 
			fmt.Println("Loaded template", index, name)
		} else {
			fmt.Println("Index = ", index, " Name = ", name)
			fmt.Println("Error Happened")
			panic(err)
		}
	}

}

func welcomeHandler(writer http.ResponseWriter, request *http.Request){
	templates["welcome"].Execute(writer.null)
}

func main() {
	fmt.Println("TODO: add some features")
	loadTemplates()
	http.HandleFunc("/", welcomeHandler)
}