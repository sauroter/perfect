package main

import (
	"fmt"
	"log"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	//max, err := strconv.ParseInt(os.Args[1], 10, 32)
	//if err != nil {
	//	log.Fatal(err)
	//}
	max := 200000
	maker(int(max))

}

func maker(max int) {
	numWorkers := runtime.NumCPU()
	perfect := make(chan int, numWorkers)
	numbers := make(chan int, max)
	done := make(chan struct{}, numWorkers)
	for i := 1; i <= max; i++ {
		numbers <- i
	}
	close(numbers)
	for i := 0; i < numWorkers; i++ {
		log.Println("Worker ", i)
		go worker(numbers, perfect, done)
	}
	numDone := 0
	for {
		select {

		case number := <-perfect:
			{
				fmt.Println(number)
			}
		case <-done:
			{
				log.Println("Worker ", numDone, " done")
				numDone += 1
				if numDone == numWorkers {
					return
				}
			}

		}
	}

}

func worker(numbers, perfect chan int, done chan struct{}) {
	for num := range numbers {

		sum := 0
		for i := 1; i < num; i++ {

			if (num % i) == 0 {
				sum = sum + i
			}

		}
		//fmt.Println(num)
		//fmt.Println(sum)
		if sum == num {
			perfect <- num
		}

	}
	done <- struct{}{}
}
