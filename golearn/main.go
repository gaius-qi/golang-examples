package main

import (
	"fmt"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/linear_models"
)

func main() {
	data, err := base.ParseCSVToInstances("./datasets/exams.csv", true)
	if err != nil {
		panic(err)
	}

	train, test := base.InstancesTrainTestSplit(data, 0.2)
	lr := linear_models.NewLinearRegression()
	if err := lr.Fit(train); err != nil {
		panic(err)
	}

	predictions, err := lr.Predict(test)
	if err != nil {
		panic(err)
	}

	fmt.Println("Linear Regression (information gain)")
	cf, err := evaluation.GetConfusionMatrix(test, predictions)
	if err != nil {
		panic(fmt.Errorf("Unable to get confusion matrix: %s", err.Error()))
	}

	fmt.Println(evaluation.GetSummary(cf))
	fmt.Println("----------------------")
	fmt.Println(evaluation.GetAccuracy(cf))
	fmt.Println("----------------------")
	fmt.Println(evaluation.ShowConfusionMatrix(cf))
	fmt.Println("----------------------")
	fmt.Println(evaluation.GetF1Score("177", cf))
	fmt.Println(evaluation.GetF1Score("152", cf))
	fmt.Println("----------------------")
}
