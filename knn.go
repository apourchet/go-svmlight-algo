package svmalgo

import (
	. "github.com/apourchet/go-svmlight"
)

type SimMatrix struct {
	Matrix   [][]float64
	FileName string
}

type Neighbor struct {
	From, To  SVMInstance
	Sim       float64
	NormedSim float64
}

type Centroid SVMInstance

var IsSimMatrixSet = false

func CalculateSims(instance *SVMInstance, file *SVMFile) *SimMatrix {
	arr := make([][]float64, 1)
	// TODO implement this

	ret := SimMatrix{arr, file.FileName}
	return &ret
}

// This has space complexity n^2. Be aware of that before it 
// starts making a 1,000,000 x 1,000,000 matrix of float64's
func ConstructSimMatrix(file *SVMFile) *SimMatrix {
	arr := make([][]float64, len(file.Instances))
	// TODO implement this

	ret := SimMatrix{arr, file.FileName}
	return &ret
}

func SetSimMatrix(mat *SimMatrix) {

}

func ExportSimMatrix(fileName string) {

}

func ImportSimMatrix(fileName string) *SimMatrix {
	arr := [][]float64{}

	ret := SimMatrix{arr, fileName}
	return &ret
}

func Knn(instance *SVMInstance, file *SVMFile) []SVMInstance {
	return []SVMInstance{}
}

func BuildCentroids(file *SVMFile) []Centroid {
	arr := []Centroid{}
	return arr
}
