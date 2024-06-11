package main

import (
	treesearch "github.com/brharrelldev/take-home-challenge/tree-search"
)

// see tree for implementation of the tree.
func main() {

	bst := &treesearch.BinarySearchTree{}

	bst.Insert(1, []byte("test"))
	bst.Insert(2, []byte("test"))
	bst.Insert(5, []byte("test"))
	bst.Insert(8, []byte("test"))
	bst.Insert(3, []byte("test"))
	bst.Insert(4, []byte("test"))
	bst.Insert(1, []byte("test"))

	bst.String()

}
