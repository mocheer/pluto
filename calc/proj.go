package calc

import "github.com/go-spatial/proj"

//
func ProjInverse(x, y float64) ([]float64, error) {
	return proj.Inverse(proj.WorldMercator, []float64{x, y})
}
