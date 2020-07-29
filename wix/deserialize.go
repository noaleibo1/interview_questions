package main

import (
	"strings"
)

type Btree struct {
	Value string
	Left  *Btree
	Right *Btree
}

// Add getters and setters
// Test empty complexTree
// convert value to bit

func Deserialize(str string) *Btree {
	// Regex not working well :(
	if !isValidDeserializeInput(str) {
		panic(ErrorInvalidTree)
	}
	if str != "" {
		str = str[1 : len(str)-1]
	}
	array := strings.Split(str, " ")
	return RecursivelyBuildTreeFromArray(&Btree{}, array, 0)
}

// Regex not working :(
func isValidDeserializeInput(str string) bool {
	if str == "" {
		return true
	} else if str[0] != '[' || str[len(str)-1] != ']' {
		return false
	}
	return true
	// pattern := regexp.MustCompile(RegexForDeserializationInput)
	// return pattern.Match([]byte(str))
}

func RecursivelyBuildTreeFromArray(tree *Btree, array []string, index int) *Btree {
	if index > len(array)-1 || array[index] == "" {
		return nil
	}
	tree = &Btree{
		Value: string(array[index]),
		Left:  RecursivelyBuildTreeFromArray(&Btree{}, array, 2*index+1),
		Right: RecursivelyBuildTreeFromArray(&Btree{}, array, 2*index+2),
	}
	return tree
}


