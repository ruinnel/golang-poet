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

func (f *FileBuilder) Build() string {
	var lines []string
	// package
	lines = append(lines, f.pkg.String())

	// imports
	if f.imports != nil {
		lines = append(lines, f.imports.Build())
	}

	// const
	lines = append(lines, f.constants.Build())

	// interface
	lines = append(lines, f.interfaces.String())

	// struct
	lines = append(lines, f.structs.String())

	// variable
	lines = append(lines, f.variables.String())

	// functions
	lines = append(lines, f.functions.String())

	return strings.Join(lines, "\n\n")
}

func (f *FileBuilder) Write(skipFormat bool) error {
	code := []byte(f.Build())
	if skipFormat {
		return ioutil.WriteFile(f.name, code, 0644)
	} else {
		formatted, err := format.Source(code)
		if err != nil {
			log.Println(string(code))
			log.Panicf("code format fail - %v", err)
		}
		return ioutil.WriteFile(f.name, formatted, 0644)
	}
}

func (f *FileBuilder) WithPackage(pkg string) *FileBuilder {
	f.pkg = types.Package(pkg)
	return f
}

func (f *FileBuilder) WithImports(imports *ImportBuilder) *FileBuilder {
	f.imports = imports
	return f
}

func (f *FileBuilder) WithConstants(constants *ConstantBuilder) *FileBuilder {
	f.constants = constants
	return f
}

func (f *FileBuilder) WithInterfaces(interfaces types.Interfaces) *FileBuilder {
	f.interfaces = interfaces
	return f
}

func (f *FileBuilder) AddInterface(intf types.Interface) *FileBuilder {
	f.interfaces = append(f.interfaces, intf)
	return f
}

func (f *FileBuilder) WithStructs(structs types.Structs) *FileBuilder {
	f.structs = structs
	return f
}

func (f *FileBuilder) AddStruct(strt types.Struct) *FileBuilder {
	f.structs = append(f.structs, strt)
	return f
}

func (f *FileBuilder) WithVariables(variables types.Variables) *FileBuilder {
	f.variables = variables
	return f
}

func (f *FileBuilder) AddVariable(variable types.Variable) *FileBuilder {
	f.variables = append(f.variables, variable)
	return f
}

func (f *FileBuilder) WithFunctions(functions types.Functions) *FileBuilder {
	f.functions = functions
	return f
}

func (f *FileBuilder) AddFunction(fun types.Function) *FileBuilder {
	f.functions = append(f.functions, fun)
	return f
}
