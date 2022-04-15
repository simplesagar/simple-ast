package simpleast

import (
	"flag"
	"fmt"
	"os"
)

var outfile string

func init() {
	flag.Usage = usage
	flag.StringVar(&outfile, "out", "", "write output to a file instead of stdout")
}

func usage() {
	fmt.Fprintf(os.Stderr, `
Usage of simpleast:

simpleast [flags] file

simpleast accepts the following flags:
`)
	flag.PrintDefaults()
	os.Exit(2)
}

// Main1 called by cmd/simpleast
func Main() int {
	return 1
}

// Parse converts input to ast
func Parse(input string) (string, error) {
	return "", nil
}
