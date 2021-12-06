package main

// Day 1: Sonar Sweep
// Part 1: Count number of times a depth measurement increases from the previous measurement
// Part 1: Answer is total number of measurements larger than the previous measurement
// Part 2: Count number of times the sum of 3 measurements in a sequence increases from the previous sum
// Part 2: Answer is the total number of sums larger than the previous sum

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readLines(path string) ([]int, error) {
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		num, _ := strconv.Atoi(line)

		lines = append(lines, num)
	}
	return lines, scanner.Err()
}

func sum(input []int) int {
	sum := 0

	for i := range input {
		sum += input[i]
	}

	return sum
}

func main() {

	lines, err := readLines("input.txt")
	check(err)

	var count int
	tmp := sum(lines[:3])

	for i := 1; i < len(lines)-2; i++ {
		win := sum(lines[i:(i + 3)])
		// fmt.Println("previous:", tmp, "current:", win)

		if win > tmp {
			count++
		}

		tmp = win
	}
	fmt.Println("Total increased:", count)
}

// Part 1 Code
//
// 	for i := 1; i < len(lines); i++ {
// 	// fmt.Println("Current line is", i, "value: ", lines[i] )
// 	// fmt.Println(lines[i], "-", tmp)
// 		result := lines[i] - tmp
// 			if result > 0 {
// 				count++
// 			}
// 		tmp = lines[i]
// 	}
// 	fmt.Println("Total increased:", count )
