package main

import (
	"fmt"
	"math/big"
)

//Arcs holds an arcset which is indexed by parents and then has an array of children. Children dominate parents in our trees.
type Arcs map[int][]int

type BinaryTree *Node

type Node struct {
	vertex          int
	children        []*Node
	canonicalNumber *big.Int
}

func main() {
	//standard data structure for edges. Parent, then children array
	testData := Arcs{
		1: []int{2, 3},
		2: []int{6, 5},
		3: []int{7},
		5: []int{9, 4},
	}

	testData2 := Arcs{
		0:  []int{17, 23},
		17: []int{111, 14},
		23: []int{7},
		14: []int{34, 42},
	}

	tree := newTree(testData)
	scoreNode(tree)
	fmt.Println(tree.canonicalNumber)
	tree2 := newTree(testData2)
	scoreNode(tree2)
	fmt.Println(tree2.canonicalNumber)
}

/*
func encodeTree(arcSet []Arc) Node {
	arcSet, leaves := findLeaves(arcSet)
	return tree
}*/

func scoreTree(tree BinaryTree) {
	scoreNode(tree)
}

func scoreNode(inputNode *Node) {
	if inputNode.children == nil {
		inputNode.canonicalNumber = big.NewInt(2)
	} else {
		for _, child := range inputNode.children {
			scoreNode(child)
		}
		if len(inputNode.children) == 2 {
			concatCannonical(inputNode.children[0].canonicalNumber, inputNode.children[1].canonicalNumber, inputNode.canonicalNumber)
		} else {
			concatCannonical(inputNode.children[0].canonicalNumber, big.NewInt(0), inputNode.canonicalNumber)
		}
	}
}

func newTree(edges Arcs) BinaryTree {
	root := findRoot(edges)
	rootNode := makeNode(root, edges)
	tree := BinaryTree(rootNode)

	return tree
}

func makeNode(input int, edges Arcs) *Node {
	var newNode Node
	newNode.vertex = input
	newNode.canonicalNumber = big.NewInt(0)
	if edges[input] == nil {
		return &newNode
	}
	for _, vertex := range edges[input] {
		newNode.children = append(newNode.children, makeNode(vertex, edges))
	}
	return &newNode
}

func findRoot(arcSet Arcs) int {
	var root int
	for parent := range arcSet {
		rootFlag := true
		for _, checkArc := range arcSet {
			for i := range checkArc {
				if checkArc[i] == parent {
					rootFlag = false
					break
				}
			}
		}
		if rootFlag == true {
			root = parent
		}
	}

	return root
}

func concatCannonical(x, y, out *big.Int) {
	diff := new(big.Int)
	diff.Sub(x, y)
	if diff.Sign() == -1 {
		concatHelper(y, x, out)
	} else {
		concatHelper(y, x, out)
	}
}

func concatHelper(x, y, z *big.Int) {
	shiftBy(x, int64(y.BitLen()))
	z.Add(x, y)
	shiftBy(z, int64(1)) //appends zero to the end of the bitstring
	shift := int64(int64(z.BitLen()))
	front := big.NewInt(1)
	shiftBy(front, shift)
	z.Add(z, front)
}

func shiftBy(x *big.Int, shift int64) {
	factor := new(big.Int)
	factor.Exp(big.NewInt(2), big.NewInt(shift), nil)
	x.Mul(x, factor)
}
