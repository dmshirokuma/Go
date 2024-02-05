package main

import (
	"fmt"

	"sample.com/fluentbuilder/sample"
)

func main() {
	sampleData := sample.Builder().Id(12345).Name("SampleName").Build()
	fmt.Println(sampleData)
}
