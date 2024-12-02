/*
 * Template for Advent of Code tasks completed in Go
 * https://adventofcode.com/
 */
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Check if the input file is provided
	if len(os.Args) < 2 {
		fmt.Println("Missing input file")
		return
	}

	// Open the input file
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the lines from the file
	lines, err := readLines(file)
	if err != nil {
		fmt.Println("Error reading lines:", err)
		return
	}

	// Run both tasks
	task1(lines)
	task2(lines)
}

func readLines(file *os.File) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func task1(lines []string) {

}

func task2(lines []string) {

}
