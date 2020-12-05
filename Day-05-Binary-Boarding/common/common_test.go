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

func TestShouldSplitByAnyOfDelimiter(t *testing.T) {
	// given
	input := "a:b.c-d e\nf,g"

	// when
	array := SplitAny(input, []string{":", ".", "-", " ", "\n", ","})

	// then
	assert.Equal(t, []string{"a", "b", "c", "d", "e", "f", "g"}, array)
}

func TestShouldCheckIfArraysHaveAllSameElements(t *testing.T) {
	// given
	a1 := []string{"a", "b", "c", "d"}
	a2 := []string{"b", "a", "d", "c"}
	a3 := []string{"a", "d", "c"}
	a4 := []string{"a", "d", "c", "h"}

	// then
	assert.True(t, HaveSameElements(a1, a2))
	assert.True(t, HaveSameElements(a2, a1))
	assert.True(t, HaveSameElements(a1, a1))
	assert.True(t, HaveSameElements(a2, a2))
	assert.False(t, HaveSameElements(a1, a3))
	assert.False(t, HaveSameElements(a3, a1))
	assert.False(t, HaveSameElements(a1, a4))
	assert.False(t, HaveSameElements(a4, a1))
	assert.False(t, HaveSameElements(a3, a4))
	assert.False(t, HaveSameElements(a4, a3))
}

func TestShouldCheckIfArrayContainsElement(t *testing.T) {
	// given
	array := []string{"a", "b", "c", "d"}

	// then
	assert.True(t, ArrayContains(array, "a"))
	assert.True(t, ArrayContains(array, "b"))
	assert.True(t, ArrayContains(array, "c"))
	assert.True(t, ArrayContains(array, "d"))
	assert.False(t, ArrayContains(array, "e"))
	assert.False(t, ArrayContains(array, ""))
	assert.False(t, ArrayContains(array, " "))
}

func TestShouldCheckIfIntArrayContainsElement(t *testing.T) {
	// given
	array := []int{1, 2, 3, 4}

	// then
	assert.True(t, IntArrayContains(array, 1))
	assert.True(t, IntArrayContains(array, 2))
	assert.True(t, IntArrayContains(array, 3))
	assert.True(t, IntArrayContains(array, 4))
	assert.False(t, IntArrayContains(array, 5))
	assert.False(t, IntArrayContains(array, 0))
	assert.False(t, IntArrayContains(array, -1))
}

func TestShouldGetMapKeys(t *testing.T) {
	// given
	inputMap := map[string]string{"key": "value", "k2": "v2"}

	// then
	assert.Equal(t, MapKeys(inputMap), []string{"key", "k2"})
}

func TestShouldCheckIfMapContainsAllKeysFromList(t *testing.T) {
	// given
	inputMap := map[string]string{"key": "value", "k2": "v2"}
	checkKeys := []string{"key", "k2"}
	checkKeys2 := []string{"k2", "key"}
	checkKeys3 := []string{"key"}
	checkKeys4 := []string{"k2"}
	var checkKeys5 []string
	checkKeys6 := []string{""}

	// then
	assert.True(t, HaveAllKeys(inputMap, checkKeys))
	assert.True(t, HaveAllKeys(inputMap, checkKeys2))
	assert.False(t, HaveAllKeys(inputMap, checkKeys3))
	assert.False(t, HaveAllKeys(inputMap, checkKeys4))
	assert.False(t, HaveAllKeys(inputMap, checkKeys5))
	assert.False(t, HaveAllKeys(inputMap, checkKeys6))
}

func TestShouldCompareIntValues(t *testing.T) {
	assert.Equal(t, CompareInt(1, 2), -1)
	assert.Equal(t, CompareInt(2, 1), 1)
	assert.Equal(t, CompareInt(2, 2), 0)
}

func TestShouldCompareStringValues(t *testing.T) {
	assert.Equal(t, CompareStrings("b", "a"), 1)
	assert.Equal(t, CompareStrings("a", "b"), -1)
	assert.Equal(t, CompareStrings("aa", "ab"), -1)
	assert.Equal(t, CompareStrings("a", "a"), 0)
	assert.Equal(t, CompareStrings("a", "1"), 1)
	assert.Equal(t, CompareStrings("1", "/"), 1)
	assert.Equal(t, CompareStrings("ala", "ale"), -1)
}
