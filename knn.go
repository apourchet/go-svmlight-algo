package svmalgo

import (
	. "github.com/apourchet/go-svmlight"
)

func CalculateSimilarities(instance *SVMInstance, file SVMFile) []float64 {
	return []float64{}
}

func ConstructSimilarityMatrix(file *SVMFile) *[][]float64 {
	return [][]float64{}
}

func SetSimilarityMatrix(mat *[][]float64) {

}

func ExportSimilarityMatrix(fileName string) {

}

func Knn(instance *SVMInstance, file *SVMFile) []SVMInstance {
	return []SVMInstance{}
}

func BuildCentroids(file *SVMFile) []SVMInstance {
	return []SVMInstance{}
}
