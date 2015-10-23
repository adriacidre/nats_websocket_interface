package main

import (
	"flag"
	"log"

	"./natsws"
)

func main() {
	flag.Parse()
	log.SetFlags(0)

	n := natsws.Conn{
		Schema: "ws",
		Host:   "localhost:8081",
	}
	n.Subscribe("supu", func(body []byte) {
		println(string(body))
	})
}
