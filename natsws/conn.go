package natsws

import (
	"bytes"
	"crypto/tls"
	"flag"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

type Conn struct {
	Schema string
	Host   string
}

type callback func(body []byte)

func (ws *Conn) Publish(subject string, payload []byte) {
	url := "http://" + ws.Host + "/" + subject
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		println("ERROR")
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Do(req)
	if err != nil {
		println(err.Error())
		return
	}
	println(resp.Status)
}

func (ws *Conn) Subscribe(subject string, fn callback) {
	addr := flag.String("addr", ws.Host, "http service address")
	u := url.URL{Scheme: "ws", Host: *addr, Path: "/supu"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	go func() {
		defer c.Close()
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
			}
			fn(message)
		}
	}()

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for t := range ticker.C {
		err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
