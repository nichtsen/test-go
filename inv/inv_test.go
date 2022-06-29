package inv

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
)

func TestRandom(t *testing.T) {
	var tmp chan int
	tmp = make(chan int, 5)
	ws := sync.WaitGroup{}
	for i:=0; i<5; i++ {
		r := rand.Intn(5)
		ws.Add(1)
		go func(arg int, rec chan int) {
			res := arg*2
			rec <- res
			ws.Done()	
		}(r, tmp)
	}
	go receive(tmp)
	ws.Wait()
}

func TestEnd(t *testing.T) {
	var tmp chan int
	tmp = make(chan int, 5)
	for i:=0; i<5; i++ {
		r := rand.Intn(5)
		go func(arg int, rec chan int) {
			res := arg*2
			rec <- res
		}(r, tmp)
	}
	end := make(chan struct{})
	go receiveWithEnd(tmp, end)
	<- end

}

func receiveWithEnd(rec chan int, end chan struct{}) {
	fmt.Println("receive routine")
	cnt := 0
	for {
		select {
		case res :=  <-rec:
			fmt.Printf("%v\n", res)
			cnt++
			if cnt == 5 {
				fmt.Println("receive routine end")
				end <- struct{}{}
				return
			}

		}
	}
} 

func receive(rec chan int) {
	fmt.Println("receive routine")
	cnt := 0
	loop:
	for {
		select {
		case res :=  <-rec:
			fmt.Printf("%v\n", res)
			cnt++
			if cnt == 5 {
				break loop
			}

		}
	}
	fmt.Println("receive routine end")
} 


					


