package main

import (
	"github.com/ruinnel/golang-poet/poet"
	"log"
	"reflect"
)

func main() {

	intType := reflect.TypeOf(1)
	stringType := reflect.TypeOf("")
	structName := "TestStruct"

	file := poet.NewFileSpec("./.out/test.go")

	file.AddImport(poet.ImportSpec{Package: "fmt"})
	file.AddImport(poet.ImportSpec{Package: "github.com/ruinnel/go-poet", Alias: "poet"})

	constants := poet.Constants{
		poet.Constant{Name: "MaxValue", Value: poet.Value{Value: 100}},
		poet.Constant{Name: "MinValue", Value: poet.Value{Value: 1 + 1i}},
		poet.Constant{Name: "DefaultName", Value: poet.Value{Value: "test"}},
		poet.Constant{Name: "DefaultExp", Value: poet.Value{Value: poet.Expression("1 + 1")}},
	}

	genericType1 := poet.ParameterizedType{Symbol: "T1", Approximation: false, Type: reflect.TypeOf(1.0)}
	interfaces := poet.Interfaces{
		poet.Interface{
			Name:               "Interface",
			ParameterizedTypes: []poet.ParameterizedType{genericType1},
			GenericTypes: []poet.GenericType{
				{Approximation: false, Type: reflect.TypeOf(1)},
				{Approximation: true, Type: reflect.TypeOf(1.9)},
			},
			Functions: poet.Functions{
				poet.Function{
					Interface: true,
					Receiver:  nil,
					Name:      "testFunc",
					Arguments: poet.Arguments{
						poet.Argument{Name: "arg1", Type: genericType1},
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
	}

	structs := poet.Structs{
		poet.Struct{
			Name: structName,
			ParameterizedTypes: []poet.ParameterizedType{
				genericType1,
				{Symbol: "T2", Approximation: false, Type: reflect.TypeOf(1)}},
			Variables: poet.Variables{
				poet.Variable{StructField: true, Names: []string{"a"}, Type: genericType1},
				poet.Variable{StructField: true, Names: []string{"b"}, Type: stringType},
				poet.Variable{StructField: true, Names: []string{"c", "d", "e"}, Type: stringType},
			},
		},
	}

	variables := poet.Variables{
		poet.Variable{Names: []string{"a"}, Type: stringType, Value: poet.Value{Value: ""}},
		poet.Variable{Names: []string{"b"}, Type: nil, Value: poet.Value{Value: poet.Expression("1 + 1")}},
		poet.Variable{Names: []string{"c", "d", "e"}, Type: stringType, Value: poet.Value{Value: "s"}},
	}

	functions := poet.Functions{
		poet.Function{
			Name: "Func1",
			//Receiver: &poet.Receiver{
			//	Name: "r",
			//	Type: poet.TypeName(structName),
			//},
			ParameterizedTypes: []poet.ParameterizedType{
				{Symbol: "T1", Approximation: false, Type: reflect.TypeOf(1.0)},
				{Symbol: "T2", Approximation: false, Type: reflect.TypeOf(1)}},
			Arguments: nil,
			Returns: poet.Returns{
				{Type: reflect.TypeOf(1)},
			},
			Statements: poet.Statements{
				poet.CodeBlock{Code: "a := 1 + 1"},
				poet.CodeBlock{Code: "return a"},
			},
		},
	}

	file.
		WithConstants(constants).
		WithInterfaces(interfaces).
		WithStructs(structs).
		WithVariables(variables).
		WithFunctions(functions)

	err := file.Write(false)
	if err != nil {
		log.Fatalf("error - %v", err)
	}
}
