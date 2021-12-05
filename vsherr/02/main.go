package main

// Day 2: Dive!
// Part 1: Calculate final horizontal position and depth after following specified course
// Part 1: Answer is product of final horizontal position and final depth
// Part 2: Use 3rd metric aim to calculate vertical position

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

func readWords(path string) ([]string, error) {
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	return words, scanner.Err()
}

func product(vpos, hpos int) int {
	final := vpos * hpos

	return final
}

func main() {

	var vpos int
	var hpos int
	var aim int

	words, err := readWords("input.txt")
	check(err)

	for i := 0; i < len(words)-1; i += 2 {
		dir := words[i]
		pos, _ := strconv.Atoi(words[i+1])

		if dir == "forward" {
			hpos = hpos + pos
			vpos = vpos + (pos * aim)
			continue
		}
		if dir == "up" {
			aim = aim - pos
		} else {
			aim = aim + pos
		}
	}

	fmt.Println("Result:", product(vpos, hpos))

}
