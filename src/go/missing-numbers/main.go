// https://www.hackerrank.com/challenges/missing-numbers
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	firstList := make([]int, n)
	for i := 0; i < n; i++ {
		if scanner.Scan() {
			b, _ := strconv.Atoi(scanner.Text())
			firstList[i] = b
		}
	}
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())
	secondList := make([]int, m)
	for i := 0; i < m; i++ {
		if scanner.Scan() {
			b, _ := strconv.Atoi(scanner.Text())
			secondList[i] = b
		}
	}

	sort.Ints(firstList)
	sort.Ints(secondList)
	i1 := 0
	result := make(map[int]bool)
	for _, v := range secondList {
		if (i1 > (n - 1)) || firstList[i1] != v {
			_, exists := result[v]
			if !exists {
				fmt.Printf("%v ", v)
				result[v] = true
			}
		} else {
			i1++
		}
	}
}
