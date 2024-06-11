package treesearch

import (
	"fmt"
	"sync"
)

type Node struct {
	Key         int
	Level       int
	Data        []byte
	Count       int
	Left, Right *Node
}

type BinarySearchTree struct {
	Root *Node
	lock sync.RWMutex
}

func NewNode(item int, data []byte) *Node {

	node := &Node{}
	node.Key = item
	node.Data = data
	node.Left = nil
	node.Right = nil
	node.Level = 1
	return node

}

func (bst *BinarySearchTree) Insert(key int, data []byte) {
	bst.lock.Lock()
	defer bst.lock.Unlock()

	if bst.Root == nil {
		bst.Root = NewNode(key, data)
	} else {
		nodeInsert(bst.Root, NewNode(key, data))
	}

}

func (bst *BinarySearchTree) String() {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	fmt.Println("==================================")
	printTree(bst.Root, 0)
	fmt.Println("===================================")

}

func printTree(n *Node, level int) {

	if n != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "      "
		}
		format += "----["
		level++
		printTree(n.Left, level)
		fmt.Printf(format+"%d\n", n.Key)
		printTree(n.Right, level)

	}

}

func nodeInsert(node, newNode *Node) {

	if node.Key == newNode.Key {
		newNode.Count++
		nodeInsert(node.Left, newNode)
	}
	if newNode.Key < node.Key {
		if node.Left == nil {

			node.Left = newNode
			node.Left.Level++
		} else {
			node.Left.Level++
			nodeInsert(node.Left, newNode)
		}
	} else {
		if node.Right == nil {
			node.Right = newNode
			node.Right.Level++
		} else {
			node.Right.Level++
			nodeInsert(node.Right, newNode)
		}
	}

}

func InOrder(node *Node) {

	if node != nil {
		InOrder(node.Left)
		fmt.Printf("%+v ", node)
		InOrder(node.Right)
	}

}

func CheckDuplicateId(node *Node) int {

	if node == nil {
		return 0
	}

	if node.Count > 0 {
		return node.Level
	}

	//ugly if statement but is the best exit condition I could find for recursive traversal
	if node.Right != nil {
		return CheckDuplicateId(node.Right)

	} else {
		return CheckDuplicateId(node.Left)
	}

}
