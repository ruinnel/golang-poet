package main

import (
	"github.com/ruinnel/golang-poet/poet/builder"
	"github.com/ruinnel/golang-poet/poet/types"
	"log"
	"reflect"
)

func main() {

	intType := reflect.TypeOf(1)
	stringType := reflect.TypeOf("")
	structName := "TestStruct"

	file := builder.NewFile("./.out/test.go")

	importBuilder := builder.NewImportBuilder()
	importBuilder.AddImport("fmt").
		AddImport("github.com/ruinnel/go-poet", "poet")

	file.WithImports(importBuilder)

	constantBuilder := builder.NewConstantBuilder()
	constantBuilder.AddConstant("MaxValue", types.Value{Value: 199}).
		AddConstant("MinValue", types.Value{Value: 1 + 1i}).
		AddConstant("DefaultName", types.Value{Value: "test"}).
		AddConstant("DefaultExp", types.Value{Value: types.Expression("1 + 1")})

	file.WithConstants(constantBuilder)
	//constants := types.Constants{
	//	types.Constant{Name: "MaxValue", Value: types.Value{Value: 100}},
	//	types.Constant{Name: "MinValue", Value: types.Value{Value: 1 + 1i}},
	//	types.Constant{Name: "DefaultName", Value: types.Value{Value: "test"}},
	//	types.Constant{Name: "DefaultExp", Value: types.Value{Value: types.Expression("1 + 1")}},
	//}

	genericType1 := types.ParameterizedType{Symbol: "T1", Approximation: false, Type: reflect.TypeOf(1.0)}
	interfaces := types.Interfaces{
		types.Interface{
			Name:               "Interface",
			ParameterizedTypes: []types.ParameterizedType{genericType1},
			GenericTypes: []types.GenericType{
				{Approximation: false, Type: reflect.TypeOf(1)},
				{Approximation: true, Type: reflect.TypeOf(1.9)},
			},
			Functions: types.Functions{
				types.Function{
					Interface: true,
					Receiver:  nil,
					Name:      "testFunc",
					Arguments: types.Arguments{
						types.Argument{Name: "arg1", Type: genericType1},
						types.Argument{Name: "arg2", Type: stringType},
						types.Argument{Name: "arg3", Type: intType},
						types.Argument{Name: "arg4", Type: intType},
					},
					Returns: types.Returns{
						types.Return{Name: "a", Type: stringType},
						types.Return{Name: "b", Type: stringType},
					},
				},
			},
		},
	}

	structs := types.Structs{
		types.Struct{
			Name: structName,
			ParameterizedTypes: []types.ParameterizedType{
				genericType1,
				{Symbol: "T2", Approximation: false, Type: reflect.TypeOf(1)}},
			Fields: types.StructFields{
				types.StructField{
					Name: "a",
					Type: genericType1,
					Tags: types.StructTags{
						{Name: "tag1", Parameters: map[string]string{"param": "value", "param2": ""}},
						{Name: "tag2", Parameters: map[string]string{"param": "value", "param2": ""}},
					}},
				types.StructField{Name: "b", Type: stringType},
				types.StructField{Name: "c", Type: stringType},
			},
		},
	}

	variables := types.Variables{
		types.Variable{Names: []string{"a"}, Type: stringType, Value: types.Value{Value: ""}},
		types.Variable{Names: []string{"b"}, Type: nil, Value: types.Value{Value: types.Expression("1 + 1")}},
		types.Variable{Names: []string{"c", "d", "e"}, Type: stringType, Value: types.Value{Value: "s"}},
	}

	functions := types.Functions{
		types.Function{
			Name: "Func1",
			//Receiver: &poet.Receiver{
			//	Name: "r",
			//	Type: poet.TypeName(structName),
			//},
			ParameterizedTypes: []types.ParameterizedType{
				{Symbol: "T1", Approximation: false, Type: reflect.TypeOf(1.0)},
				{Symbol: "T2", Approximation: false, Type: reflect.TypeOf(1)}},
			Arguments: nil,
			Returns: types.Returns{
				{Type: reflect.TypeOf(1)},
			},
			Statements: types.Statements{
				types.CodeBlock{Code: "a := 1 + 1"},
				types.CodeBlock{Code: "return a"},
			},
		},
	}

	file.
		//WithConstants(constants).
		WithInterfaces(interfaces).
		WithStructs(structs).
		WithVariables(variables).
		WithFunctions(functions)

	err := file.Write(false)
	if err != nil {
		log.Fatalf("error - %v", err)
	}
}
