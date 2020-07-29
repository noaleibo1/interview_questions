package main

var treeLeftRightLeft = &Btree{
	Value: "d",
	Left:  nil,
	Right: nil,
}
var treeLeftRightRight = &Btree{
	Value: "e",
	Left:  nil,
	Right: nil,
}
var treeLeftRight = &Btree{
	Value: "c",
	Left:  treeLeftRightLeft,
	Right: treeLeftRightRight,
}
var treeLeft = &Btree{
	Value: "b",
	Left:  nil,
	Right: treeLeftRight,
}

var complexTree = &Btree{
	Value: "a",
	Left:  treeLeft,
	Right: nil,
}

var graphTreeNode = &Btree{
	Value: "g",
	Left:  nil,
	Right: nil,
}

var graphTreeLeft = &Btree{
	Value: "b",
	Left:  nil,
	Right: graphTreeNode,
}

var graphTree = &Btree{
	Value: "a",
	Left:  graphTreeLeft,
	Right: nil,
}

var singleNodeTree = &Btree{
	Value: "abc",
	Left:  nil,
	Right: nil,
}

var complexTreeSerializedHumanReadable = []string{"a", "b", "", "", "c", "", "", "", "", "d", "e"}
var complexTreeSerializedManualyForTests = "[[1100001] [1100010] [] [] [1100011] [] [] [] [] [1100100] [1100101]]"
var complexTreeSerialized = turnStringArrayToStringOfByteArrays(complexTreeSerializedHumanReadable)

var InvalidDeserializeInput = "I am not a tree"
