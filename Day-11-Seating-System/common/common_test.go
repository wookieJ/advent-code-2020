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
	assert.ElementsMatch(t, []string{"key", "k2"}, MapKeys(inputMap))
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

func TestShouldGetArrayOfArrays(t *testing.T) {
	// given
	input := "a\nb\nc\n\na\n\n1\n2"

	// when
	arrayOfArrays := GetArrayOfArrays(input, "\n\n", "\n")

	// then
	assert.ElementsMatch(t, [][]string{{"a", "b", "c"}, {"a"}, {"1", "2"}}, arrayOfArrays)
}

func TestShouldGetAllLetters(t *testing.T) {
	// given
	a1 := []string{"abc"}
	a2 := []string{"a", "b", "c"}
	a3 := []string{"ab", "ac"}
	a4 := []string{"a", "a", "a", "a"}
	a5 := []string{"b"}
	a6 := []string{"a;b", "a;b;c;de"}

	// then
	assert.ElementsMatch(t, []string{"a", "b", "c"}, SplitAndGetAll(a1, ""))
	assert.ElementsMatch(t, []string{"a", "b", "c"}, SplitAndGetAll(a2, ""))
	assert.ElementsMatch(t, []string{"a", "b", "c"}, SplitAndGetAll(a3, ""))
	assert.ElementsMatch(t, []string{"a"}, SplitAndGetAll(a4, ""))
	assert.ElementsMatch(t, []string{"b"}, SplitAndGetAll(a5, ""))
	assert.ElementsMatch(t, []string{"a", "b", "c", "de"}, SplitAndGetAll(a6, ";"))

}

func TestShouldGetCommonLetters(t *testing.T) {
	// given
	a1 := []string{"abc"}
	a2 := []string{"a", "b", "c"}
	a3 := []string{"ab", "ac"}
	a4 := []string{"a", "a", "a", "a"}
	a5 := []string{"b"}
	a6 := []string{"a;b;cd", "a;cd;g"}

	// then
	assert.ElementsMatch(t, []string{"a", "b", "c"}, SplitAndGetCommon(a1, ""))
	assert.ElementsMatch(t, []string{}, SplitAndGetCommon(a2, ""))
	assert.ElementsMatch(t, []string{"a"}, SplitAndGetCommon(a3, ""))
	assert.ElementsMatch(t, []string{"a"}, SplitAndGetCommon(a4, ""))
	assert.ElementsMatch(t, []string{"b"}, SplitAndGetCommon(a5, ""))
	assert.ElementsMatch(t, []string{"a", "cd"}, SplitAndGetCommon(a6, ";"))
}

func TestShouldGetArraysIntersection(t *testing.T) {
	assert.ElementsMatch(t, []string{"a", "b"}, GetArraysIntersection([]string{"a", "b", "c"}, []string{"a", "b", "d"}))
	assert.ElementsMatch(t, []string{"a"}, GetArraysIntersection([]string{"a"}, []string{"a"}))
	assert.ElementsMatch(t, []string{}, GetArraysIntersection([]string{"a"}, []string{}))
	assert.ElementsMatch(t, []string{}, GetArraysIntersection([]string{}, []string{}))
	assert.ElementsMatch(t, []string{}, GetArraysIntersection([]string{"a"}, []string{"b"}))
}

func TestShouldGetIntArraysIntersection(t *testing.T) {
	assert.ElementsMatch(t, []int{1, 2}, GetIntArraysIntersection([]int{1, 2, 3}, []int{1, 2, 4}))
	assert.ElementsMatch(t, []int{1}, GetIntArraysIntersection([]int{1}, []int{1}))
	assert.ElementsMatch(t, []int{}, GetIntArraysIntersection([]int{1}, []int{}))
	assert.ElementsMatch(t, []int{}, GetIntArraysIntersection([]int{}, []int{}))
	assert.ElementsMatch(t, []int{}, GetIntArraysIntersection([]int{1}, []int{2}))
}

func TestShouldConvertToInterfaceArray(t *testing.T) {
	// when
	stringArray := []string{"a", "b"}

	//then
	assert.Equal(t, []interface{}{"a", "b"}, ToInterfaceArray(stringArray))
}

func TestShouldGetMinAndMaxFromIntArray(t *testing.T) {
	// when
	min1, max1, _ := MinMax([]int{0, 1, 2, 3})
	min2, max2, _ := MinMax([]int{4, 1, -1, 3})
	_, _, err1 := MinMax(nil)
	_, _, err2 := MinMax([]int{})

	// then
	assert.Equal(t, 0, min1)
	assert.Equal(t, 3, max1)
	assert.Equal(t, -1, min2)
	assert.Equal(t, 4, max2)
	assert.Error(t, err1)
	assert.Error(t, err2)
}
