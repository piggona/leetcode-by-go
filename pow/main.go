package main

import "fmt"

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	half := myPow(x, n/2)
	if n%2 == 0 {
		return half * half
	}
	if n%2 > 0 {
		return half * half * x
	}
	return half * half / x
}

func myPowIterate(x float64, n int) float64 {
	// if n < 0 {
	// }
	return 0.0
}

func main() {
	x := 2.00000
	n := -2
	fmt.Println(myPow(x, n))
}
