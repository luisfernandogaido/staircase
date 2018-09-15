package main

import (
	"fmt"
)

func main() {
	fmt.Println(numWays(4, []int{1, 2}))
}

func numWays(n int, x []int) int {
	return sobe(n, x, 0)
}

func sobe(n int, x []int, degrau int) int {
	solucoes := 0
	for _, passos := range x {
		if degrau+passos == n {
			solucoes++
			continue
		}
		if degrau+passos < n {
			solucoes += sobe(n, x, degrau+passos)
			continue
		}
	}
	return solucoes
}
