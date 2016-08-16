package main  

import "fmt"
import "strconv"

func main() {
	lower := 10
	upper := 99
	total := 0
	digits := ""

	for i := lower; i <= upper; i++ {
		for j := lower; j<= upper; j++ {
			total = i * j
			digits = strconv.Itoa(total)
			fmt.Println(digits)
		}
	}
}