package day05

import (
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
	seeds, allRangeLists := parseSeedsAndAllRangeLists(input)

	for _, rangeList := range allRangeLists {
		destinations := []int{}
		for _, s := range seeds {
			destinations = append(destinations, calculateDestination(s, rangeList))
		}
		seeds = destinations
	}

	sort.Ints(seeds)
	lowest := seeds[0]
	return strconv.Itoa(lowest)
}

func (Puzzler) Part2(input []string) string {
	return "Part2 not yet implemented."
}

func parseSeedsAndAllRangeLists(input []string) ([]int, []RangeList) {
	seeds := parseSeeds(input[0])

	allRangeLists := []RangeList{}
	tmp := RangeList{}

	for _, line := range input[2:] {
		if strings.Contains(line, ":") {
			continue
		}

		if len(line) == 0 {
			sort.Slice(tmp, tmp.sortBySource)
			allRangeLists = append(allRangeLists, tmp)
			tmp = RangeList{}
		} else {
			tmp = append(tmp, buildRange(line))
		}
	}

	sort.Slice(tmp, tmp.sortBySource)
	allRangeLists = append(allRangeLists, tmp)

	return seeds, allRangeLists
}

func parseSeeds(line string) []int {
	seeds := []int{}
	splitByHeader := strings.Split(line, ":")
	fields := strings.Fields(splitByHeader[1])
	for _, f := range fields {
		seed, _ := strconv.Atoi(f)
		seeds = append(seeds, seed)
	}
	return seeds
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
