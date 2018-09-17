package main

import (
	"fmt"
)

func main() {

	n := 5
	jumpsAllowed := []int{1, 3, 5}

	//numWays is the original solution od the problem, wich returns only the number of ways to reach the top.
	numSolutions := numWays(n, jumpsAllowed)
	fmt.Printf("number of solutions: %v\n", numSolutions)

	//numWaysExtended shows all de paths too.
	solutions := numWaysExtended(n, jumpsAllowed)
	fmt.Println("solutions:")
	for _, solution := range solutions {
		fmt.Println(solution)
	}
}

func numWays(n int, x []int) int {
	return up(n, x, 0)
}

func up(n int, x []int, level int) int {
	solutions := 0
	for _, steps := range x {
		if level+steps == n {
			solutions++
		} else if level+steps < n {
			solutions += up(n, x, level+steps)
		}
	}
	return solutions
}
