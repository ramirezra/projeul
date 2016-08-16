package main

import "fmt"

func main() {
	sum := 0
	prev := 0
	for i := 1; i < 10; i++ {
		prev = sum
		sum = prev + i
		fmt.Println(sum)

	}
}
