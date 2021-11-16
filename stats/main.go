package main

import (
	"fmt"

	"github.com/montanaflynn/stats"
)

func main() {
	data := []float64{1, 2, 3, 4, 5, 6}
	v1, _ := stats.StandardDeviationPopulation(data)
	fmt.Println(v1)

	v2, _ := stats.StandardDeviationSample(data)
	fmt.Println(v2)

	v3, _ := stats.StandardDeviation(data)
	fmt.Println(v3)

	v4, _ := stats.PercentileNearestRank(data, 30)
	fmt.Println(v4)

	v5, _ := stats.Percentile(data, 80)
	fmt.Println(v5)

	v6, _ := stats.PopulationVariance(data)
	fmt.Println(v6)

	v7 := stats.NormFit(data)
	fmt.Println(v7)
	fmt.Println(v7[0]-3*v7[1], v7[0]+3*v7[1])
}
