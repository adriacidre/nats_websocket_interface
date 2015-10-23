package main

import (
	"flag"
	"log"
	"net/http"

	"./natsws"
)

var addr = flag.String("addr", "localhost:8081", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)

	http.HandleFunc("/", natsws.Manage)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
