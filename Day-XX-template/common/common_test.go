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
