package dict

import (
	"bytes"
	"strings"
	"testing"
)

func testDictionaryTranslate(t *testing.T, dict *Dictionary, input []byte, expected []string) {
	reader := bytes.NewBuffer(input)
	//strings.NewReader(input)
	writer := new(bytes.Buffer)

	err := dict.Translate(reader, writer, ",")
	if err != nil {
		t.Errorf("err (%s) != nil", err)
	}

	split := strings.Split(string(writer.Bytes()), ",")

	if !testSlicesEqual(split, expected) {
		t.Errorf("output == %s, want %s", split, expected)
	}
}

func TestDictionaryTranslateAligned(t *testing.T) {
	words := testMakeWordSlice(256)
	dict := NewDictionaryWords(words)

	input := []byte{0, 11, 22, 33, 44, 55, 66, 77, 88, 99, 255}

	expected := make([]string, len(input))
	for i, e := range input {
		expected[i] = words[e]
	}

	testDictionaryTranslate(t, dict, input, expected)
}

func TestDictionaryTranslateSizeSpanningMultipleBytes(t *testing.T) {
	words := testMakeWordSlice(1024)
	dict := NewDictionaryWords(words)

	input := []byte{0xac, 0xac, 0xac, 0xac}
	expected := []string{words[0x2b2], words[0x2ca], words[0x32b], words[0x000]}

	testDictionaryTranslate(t, dict, input, expected)
}

func TestDictionaryTranslateSizeLessThanOneByte(t *testing.T) {
	words := testMakeWordSlice(8)
	dict := NewDictionaryWords(words)

	input := []byte{0xac, 0xac}
	expected := []string{words[0x5], words[0x3], words[0x1], words[0x2], words[0x6], words[0x0]}

	testDictionaryTranslate(t, dict, input, expected)
}

func TestDictionaryTranslateEncounteringEOFBeforeDone(t *testing.T) {
	words := testMakeWordSlice(1024) // 10 bits
	dict := NewDictionaryWords(words)

	input := []byte{0xfc}
	expected := []string{words[0x3f0]}

	testDictionaryTranslate(t, dict, input, expected)
}
