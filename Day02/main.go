package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	safeCount := 0

	for _, line := range lines {
		levels := parseLevels(line)
		if isSafe(levels) {
			safeCount++
		}
	}

	fmt.Println("Number of safe reports:", safeCount)
}

func parseLevels(line string) []int {
	var levels []int
	for _, s := range strings.Fields(line) {
		level, _ := strconv.Atoi(s)
		levels = append(levels, level)
	}
	return levels
}

func isSafe(levels []int) bool {
	if len(levels) < 2 {
		return false
	}

	increasing := levels[1] > levels[0]
	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]
		if diff == 0 || diff < -3 || diff > 3 {
			return false
		}
		if (diff > 0) != increasing {
			return false
		}
	}
	return true
}

func task2(lines []string) {
	safeCount := 0

	for _, line := range lines {
		levels := parseLevels(line)
		if isSafe(levels) || canBeMadeSafe(levels) {
			safeCount++
		}
	}

	fmt.Println("Number of safe reports with Problem Dampener:", safeCount)
}

func canBeMadeSafe(levels []int) bool {
	for i := range levels {
		modifiedLevels := append([]int{}, levels[:i]...)
		modifiedLevels = append(modifiedLevels, levels[i+1:]...)
		if isSafe(modifiedLevels) {
			return true
		}
	}
	return false
}
