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

func getCommonBit(mat [][]int, i int) int {
	sum := 0
	for _, v := range mat {
		sum += v[i]
	}
	avg := float64(sum) / float64(len(mat))
	b := int(math.Round(avg))

	return b

}

func getCommonBitLine(mat [][]int, i int, d bool) [][]int {

	l := [][]int{}
	var b int

	if d {
		b = getCommonBit(mat, i)
	} else {
		b = int(math.Abs(float64(getCommonBit(mat, i) - 1)))
	}

	for _, v := range mat {
		if b == v[i] {
			l = append(l, v)
		}
	}

	if len(l) > 1 {
		l = getCommonBitLine(l, i+1, d)
	}

	return l
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

func binToDec(s []int) int64 {
	d, err := strconv.ParseInt((getString(s)), 2, 64)
	check(err)

	return d

}

func main() {
	bits, err := intMatrix("input.txt")
	check(err)

	in := len(bits[0])
	g := make([]int, in)
	e := make([]int, in)

	for i := 0; i < in; i++ {
		g[i] += getCommonBit(bits, i)
		e[i] += int(math.Abs(float64(getCommonBit(bits, i) - 1)))
	}

	gr := binToDec(g)
	er := binToDec(e)

	con := productint64(gr, er)
	fmt.Println("Power consumption:", con)

	og := getCommonBitLine(bits, 0, true)
	cs := getCommonBitLine(bits, 0, false)

	ogr := binToDec(og[0])
	csr := binToDec(cs[0])

	fmt.Println("Life Support:", productint64(ogr, csr))

}
