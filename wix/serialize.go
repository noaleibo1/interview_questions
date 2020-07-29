package main

import (
	"errors"
	"fmt"
)

func Serialize(tree *Btree) string {
	array, err := RecursivelyWalkTree(tree, []string{}, 0, []*Btree{})
	if err != nil {
		panic(err.Error())
	} else {
		return fmt.Sprintf("%v", array)
	}
}

func RecursivelyWalkTree(tree *Btree, array []string, index int, nodesCovered []*Btree) ([]string, error) {
	if tree == nil {
		return array, nil
	} else if nodeAlreadyVisited(tree, nodesCovered) {
		return nil, errors.New(ErrorTreeIsGraph)
	}
	nodesCovered = append(nodesCovered, tree)
	array = enlargeArrayIfNeeded(array, index)
	array[index] = fmt.Sprintf("%b", []byte(tree.Value))
	var err error
	if array, err = RecursivelyWalkTree(tree.Left, array, 2*index+1, nodesCovered); err != nil {
		return nil, err
	}
	if array, err = RecursivelyWalkTree(tree.Right, array, 2*index+2, nodesCovered); err != nil {
		return nil, err
	}
	return array, nil
}

func nodeAlreadyVisited(node *Btree, nodesCovered []*Btree) bool {
	for i := 0; i < len(nodesCovered); i++ {
		if node == nodesCovered[i] {
			return true
		}
	}
	return false
}

func enlargeArrayIfNeeded(array []string, index int) []string {
	if index < len(array) {
		return array
	} else {
		for i := len(array); i <= index; i++ {
			array = append(array, fmt.Sprintf("%b", []byte("")))
		}
	}
	return array
}