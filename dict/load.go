package dict

import (
	"bufio"
	"io"
	"math"
	"net/http"
	"os"
)

const (
	customDictPath string = "dictionary.txt"
	onlineDictUrl  string = "https://gist.github.com/bmats/9a946845ad065558d4d6/raw/4be557ed63be643afaf898197f7dcbabb37630f1/words.txt"
)

// Load a dictionary from the custom dictionary file
func newDictionaryCustom() (*Dictionary, error) {
	customFile, err := os.Open(customDictPath)
	if err != nil {
		return nil, err
	}
	defer customFile.Close()

	dict := &Dictionary{
		words: make([]string, 0, 1024),
		size:  1,
		seed:  0,
	}
	dict.readIn(customFile)
	return dict, nil
}

// Load a dictionary from the *nix builtin dictionary
func newDictionarySystem() (*Dictionary, error) {
	dictFile, err := os.Open("/usr/share/dict/words")
	if err != nil {
		return nil, err
	}
	defer dictFile.Close()

	dict := &Dictionary{
		words: make([]string, 0, 1024),
		size:  1,
		seed:  0,
	}
	dict.readIn(dictFile)
	return dict, nil
}

// Download the words for a dictionary from an online version of the *nix dictionary
func newDictionaryDownload() (*Dictionary, error) {
	resp, err := http.Get(onlineDictUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	dict := &Dictionary{
		words: make([]string, 0, 1024),
		size:  1,
		seed:  0,
	}
	dict.readIn(resp.Body)

	// Save downloaded to custom dict
	out, err := os.Create(customDictPath)
	if err == nil {
		defer out.Close()

		writer := bufio.NewWriter(out)
		defer writer.Flush()

		for _, e := range dict.words {
			writer.WriteString(e)
			writer.WriteRune('\n')
		}
	}

	return dict, nil
}

func (d *Dictionary) readIn(r io.Reader) {
	d.words = make([]string, 0, 1<<uint(d.size))

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil && err != io.EOF {
			break
		}

		d.words = append(d.words, scanner.Text())
	}

	count := len(d.words)
	d.maxSize = int(math.Logb(float64(count)))

	d.size = d.maxSize
}
