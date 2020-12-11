package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldGetFirstExample(t *testing.T) {
	// given
	input := "16\n10\n15\n5\n1\n11\n7\n19\n6\n12\n4"
	input2 := "28\n33\n18\n42\n31\n14\n46\n20\n48\n47\n24\n23\n49\n45\n19\n38\n39\n11\n1\n32\n25\n35\n8\n17\n7\n9\n4\n2\n34\n10\n3"

	// when
	result := firstPart(input)
	result2 := firstPart(input2)

	// then
	assert.Equal(t, 35, result)
	assert.Equal(t, 220, result2)
}

func TestShouldGetSecondExample(t *testing.T) {
	// given
	input := "16\n10\n15\n5\n1\n11\n7\n19\n6\n12\n4"
	input2 := "28\n33\n18\n42\n31\n14\n46\n20\n48\n47\n24\n23\n49\n45\n19\n38\n39\n11\n1\n32\n25\n35\n8\n17\n7\n9\n4\n2\n34\n10\n3"
	input3 := "10\n6\n4\n7\n1\n5"
	input4 := "4\n11\n7\n8\n1\n6\n5"
	input5 := "3\n1\n6\n2"
	input6 := "17\n6\n10\n5\n13\n7\n1\n4\n12\n11\n14"

	// when
	//result2 := secondPart(input2)
	result := secondPart(input)
	result2 := secondPart(input2)
	result3 := secondPart(input3)
	result4 := secondPart(input4)
	result5 := secondPart(input5)
	result6 := secondPart(input6)

	// then
	assert.Equal(t, 8, result)
	assert.Equal(t, 19208, result2)
	assert.Equal(t, 4, result3)
	assert.Equal(t, 7, result4)
	assert.Equal(t, 4, result5)
	assert.Equal(t, 28, result6)
}