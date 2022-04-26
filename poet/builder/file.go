package builder

import (
	"github.com/ruinnel/golang-poet/poet/types"
	"go/format"
	"io/ioutil"
	"log"
	"strings"
)

type FileBuilder struct {
	name       string
	pkg        types.Package
	imports    Builder
	constants  Builder
	interfaces types.Interfaces
	structs    types.Structs
	variables  types.Variables
	functions  types.Functions
}

func NewFile(name string) *FileBuilder {
	return &FileBuilder{name: name, pkg: "main"}
}

func (b *FileBuilder) Build() string {
	var lines []string
	// package
	lines = append(lines, b.pkg.String())

	// imports
	if b.imports != nil {
		lines = append(lines, b.imports.Build())
	}

	// const
	lines = append(lines, b.constants.Build())

	// interface
	lines = append(lines, b.interfaces.String())

	// struct
	lines = append(lines, b.structs.String())

	// variable
	lines = append(lines, b.variables.String())

	// functions
	lines = append(lines, b.functions.String())

	return strings.Join(lines, "\n\n")
}

func (b *FileBuilder) Write(skipFormat bool) error {
	code := []byte(b.Build())
	if skipFormat {
		return ioutil.WriteFile(b.name, code, 0644)
	} else {
		formatted, err := format.Source(code)
		if err != nil {
			log.Println(string(code))
			log.Panicf("code format fail - %v", err)
		}
		return ioutil.WriteFile(b.name, formatted, 0644)
	}
}

func (b *FileBuilder) WithPackage(pkg string) *FileBuilder {
	b.pkg = types.Package(pkg)
	return b
}

func (b *FileBuilder) WithImports(imports *ImportBuilder) *FileBuilder {
	b.imports = imports
	return b
}

func (b *FileBuilder) WithConstants(constants *ConstantBuilder) *FileBuilder {
	b.constants = constants
	return b
}

func (b *FileBuilder) WithInterfaces(interfaces types.Interfaces) *FileBuilder {
	b.interfaces = interfaces
	return b
}

func (b *FileBuilder) AddInterface(intf types.Interface) *FileBuilder {
	b.interfaces = append(b.interfaces, intf)
	return b
}

func (b *FileBuilder) WithStructs(structs types.Structs) *FileBuilder {
	b.structs = structs
	return b
}

func (b *FileBuilder) AddStruct(strt types.Struct) *FileBuilder {
	b.structs = append(b.structs, strt)
	return b
}

func (b *FileBuilder) WithVariables(variables types.Variables) *FileBuilder {
	b.variables = variables
	return b
}

func (b *FileBuilder) AddVariable(variable types.Variable) *FileBuilder {
	b.variables = append(b.variables, variable)
	return b
}

func (b *FileBuilder) WithFunctions(functions types.Functions) *FileBuilder {
	b.functions = functions
	return b
}

func (b *FileBuilder) AddFunction(fun types.Function) *FileBuilder {
	b.functions = append(b.functions, fun)
	return b
}
