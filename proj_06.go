/* The sum of the squares of the first ten natural numbers is,

1^2 + 2^2 + ... + 10^2 = 385
The square of the sum of the first ten natural numbers is,

(1 + 2 + ... + 10)^2 = 55^2 = 3025
Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.

Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
*/

package main

import "fmt"

func sumofquares(a int, b int) int {
	sum := 0
	for i := a; i <= b; i++ {
		sum += (i * i)
	}
	return sum
}

func squareofsums(a int, b int) int {
	total := 0
	for j := a; j <= b; j++ {
		total += j
	}
	return total * total
}

func main() {
	lower := 1
	upper := 100
	fmt.Println(sumofquares(lower, upper))
	fmt.Println(squareofsums(lower, upper))
	fmt.Println((squareofsums(lower, upper) - (sumofquares(lower, upper))))
}
