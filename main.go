package main

import (
	"ac23/day01"
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("== sample input ====")
	input := readInput("day01/sample")
	result := day01.Part1(input)
	fmt.Println(result)

	fmt.Println("== puzzle input ====")
	input = readInput("day01/input")
	result = day01.Part1(input)
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
