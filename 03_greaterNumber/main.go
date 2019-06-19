package main

import "fmt"

func main() {
	var num1 int
	var num2 int
	fmt.Print("Enter first no. :")
	fmt.Scan(&num1)
	fmt.Print("Enter Second no. :")
	fmt.Scan(&num2)
	if num1 > num2 {
		fmt.Println(num1, "is greater")
	} else {
		fmt.Println(num2, "is greater")
	}
}
