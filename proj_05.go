/*
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/

package main  

import "fmt"

func divcheck(value int) bool {
	counter := 0
	for i := 1; i<= 20; i++ {
		if(value % i == 0){
			counter ++
		}
	}	
	if counter == 20 {
		return true
	} else {
		return false
	}

}

func main() {
	var result bool
	for j :=0; j<=250000000;j++{
		result = divcheck(j)
		if result == true {
			fmt.Println(j)
		}
	}
}