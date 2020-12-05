package common

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestShouldFatalIfInputFileError(t *testing.T) {
	// when
	_, err := getFile("/not/existing/file")

	// then
	assert.Error(t, err)
}

func TestShouldGetInputFromFile(t *testing.T) {
	// given
	content := []byte("Test file")
	tmpfile, err := ioutil.TempFile("", "test")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}

	// when
	input, err := getFile(tmpfile.Name())

	// then
	assert.Nil(t, err)
	assert.Equal(t, "Test file", input)
}

func TestShouldGetIntValuesFromStringInput(t *testing.T) {
	// given
	inputLines := "145\n20\n23\n90\n0\n2\n99999\n6"
	inputCommas := "145,20,23,90,0,2,99999,6"

	// when
	intValuesLines := GetIntArrayFromStringInput(inputLines, "\n")
	intValuesCommas := GetIntArrayFromStringInput(inputCommas, ",")

	// then
	assert.Equal(t, []int{145, 20, 23, 90, 0, 2, 99999, 6}, intValuesLines)
	assert.Equal(t, []int{145, 20, 23, 90, 0, 2, 99999, 6}, intValuesCommas)
}

func TestShouldGetStringValuesFromStringInput(t *testing.T) {
	// given
	inputLines := "145\n20\ntest\nelo"
	inputCommas := "145,20,test,elo"

	// when
	stringValuesLines := GetStringArrayFromStringInput(inputLines, "\n")
	stringValuesCommas := GetStringArrayFromStringInput(inputCommas, ",")

	// then
	assert.Equal(t, []string{"145", "20", "test", "elo"}, stringValuesLines)
	assert.Equal(t, []string{"145", "20", "test", "elo"}, stringValuesCommas)
}

func TestShouldGetMapFromStringInput(t *testing.T) {
	// given
	inputLines := "key1:val1\nkey2:val2\n\nk1:v1\n\nkey:value a:b c:d"

	// when
	inputMap := GetArrayOfMapsFromString(inputLines, "\n\n", []string{" ", "\n"}, ":")

	// then
	assert.Equal(t, []map[string]string{{"key1": "val1", "key2": "val2"}, {"k1": "v1"}, {"key": "value", "a": "b", "c": "d"}}, inputMap)
}
