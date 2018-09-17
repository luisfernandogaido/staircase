package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var (
		n               int
		jumpsAllowedStr string
	)
	flag.IntVar(&n, "n", 4, "Total steps")
	flag.StringVar(&jumpsAllowedStr, "j", "1,3,5", "Jumps allowed separared by comma")
	flag.Parse()
	jumpsAllowed := make([]int, 0)
	jumps := strings.Split(jumpsAllowedStr, ",")
	for _, j := range jumps {
		ja, err := strconv.Atoi(j)
		if err != nil {
			panic("all jumps must be integers")
		}
		jumpsAllowed = append(jumpsAllowed, ja)
	}
	fmt.Println("total steps:", n)
	fmt.Println("jumps allowed:", jumpsAllowed)
	solutions := numWays(n, jumpsAllowed)
	fmt.Println("solutions:", len(solutions))
	for _, solution := range solutions {
		sum := 0
		for _, step := range solution {
			sum += step
		}
		fmt.Printf("%v = %v\n", solution, sum)
	}
}

func numWays(n int, x []int) [][]int {
	validatePositives(n, x)
	x = removeRepetead(x)
	return up(n, x, 0, []int{})
}

func up(n int, x []int, level int, made []int) [][]int {
	solutions := make([][]int, 0)
	for _, steps := range x {
		if level+steps == n {
			solutions = append(solutions, append(made, steps))
		} else if level+steps < n {
			possibleSolutions := up(n, x, level+steps, append(made, steps))
			for _, s := range possibleSolutions {
				if sumSteps(s) == n {
					solutions = append(solutions, s)
				}
			}
		}
	}
	return solutions
}

func sumSteps(steps []int) int {
	sum := 0
	for _, step := range steps {
		sum += step
	}
	return sum
}

func removeRepetead(s []int) []int {
	m := make(map[int]bool)
	uniques := make([]int, 0)
	for _, e := range s {
		if _, ok := m[e]; !ok {
			uniques = append(uniques, e)
			m[e] = true
		}
	}
	return uniques
}

func validatePositives(n int, x []int) {
	if n <= 0 {
		panic("n must be a positive integer")
	}
	for _, steps := range x {
		if steps <= 0 {
			panic("all steps must be positive integers")
		}
	}
}
