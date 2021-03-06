package types

type StructField struct {
	Name string
	Type Type
	Tags StructTags
}

type StructFields []StructField

func (v StructField) String() string {
	return GetTemplate().StructField(v)
}

func (v StructFields) String() string {
	return GetTemplate().StructFields(v)
}
