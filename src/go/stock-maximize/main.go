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
		testCaseMax := make([]int, t)
		for i := 0; i < t; i++ {
			if scanner.Scan() {
				n, _ := strconv.Atoi(scanner.Text())
				testCase := make([]int, n)
				m, mi := 0, 0
				for j := 0; j < n; j++ {
					if scanner.Scan() {
						c, _ := strconv.Atoi(scanner.Text())
						testCase[j] = c
						if c > m {
							m = c
							mi = j
						}
					}
				}
				testCaseMax[i] = mi
				testCases[i] = testCase
			}
		}

		for j, testCase := range testCases {
			var profit int64
			l := len(testCase) - 1

			profit = getProfit(0, testCaseMax[j], testCase)
			s := testCaseMax[j] + 1
			e := l
			for {
				if s > e {
					break
				}
				maxIdx := getMax(s, e, testCase)
				profit = profit + getProfit(s, maxIdx, testCase)
				if maxIdx == l {
					break
				}
				s = maxIdx + 1
			}
			if profit < 0 {
				profit = 0
			}
			fmt.Println(profit)
		}
	}
}

func getProfit(startingIdx int, endingIdx int, testCase []int) int64 {
	sharePrice := testCase[endingIdx]
	cost, sharesBought := 0, 0
	for k := startingIdx; k < endingIdx; k++ {
		cost = cost + testCase[k]
		sharesBought++
	}
	return int64(((sharePrice * sharesBought) - cost))
}

func getMax(startingIdx int, endingIdx int, testCase []int) int {
	m, mi := 0, 0
	for k := startingIdx; k <= endingIdx; k++ {
		v := testCase[k]
		if v > m {
			m = v
			mi = k
		}
	}
	return mi
}
