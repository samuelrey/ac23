package main

import (
	"ac23/day01"
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	p := getPuzzler()
	fmt.Println("== sample input ====")
	input := readInput("day01/sample")
	result := p.Part1(input)
	fmt.Println(result)

	fmt.Println("== puzzle input ====")
	input = readInput("day01/input")
	result = p.Part1(input)
	fmt.Println(result)
}

func readInput(filename string) []string {
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

type puzzler interface {
	Part1([]string) string
	Part2([]string) string
}

func getPuzzler() puzzler {
	return day01.Puzzler{}
}
