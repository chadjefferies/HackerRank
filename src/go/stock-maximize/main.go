// https://www.hackerrank.com/challenges/stockmax
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
			if scanner.Scan() {
				n, _ := strconv.Atoi(scanner.Text())
				testCase := make([]int, n)
				for j := 0; j < n; j++ {
					if scanner.Scan() {
						c, _ := strconv.Atoi(scanner.Text())
						testCase[j] = c
					}
				}
				testCases[i] = testCase
			}
		}

		for _, testCase := range testCases {
			var profit int64 = 0
			l := len(testCase) - 1
			s, e := 0, l
			for s <= e {
				mi := getMaxInRange(s, e, testCase)
				profit += getProfit(s, mi, testCase)
				if mi == l {
					break
				}
				s = mi + 1
			}
			fmt.Println(profit)
		}
	}
}

func getProfit(start int, end int, testCase []int) int64 {
	sharePrice := testCase[end]
	cost, sharesBought := 0, 0
	for k := start; k <= end; k++ {
		cost += testCase[k]
		sharesBought++
	}
	return int64(((sharePrice * sharesBought) - cost))
}

func getMaxInRange(start int, end int, testCase []int) int {
	m, mi := 0, 0
	for k := start; k <= end; k++ {
		v := testCase[k]
		if v > m {
			m, mi = v, k
		}
	}
	return mi
}
