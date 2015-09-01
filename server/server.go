package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	PORT = ":8081"
	CERT = "cert.pem"
	KEY  = "key.pem"
)

func addUser(w http.ResponseWriter, r *http.Request) {

}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "We're secured!")
}

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/add/user", addUser)
	if err := http.ListenAndServeTLS(PORT, CERT, KEY, nil); err != nil {
		log.Fatal(err)
	}

}
