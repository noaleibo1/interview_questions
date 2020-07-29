package main

import "fmt"

func areTwoTreesEqual(tree1 *Btree, tree2 *Btree) bool {
	if (tree1 == nil && tree2 != nil) || (tree1 != nil && tree2 == nil) {
		return false
	} else if tree1 == nil && tree2 == nil {
		return true
	} else if tree1.Value != tree2.Value {
		return false
	}
	return areTwoTreesEqual(tree1.Left, tree2.Left) && areTwoTreesEqual(tree1.Right, tree2.Right)
}

func turnStringArrayToStringOfByteArrays(str []string) string{
	var result []string
	for _, s :=range str {
		result = append(result, fmt.Sprintf("%b",[]byte(s)))
	}
	return fmt.Sprintf("%s", result)
}