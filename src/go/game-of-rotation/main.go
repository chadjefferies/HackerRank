// https://www.hackerrank.com/challenges/game-of-rotation
package main

import (
	"bufio"
	"fmt"
	"os"
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

		pmean := 0
		for k := 0; k < n; k++ {
			p := 0
			for i, v := range inputs {
				fmt.Println(p)
				p = p + ((i + 1) * v)
			}
			if p > pmean {
				pmean = p
				//fmt.Println(pmean)
				inputs = append(inputs, inputs[0])
				inputs = append(inputs[:0], inputs[1:]...)
				fmt.Println(inputs)
				continue
			}

			inputs = append(inputs, inputs[0])
			inputs = append(inputs[:0], inputs[1:]...)
		}

		fmt.Println(pmean)
	}
}

/*
// https://www.hackerrank.com/challenges/game-of-rotation
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err == nil {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Split(bufio.ScanWords)
		inputs := make([]int64, n)

		for j := 0; j < n; j++ {
			if scanner.Scan() {
				b, _ := strconv.ParseInt(scanner.Text(), 10, 64)
				inputs[j] = b
			}
		}

		//		maxIdx := n - 1
		startingIndex := 1
		//		var maxDiff int64 = 0
		//		for k := 0; k < n; k++ {
		//			d := (inputs[maxIdx]) - (inputs[k])
		//			if d > maxDiff {
		//				maxDiff = d
		//				startingIndex = k
		//			}
		//			if maxIdx == (n - 1) {
		//				maxIdx = 0
		//			} else {
		//				maxIdx++
		//			}
		//		}

		relativeIdx := startingIndex
		var pmean int64 = 0
		for i := 0; i < n; i++ {
			v := inputs[relativeIdx]
			fmt.Printf("%d ", pmean)
			pmean = pmean + (int64(i+1) * v)
			fmt.Printf("%d ", v)

			if relativeIdx == (n - 1) {
				relativeIdx = 0
			} else {
				relativeIdx++
			}
		}

		fmt.Println()
		fmt.Println(pmean)

		//		inputs = append(inputs, inputs[0])
		//		inputs = append(inputs[:0], inputs[1:]...)
	}
}
*/
