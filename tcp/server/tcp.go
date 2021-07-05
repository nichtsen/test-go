package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/topfreegames/pitaya/acceptor"
)

func Serve(addr string, cert ...string) {
	tpcL := acceptor.NewTCPAcceptor(addr, cert...)
	fmt.Printf("tpcl")
	defer tpcL.Stop()
	go tpcL.ListenAndServe()
	conn := tpcL.GetConnChan()

	// Gracefully shutdown server
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
forLoop:
	for {
		select {
		case pchan := <-conn:
			b, err := pchan.GetNextMessage()
			if err != nil {
				fmt.Println("error: ", err)
				break
			}
			fmt.Println("get msg: ", string(b))
		case <-done:
			fmt.Println("try to shutdown server ...")
			tpcL.Stop()
			fmt.Println("successfully shutdown")
			break forLoop
		}
	}
}

func main() {
	const (
		addr = "127.0.0.1:35766"
	)
	Serve(addr, "server.crt", "server.key")
}
