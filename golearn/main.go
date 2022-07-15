package main

import (
	"fmt"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/linear_models"
)

func main() {
	train, test := base.InstancesTrainTestSplit(nil, 0.2)
	lr := linear_models.NewLinearRegression()
	if err := lr.Fit(train); err != nil {
		panic(err)
	}
	fmt.Println(lr.Predict(test))
}
