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

			x, result := 0, "NO"
			if c < a {
				for ok := true; ok; {
					z := swapJugs(x, a, b)
					if z == 0 {
						break
					}
					ok = c%z != 0
					if !ok {
						result = "YES"
						break
					}
					x = z
				}
			}
			fmt.Println(result)
		}
	}
}

func swapJugs(x int, a int, b int) int {
	// take remaining in a and dump in b
	br := b - x
	r := a
	for r >= b {
		r = r - br
		// dump out b
		br = b
	}
	return r
}
