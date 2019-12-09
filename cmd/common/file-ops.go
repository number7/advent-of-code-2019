package common

import (
	"fmt"
	"os"
)

func MustOpenFile(file string) *os.File {
	datafile, err := os.Open(file)
	if err != nil {
		panic(fmt.Sprintf("cant read your goddamned file %v", file))
	}
	return datafile
}
