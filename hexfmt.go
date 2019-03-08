package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// HexFmt reads a binary file and formats the output to a hex string
func HexFmt(fileName string, numElementsPerLine int) string {
	binFile, err := os.Open(fileName)
	checkError(err)
	defer binFile.Close()

	bin, err := ioutil.ReadAll(binFile)
	checkError(err)

	var sb strings.Builder
	for i, b := range bin {
		if i > 0 && i%numElementsPerLine == 0 {
			sb.WriteString("\n")
		}
		s := fmt.Sprintf("0x%02X, ", b)
		sb.WriteString(s)
	}
	return sb.String()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
