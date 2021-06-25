package tcp

import (
	"context"
	"crypto/tls"
	"log"
	"net"
	"time"
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

	conn, err := dialer.DialContext(ctx, "tcp", addr)
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()
	if _, err := conn.Write([]byte("client test")); err != nil {
		log.Println(err)
		return
	}
	log.Println("write to server successfully")
}
