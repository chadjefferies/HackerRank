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
			profit := 0
			sharesBought := 0
			for k := 0; k < len(testCase); k++ {
				price := testCase[k]
				sellPrice := price * sharesBought
				if sellPrice == 0 {
					sellPrice = price
				}
				fmt.Printf("current: %d profit: %d sellPrice: %d\n", price, profit, sellPrice)

				if k < len(testCase)-1 {
					// look ahead to next item
					nextPrice := testCase[k+1]

					// what we could sell the next one for if we buy the current
					nextSellPrice := nextPrice * (sharesBought + 2)

					// buy current, sell next
					profitSellNext := (profit - price) + nextSellPrice

					// sell current
					profitSellCurrent := profit + sellPrice

					fmt.Printf("next price: %d next sell price: %d profit from selling next: %d profit from selling current: %d\n", nextPrice, nextSellPrice, profitSellNext, profitSellCurrent)

					if profitSellCurrent > profitSellNext {
						profit = profitSellCurrent
						sharesBought++
					}
				} else {
					// we are at the end of the list

					// sell current
					profit = profit + sellPrice
				}
			}
			fmt.Println(profit)
		}
	}
}

func buyOne(profit int, price int) int {
	return profit - price
}

func sellShares(profit int, numShares int, price int) int {
	return profit + (price * numShares)
}
