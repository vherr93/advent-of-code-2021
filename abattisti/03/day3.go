package main

import (
	"fmt"
	"os"
	"bufio"
	"math"
	"strconv"
	"strings"

)

func check(e error)  {
	if e != nil{
		fmt.Println(e)
		os.Exit(1)
	}
}

func createReport(fPath string)([][]int, error) {

	file, err := os.Open(fPath)
	check(err)
	defer file.Close()

	report := [][] int {}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {

		tmpStr := scanner.Text()
		tmpSlice := []int {}

		for i := 0; i < len(tmpStr); i++ {
			tmpNum, err := strconv.Atoi(string(tmpStr[i]))
			check(err)
			tmpSlice = append(tmpSlice, tmpNum)
		}
	
		report = append (report, tmpSlice)
		
	}

	return report, scanner.Err()
}

func mostCommonBit(report [][]int, index int)(int){

	sum := 0

	for _, v := range report {
		sum += v[index]
	}

	avg := float64(sum) / float64(len(report))
	mcb := int(math.Round(avg))
	
	return(mcb)

}

func binaryToInt (bin []int)(int){
	
	tmpSlice := []string{}

	for i := range bin{
        tmpSlice = append(tmpSlice, strconv.Itoa(bin[i]))
	}

	tmpStr := strings.Join(tmpSlice, "")

	num,err := strconv.ParseInt(tmpStr,2,32)
	check(err)

	return int(num)
}

func filterReport(report [][]int, index int,  mCommon bool)([]int){

}

func main(){

	//put diagnostic report into 2d slice
	fPath := "input.txt"
	report, err := createReport(fPath)
	check(err)

	bitLen := len(report[0])
	gR := make([]int, bitLen)
	eR := make([]int, bitLen)

	//get gamma and epsilon values
	for i :=0; i < bitLen; i++ {
		mcb := mostCommonBit(report,i)
		gR[i] += mcb
		eR[i] += int(math.Abs(float64(mcb - 1)))
	}

	powerConsumption := binaryToInt(gR) * binaryToInt(eR)
	fmt.Println("Power Consumption: ", powerConsumption)
	



	


	

	

}