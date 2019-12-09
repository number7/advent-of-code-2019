package main

import (
	"bufio"
	"fmt"
	"gkwkr/advent-of-code-2019/cmd/common"
	"gkwkr/advent-of-code-2019/cmd/module-mass/computer"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		panic("You need to give a path to a file to read on the command line")
	}
	datafile := common.MustOpenFile(os.Args[1])
	defer datafile.Close()

	lineScanner := bufio.NewScanner(datafile)

	total := 0
	for i := 1; lineScanner.Scan(); i++ {
		mass, err := strconv.Atoi(lineScanner.Text())
		if err != nil {
			fmt.Printf("Line %v has a non-integer value.\n", i)
		}
		gas := computer.ComputeGas(mass)
		fmt.Printf("Gas for line %v value of %v: %v\n", i, lineScanner.Text(), gas)
		total += gas
	}
	fmt.Printf("TOTAL: %v\n", total)
}
