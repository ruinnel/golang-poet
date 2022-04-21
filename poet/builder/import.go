package builder

import (
	"github.com/ruinnel/golang-poet/poet/types"
)

type ImportBuilder struct {
	imports types.Imports
}

func NewImportBuilder() *ImportBuilder {
	var imports types.Imports
	return &ImportBuilder{imports}
}

func (b *ImportBuilder) AddImport(packageName string, args ...string) *ImportBuilder {
	alias := ""
	if len(args) > 0 {
		alias = args[0]
	}
	b.imports = append(b.imports, types.Import{Package: packageName, Alias: alias})
	return b
}

func (b *ImportBuilder) WithImports(imports ...types.Import) *ImportBuilder {
	b.imports = imports
	return b
}

func (b ImportBuilder) Build() string {
	return b.imports.String()
}
