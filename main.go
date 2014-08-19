package main

import (
	"fmt"
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
  http.Handle("/", fs)

  log.Println("Listening...")
  var h Hello
  http.ListenAndServe(":3000", h)
}

// func main() {
//   http.Handle("/assets/", http.FileServer(http.Dir("./assets")))
//   if err := http.ListenAndServe(":8080", nil); err != nil {
//     log.Fatal("ListenAndServe: ", err)
//   }
//   var h Hello
//   http.ListenAndServe("localhost:4000", h)
// }
