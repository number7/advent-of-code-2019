package main

import (
	"bufio"
	"fmt"
	"github.com/number7/advent-of-code-2019/cmd/common"
	"os"
	"strings"
)

var directions = map[string]string{
	"L": "Left",
	"R": "Right",
	"U": "Up",
	"D": "Down",
}

func main() {
	if len(os.Args) != 2 {
		panic("You need to give a path to a file to read on the command line")
	}
	datafile := common.MustOpenFile(os.Args[1])
	defer datafile.Close()

	lineScanner := bufio.NewScanner(datafile)

	wireOne := ReadLine(lineScanner)
	wireTwo := ReadLine(lineScanner)

	fmt.Printf("Line one: %v\n\n", wireOne)
	fmt.Printf("Line two: %v\n\n", wireTwo)
}

func ReadLine(scanner *bufio.Scanner) []string {
	scanner.Scan()
	return SplitDirections(scanner.Text())
}

func SplitDirections(str string) []string {
	return strings.Split(str, ",")
}
