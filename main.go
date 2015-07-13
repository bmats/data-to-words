package main

import (
	"flag"
	"github.com/bmats/data-to-words/dict"
	"io"
	"os"
	"strings"
)

var (
	sizeParam  = flag.Uint("size", 0, "The dictionary size. Use 0 to select the maximum")
	seedParam  = flag.Int("seed", 0, "Seed for choosing words")
	delimParam = flag.String("delim", " ", "Word delimiter")
)

func main() {
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
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Stderr.WriteString("\n")
	}
	os.Stdout.WriteString("\n")
}
