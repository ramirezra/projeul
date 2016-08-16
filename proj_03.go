package main  

import "fmt"


func main() {
	value := 600851475143
	
	for i := 1; i<value; i++ {
		if value % i == 0 {
			counter := 0
			for j := 1; j<i; j++ {
				if i % j == 0 {
					counter++
				}
			}
			if counter < 2 {
				fmt.Printf("%d - True\n", i)
			}
		}
	} 
}