package data_structure

import (
	"fmt"
	"math"
)

type binaryNode struct {
	key       int
	leftNode  *binaryNode
	rightNode *binaryNode
}

type binaryTree struct {
	root *binaryNode
}

func NewBinaryTree(key int) *binaryTree {
	binaryTree := &binaryTree{}
	binaryTree.init(key)
	return binaryTree
}

func (binaryTree *binaryTree) init(key int) *binaryTree {
	binaryTree.root = &binaryNode{
		key:       key,
		leftNode:  nil,
		rightNode: nil,
	}
	return binaryTree
}

func (binaryTree *binaryTree) append(binaryNodeX *binaryNode, key int) *binaryTree {

	if binaryNodeX.key > key {
		if binaryNodeX.leftNode == nil {
			binaryNodeX.leftNode = &binaryNode{
				key:       key,
				leftNode:  nil,
				rightNode: nil,
			}
		} else {
			return binaryTree.append(binaryNodeX.leftNode, key)
		}
	}

	if binaryNodeX.key < key {
		if binaryNodeX.rightNode == nil {
			binaryNodeX.rightNode = &binaryNode{
				key:       key,
				leftNode:  nil,
				rightNode: nil,
			}
		} else {
			return binaryTree.append(binaryNodeX.rightNode, key)
		}
	}

	return binaryTree
}

func (binaryTree *binaryTree) Insert(key int) {
	if binaryTree.root == nil {
		binaryTree.init(key)
		return
	}
	binaryTree.append(binaryTree.root, key)
}

func (binaryTree *binaryTree) PreOrder() {
	binaryTree.xPreOrder(binaryTree.root)
}

func (binaryTree *binaryTree) InOrder() {
	binaryTree.xInOrder(binaryTree.root)
}

func (binaryTree *binaryTree) PostOrder() {
	binaryTree.xPostOrder(binaryTree.root)
}

func (binaryTree *binaryTree) PrintTree() {
	maxLevel := binaryTree.maxLevel(binaryTree.root)
	var nodeList []*binaryNode
	nodeList = append(nodeList, binaryTree.root)
	binaryTree.xPrintTree(nodeList, 1, maxLevel)
}

func (binaryTree *binaryTree) xPreOrder(binaryNodeX *binaryNode) {
	if binaryNodeX == nil {
		return
	}
	fmt.Print(binaryNodeX.key, " ")
	binaryTree.xPreOrder(binaryNodeX.leftNode)
	binaryTree.xPreOrder(binaryNodeX.rightNode)
}

func (binaryTree *binaryTree) xInOrder(binaryNodeX *binaryNode) {
	if binaryNodeX == nil {
		return
	}
	binaryTree.xInOrder(binaryNodeX.leftNode)
	fmt.Print(binaryNodeX.key, " ")
	binaryTree.xInOrder(binaryNodeX.rightNode)
}

func (binaryTree *binaryTree) xPostOrder(binaryNodeX *binaryNode) {
	if binaryNodeX == nil {
		return
	}
	binaryTree.xPostOrder(binaryNodeX.leftNode)
	binaryTree.xPostOrder(binaryNodeX.rightNode)
	fmt.Print(binaryNodeX.key, " ")
}

func (binaryTree *binaryTree) xPrintTree(nodeList []*binaryNode, level int, maxLevel int) {

	if len(nodeList) == 0 || isNull(nodeList) {
		return
	}

	floor := maxLevel - level
	endLine := int(math.Pow(2, math.Max(float64(floor-1), 0)))
	firstSpace := int(math.Pow(2, float64(floor))) - 1
	betweenSpace := int(math.Pow(2, float64(floor+1))) - 1

	printWhite(firstSpace)

	var nodeListX []*binaryNode
	for _, binaryNodeX := range nodeList {
		if binaryNodeX != nil {
			fmt.Print(binaryNodeX.key)
			nodeListX = append(nodeListX, binaryNodeX.leftNode)
			nodeListX = append(nodeListX, binaryNodeX.rightNode)
		} else {
			nodeListX = append(nodeListX, nil)
			nodeListX = append(nodeListX, nil)
			fmt.Print(" ")
		}
		printWhite(betweenSpace)
	}

	fmt.Println("")

	for i := 1; i <= endLine; i++ {
		for j := 0; j < len(nodeList); j++ {

			printWhite(firstSpace - i)

			if nodeList[j] == nil {
				printWhite(2*endLine*i + 1)
				continue
			}

			if nodeList[j].leftNode != nil {
				fmt.Print("/")
			} else {
				printWhite(1)
			}

			printWhite(2*i - 1)

			if nodeList[j].rightNode != nil {
				fmt.Print("\\")
			} else {
				printWhite(1)
			}

			printWhite(2*endLine - i)
		}
		fmt.Println("")
	}
	binaryTree.xPrintTree(nodeListX, level+1, maxLevel)
}

func (binaryTree *binaryTree) maxLevel(binaryNodeX *binaryNode) int {
	if binaryNodeX == nil {
		return 0
	}
	return int(math.Max(float64(binaryTree.maxLevel(binaryNodeX.leftNode)), float64(binaryTree.maxLevel(binaryNodeX.rightNode)))) + 1
}

func isNull(nodeList []*binaryNode) bool {
	for _, binaryNodeX := range nodeList {
		if binaryNodeX != nil {
			return false
		}
	}
	return true
}

func printWhite(count int) {
	for i := 0; i < count; i++ {
		fmt.Print(" ")
	}
}
