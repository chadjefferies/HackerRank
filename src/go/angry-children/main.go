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
		fairness := inputs[n-1] - inputs[0]
		for j := 0; j < n; j++ {
			if j > (n - k) {
				break
			}
			f := inputs[j+(k-1)] - inputs[j]
			if f < fairness {
				fairness = f
			}
			if fairness == 0 {
				break
			}
		}
		fmt.Println(fairness)
	}
}
