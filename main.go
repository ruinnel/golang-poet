package main

import (
	"github.com/ruinnel/golang-poet/poet"
	"log"
	"reflect"
)

func main() {

	intType := reflect.TypeOf(1)
	stringType := reflect.TypeOf("")

	file := poet.File{
		Name:    "./.out/test.go",
		Package: poet.Package{Name: "main"},
		Imports: poet.Imports{
			poet.Import{Package: "fmt"},
			poet.Import{Package: "github.com/ruinnel/go-poet", Alias: "poet"},
		},
		Constants: poet.Constants{
			poet.Constant{Name: "MaxValue", Value: poet.Value{Value: 100}},
			poet.Constant{Name: "MinValue", Value: poet.Value{Value: 1 + 1i}},
			poet.Constant{Name: "DefaultName", Value: poet.Value{Value: "test"}},
			poet.Constant{Name: "DefaultExp", Value: poet.Value{Value: poet.Expression("1 + 1")}},
		},
		Interfaces: poet.Interfaces{
			poet.Interface{
				Name: "Interface",
				Functions: poet.Functions{
					poet.Function{
						Interface: true,
						Receiver:  nil,
						Name:      "testFunc",
						Arguments: poet.Arguments{
							poet.Argument{Name: "arg1", Type: stringType},
							poet.Argument{Name: "arg2", Type: stringType},
							poet.Argument{Name: "arg3", Type: intType},
							poet.Argument{Name: "arg4", Type: intType},
						},
						Returns: poet.Returns{
							poet.Return{Name: "a", Type: stringType},
							poet.Return{Name: "b", Type: stringType},
						},
					},
				},
			},
		},
		Structs: poet.Structs{
			poet.Struct{Name: "TestStruct", Variables: poet.Variables{
				poet.Variable{StructField: true, Names: []string{"a"}, Type: &stringType},
				poet.Variable{StructField: true, Names: []string{"b"}, Type: &stringType},
				poet.Variable{StructField: true, Names: []string{"c", "d", "e"}, Type: &stringType},
			}},
		},
		Variables: poet.Variables{
			poet.Variable{Names: []string{"a"}, Type: &stringType, Value: poet.Value{Value: ""}},
			poet.Variable{Names: []string{"b"}, Type: nil, Value: poet.Value{Value: poet.Expression("1 + 1")}},
			poet.Variable{Names: []string{"c", "d", "e"}, Type: &stringType, Value: poet.Value{Value: "s"}},
		},
		Functions: poet.Functions{
			poet.Function{
				Name:      "Func1",
				Arguments: nil,
				Returns:   nil,
				Statements: poet.Statements{
					poet.CodeBlock{Code: "a := 1 + 1"},
					poet.CodeBlock{Code: "return a"},
				},
			},
		},
	}
	err := file.Write(false)
	if err != nil {
		log.Fatalf("error - %v", err)
	}
}
