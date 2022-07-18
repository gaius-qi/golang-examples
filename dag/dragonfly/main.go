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
	for i := 0; i < 50000; i++ {
		if err := d.AddVertex(fmt.Sprint(i), i); err != nil {
			panic(err)
		}
		ids = append(ids, fmt.Sprint(i))
	}

	for i := 0; i < len(ids)-1; i++ {
		if err := d.AddEdge(ids[i], ids[i+1]); err != nil {
			panic(err)
		}
	}

	startTime := time.Now()
	if err := d.AddEdge(ids[4], ids[400]); err != nil {
		panic(err)
	}

	elapsedTime := time.Since(startTime) / time.Millisecond
	fmt.Printf("100000 in %d ms", elapsedTime)
}
