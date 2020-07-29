package main

import "fmt"

func main() {
	A := &Node{
		value:  "A",
		father: nil,
	}
	B:= &Node{
		value:  "B",
		father: A,
	}
	// C:= &Node{
	// 	value:  "C",
	// 	father: A,
	// }
	D:= &Node{
		value:  "D",
		father: B,
	}
	E:= &Node{
		value:  "E",
		father: D,
	}
	// F:= &Node{
	// 	value:  "F",
	// 	father: D,
	// }

	node := findFirstAncestor(E,E)
	fmt.Println(node)
}

type Node struct {
	value string
	father *Node
}

func findFirstAncestor(node1 *Node, node2 *Node) *Node {
	var ancestorsForNode1 []*Node
	var ancestorsForNode2 []*Node

	currentAncestor1 := node1.father
	currentAncestor2 := node2.father
	for currentAncestor2 != nil || currentAncestor1 != nil {
		if currentAncestor1 != nil {
			ancestorsForNode1 = append(ancestorsForNode1, currentAncestor1)
			if nodeInNodeList(currentAncestor1, ancestorsForNode2) {
				return currentAncestor1
			}
			currentAncestor1 = currentAncestor1.father
		}
		if currentAncestor2 != nil {
			ancestorsForNode2 = append(ancestorsForNode2, currentAncestor2)
			if nodeInNodeList(currentAncestor2, ancestorsForNode1) {
				return currentAncestor2
			}
			currentAncestor2 = currentAncestor2.father
		}
	}
	return nil
}

func nodeInNodeList(node *Node, list []*Node) bool {
	for i := 0; i < len(list); i++ {
		if list[i] == node {
			return true
		}
	}
	return false
}
