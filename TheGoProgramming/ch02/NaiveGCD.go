// NaiveGCD

package main

import "fmt"

func main() {
	var a = 4
	var b = 10
	best := 0
	for d := 1; d <= a+b; d++ {
		if ((d % a) == 0) && ((d % b) == 0) {
			best = d
		}
	}
	fmt.Printf("GCD = %d\n", best)
}
