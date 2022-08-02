package main

import (
	"dsa/src/data-structures/binary-tree"
	"fmt"
)

func main() {
	// bst := &bTree.BinaryTree{Root: &bTree.Node{Value: 1}}
	// bst.InsertIterative(7)
	// bst.InsertIterative(10)
	// bst.InsertIterative(3)
	// bst.InsertIterative(2)
	// bst.InsertIterative(5)
	// bst.InsertIterative(12)

	//      7
	//   3     10
	// 2   5      12
	// bst.Root = DeleteNodeRecursive(bst.Root, 2)
	// fmt.Println(bst.Root.MaxValue())
	// bTree.Insert(bst.Root, 2)
	// bTree.Insert(bst.Root, 3)
	// bTree.Insert(bst.Root, 4)
	// bTree.Insert(bst.Root, 5)
	// bTree.Insert(bst.Root, 6)
	// bTree.Insert(bst.Root, 7)
	// bTree.Insert(bst.Root, 8)
	// BST := bTree.BTreeToBST(bst.Root)
	// fmt.Println(bTree.InorderIterative(BST.Root))
	var node *bTree.AVLNode

	node = bTree.InsertAVL(node, 1)
	node = bTree.InsertAVL(node, 2)
	node = bTree.InsertAVL(node, 3)
	fmt.Println(bTree.AVLInorderIterative(node))
	fmt.Printf("%+v", node)
}
