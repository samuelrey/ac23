package day12

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Puzzler struct {
}

func (Puzzler) Part1(input []string) string {
	datum := "????.######..#####. 1,6,5"
	springs, damages := parseSpringDamage(datum)
	fmt.Println(springs, damages)

	control := "????"
	count := 1
	var dfs func(string, int) int
	dfs = func(control string, count int) int {
		r := regexp.MustCompile(`\?`)
		matches := r.FindAllStringIndex(control, -1)
		if count == 0 && len(matches) == 0 {
			return 1
		}

		if count == 0 || len(matches) == 0 {
			return 0
		}

		sum := 0
		for _, match := range matches {
			index := match[0]
			tmp := control[:index] + "#" + control[index+1:]
			sum += dfs(tmp, count-1)

			tmp = control[:index] + "." + control[index+1:]
			sum += dfs(tmp, count)
		}

		return sum
	}
	fmt.Println(dfs(control, count))
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
