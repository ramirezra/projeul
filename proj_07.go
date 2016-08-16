package main

import "fmt"

func isprime(a int) int {
	counter := 0
	for i := 2; i <= a; i++ {
		if a%i == 0 {
			counter++
		}

	}
	return counter
}

func main() {
	counter := 0
	value := 0

	for i := 1; counter <= 10001; i++ {
		if isprime(i) == 1 {
			value = i
			fmt.Printf("%d - %d \n", counter, value)
			counter++
		}
	}

}
