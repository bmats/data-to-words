package dict

import (
	"io/ioutil"
	"os"
	"runtime"
	"strings"
	"testing"
)

func TestNewDictionaryCustom(t *testing.T) {
	words := []string{"the", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"}
	data := strings.Join(words, "\n")
	err := ioutil.WriteFile(customDictPath, []byte(data), 0644)
	if err != nil {
		t.Errorf("write dictionary err (%s) != nil", err)
	}
	defer os.Remove(customDictPath)

	dict, err := newDictionaryCustom()
	if err != nil {
		t.Errorf("err (%s) != nil", err)
	}

	if !testSlicesEqual(words, dict.words) {
		t.Error("words != dict.words")
	}
}

func TestNewDictionarySystem(t *testing.T) {
	if runtime.GOOS != "linux" && runtime.GOOS != "darwin" {
		t.Logf("OS %s may not have system dictionary, so test skipped")
		return
	}

	dict, err := newDictionarySystem()
	if err != nil {
		t.Errorf("err (%s) != nil", err)
	}

	if len(dict.words) == 0 {
		t.Error("len(dict.words) == 0")
	}
}

func TestNewDictionaryDownload(t *testing.T) {
	// Make sure custom dict doesn't already exist
	os.Remove(customDictPath)

	dict, err := newDictionaryDownload()
	if err != nil {
		t.Errorf("err (%s) != nil", err)
	}

	if len(dict.words) == 0 {
		t.Error("len(dict.words) == 0")
	}

	customDict, err := newDictionaryCustom()
	if err != nil {
		t.Errorf("custom err (%s) != nil", err)
	}
	defer os.Remove(customDictPath)

	if !testSlicesEqual(dict.words, customDict.words) {
		t.Error("downloaded dict != saved dict")
	}
}
