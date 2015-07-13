package dict

import (
	"errors"
	"math"
)

type Dictionary struct {
	words   []string
	size    int
	maxSize int
	seed    int
}

func NewDictionary() (*Dictionary, error) {
	dict, err := newDictionaryCustom()
	if err == nil {
		return dict, nil
	}

	dict, err = newDictionarySystem()
	if err == nil {
		return dict, nil
	}

	dict, err = newDictionaryDownload()
	if err == nil {
		return dict, nil
	}

	return nil, errors.New("Cannot find a dictionary")
}

func NewDictionaryWords(words []string) *Dictionary {
	count := len(words)
	maxSize := int(math.Logb(float64(count)))

	dict := &Dictionary{
		words:   words,
		size:    maxSize,
		maxSize: maxSize,
		seed:    0,
	}

	return dict
}

func (d *Dictionary) Word(val int) string {
	val += d.seed

	// Use the step and seed to get the index
	count := len(d.words)
	step := count / (1 << uint(d.size))
	val = int(val * step)

	if val >= count {
		val %= count
	}

	return d.words[val]
}

func (d *Dictionary) Size() int {
	return d.size
}

func (d *Dictionary) SetSize(size int) {
	if size <= 0 || size > d.maxSize {
		size = d.maxSize
	}

	d.size = size
}

func (d *Dictionary) MaxSize() int {
	return d.maxSize
}

func (d *Dictionary) Seed() int {
	return d.seed
}

func (d *Dictionary) SetSeed(seed int) {
	d.seed = seed
}
