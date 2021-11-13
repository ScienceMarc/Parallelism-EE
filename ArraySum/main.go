package main

import (
	"fmt"
	"runtime"
	"sync"
)

func multiSum(arr []int) int {
	//Creating a waitgroup, allows us to pause a thread until all goroutines are done.
	var wg sync.WaitGroup
	n := runtime.NumCPU()        //Represents the number of logical threads available
	results := make(chan int, n) //Channels act as buffers which allow us to move data from one goroutine to another

	//Determine how large each subdivision of the array needs to be
	sliceLength := float64(len(arr)) / float64(n)
	for i := 0.0; i < float64(n); i++ { //Do n times
		wg.Add(1) //Add 1 to the waitgroup for each new goroutine spawned
		//Spawn goroutine, similar to the single threaded example
		go func(subArray []int, resultChannel chan int) {
			defer wg.Done()
			sum := 0
			for i := 0; i < len(subArray); i++ {
				sum += subArray[i]
			}
			resultChannel <- sum
		}(arr[int(sliceLength*i):int(sliceLength*(i+1))], results)
	}
	wg.Wait() //Wait for all goroutines to execute wg.Done()

	//perform a normal sum for each of the results of the subdivisions
	result := 0
	for i := 0; i < n; i++ {
		result += <-results
	}
	return result
}

func sum(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}

func main() {
	longArray := make([]int, 1<<30)
	for i := 0; i < 1<<30; i++ {
		longArray[i] = 1
	}
	fmt.Println("allocated array")

	fmt.Println(multiSum(longArray))
}
