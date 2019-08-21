package main

import (
	"os"
	"flag"
	"lexgo/scanner"
)

func main() {
	fptr := flag.String("fpath", "source.go", "file path to read from")
	dirptr := flag.String("outputdir", "", "file path to write to")
	flag.Parse()

	if fptr == nil {
		os.Exit(1)
	}
	if dirptr == nil {
		dir, err := os.Getwd()
		if err != nil {
			os.Exit(1)
		}
		dirptr = &dir
	}

	scanner.SourcefileWalk(*fptr, *dirptr)
}
