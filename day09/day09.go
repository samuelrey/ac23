package day09

import "fmt"

type Puzzler struct {
}

func (Puzzler) Part1(input []string) string {
	x := foo([]int{1, 3, 6, 10, 15, 21})
	fmt.Println(x)
	return "Part1 not yet implemented."
}

func (Puzzler) Part2(input []string) string {
	return "Part2 not yet implemented."
}

func foo(measurements []int) int {
	if checkAllZero(measurements) {
		return 0
	}
	n := len(measurements)

	differences := make([]int, n-1)
	for i := n - 1; i > 0; i-- {
		j := i - 1

		differences[j] = measurements[i] - measurements[j]
	}

	x := foo(differences)

	return measurements[n-1] + x
}

func checkAllZero(measurements []int) bool {
	for _, m := range measurements {
		if m != 0 {
			return false
		}
	}

	return true
}
