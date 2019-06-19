package main

import "fmt"

func main() {
	fmt.Println("even nos. between 0 and 100 are :")
	for i := 2; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
