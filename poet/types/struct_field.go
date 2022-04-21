package types

type StructField struct {
	Name string
	Type Type
	Tags StructTags
}

type StructFields []StructField

func (v StructFields) String() string {
	return GetTemplate().StructField(v)
}
