package main

import (
	"fmt"
	"math/big"
)

//Arcs holds an arcset which is indexed by parents and then has an array of children. Children dominate parents in our trees.
type Arcs map[int][]int

type BinaryTree struct {
	root *Node
}

type Node struct {
	vertex          int
	children        []*Node
	canonicalNumber big.Int
}

func main() {
	//standard data structure for edges. Parent, then children array
	testData := Arcs{
		1: []int{2, 3},
		2: []int{6, 5},
		3: []int{7},
		5: []int{9, 2},
	}

	fmt.Println(newTree(testData))
}

/*
func encodeTree(arcSet []Arc) Node {
	arcSet, leaves := findLeaves(arcSet)
	return tree
}*/

func newTree(edges Arcs) BinaryTree {
	root, children := findRoot(edges)
	var rootNode Node //declare root node
	rootNode.vertex = root

	rootNode.children = makeNode(children, edges)

	fmt.Println(children)
	//rootNode.children = children
	tree := BinaryTree{&rootNode}

	return tree
}

func makeNode(input int, edges Arcs) *Node {
	var newNode Node
	newNode.vertex = input
	for vertex := range edges[input] {
		newNode.children = makeNode(edges[input], edges)
	}
	nodes = append(nodes, &newNode)
	fmt.Println(vertex)

	return &newNode
}

func findRoot(arcSet Arcs) (int, []int) {
	var root int
	var children []int
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
			children = arcSet[root]
		}
	}

	return root, children
}

/*
func findLeaves(arcSet []Arc) ([]Arc, []Arc) {
	fmt.Println(arcSet)
	var leaves []Arc
	var notLeaves []Arc
	for _, arc := range arcSet {
		leafFlag := true
		for _, checkArc := range arcSet {
			if arc.from == checkArc.to {
				leafFlag = false
			}

		}
		if leafFlag == true {
			leaves = append(leaves, arc)
		} else {
			notLeaves = append(notLeaves, arc)
		}
	}
	return notLeaves, leaves
}
*/
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
	fmt.Println(z)
	shiftBy(z, int64(1)) //appends zero to the end of the bitstring
	fmt.Println(z)
	shift := int64(int64(z.BitLen()))
	front := big.NewInt(1)
	shiftBy(front, shift)
	z.Add(z, front)
	fmt.Println(z)
}

func shiftBy(x *big.Int, shift int64) {
	factor := new(big.Int)
	factor.Exp(big.NewInt(2), big.NewInt(shift), nil)
	x.Mul(x, factor)
}
