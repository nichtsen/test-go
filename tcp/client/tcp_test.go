package main

import "testing"

const (
	addr = "127.0.0.1:35766"
)

func TestDial(t *testing.T) {
	Dial(addr)
}
