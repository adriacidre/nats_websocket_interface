package main

import (
	"./natsws"
)

func main() {
	n := natsws.Conn{
		Schema: "ws",
		Host:   "localhost:8081",
	}
	n.Publish("supu", []byte("kkk"))
}
