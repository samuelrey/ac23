package day09

import (
	"strconv"
	"strings"
)

type Puzzler struct {
}

func (Puzzler) Part1(input []string) string {
	numbers := make([][]int, len(input))
	for i, line := range input {
		fields := strings.Fields(line)
		asNum := make([]int, len(fields))
		for j, f := range fields {
			n, _ := strconv.Atoi(f)
			asNum[j] = n
		}
		numbers[i] = asNum
	}

	total := 0
	for _, nums := range numbers {
		total = total + foo(nums)
	}

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
