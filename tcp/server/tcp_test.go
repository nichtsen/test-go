package main

import (
	"testing"
)

const (
	addr = "127.0.0.1:35766"
)

func TestServe(t *testing.T) {
	Serve(addr)
}

func TestTls(t *testing.T) {
	Serve(addr, "server.crt", "server.key")
}
