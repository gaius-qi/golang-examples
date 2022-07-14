package main

import (
	"fmt"

	"github.com/heimdalr/dag"
)

func main() {

	// initialize a new graph
	d := dag.NewDAG()

	// init three vertices
	v1, _ := d.AddVertex(1)
	v2, _ := d.AddVertex(2)
	v3, _ := d.AddVertex(3)
	v4, _ := d.AddVertex(4)

	// add the above vertices and connect them with two edges
	if err := d.AddEdge(v1, v2); err != nil {
		panic(err)
	}

	if err := d.AddEdge(v1, v3); err != nil {
		panic(err)
	}

	if err := d.AddEdge(v3, v4); err != nil {
		panic(err)
	}

	if err := d.AddEdge(v4, v1); err != nil {
		panic(err)
	}

	// describe the graph
	fmt.Print(d.String())
}
