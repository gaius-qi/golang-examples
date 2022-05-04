package main

import (
	"fmt"

	"github.com/xlab/treeprint"
)

func main() {
	// to add a custom root name use `treeprint.NewWithRoot()` instead
	tree1 := treeprint.NewWithRoot("tree1")

	// create a new branch in the root
	one := tree1.AddBranch("one")

	// add some nodes
	one.AddNode("subnode1").AddNode("subnode2")

	// create a new sub-branch
	one.AddBranch("two").
		AddNode("subnode1").AddNode("subnode2"). // add some nodes
		AddBranch("three").                      // add a new sub-branch
		AddNode("subnode1").AddNode("subnode2")  // add some nodes too

	// add one more node that should surround the inner branch
	one.AddNode("subnode3")

	// add a new node to the root
	tree1.AddNode("outernode")

	fmt.Println(tree1.String())

	tree2 := treeprint.NewWithRoot("tree2")

	tree2.AddBranch("three")

	fmt.Println(tree2.String())

	fmt.Printf("%s %s", tree1.String(), tree2.String())
}
