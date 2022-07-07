package main

import (
	"fmt"
)

func factor(n int) {
	for i:=2; i* i <= n; i++ {
		if n % i == 0 {	
			fmt.Printf("%d ", i)
			factor(n/i)
			return 
		}
	}
	fmt.Printf("%d", n)
}

func main() {
	var num int
	fmt.Scanf("%d", &num)
	factor(num)
}
