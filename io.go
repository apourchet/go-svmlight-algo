package svmalgo

import (
	"bufio"
	"bytes"
	"os"
)

type ReadFunction func(line string, lineNumber int)

func WriteFile(output *bytes.Buffer, fileName string) {
	sysfile, _ := os.Create(fileName)
	defer func() {
		if err := sysfile.Close(); err != nil {
			panic(err)
		}
	}()
	sysfile.Write(output.Bytes())
}

func ReadLines(fileName string, readFunction ReadFunction) {
	fi, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()
	reader := bufio.NewReader(fi)
	lineNumber := 0
	for buffer, _, err := reader.ReadLine(); err == nil; buffer, _, err = reader.ReadLine() {
		readFunction(string(buffer), lineNumber)
		lineNumber++
	}
}
