package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	var leftList, rightList []int
	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) != 2 {
			fmt.Println("Invalid input format")
			return
		}
		leftNum, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error parsing number:", err)
			return
		}
		rightNum, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("Error parsing number:", err)
			return
		}
		leftList = append(leftList, leftNum)
		rightList = append(rightList, rightNum)
	}

	// Sort both lists
	sort.Ints(leftList)
	sort.Ints(rightList)

	// Calculate the total distance
	totalDistance := 0
	for i := 0; i < len(leftList); i++ {
		distance := abs(leftList[i] - rightList[i])
		totalDistance += distance
	}

	fmt.Println("Total distance:", totalDistance)
}

func task2(lines []string) {
	var leftList, rightList []int
	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) != 2 {
			fmt.Println("Invalid input format")
			return
		}
		leftNum, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error parsing number:", err)
			return
		}
		rightNum, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("Error parsing number:", err)
			return
		}
		leftList = append(leftList, leftNum)
		rightList = append(rightList, rightNum)
	}

	// Create a map to count occurrences in the right list
	rightCount := make(map[int]int)
	for _, num := range rightList {
		rightCount[num]++
	}

	// Calculate the similarity score
	similarityScore := 0
	for _, num := range leftList {
		similarityScore += num * rightCount[num]
	}

	// Print the similarity score
	fmt.Println("Similarity score:", similarityScore)
}

// Helper function to calculate the absolute value
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
