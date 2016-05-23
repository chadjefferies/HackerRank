// https://www.hackerrank.com/challenges/die-hard-3
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var t int
	_, err := fmt.Scan(&t)
	if err == nil {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Split(bufio.ScanWords)
		testCases := make([][]int, t)
		for i := 0; i < t; i++ {
			testCase := make([]int, 3)
			for j := 0; j < 3; j++ {
				if scanner.Scan() {
					c, _ := strconv.Atoi(scanner.Text())
					testCase[j] = c
				}
			}
			testCases[i] = testCase
		}

		for _, testCase := range testCases {
			a, b := 0, 0
			if testCase[0] > testCase[1] {
				a = testCase[0]
				b = testCase[1]
			} else {
				a = testCase[1]
				b = testCase[0]
			}
			c := testCase[2]
			result := "NO"

			r, i, k := 1, 1, 0
			for r > 0 && r != c {
				pr := (i * a) + ((-i) * b)
				if pr < r {
					i++
				}
				k++
			}
			fmt.Println(i)
			fmt.Println(k)
			fmt.Println(r)

			//			r := a - b
			//			for r >= b {
			//				r = r - b
			//			}
			//			if r > 0 && c%r == 0 {
			//				result = "YES"
			//			} else if c%(b-r) == 0 {
			//				result = "YES"
			//			}
			fmt.Println(result)
		}
	}
}
