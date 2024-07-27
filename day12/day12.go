package day12

import (
	"fmt"
	"strconv"
	"strings"
)

type Puzzler struct {
}

func (Puzzler) Part1(input []string) string {
	datum := ". 1"
	springs, damages := parseSpringDamage(datum)
	fmt.Println(springs, damages)

	i, j := 0, 0
	for i < len(springs) && j < len(damages) {
		k := damages[j]
		for i < len(springs) && k > 0 {
			
		}
	}
	return "Part1 not yet implemented."
}

func (Puzzler) Part2(input []string) string {
	return "Part2 not yet implemented."
}

func parseSpringDamage(input string) (string, []int) {
	fields := strings.Fields(input)
	springs := fields[0]
	damageStr := fields[1]
	damageVals := strings.Split(damageStr, ",")
	damages := make([]int, len(damageVals))
	for i, d := range damageVals {
		val, _ := strconv.Atoi(d)
		damages[i] = val
	}

	return springs, damages
}
