package main

import (
	"fmt"
	"time"

	"d7y.io/dragonfly/v2/pkg/dag"
)

func main() {
	d := dag.NewDAG()

	// init three vertices
	var ids []string
	for i := 0; i < 1000000; i++ {
		if err := d.AddVertex(fmt.Sprint(i), i); err != nil {
			panic(err)
		}
		ids = append(ids, fmt.Sprint(i))
	}

	edgeCount := 10
	for i := 0; i < len(ids)-1; i++ {
		if i+edgeCount > len(ids)-1 {
			break
		}

		for n := 1; n < edgeCount; n++ {
			if err := d.AddEdge(ids[i], ids[i+n]); err != nil {
				panic(err)
			}
		}
	}

	startTime := time.Now()
	if err := d.AddEdge(ids[4], ids[400]); err != nil {
		panic(err)
	}

	elapsedTime := time.Since(startTime) / time.Millisecond
	fmt.Printf("100000 in %d ms", elapsedTime)
}
