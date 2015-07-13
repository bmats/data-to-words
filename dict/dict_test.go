package dict

import (
	"math"
	"testing"
)

func testMakeWordSlice(length int) []string {
	letters := make([]string, length)
	for i := range letters {
		letters[i] = string('🌀' + i)
	}

	return letters
}

func testSlicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i, e := range a {
		if e != b[i] {
			return false
		}
	}

	return true
}

func TestNewDictionary(t *testing.T) {
	_, err := NewDictionary()

	if err != nil {
		t.Errorf("err (%s) != nil", err)
	}
}

func TestNewDictionaryWords(t *testing.T) {
	words := testMakeWordSlice(30)
	dict := NewDictionaryWords(words)

	if !testSlicesEqual(dict.words, words) {
		t.Error("dict.words != words")
	}

	expectedSize := int(math.Logb(float64(len(words))))
	if dict.size != expectedSize {
		t.Errorf("dict.size == %d, want %d", dict.size, expectedSize)
	}

	if dict.size != dict.maxSize {
		t.Error("dict.size != dict.maxSize")
	}
}

func TestDictionaryWord(t *testing.T) {
	words := testMakeWordSlice(64)

	dict := NewDictionaryWords(words)

	for i, l := range words {
		if word := dict.Word(i); word != l {
			t.Errorf("Word(%d) == %s, want %s", i, word, l)
		}
	}
}

func TestDictionaryWordWithSize(t *testing.T) {
	words := testMakeWordSlice(64)

	dict := NewDictionaryWords(words)
	dict.SetSize(4)

	for i := 0; i < 16; i++ {
		l := words[i*4]

		if word := dict.Word(i); word != l {
			t.Errorf("Word(%d) == %s, want %s", i, word, l)
		}
	}
}

func TestDictionaryWordWithSeed(t *testing.T) {
	words := testMakeWordSlice(16)

	dict := NewDictionaryWords(words)
	dict.SetSeed(2485)

	for i, _ := range words {
		seedIndex := (i + dict.Seed()) % len(words)
		l := words[seedIndex]

		if word := dict.Word(i); word != l {
			t.Errorf("Word(%d) == %s, want %s", i, word, l)
		}
	}
}

func TestDictionaryDefaultSize(t *testing.T) {
	words := testMakeWordSlice(50)
	dict := NewDictionaryWords(words)

	expectedSize := int(math.Logb(float64(len(words))))

	if dict.Size() != expectedSize {
		t.Errorf("Size() == %d, want %d", dict.Size(), expectedSize)
	}
}

func TestDictionarySetSize(t *testing.T) {
	dict, err := NewDictionary()
	if err != nil {
		t.Errorf("err (%s) != nil", err)
	}

	const newSize int = 3
	dict.SetSize(newSize)

	if dict.Size() != newSize {
		t.Errorf("Size() == %d, want %d", dict.Size(), newSize)
	}
}

func TestDictionarySetSeed(t *testing.T) {
	dict, err := NewDictionary()
	if err != nil {
		t.Errorf("err (%s) != nil", err)
	}

	const newSeed int = 2485
	dict.SetSeed(newSeed)

	if dict.Seed() != newSeed {
		t.Errorf("Seed() == %d, want %d", dict.Seed(), newSeed)
	}
}
