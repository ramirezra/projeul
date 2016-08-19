/*
 The four adjacent digits in the 1000-digit number that have the greatest product are 9 × 9 × 8 × 9 = 5832.
 Find the thirteen adjacent digits in the 1000-digit number that have the greatest product. What is the value of this product?

 To Do:
 1. Convert big number to slice or map. (Range, strconv?)
 2. Loop through slice 4 at a time, generating product (product = a * b * c *d). Return to a variable.
 3. Compare product to answer variable . If >  answer, store as new answer; otherwise discard.
 4. Print answer after loop is complete.
*/

package main

import (
	"fmt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func numfinder(n rune)int int int int int{
	a := int((bignumber[n-4])-'0')
	b := int((bignumber[n-3])-'0')
	c := int((bignumber[n-2])-'0')
	d := int((bignumber[n-1])-'0')
	e := a * b * c * d
}
func main() {
/*	data, err := ioutil.ReadFile("proj_08.txt")
	check(err)
	fmt.Print(string(data))

	f, err := os.Open("proj_08.txt")
	check(err)

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	var test int64
	b :=[]bytes */
	bignumber := "7316717653133062491922511967442657474235534919493496983520312774506326239578318016984801869478851843858615607891129494954595017379583319528532088055111254069874715852386305071569329096329522744304355766896648950445244523161731856403098711121722383113622298934233803081353362766142828064444866452387493035890729629049156044077239071381051585930796086670172427121883998797908792274921901699720888093776657273330010533678812202354218097512545405947522435258490771167055601360483958644670632441572215539753697817977846174064955149290862569321978468622482839722413756570560574902614079729686524145351004748216637048440319989000889524345065854122758866688116427171479924442928230863465674813919123162824586178664583591245665294765456828489128831426076900422421902267105562632111110937054421750694165896040807198403850962455444362981230987879927244284909188845801561660979191338754992005240636899125607176060588611646710940507754100225698315520005593572972571636269561882670428252483600823257530420752963450"
		fmt.Println(len(bignumber))
		b := 4
		for i, _ := range bignumber[(b-4):b] {
				//fmt.Println(i, int(c)-'0')
				a,b,c,d,e := findnumber(bignumber[i+4])
				fmt.Printf("%d \t %d\t %d\t %d\t %d\t %d\t \n",i,a,b,c,d,e)
			//	z := (int(bignumber[i-4])-'0') * (int(bignumber[i-3])-'0') * (int(bignumber[i-2])-'0') * (int(bignumber[i-1]))
			//	fmt.Println(z)
			}
}
