package main

import (
	"ac23/day01"
	"ac23/day02"
	"bufio"
	"fmt"
	"log"
	"os"
)

type Puzzler interface {
	Part1([]string) string
	Part2([]string) string
}

func main() {
	puzzlerByID := map[string]Puzzler{
		"day01": day01.Puzzler{},
		"day02": day02.Puzzler{},
	}

	day := "day02"
	target := "sample"

	puzzler := puzzlerByID[day]

	input := readInput(day, target)

	displayResult(day, "Part1", puzzler.Part1(input))
	// displayResult(day, "Part2", puzzler.Part2(input))
}

func readInput(day, target string) []string {
	filename := fmt.Sprintf("%s/%s", day, target)
	fmt.Printf("Reading %s\n", filename)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func displayResult(day string, part string, result string) {
	fmt.Printf("%s/%s: %s\n", day, part, result)
}
