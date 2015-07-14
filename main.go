package main

import (
	"flag"
	"fmt"
	"github.com/bmats/data-to-words/dict"
	"io"
	"os"
	"strings"
)

var (
	sizeParam  = flag.Uint("size", 0, "Dictionary size. 0 to select the maximum")
	seedParam  = flag.Int("seed", 0, "Seed for choosing words")
	delimParam = flag.String("delim", " ", "Word delimiter")
)

func usage() {
	fmt.Println("data-to-words 0.1.0")
	fmt.Printf("Usage: %s [options] [data string]\n\n", os.Args[0])

	fmt.Println("Pipe some bytes to me or pass bytes to me as an argument.\n")

	fmt.Println("Options:")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()

	dict, err := dict.NewDictionary()
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Stderr.WriteString("\n")
		return
	}

	dict.SetSize(int(*sizeParam))
	dict.SetSeed(*seedParam)

	var reader io.Reader

	// Use argument as input if possible, otherwise, use stdin
	text := flag.Arg(0)
	if len(text) > 0 {
		reader = strings.NewReader(text)
	} else {
		reader = os.Stdin
	}

	err = dict.Translate(reader, os.Stdout, *delimParam)
	if err == nil {
		fmt.Println()
	} else {
		fmt.Fprintln(os.Stderr, err)
	}
}
