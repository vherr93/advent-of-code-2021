package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"

)

func check(e error)  {
	if e != nil{
		os.Exit(1)
	}
}

//parse file of ints into an array which is then returned. 
func readFile(path string) ([]int, error){

	//read in file, check for errors, close file
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	//read through file, place ints into array
	scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanWords)
    var depthReadings []int
    for scanner.Scan() {
        x, err := strconv.Atoi(scanner.Text())
        check(err)
        depthReadings = append(depthReadings, x)
    }

    return depthReadings, scanner.Err()
}

//loop through array, record num times new reading is larger than previous
func measureDepth(depthReadings []int)(int){

	count := 0
	const wSize = 3
	
	for i := 0; i < len(depthReadings) - wSize; i+= 1 {

		if (depthReadings[i + wSize] > depthReadings[i]){
			count += 1
		}
	}
	
	return count
}

func main (){

	depthReadings,err := readFile("input.txt")
	check(err)
	//dlen := len(depthReadings)
	//fmt.Println(len(depthReadings))
	numIncreases := measureDepth(depthReadings)
	fmt.Println(numIncreases)


}


