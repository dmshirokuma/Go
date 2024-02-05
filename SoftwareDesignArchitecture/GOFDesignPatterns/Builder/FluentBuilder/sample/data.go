package sample

type sampleData struct {
	Id   int
	Name string
}

func newSampleData(builder *sampleDataBuilderParams) *sampleData {
	return &sampleData{
		Id:   builder.id,
		Name: builder.name,
	}
}

type sampleDataBuilderParams struct {
	id   int
	name string
}

func Builder() *sampleDataBuilderParams {
	return &sampleDataBuilderParams{}
}

func (b *sampleDataBuilderParams) Id(newId int) *sampleDataBuilderParams {
	b.id = newId
	return b
}

func (b *sampleDataBuilderParams) Name(newName string) *sampleDataBuilderParams {
	b.name = newName
	return b
}

func (b *sampleDataBuilderParams) Build() *sampleData {
	return newSampleData(b)
}
