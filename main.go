package main

import "fmt"


type I interface {
	Read()
	rea()
}

type II interface {
	write()
}

type T struct {}
type B struct {}

func (t *T) Read() {}
func (t *B) Read() {}
func (t *B) rea() {}
func (t *B) write() {fmt.Println("write successfully")}



func main() {
	var b I
	b = &B{}
	c := b.(II)
	c.write()
}
