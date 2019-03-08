package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fileName := flag.String("f", "", "file name")
	numElementsPerLine := flag.Int("n", 16, "number of elements per line")
	flag.Parse()

	if *fileName == "" || *numElementsPerLine < 1 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	result := HexFmt(*fileName, *numElementsPerLine)
	fmt.Println(result)
}
