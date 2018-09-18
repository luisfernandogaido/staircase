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

var solutions = make([][]int, 0)

func numWays(n int, x []int) [][]int {
	validatePositives(n, x)
	x = removeRepetead(x)
	up(n, x, 0, []int{})
	return solutions
}

func up(n int, x []int, level int, made []int) {
	for _, steps := range x {
		if level+steps == n {
			solutions = append(solutions, append(made, steps))
		} else if level+steps < n {
			up(n, x, level+steps, append(made, steps))
		}
	}
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

func faca() {
	numeros := []int{0, 1, 3, 5}
	mapa := make(map[string]bool)
	for _, numero1 := range numeros {
		for _, numero2 := range numeros {
			for _, numero3 := range numeros {
				for _, numero4 := range numeros {
					for _, numero5 := range numeros {
						for _, numero6 := range numeros {
							for _, numero7 := range numeros {
								for _, numero8 := range numeros {
									for _, numero9 := range numeros {
										for _, numero10 := range numeros {
											if numero1+numero2+numero3+numero4+numero5+
												numero6+numero7+numero8+numero9+numero10 == 10 {
												chave := make([]string, 0)
												if numero1 != 0 {
													chave = append(chave, strconv.Itoa(numero1))
												}
												if numero2 != 0 {
													chave = append(chave, strconv.Itoa(numero2))
												}
												if numero3 != 0 {
													chave = append(chave, strconv.Itoa(numero3))
												}
												if numero4 != 0 {
													chave = append(chave, strconv.Itoa(numero4))
												}
												if numero5 != 0 {
													chave = append(chave, strconv.Itoa(numero5))
												}
												if numero6 != 0 {
													chave = append(chave, strconv.Itoa(numero6))
												}
												if numero7 != 0 {
													chave = append(chave, strconv.Itoa(numero7))
												}
												if numero8 != 0 {
													chave = append(chave, strconv.Itoa(numero8))
												}
												if numero9 != 0 {
													chave = append(chave, strconv.Itoa(numero9))
												}
												if numero10 != 0 {
													chave = append(chave, strconv.Itoa(numero10))
												}
												mapa[strings.Join(chave, " ")] = true
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	for k := range mapa {
		fmt.Println("[" + k + "]")
	}
	fmt.Println(len(mapa))
}
