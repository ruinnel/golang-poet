package poet

type Package struct {
	Name string
}

func (p Package) String() string {
	return GetTemplate().Package(p)
}
