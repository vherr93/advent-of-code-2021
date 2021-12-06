package main

// Day 3: Binary Diagnostic
// Part 1: Calculate most common & least common bit in a list of binary numbers
// Part 1: Answer is the decimal product of above

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// 2d slice
// dy is number of rows
// dx is bits in each row
// scanner

func intMatrix(path string) ([][]int, error) {
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	mat := [][]int{}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {

		s := scanner.Text()
		var tmp []int

		for i := 0; i < len(s); i++ {
			n, err := strconv.Atoi(string(s[i]))
			check(err)
			tmp = append(tmp, n)
		}

		mat = append(mat, tmp)
	}

	return mat, scanner.Err()
}

// func getCommon(mat [][]int, i int) int {

// 	var tmp []int
// 	// for each bit slice
// 	for _, v := range mat[i] {
// 		// append val at i position to tmp
// 		tmp = append(tmp, v)

// 		fmt.Println(tmp)
// 		fmt.Println(len(tmp))
// 		i++
// 	}
// 	return i
// }

func getCommon(mat [][]int, i int) int {
	sum := 0
	for _, v := range mat {
		sum += v[i]
	}
	avg := float64(sum) / float64(len(mat))
	b := int(math.Round(avg))

	// create new slice where v[i] = b
	// return this slice

	fmt.Println(b)
	return b

}

func getString(sli []int) string {
	var s []string
	for i := range sli {
		n := sli[i]
		t := strconv.Itoa(n)
		s = append(s, t)
	}
	d := strings.Join(s, "")
	return d
}

func productint64(g, e int64) int64 {
	final := g * e

	return final
}

// func binToDec {

// }

func main() {
	bits, err := intMatrix("input.txt")
	check(err)

	in := len(bits[0])

	// binary seq of most common val in index of slices in bits arr
	g := make([]int, in)
	// least common bit
	e := make([]int, in)
	// calculate decimal of g & e

	for i := 0; i < in; i++ {
		g[i] += getCommon(bits, i)
		// returns parsed 2d slice, pass this now instead
		e[i] += int(math.Abs(float64(getCommon(bits, i) - 1)))
	}

	// at this point we want 2 row numbers

	fmt.Println(g)
	fmt.Println(e)

	gr, err := strconv.ParseInt((getString(g)), 2, 64)
	check(err)
	er, err := strconv.ParseInt((getString(e)), 2, 64)
	check(err)
	//fmt.Println(gr)
	//fmt.Println(er)

	con := productint64(gr, er)
	fmt.Println("Power consumption:", con)

}
