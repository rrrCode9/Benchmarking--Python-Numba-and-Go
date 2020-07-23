// golang 1.14.4

package main

import (
	"fmt"
	"time"
)

func prime(n int) []int {
	n++
	PrimeNumbers := []int{}
	for i := 2; i < n; i++ {

		c := 2
		for {

			if i%c == 0 {

				break
			}

			c++

			if c == i {

				PrimeNumbers = append(PrimeNumbers, i)
				break

			}

		}

	}
	// fmt.Println(PrimeNumbers)
	return PrimeNumbers
}

func main() {

	TestCount := 10

	for i := 0; i < TestCount; i++ {
		start := time.Now()
		prime(100000)

		fmt.Printf("ExCount #%v | Execution time = %vs\n", i+1, time.Now().Sub(start).Seconds())

	}

}
