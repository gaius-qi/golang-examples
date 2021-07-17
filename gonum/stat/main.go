package main

import (
	"fmt"
	"math"
	"sort"

	"gonum.org/v1/gonum/stat"
)

func main() {
	xs := []float64{
		32.32, 56.98, 21.52, 44.32,
		55.63, 13.75, 43.47, 43.34,
		12.34,
	}
	fmt.Print("初始数组：")
	fmt.Println(xs)

	//平均值
	mean := stat.Mean(xs, nil)
	fmt.Print("平均值：")
	fmt.Println(mean)

	//加权平均值
	wmean := stat.Mean(xs, []float64{0.1, 0.2, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1})
	fmt.Print("加权平均值：")
	fmt.Println(wmean)

	//标准差
	variance := stat.Variance(xs, nil)
	fmt.Print("标准差：")
	fmt.Println(variance)

	//方差
	sqrt := math.Sqrt(variance)
	fmt.Print("方差：")
	fmt.Println(sqrt)

	//中位数
	sort.Float64s(xs) //数组排序
	fmt.Print("排序后数组：")
	fmt.Println(xs)

	media := stat.Quantile(0.5, stat.Empirical, xs, nil)
	fmt.Print("中位数：")
	fmt.Println(media)
}
