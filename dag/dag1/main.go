package main

import (
	"fmt"
	"time"

	"github.com/heimdalr/dag"
)

func main() {

	// initialize a new graph
	d := dag.NewDAG()

	// init three vertices
	var ids []string
	for i := 0; i < 2000; i++ {
		id, _ := d.AddVertex(i)
		ids = append(ids, id)
	}

	for i := 0; i < len(ids)-1; i++ {
		if err := d.AddEdge(ids[i], ids[i+1]); err != nil {
			panic(err)
		}
	}

	startTime := time.Now()
	if err := d.AddEdge(ids[4], ids[100]); err != nil {
		panic(err)
	}

	elapsedTime := time.Since(startTime) / time.Millisecond
	fmt.Printf("2000 in %d ms", elapsedTime)
}
