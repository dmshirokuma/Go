package main

import (
	"fmt"

	"sample.com/fluentbuilder/dto"
)

func main() {
	result := dto.NewSampleDTOBuilder().
		Data1("data1").
		Data2("data2").
		Data3("data3").
		Build()
	fmt.Println(result)
}
