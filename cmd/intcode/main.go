package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const (
	OpcodeTerminate = 99
	OpcodeAdd       = 1
	OpcodeMultiply  = 2
	OffsetOpCode    = 0
	OffsetOperand1  = 1
	OffsetOperand2  = 2
	OffsetLocation  = 3
)

func main() {
	if len(os.Args) != 2 {
		panic("You need to give a path to a file to read on the command line")
	}
	datafile, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(fmt.Sprintf("Unable to read in file %v\nError returned: %v", os.Args[1]))
	}
	dataArrayString := strings.Split(string(datafile), ",")
	dataArray := convertToIntegerArray(dataArrayString)
	printArray(dataArray, "Values Before")

	for i := 0; true; i++ {
		var total int
		opcode := dataArray[i+OffsetOpCode]
		if opcode == OpcodeTerminate {
			break
		}

		op1Position := dataArray[i+OffsetOperand1]
		op2Position := dataArray[i+OffsetOperand2]
		locationPosition := dataArray[i+OffsetLocation]

		op1Value := dataArray[op1Position]
		op2Value := dataArray[op2Position]
		//location := dataArray[locationPosition]

		fmt.Printf("Operation %v: Opcode: %v\tOperand 1 Position: %v\tOperand 1: %v\tOperand 2 Position: %v\tOperand 2: %v\tLocation: %v\n", i, opcode, op1Position, op1Value, op2Position, op2Value, locationPosition)

		switch opcode {
		case OpcodeAdd:
			total = op1Value + op2Value
		case OpcodeMultiply:
			total = op1Value * op2Value
		}
		dataArray[locationPosition] = total
		i += 3
	}

	printArray(dataArray, "Values After")
}

func convertToIntegerArray(data []string) []int {
	intArray := make([]int, len(data)+1)
	for i := 0; i < len(data); i++ {
		value, _ := strconv.Atoi(data[i])
		intArray[i] = value
	}
	return intArray
}

func printArray(slice []int, header string) {
	fmt.Printf("Name: %v\n", header)
	for i := 0; i < len(slice); i++ {
		fmt.Printf("Position %v,%v,(%v)\n", i, slice[i], getAction(i))
	}
	fmt.Print("\n")
}

func getAction(i int) string {
	action := "no-op"
	switch i % 4 {
	case 0:
		action = "opcode"
	case 1:
		action = "operand 1"
	case 2:
		action = "operand 2"
	case 3:
		action = "location"
	}
	return action
}
