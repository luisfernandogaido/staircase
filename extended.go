package main

func numWaysExtended(n int, x []int) [][]int {
	return upExtended(n, x, 0, []int{})
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
