package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	port = ":8081"
	cert = "cert.pem"
	key  = "key.pem"
)

// user
func user(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Print("Get User")
	case "POST":
		fmt.Print("POST User")
	case "PUT":
		fmt.Print("PUT User")
	}
	if err := r.ParseForm(); err != nil {
		fmt.Printf("Error parsing add user form:%s", err.Error())
	}
}

// root call `/`
func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "We're secured!")
}

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/user", user)
	if err := http.ListenAndServeTLS(port, cert, key, nil); err != nil {
		log.Fatal(err)
	}

}
