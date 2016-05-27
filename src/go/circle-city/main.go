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

		for _, testCase := range testCases {
			d := testCase[0]
			k := testCase[1]
			fmt.Printf("d:%v k:%v\n", d, k)
			fmt.Println(math.Sqrt(float64(d)))
		}
	}
}
