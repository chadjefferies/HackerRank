// https://www.hackerrank.com/challenges/circle-city
package main

import (
	"fmt"
	"math"
)

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var d, k, pts float64
		fmt.Scan(&d, &k)
		r := math.Floor(math.Sqrt(d))
		for x := r; x > 0; x-- {
			y, f := math.Modf(math.Sqrt(d - (x * x)))
			if f == 0 {
				if y == 0 {
					pts += 4 // only 4 possible values (x,0) (0,x) (-x,0) (0,-x)
				} else if y <= x {
					pts += 8
				}
			}
		}
		if pts <= k {
			fmt.Println("possible")
		} else {
			fmt.Println("impossible")
		}
	}
}
