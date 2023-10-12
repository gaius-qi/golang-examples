package main

import (
	"fmt"

	"github.com/gogo/protobuf/proto"
	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
)

func main() {
	temps := prometheus.NewSummary(prometheus.SummaryOpts{
		Name:       "temps",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	})

	for i := 0; i < 200; i++ {
		temps.Observe(float64(i))
	}

	metric := &dto.Metric{}
	temps.Write(metric)
	fmt.Println(proto.MarshalTextString(metric))
}
