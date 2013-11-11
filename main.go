package svmalgo

import (
	"fmt"
	"github.com/apourchet/go-svmlight"
)

func TestFunction() {
	f := svmlight.ParseSVMFile("data/boxes.train")
	fmt.Println(f.CountLabels(-1))
}
