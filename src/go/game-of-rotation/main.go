// https://www.hackerrank.com/challenges/game-of-rotation
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err == nil {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Split(bufio.ScanWords)
		inputs := make([]int, n)
		for j := 0; j < n; j++ {
			if scanner.Scan() {
				b, _ := strconv.Atoi(scanner.Text())
				inputs[j] = b
			}
		}

		sort.Ints(inputs)
		fmt.Println(inputs)
		pmean := 0
		for i, v := range inputs {
			if v < 0 {
				continue
			}
			pmean = pmean + ((i + 1) * v)
		}

		//fmt.Println(maxAt)

		/*
			for k := 0; k < n; k++ {
				//inputs = append(inputs, inputs[0])
				//inputs = append(inputs[:0], inputs[1:]...)
				p := 0
				for i, v := range inputs {
					p = p + ((i + 1) * v)
				}
				if p > pmean {
					pmean = p
					break
				}
			}
		*/
		fmt.Println(pmean)
	}
}
