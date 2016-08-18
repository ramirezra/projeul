package main

import (
	"fmt"
	"math"
)

func pythagoras(a int, b int) int {
	csquared := (a * a) + (b * b)
	return csquared

}
func checksum(a int, b int, c int) int {
	return a + b + c
}
func main() {
	upper := 500
	for a := 0; a < upper; a++ {
		for b := 0; b < upper; b++ {
			csquared := pythagoras(a, b)
			c := int(math.Trunc(math.Sqrt(float64(csquared))))
			d := checksum(a, b, c)
			if (a*a)+(b*b) == csquared && d == 1000 {
				fmt.Printf("(%d)^2 + (%d)^2 = %d | %d\n", a, b, c, d)
			}
		}
	}

}

// To Do: Find way to fix floating point issue to return just one response "Natual Number".
