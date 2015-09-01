package main

import (
	"crypto/tls"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile)

	conf := &tls.Config{
		InsecureSkipVerify: true,
	}
	log.Println("Connecting..")
	conn, err := tls.Dial("tcp", "127.0.0.1:8083", conf)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	log.Println("Sending hi")
	n, err := conn.Write([]byte("hello \n"))
	if err != nil {
		log.Println(n, err)
		return
	}

	buf := make([]byte, 100)
	n, err = conn.Read(buf)
	if err != nil {
		log.Println(n, err)
		return
	}
	log.Println(string(buf[:n]))

}
