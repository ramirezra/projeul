package main

import "fmt"

func main() {
	first := 0
	second := 1
	sum := 0
	total := 0

	for sum < 4000000 {
		sum = first + second
		first = second
		second = sum

		if sum % 2 == 0 {
			total += sum
			fmt.Println(total)
		}
	}
}