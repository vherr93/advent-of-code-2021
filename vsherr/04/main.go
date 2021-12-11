package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// create slice of comma separated numbers
func readCsv(path string) ([]int, error) {
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	var nums []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		val := strings.Split(line, ",")
		for _, i := range val {
			n, _ := strconv.Atoi(i)

			nums = append(nums, n)
		}
	}
	return nums, scanner.Err()
}

func createBingoBoards(path string) ([]*Board, error) {
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	var bingoBangoBongo []*Board

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	tmpb := new(Board)
	// needs to scan 5 lines, then if encounter /n/n, start over until EOF
	for scanner.Scan() {

		s := scanner.Text()
		// fmt.Println(s)
		if s == "" {
			bingoBangoBongo = append(bingoBangoBongo, tmpb)
			tmpb = new(Board)
		} else {
			tmpb.num = append(tmpb.num, strings.Split(s, " "))
		}

	}

	return bingoBangoBongo, scanner.Err()
}

type Board struct {
	num [][]string
}

// type BingoBoards struct {
// 	boards []Board
// }

// func (boards *BingoBoards) AddBoard(board Board) []Board {
// 	boards.boards = append(boards.boards, board)
// 	return boards.boards
// }

func main() {

	t, err := readCsv("numbers.txt")
	check(err)
	fmt.Println(t)

	b, err := createBingoBoards("boards.txt")
	check(err)
	for _, v := range b {
		fmt.Println(v.num)
	}

}
