package sched

import (
	"fmt"
	"runtime"
)

func Routine() {
	go say("bar")
	say("foo")
}

func say(s string) {
	for i := 0; i < 3; i++ {
		a := i + 10
		fmt.Printf("ready to say %s\n", s)
		runtime.Gosched()
		fmt.Println(s, " | ", a)
	}
}
