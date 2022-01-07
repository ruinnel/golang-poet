package poet

import (
	"go/format"
	"io/ioutil"
	"log"
	"strings"
)

type File struct {
	Name       string
	Package    Package
	Imports    Imports
	Constants  Constants
	Interfaces Interfaces
	Structs    Structs
	Variables  Variables
	Functions  Functions
}

func (f File) String() string {
	var lines []string
	// package
	lines = append(lines, f.Package.String())

	// imports
	lines = append(lines, f.Imports.String())

	// const
	lines = append(lines, f.Constants.String())

	// interface
	lines = append(lines, f.Interfaces.String())

	// struct
	lines = append(lines, f.Structs.String())

	// variable
	lines = append(lines, f.Variables.String())

	// functions
	lines = append(lines, f.Functions.String())

	return strings.Join(lines, "\n")
}

func (f File) Write(skipFormat bool) error {
	code := []byte(f.String())
	if skipFormat {
		return ioutil.WriteFile(f.Name, code, 0644)
	} else {
		formatted, err := format.Source(code)
		if err != nil {
			log.Println(string(code))
			log.Panicf("code format fail - %v", err)
		}
		return ioutil.WriteFile(f.Name, formatted, 0644)
	}
}
