package types

import (
	"go/format"
	"io/ioutil"
	"log"
	"strings"
)

type FileSpec struct {
	name       string
	pkg        Package
	imports    Imports
	constants  Constants
	interfaces Interfaces
	structs    Structs
	variables  Variables
	functions  Functions
}

func NewFileSpec(name string) *FileSpec {
	return &FileSpec{name: name, pkg: "main"}
}

func (f *FileSpec) String() string {
	var lines []string
	// package
	lines = append(lines, f.pkg.String())

	// imports
	lines = append(lines, f.imports.String())

	// const
	lines = append(lines, f.constants.String())

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

func (f *FileSpec) Write(skipFormat bool) error {
	code := []byte(f.String())
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

func (f *FileSpec) WithPackage(pkg string) *FileSpec {
	f.pkg = Package(pkg)
	return f
}

func (f *FileSpec) WithImports(imports Imports) *FileSpec {
	f.imports = imports
	return f
}

func (f *FileSpec) AddImport(imp ImportSpec) *FileSpec {
	f.imports = append(f.imports, imp)
	return f
}

func (f *FileSpec) WithConstants(constants Constants) *FileSpec {
	f.constants = constants
	return f
}

func (f *FileSpec) AddConstant(constant Constant) *FileSpec {
	f.constants = append(f.constants, constant)
	return f
}

func (f *FileSpec) WithInterfaces(interfaces Interfaces) *FileSpec {
	f.interfaces = interfaces
	return f
}

func (f *FileSpec) AddInterface(intf Interface) *FileSpec {
	f.interfaces = append(f.interfaces, intf)
	return f
}

func (f *FileSpec) WithStructs(structs Structs) *FileSpec {
	f.structs = structs
	return f
}

func (f *FileSpec) AddStruct(strt Struct) *FileSpec {
	f.structs = append(f.structs, strt)
	return f
}

func (f *FileSpec) WithVariables(variables Variables) *FileSpec {
	f.variables = variables
	return f
}

func (f *FileSpec) AddVariable(variable Variable) *FileSpec {
	f.variables = append(f.variables, variable)
	return f
}

func (f *FileSpec) WithFunctions(functions Functions) *FileSpec {
	f.functions = functions
	return f
}

func (f *FileSpec) AddFunction(fun Function) *FileSpec {
	f.functions = append(f.functions, fun)
	return f
}
