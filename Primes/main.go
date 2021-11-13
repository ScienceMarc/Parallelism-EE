package main

import (
	"fmt"
	"math"
	"sort"
	"sync"
)

func isPrime(n uint64) bool {
	if n == 0 || n == 1 { //0 and 1 are not prime and are edge cases.
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ { //i = 2 -> sqrt(n)
		if n%uint64(i) == 0 { //If dividing n by i has no remainder, it is not prime.
			return false
		}
	}
	return true
}

func findPrimes(x uint64) []uint64 {
	primes := make([]uint64, 0)
	for i := uint64(0); i < x; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

func findPrimesParallel(x uint64) []uint64 {
	primesChannel := make(chan uint64, x) //Channels can be accessed by multiple threads simultaneously
	var wg sync.WaitGroup
	for i := uint64(0); i < x; i++ {
		wg.Add(1)
		go func(i uint64, wg *sync.WaitGroup) { //Each primality check is done in a seperate goroutine
			defer wg.Done()
			if isPrime(i) {
				primesChannel <- i
			}
		}(i, &wg)
	}
	wg.Wait() //Wait for all numbers to finish.
	close(primesChannel)
	primes := make([]uint64, 0)
	for p := range primesChannel { //Move primes in channel to slice
		primes = append(primes, p)
	}
	sort.Slice(primes, func(i, j int) bool { return primes[i] < primes[j] })
	return primes
}

func main() {
	n := uint64(10)
	hasError := false
	parallelPrimes := findPrimesParallel(n)
	primes := findPrimes(n)
	for !hasError {
		parallelPrimes = findPrimesParallel(n)
		primes = findPrimes(n)

		for i := 0; i < len(parallelPrimes)-1; i++ {
			if parallelPrimes[i] != primes[i] {
				fmt.Println("error")
				hasError = true
			}
		}
	}

	fmt.Println(parallelPrimes)
	fmt.Println(primes)
}
