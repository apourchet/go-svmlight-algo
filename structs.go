package svmalgo

import (
	. "github.com/apourchet/go-svmlight"
)

type SimMatrix struct {
	Matrix   [][]float64
	FileName string
}

type Neighbor struct {
	From, To SVMInstance
	Sim      float64
}

type NeighborList struct {
	Neighbors []Neighbor
}

func (l NeighborList) Len() int {
	return len(l.Neighbors)
}

func (l NeighborList) Less(i, j int) bool {
	return l.Neighbors[i].Sim < l.Neighbors[j].Sim
}

func (l NeighborList) Swap(i, j int) {
	tmp := l.Neighbors[i]
	l.Neighbors[i] = l.Neighbors[j]
	l.Neighbors[j] = tmp
}

type CompFunction func(instance SVMInstance, attribute int) bool

type InfoGainFunc func(instances []SVMInstance, attributes []int) int
