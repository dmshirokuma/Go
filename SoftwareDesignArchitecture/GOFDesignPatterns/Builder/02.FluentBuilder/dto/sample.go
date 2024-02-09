package dto

type sampleDTO struct {
	Data1, Data2, Data3 string
}

type sampleDTOBuilder struct {
	sampleDTO *sampleDTO
}

// interfaceを使用することによって警告を避けると共にsampleDTOそのものを非公開にできる=builderの使用を強制できる
type SampleDTOBuilder interface {
	Data1(data1 string) *sampleDTOBuilder
	Data2(data2 string) *sampleDTOBuilder
	Data3(data3 string) *sampleDTOBuilder
	Build() *sampleDTO
}

func NewSampleDTOBuilder() SampleDTOBuilder {
	return &sampleDTOBuilder{&sampleDTO{}}
}

func (b *sampleDTOBuilder) Build() *sampleDTO {
	return b.sampleDTO
}

func (b *sampleDTOBuilder) Data1(data1 string) *sampleDTOBuilder {
	b.sampleDTO.Data1 = data1
	return b
}

func (b *sampleDTOBuilder) Data2(data2 string) *sampleDTOBuilder {
	b.sampleDTO.Data2 = data2
	return b
}

func (b *sampleDTOBuilder) Data3(data3 string) *sampleDTOBuilder {
	b.sampleDTO.Data3 = data3
	return b
}
