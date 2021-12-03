package main

// Day 1: Sonar Sweep
// Count number of times a depth measurement increases from the previous measurement 
// Answer is total number of measurements larger than the previous measurement

import (
	"fmt"
	"os"
	"bufio"
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
		num, _ := strconv.Atoi(line)
		
		lines = append(lines, num)
	}
	return lines, scanner.Err()
}

func main() {
	
	lines, err := readLines("input.txt")
	check(err)
	// fmt.Println(lines)

	var count int
	var tmp int
	for i, line := range lines {
		// fmt.Println("Current line value: ", line )
		// fmt.Println(line, "-", tmp)
		if i > 0 {
		result := line - tmp 	
			if result > 0 {
				count++
			}
		}
		tmp = line
	}
	fmt.Println("Total increased: ", count )
}
