package tcp

import (
	"fmt"
	"time"

	"github.com/topfreegames/pitaya/acceptor"
)

func Serve(addr string, cert ...string) {
	tpcL := acceptor.NewTCPAcceptor(addr, cert...)
	defer tpcL.Stop()
	go tpcL.ListenAndServe()
	conn := tpcL.GetConnChan()
	select {
	case pchan := <-conn:
		b, err := pchan.GetNextMessage()
		if err != nil {
			fmt.Println("error: ", err)
			break
		}
		fmt.Println("get msg: ", string(b))
	case <-time.After(time.Second * 100):
		fmt.Println("timeout, exit")
	}
}
