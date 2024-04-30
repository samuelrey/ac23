package day05

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Puzzler struct {
}

type Range struct {
	source   int
	distance int
	length   int
}

type RangeList []Range

func (rangeList RangeList) sortBySource(i, j int) bool {
	return rangeList[i].source < rangeList[j].source
}

func (Puzzler) Part1(input []string) string {
	input = []string{"50 98 2", "52 50 48"}
	rangeList := RangeList{}
	for _, i := range input {
		rangeList = append(rangeList, buildRange(i))
	}

	sort.Slice(rangeList, rangeList.sortBySource)

	seeds := []int{79, 14, 55, 13}
	destinations := []int{}
	for _, s := range seeds {
		destinations = append(destinations, calculateDestination(s, rangeList))
	}
	fmt.Println(seeds)
	fmt.Println(destinations)

	return "Part1 not yet implemented."
}

func (Puzzler) Part2(input []string) string {
	return "Part2 not yet implemented."
}

func buildRange(line string) Range {
	fields := strings.Fields(line)
	numbers := []int{}

	for _, f := range fields {
		n, _ := strconv.Atoi(f)
		numbers = append(numbers, n)
	}

	return Range{
		source:   numbers[1],
		distance: numbers[0] - numbers[1],
		length:   numbers[2],
	}
}

func calculateDestination(seed int, rangeList RangeList) int {
	for _, rnge := range rangeList {
		if seed < rnge.source {
			return seed
		}

		if seed >= rnge.source && seed <= rnge.source+rnge.length-1 {
			return seed + rnge.distance
		}
	}

	return seed
}
