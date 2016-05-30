// https://www.hackerrank.com/challenges/circle-city
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var t int
	_, err := fmt.Scan(&t)
	if err == nil {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Split(bufio.ScanWords)
		testCases := make([][]int64, t)
		for i := 0; i < t; i++ {
			testCase := make([]int64, 2)
			for j := 0; j < 2; j++ {
				if scanner.Scan() {
					c, _ := strconv.ParseInt(scanner.Text(), 0, 64)
					testCase[j] = int64(c)
				}
			}
			testCases[i] = testCase
		}
		points := make(map[int64]int64)
		for _, testCase := range testCases {
			result := "impossible"
			d := testCase[0]
			k := testCase[1]
			s, exists := points[d]
			if !exists {
				s = 0
				r, _ := math.Modf(math.Sqrt(float64(d)))
				for x := int64(r); x > 0; x-- {
					l, lm := math.Modf(math.Sqrt(float64(d - (x * x))))
					if lm == 0 {
						if l == 0 {
							s += 4
						} else if int64(l) <= x {
							s += 8
						}
					}
				}

				if s <= k {
					points[d] = s
				}
			}
			if k == 0 && s > 0 {
				result = "impossible"
			} else if s == 0 && k == 0 || s <= k {
				result = "possible"
			}
			fmt.Println(result)
		}
	}
}
