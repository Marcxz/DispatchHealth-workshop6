package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	numbers := generateRandomNumbers(100000000)
	// Single core way
	t := time.Now()
	sum := Add(numbers)
	fmt.Printf("Sequential Add, Sum: %d, Time Taken: %s\n", sum, time.Since(t))

	// Goroutine way
	t = time.Now()
	sum = AddConcurrent(numbers)
	fmt.Printf("Concurrent, Add, Sum: %d, Time Taken: %s\n", sum, time.Since(t))
}

// Add - sequential code add numbers
func Add(numbers []int) int64 {
	var sum int64
	for _, n := range numbers {
		sum += int64(n)
	}
	return sum
}

// AddConcurrent - concurrent code to add numbers
func AddConcurrent(numbers []int) int64 {

	// Utilize all cores on machine
	numOfCores := runtime.NumCPU()
	runtime.GOMAXPROCS(numOfCores)

	var sum int64
	max := len(numbers)

	sizeOfParts := max / numOfCores
	var wg sync.WaitGroup

	for i := 0; i < numOfCores; i++ {
		// Divide the input into parts
		start := i * sizeOfParts
		end := start + sizeOfParts
		part := numbers[start:end]

		// Run computation for each part in separate goroutine
		wg.Add(1)
		go func(nums []int) {
			defer wg.Done()

			var partSum int64

			// Calculate sum for each part
			for _, n := range nums {
				partSum += int64(n)
			}

			// Add sum of each part to cummulative sum
			atomic.AddInt64(&sum, partSum)
		}(part)
	}
	wg.Wait()
	return sum
}

func generateRandomNumbers(length int) []int {
	arr := make([]int, length)
	for i := 0; i < length; i++ {
		arr[i] = rand.Intn(10000000000)
	}
	return arr
}