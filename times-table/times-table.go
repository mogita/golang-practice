package main

import (
	"fmt"
)

func main() {
	basicTimesTable()
}

func basicTimesTable() {
	for a := range [9]int{} {
		for b := range [9]int{} {
			if a < b {
				continue
			}
			fmt.Printf("%dx%d=%d ", b+1, a+1, (b+1)*(a+1))
			if b == a {
				fmt.Printf("\n")
			}
		}
	}
}
