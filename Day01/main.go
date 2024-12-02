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
	task1()
	task2()
}

func task1() {
	// Open the input file
	file, err := os.Open("Input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the input file
	var leftList, rightList []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
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

	// Print the total distance
	fmt.Println("Total distance:", totalDistance)
}

func task2() {
	// Task 2 implementation goes here
}

// Helper function to calculate the absolute value
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
