package main

import (
	"context"
	"crypto/tls"
	"log"
	"net"
	"time"

	"github.com/topfreegames/pitaya/conn/codec"
	"github.com/topfreegames/pitaya/conn/message"
	"github.com/topfreegames/pitaya/conn/packet"
)

func Dial(addr string) {
	var d net.Dialer

	dialer := &tls.Dialer{
		NetDialer: &d,
		Config: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	log.Println("start dialContext")

	conn, err := dialer.DialContext(ctx, "tcp", addr)
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}

	log.Println("end dialContext")
	msg := message.New(true)
	msg.Type = message.Request
	msg.ID = 777
	msg.Data = []byte("client hello")
	msg.Route = "not implemented route"

	encoder := message.NewMessagesEncoder(false)

	cipherMsg, err := encoder.Encode(msg)
	if err != nil {
		log.Fatal(err)
	}
	p := codec.NewPomeloPacketEncoder()
	pck, err := p.Encode(packet.Handshake, cipherMsg)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("codec succ")

	defer conn.Close()
	if _, err := conn.Write(pck); err != nil {
		log.Println(err)
		return
	}
	log.Println("write to server successfully")
}

func main() {
	const (
		addr = "127.0.0.1:35766"
	)
	Dial(addr)
}
