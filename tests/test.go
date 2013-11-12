package main

import (
	. "../../go-svmlight-algo"
	"fmt"
	. "github.com/apourchet/go-svmlight"
)

func main() {
	fmt.Println("Testing has begun.")
	f := ParseSVMFile("data/boxes.train")
	ns := Knn(f.Instances[0], f, 4)
	pred := KnnPredict(ns)
	fmt.Println(pred)
}
