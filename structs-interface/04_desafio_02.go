package main

import (
	"fmt"
)

func main() {

	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			if i%5 == 0 {
				fmt.Printf("Pin Pan \n")
			} else {
				fmt.Printf("Pin \n")
			}
		} else if i%5 == 0 {
			fmt.Printf("Pan \n")
		} else {
			fmt.Printf("%v \n", i)
		}
	}
}
