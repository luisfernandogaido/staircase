package main

import (
	"fmt"
)

func main() {
	fmt.Println(numWays(50, []int{1, 2, 5}))
}

func numWays(n int, x []int) int {
	return up(n, x, 0)
}

func up(n int, x []int, level int) int {
	solutions := 0
	for _, steps := range x {
		if level+steps == n {
			solutions++
			continue
		}
		if level+steps < n {
			solutions += up(n, x, level+steps)
			continue
		}
	}
	return solutions
}
