package kriging

import "github.com/liuvigongzuoshi/go-kriging/ordinarykriging"

// TrainSpherical 训练数据集 使用 spherical 模型对数据集进行训练，返回的是一个variogram对象。
func TrainSpherical(values, x, y []float64) (*ordinarykriging.Variogram, error) {
	return ordinarykriging.NewOrdinary(values, x, y).Train(ordinarykriging.Spherical, 0, 100)
}

// TrainGaussian 训练数据集 使用 gaussian 模型模型对数据集进行训练，返回的是一个variogram对象。
func TrainGaussian(values, x, y []float64) (*ordinarykriging.Variogram, error) {
	return ordinarykriging.NewOrdinary(values, x, y).Train(ordinarykriging.Gaussian, 0, 100)
}

// TrainExponential 训练数据集 使用 exponential 模型模型对数据集进行训练，返回的是一个variogram对象。
func TrainExponential(values, x, y []float64) (*ordinarykriging.Variogram, error) {
	return ordinarykriging.NewOrdinary(values, x, y).Train(ordinarykriging.Exponential, 0, 100)
}
