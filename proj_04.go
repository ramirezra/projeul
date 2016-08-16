/*
	A palindromic number reads the same both ways.
	The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

	Find the largest palindrome made from the product of two 3-digit numbers.
*/

package main

import (
	"fmt"
	"strconv"
)

func reverse(value string) string {
	data := []rune(value)
	result := []rune{}

	for i := len(data) - 1; i >= 0; i-- {
		result = append(result, data[i])
	}
	return string(result)
}

func ispallendrome(s string) bool {
	pallen := s[0 : len(s)/2]
	drome := reverse(s)[0 : len(s)/2]
	result := (pallen == drome)
	return bool(result)
}

func main() {
	lower := 100
	upper := 999
	value := 0

	for i := lower; i <= upper; i++ {
		for j := lower; j <= upper; j++ {
			value = i * j

			if ispallendrome(strconv.Itoa(value)) {
				fmt.Printf("%d - %t\n", value, ispallendrome(strconv.Itoa(value)))
			}
		}
	}
}
