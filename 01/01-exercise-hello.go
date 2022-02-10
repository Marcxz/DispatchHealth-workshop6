package main

import (
	"fmt"
	"time"
)

func main() {
	// Direct call
	
	fun("direct call")

	// TODO: write goroutine with differents variants for function call.
	
	// goroutine function call
	
	// go fun("goroutine-1")

	// goroutine with anonymous function
	
	// go func() {
	// 	fun("goroutine-2")
	// }()
	
	// gorouting with function value call
	
	// fv := fun
	// go fv("goroutine-3")
	
	// wait for goroutines to end 
	// fmt.Println("wait for goroutines...")
	// time.Sleep(100 * time.Millisecond)

	fmt.Println("done")
}


// Single core function
func fun(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(1 * time.Millisecond)
	}
}
