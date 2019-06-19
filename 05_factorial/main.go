package main

import "fmt"

func main() {
	var x int
	fmt.Print("enter any no.")
	fmt.Scan(&x)
	fmt.Print(factorial(x))
}
func factorial(y int) int {
	if y == 0 {
		return 1
	}
	return y * factorial(y-1)
}
