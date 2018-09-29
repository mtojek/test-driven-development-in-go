package test_driven_development_in_go

import (
	"math"
	"sort"
)

func divisorsOf(n int) []int {
	var divisors []int
	for i := 1; i < int(math.Sqrt(float64(n)) + 1); i++ {
		if n % i == 0 {
			divisors = append(divisors, i)
			if i != n / i {
				divisors = append(divisors, n / i)
			}
		}
	}
	sort.Ints(divisors)
	return divisors
}