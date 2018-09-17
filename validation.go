package main

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
		panic("n must be positive integer")
	}
	for _, steps := range x {
		if steps <= 0 {
			panic("all steps must be positive integer")
		}
	}
}
