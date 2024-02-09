package main

import (
	"fmt"
	"sample/normaibuilder/dto"
)

func main() {
	builder := dto.NewSampleDTOBuilder()
	builder.Data1("data1")
	builder.Data2("data2")
	builder.Data3("data3")
	result := builder.Build()

	fmt.Println(result)
}
