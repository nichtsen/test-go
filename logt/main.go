package main

import (
	"github.com/ngaut/log"
	"io"
	"os"
)

func main() {
	f, err := os.OpenFile("test.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	writer := io.MultiWriter(os.Stdout, f)
	log.SetOutput(writer)
	log.SetRotateByHour()
	log.Info("info")
	log.Warn("info")
	log.Error("error")
}
