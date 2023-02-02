package main

import . "github.com/nichtsen/test-go/polymorphy"

func main() {
	c := new(Client)
	c.Run("red", "green", "blue")
}
