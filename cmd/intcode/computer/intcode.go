package computer

import "fmt"

const (
	OpcodeTerminate = 99
	OpcodeAdd       = 1
	OpcodeMultiply  = 2
	OffsetOpCode    = 0
	OffsetOperand1  = 1
	OffsetOperand2  = 2
	OffsetLocation  = 3
)

func Compute(dataArray []int, noun int, verb int, printData bool) []int {
	newArray := make([]int, len(dataArray))
	copy(newArray, dataArray)
	newArray[1] = noun
	newArray[2] = verb
	if printData {
		printArray(newArray, "Values Before")
	}

	for i := 0; true; i += 4 {
		var total int
		opcode := newArray[i+OffsetOpCode]
		if opcode == OpcodeTerminate {
			break
		}

		op1Position := newArray[i+OffsetOperand1]
		op2Position := newArray[i+OffsetOperand2]
		location := newArray[i+OffsetLocation]

		op1Value := newArray[op1Position]
		op2Value := newArray[op2Position]

		if printData {
			fmt.Printf("Operation %v: Opcode: %v\tOperand 1 Position: %v\tOperand 1: %v\tOperand 2 Position: %v\tOperand 2: %v\tLocation: %v\n", i, opcode, op1Position, op1Value, op2Position, op2Value, location)
		}

		switch opcode {
		case OpcodeAdd:
			total = op1Value + op2Value
		case OpcodeMultiply:
			total = op1Value * op2Value
		}
		newArray[location] = total
	}

	if printData {
		printArray(newArray, "Values After")
	}
	return newArray
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
