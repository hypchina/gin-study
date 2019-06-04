package data_structure

import (
	"testing"
)

func TestNewBinaryTree(t *testing.T) {

	binaryTree := NewBinaryTree(4)
	binaryTree.Insert(2)
	binaryTree.Insert(5)
	binaryTree.Insert(7)
	binaryTree.Insert(6)
	binaryTree.Insert(8)

	/*binaryTree.PreOrder()
	fmt.Println("---preOrder--")

	binaryTree.InOrder()
	fmt.Println("---inOrder--")*/

	binaryTree.PrintTree()
	//fmt.Println("---postOrder--")
}
