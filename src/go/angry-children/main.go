// https://www.hackerrank.com/challenges/angry-children
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var n, k int
	_, err := fmt.Scan(&n, &k)
	if err == nil {
		scanner := bufio.NewScanner(os.Stdin)
		var b int
		inputs := make([]int, n)
		for i, _ := range inputs {
			if !scanner.Scan() {
				break
			}
			s := scanner.Text()
			b, _ = strconv.Atoi(s)
			inputs[i] = b
		}
		sort.Ints(inputs)
		answers := make([]int, n-(k-1))
		var j int
		for j = 0; j < n; j++ {
			if j > (n - k) {
				break
			}
			answers[j] = inputs[j+(k-1)] - inputs[j]
			if answers[j] == 0 {
				break
			}
		}
		var fairness int
		for p := 0; p < j; p++ {
			if p == 0 {
				fairness = answers[p]
				continue
			}
			if answers[p] < fairness {
				fairness = answers[p]
			}
			if fairness == 0 {
				break
			}
		}
		fmt.Println(fairness)
	}
}
