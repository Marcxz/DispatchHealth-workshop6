package main

import "fmt"

func main() {
	// Example 1 from using channel
	ch := make(chan int)
	go func(a, b int) {
		c := a + b
		ch <- c
	}(1, 2)
	r := <- ch
	fmt.Printf("computed value %v \n", r)

	// Example 2 Range channels
	// ch := make(chan int)
	// go func() {
	// 	for i:= 0; i < 6; i++ {
	// 		ch <- i
	// 	}
	// 	close(ch)
	// }()
	// for v := range ch {
	// 	fmt.Println(v)
	// }

	// Example 3 buffered channels with capacity
	// [0,1,2,3,4,5]
	// ch := make(chan int, 6)
	// go func() {
	// 	defer close(ch)

	// 	for i := 0; i < 6; i++ {
	// 		fmt.Printf("Sending : %d\n", i)
	// 		ch <- i
	// 	}
	// }()
	// for v := range ch {
	// 	fmt.Printf("Received: %v \n", v)
	// }
}
