package main

import (
	"os"
	"runtime"
	"strconv"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	strconv.ParseInt(os.Args[1], 10, 32)

}

func maker(max int) {
	numWorkers := runtime.NumCPU()
	perfect := make(chan int, numWorkers)
	numbers := make(chan int, max)
	done := make(chan struct{}, numWorkers)
	for i := 2; i <= max; i++ {
		numbers <- i
	}
	close(numbers)
	for i := 0; i < numWorkers; i++ {
		go worker(numbers, perfect, done)
	}

}

func worker(numbers, perfect chan int, done chan struct{}) {
	for num := range numbers {

		sum := 0
		for i := 0; i <= num; i++ {

			if (num % i) == 0 {
				sum += i
			}
		}
		if sum == num {
			perfect <- num
		}

	}

}
