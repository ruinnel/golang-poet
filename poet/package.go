package poet

type Package string

func (p Package) String() string {
	return GetTemplate().Package(p)
}
