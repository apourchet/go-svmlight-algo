package svmalgo

import (
	"bytes"
	"fmt"
	. "github.com/apourchet/go-svmlight"
	"sort"
)

type Centroid SVMInstance

func DotProduct(instance1, instance2 SVMInstance) float64 {
	sum := 0.
	for featureId, value := range instance1.Features {
		sum += float64(value) * float64(instance2.Features[featureId])
	}
	return sum
}

func CalculateSims(instance SVMInstance, file *SVMFile) *SimMatrix {
	arr := make([][]float64, 1)
	arr[0] = make([]float64, len(file.Instances))
	for index, otherInstance := range file.Instances {
		instancePtr := &(otherInstance)
		instancePtr.Norm()
		arr[0][index] = DotProduct(instance, *instancePtr)
	}
	ret := SimMatrix{arr, file.FileName}
	return &ret
}

// This has space complexity n^2. Be aware of that before it 
// starts making a 1,000,000 x 1,000,000 matrix of float64's
func ConstructSimMatrix(file *SVMFile) *SimMatrix {
	arr := make([][]float64, len(file.Instances))
	for _, instance := range file.Instances {
		arr = append(arr, CalculateSims(instance, file).Matrix[0])
	}
	ret := SimMatrix{arr, file.FileName}
	return &ret
}

func ExportSimMatrix(mat *SimMatrix, fileName string) {
	output := bytes.NewBufferString("")
	for _, arr := range mat.Matrix {
		for _, sim := range arr {
			output.WriteString(fmt.Sprintf("%f ", sim))
		}
		output.WriteString("\n")
	}
	WriteFile(output, fileName)
}

func ImportSimMatrix(fileName string) *SimMatrix {
	arr := [][]float64{}

	ret := SimMatrix{arr, fileName}
	return &ret
}

func Knn(instance SVMInstance, file *SVMFile, k int) []Neighbor {
	mat := CalculateSims(instance, file)
	return KnnCached(mat, instance, file, k)
}

func KnnCached(mat *SimMatrix, instance SVMInstance, file *SVMFile, k int) []Neighbor {
	neighbors := NeighborList{}
	for i, sim := range mat.Matrix[0] {
		newNeighbor := Neighbor{instance, file.Instances[i], sim}
		neighbors.Neighbors = append(neighbors.Neighbors, newNeighbor)
	}
	sort.Sort(neighbors)
	return neighbors.Neighbors[:k]
}

func KnnPredict(neighbors []Neighbor) int {
	labelCounts := make(map[int]int)
	maxLabelCount := 0
	maxLabel := -1
	for _, n := range neighbors {
		if _, pres := labelCounts[n.To.Label]; pres {
			labelCounts[n.To.Label]++
		} else {
			labelCounts[n.To.Label] = 1
		}
	}
	for label, count := range labelCounts {
		if count > maxLabelCount {
			maxLabel = label
			maxLabelCount = count
		}
	}
	return maxLabel
}

func BuildCentroids(file *SVMFile) []Centroid {
	arr := []Centroid{}
	labelCounts := make(map[int]int)
	for _, i := range file.Instances {
		if _, pres := labelCounts[i.Label]; pres {
			labelCounts[i.Label]++
		} else {
			labelCounts[i.Label] = 1
		}
	}
	return arr
}
