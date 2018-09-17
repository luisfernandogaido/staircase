package main

import (
	"fmt"
)

func main() {
	n := 4
	jumpsAllowed := []int{1, 2}

	//numWays is the original solution of the problem, wich returns only the number of ways to reach the top.
	numSolutions := numWays(n, jumpsAllowed)
	fmt.Printf("number of solutions: %v\n", numSolutions)

	//numWaysExtended shows all the paths too.
	solutions := numWaysExtended(n, jumpsAllowed)
	fmt.Println("solutions:", len(solutions))
	for _, solution := range solutions {
		sum := 0
		for _, step := range solution {
			sum += step
		}
		fmt.Printf("%v: %v\n", solution, sum)
	}
}

func numWays(n int, x []int) int {
	validatePositives(n, x)
	return up(n, removeRepetead(x), 0)
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

func numWaysExtended(n int, x []int) [][]int {
	validatePositives(n, x)
	return upExtended(n, removeRepetead(x), 0, []int{})
}

func upExtended(n int, x []int, level int, made []int) [][]int {
	solutions := make([][]int, 0)
	for _, steps := range x {
		if level+steps == n {
			solutions = append(solutions, append(made, steps))
		} else if level+steps < n {
			solutions = append(solutions, upExtended(n, x, level+steps, append(made, steps))...)
		}
	}
	return solutions
}
