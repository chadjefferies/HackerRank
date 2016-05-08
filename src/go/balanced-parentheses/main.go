// https://www.hackerrank.com/challenges/balanced-parentheses
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t int
	_, err := fmt.Scan(&t)
	if err == nil {
		scanner := bufio.NewScanner(os.Stdin)
		testCases := make([]string, t)
		for i := 0; i < t; i++ {
			if !scanner.Scan() {
				break
			}
			testCases[i] = scanner.Text()
		}

		lookup := map[rune]rune{
			'{': '}',
			'[': ']',
			'(': ')',
		}
		var stack []rune
		for _, testCase := range testCases {
			var result string
			for _, r := range testCase {
				if _, exists := lookup[r]; exists {
					stack = append(stack, r)
					result = "NO"
				} else {
					l := len(stack)
					if l == 0 {
						result = "NO"
						break
					}
					var x rune
					stack, x = stack[:l-1], stack[l-1]
					if r != lookup[x] {
						result = "NO"
						break
					} else {
						result = "YES"
					}
				}
			}
			fmt.Println(result)
		}
	}
}
