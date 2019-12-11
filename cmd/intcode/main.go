package main

import (
	"fmt"
	"github.com/number7/advent-of-code-2019/cmd/intcode/computer"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const TargetValue = 19690720

func main() {
	if len(os.Args) != 2 {
		panic("You need to give a path to a file to read on the command line")
	}
	datafile, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(fmt.Sprintf("Unable to read in file %v\nError returned: %v", os.Args[1], err))
	}
	dataArrayString := strings.Split(string(datafile), ",")
	dataArray := convertToIntegerArray(dataArrayString)

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			newArray := computer.Compute(dataArray, i, j, false)

			if newArray[0] == TargetValue {
				fmt.Printf("NOUN: %v\tVERB: %v\t\tInputValue: %v\n", i, j, 100*i+j)
				break
			}
		}
	}
}

func convertToIntegerArray(data []string) []int {
	intArray := make([]int, len(data))
	for i := 0; i < len(data); i++ {
		value, _ := strconv.Atoi(data[i])
		intArray[i] = value
	}
	return intArray
}
