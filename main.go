package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Hello struct{}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

func main() {
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/", rootHandler)

	log.Println("Listening...")

	http.ListenAndServe(":3000", nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("views/layout.html.tmpl")
	t.Execute(w, nil)
}

// func main() {
//   http.Handle("/assets/", http.FileServer(http.Dir("./assets")))
//   if err := http.ListenAndServe(":8080", nil); err != nil {
//     log.Fatal("ListenAndServe: ", err)
//   }
//   var h Hello
//   http.ListenAndServe("localhost:4000", h)
// }
