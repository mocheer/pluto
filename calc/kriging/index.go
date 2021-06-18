package kriging

import "github.com/liuvigongzuoshi/go-kriging/ordinarykriging"

// Train 训练数据集
func Train(values, x, y []float64) (*ordinarykriging.Variogram, error) {
	sigma2 := 0.0
	alpha := 100.0
	ordinaryKriging := ordinarykriging.NewOrdinary(values, x, y)
	return ordinaryKriging.Train(ordinarykriging.Spherical, sigma2, alpha)
}
