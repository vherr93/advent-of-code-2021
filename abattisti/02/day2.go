package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"

)

func check(e error)  {
	if e != nil{
		fmt.Println(e)
		os.Exit(1)
	}
}

func main(){

	horPos := 0
	depth :=0
	aim := 0

	fpath := "input.txt"

	file, err := os.Open(fpath)
	check(err)
	defer file.Close()

	//read through file, place ints into array
	scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanWords)
    for scanner.Scan() {
        command := scanner.Text()
		scanner.Scan()
		value, err := strconv.Atoi(scanner.Text())
		check(err)

		if (command == "forward"){
			horPos += value
			depth += aim * value
		} else if (command == "down"){
			aim += value
		} else if (command == "up"){
			aim -= value
		}
    }

	fmt.Println(horPos)
	fmt.Println(depth)

	fmt.Println(depth * horPos)


}