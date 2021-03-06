package types

import (
	"bytes"
	"log"
	"text/template"
)

const packageTemplate = `package {{.}}`
const singleImportTemplate = `import {{.Name}}`
const multipleImportTemplate = `import (
{{- range $index, $import := . }}
  {{if $import.Alias}}{{printf "%s " $import.Alias}}{{end -}}
  "{{$import.Package}}"{{end}}
)
`

const singleConstantTemplate = `const {{.Name}} = {{.Value}}`
const multipleConstantTemplate = `const (
{{range $index, $con := . -}}
  {{printf "  %s = %s\n" $con.Name $con.Value -}}
{{- end -}}
)`

const singleArgumentTemplate = `{{.Name}} {{.Type}}`
const multipleArgumentTemplate = `{{$args := .}}{{- range $index, $ret := $args -}}{{if $index}}, {{end}}{{.Name}} {{- if $args.NeedType $index }} {{.Type}}{{end}}{{end}}`

const singleReturnTemplate = `{{if .Name}}({{.Name}}{{end}} {{.Type}}{{if .Name}}){{end}}`
const multipleReturnTemplate = `{{$returns := .}}({{- range $index, $ret := $returns -}}{{if $index}}, {{end}}{{.Name}} {{- if $returns.NeedType $index }} {{.Type}}{{end}}{{end}})`

const variableTemplate = `var {{ range $index, $name := .Names -}}
{{- if $index}}, {{end}}{{$name -}}{{- end -}}
{{- if .Type}} {{.Type}}{{end}}
{{- if .Value}} = {{.Value}}{{end -}}
`

const structTagTemplate = "`" + `
{{- range $index, $tag := . -}}
{{if $index}} {{end}}{{$tag.Name}}:"{{$tag.Parameters}}"
{{- end -}}
` + "`"

const structFieldTemplate = `
{{- .Name}} {{- if .Type}} {{.Type}}{{end}}{{if .Tags}} {{.Tags}}{{end}}
`

const structTemplate = `type {{.Name}}{{.ParameterizedTypes}} struct {
{{.Fields}}
}`

const functionTemplate = `
{{- if not .Interface }}func {{end -}}
{{- if .Receiver }}({{.Receiver.Name}} {{.Receiver.Type}}){{end -}}
{{.Name}}{{.ParameterizedTypes}}({{.Arguments}}) {{.Returns -}}
{{- if .Statements}} {
{{.Statements}}
}
{{end -}}
`

const interfaceTemplate = `
type {{.Name}}{{.ParameterizedTypes}} interface {
  {{ .GenericTypes }}
  {{ .Functions }}
}`

const genericTypeTemplate = `{{- range $index, $type := . -}}
  {{- if $index}} |{{end}}{{- if $type.Approximation}}~{{end}}{{$type.Type.Name -}}
{{- end }}`

const singleParameterizedTypeTemplate = `[{{.Symbol}} {{.Type}}]`
const multipleParameterizedTypeTemplate = `[{{$types := .}}{{- range $index, $ret := $types -}}{{if $index}}, {{end}}{{.Symbol}} {{- if $types.NeedType $index }} {{.Type}}{{end}}{{end}}]`

type Template struct {
	packageTemplate                   *template.Template
	singleImportTemplate              *template.Template
	multipleImportTemplate            *template.Template
	singleConstantTemplate            *template.Template
	multipleConstantTemplate          *template.Template
	singleArgumentTemplate            *template.Template
	multipleArgumentTemplate          *template.Template
	singleReturnTemplate              *template.Template
	multipleReturnTemplate            *template.Template
	structTagTemplate                 *template.Template
	structFieldTemplate               *template.Template
	structTemplate                    *template.Template
	variableTemplate                  *template.Template
	functionTemplate                  *template.Template
	interfaceTemplate                 *template.Template
	genericTypeTemplate               *template.Template
	singleParameterizedTypeTemplate   *template.Template
	multipleParameterizedTypeTemplate *template.Template
}

var instance *Template

func init() {
	packageTmpl := template.New("package")
	packageTmpl, err := packageTmpl.Parse(packageTemplate)
	if err != nil {
		log.Panicf("package template parse fail. %v", err)
	}

	singleImportTmpl := template.New("singleImportTemplate")
	singleImportTmpl, err = singleImportTmpl.Parse(singleImportTemplate)
	if err != nil {
		log.Panicf("single import template parse fail. %v", err)
	}

	multipleImportTmpl := template.New("multipleImportTemplate")
	multipleImportTmpl, err = multipleImportTmpl.Parse(multipleImportTemplate)
	if err != nil {
		log.Panicf("multiple import template parse fail. %v", err)
	}

	singleConstantTmpl := template.New("singleConstantTemplate")
	singleConstantTmpl, err = singleConstantTmpl.Parse(singleConstantTemplate)
	if err != nil {
		log.Panicf("single constant template parse fail. %v", err)
	}

	multipleConstantTmpl := template.New("multipleConstantTemplate")
	multipleConstantTmpl, err = multipleConstantTmpl.Parse(multipleConstantTemplate)
	if err != nil {
		log.Panicf("multiple constant template parse fail. %v", err)
	}

	singleArgumentTmpl := template.New("singleArgumentTemplate")
	singleArgumentTmpl, err = singleArgumentTmpl.Parse(singleArgumentTemplate)
	if err != nil {
		log.Panicf("single argument template parse fail. %v", err)
	}

	multipleArgumentTmpl := template.New("multipleArgumentTemplate")
	multipleArgumentTmpl, err = multipleArgumentTmpl.Parse(multipleArgumentTemplate)
	if err != nil {
		log.Panicf("multiple argument template parse fail. %v", err)
	}

	singleReturnTmpl := template.New("singleReturnTemplate")
	singleReturnTmpl, err = singleReturnTmpl.Parse(singleReturnTemplate)
	if err != nil {
		log.Panicf("single return template parse fail. %v", err)
	}

	multipleReturnTmpl := template.New("multipleReturnTemplate")
	multipleReturnTmpl, err = multipleReturnTmpl.Parse(multipleReturnTemplate)
	if err != nil {
		log.Panicf("multiple return template parse fail. %v", err)
	}

	structTagTmpl := template.New("structTag")
	structTagTmpl, err = structTagTmpl.Parse(structTagTemplate)
	if err != nil {
		log.Panicf("struct tag template parse fail. %v", err)
	}

	structFieldTmpl := template.New("structField")
	structFieldTmpl, err = structFieldTmpl.Parse(structFieldTemplate)
	if err != nil {
		log.Panicf("struct field template parse fail. %v", err)
	}

	structTmpl := template.New("struct")
	structTmpl, err = structTmpl.Parse(structTemplate)
	if err != nil {
		log.Panicf("struct template parse fail. %v", err)
	}

	variableTmpl := template.New("variable")
	variableTmpl, err = variableTmpl.Parse(variableTemplate)
	if err != nil {
		log.Panicf("variable template parse fail. %v", err)
	}

	functionTmpl := template.New("function")
	functionTmpl, err = functionTmpl.Parse(functionTemplate)
	if err != nil {
		log.Panicf("function template parse fail. %v", err)
	}

	interfaceTmpl := template.New("interface")
	interfaceTmpl, err = interfaceTmpl.Parse(interfaceTemplate)
	if err != nil {
		log.Panicf("interface template parse fail. %v", err)
	}

	genericTypeTmpl := template.New("genericType")
	genericTypeTmpl, err = genericTypeTmpl.Parse(genericTypeTemplate)
	if err != nil {
		log.Panicf("genericType template parse fail. %v", err)
	}

	singleParameterizedTypeTmpl := template.New("singleParameterizedTypeTemplate")
	singleParameterizedTypeTmpl, err = singleParameterizedTypeTmpl.Parse(singleParameterizedTypeTemplate)
	if err != nil {
		log.Panicf("single ParameterizedType template parse fail. %v", err)
	}

	multipleParameterizedTypeTmpl := template.New("multipleParameterizedTypeTemplate")
	multipleParameterizedTypeTmpl, err = multipleParameterizedTypeTmpl.Parse(multipleParameterizedTypeTemplate)
	if err != nil {
		log.Panicf("multiple ParameterizedType template parse fail. %v", err)
	}

	instance = &Template{
		packageTemplate:                   packageTmpl,
		singleImportTemplate:              singleImportTmpl,
		multipleImportTemplate:            multipleImportTmpl,
		singleConstantTemplate:            singleConstantTmpl,
		multipleConstantTemplate:          multipleConstantTmpl,
		singleArgumentTemplate:            singleArgumentTmpl,
		multipleArgumentTemplate:          multipleArgumentTmpl,
		singleReturnTemplate:              singleReturnTmpl,
		multipleReturnTemplate:            multipleReturnTmpl,
		structTagTemplate:                 structTagTmpl,
		structFieldTemplate:               structFieldTmpl,
		structTemplate:                    structTmpl,
		variableTemplate:                  variableTmpl,
		functionTemplate:                  functionTmpl,
		interfaceTemplate:                 interfaceTmpl,
		genericTypeTemplate:               genericTypeTmpl,
		singleParameterizedTypeTemplate:   singleParameterizedTypeTmpl,
		multipleParameterizedTypeTemplate: multipleParameterizedTypeTmpl,
	}
}

func GetTemplate() *Template {
	return instance
}

func (t Template) Package(pkg Package) string {
	code := bytes.NewBufferString("")
	err := t.packageTemplate.Execute(code, string(pkg))
	if err != nil {
		log.Panicf("package template execute fail. %v", err)
	}

	return code.String()
}

func (t Template) Import(imp Import) string {
	code := bytes.NewBufferString("")
	err := t.singleImportTemplate.Execute(code, imp)
	if err != nil {
		log.Panicf("import template execute fail. %v", err)
	}

	return code.String()
}

func (t Template) Imports(imps Imports) string {
	code := bytes.NewBufferString("")
	var err error
	if len(imps) == 1 {
		err = t.singleImportTemplate.Execute(code, imps[0])
	} else if len(imps) > 1 {
		err = t.multipleImportTemplate.Execute(code, imps)
	}
	if err != nil {
		log.Panicf("import template execute fail. %v", err)
	}

	return code.String()
}

func (t Template) Constant(con Constant) string {
	code := bytes.NewBufferString("")
	err := t.singleConstantTemplate.Execute(code, con)
	if err != nil {
		log.Panicf("constant template execute fail. %v", err)
	}

	return code.String()
}

func (t Template) Constants(cons Constants) string {
	code := bytes.NewBufferString("")
	var err error
	if len(cons) == 1 {
		err = t.singleConstantTemplate.Execute(code, cons[0])
	} else if len(cons) > 1 {
		err = t.multipleConstantTemplate.Execute(code, cons)
	}
	if err != nil {
		log.Panicf("constant template execute fail. %v", err)
	}

	return code.String()
}

func (t Template) Argument(arg Argument) string {
	code := bytes.NewBufferString("")
	err := t.singleArgumentTemplate.Execute(code, arg)
	if err != nil {
		log.Panicf("argument template execute fail. %v", err)
	}

	return code.String()
}

func (t Template) Arguments(args Arguments) string {
	code := bytes.NewBufferString("")
	var err error
	if len(args) == 1 {
		err = t.singleArgumentTemplate.Execute(code, args[0])
	} else if len(args) > 1 {
		err = t.multipleArgumentTemplate.Execute(code, args)
	}
	if err != nil {
		log.Panicf("argument template execute fail. %v", err)
	}

	return code.String()
}

func (t Template) Return(ret Return) string {
	code := bytes.NewBufferString("")
	err := t.singleReturnTemplate.Execute(code, ret)
	if err != nil {
		log.Panicf("return template execute fail. %v", err)
	}

	return code.String()
}

func (t Template) Returns(returns Returns) string {
	code := bytes.NewBufferString("")
	var err error
	if len(returns) == 1 {
		err = t.singleReturnTemplate.Execute(code, returns[0])
	} else if len(returns) > 1 {
		err = t.multipleReturnTemplate.Execute(code, returns)
	}
	if err != nil {
		log.Panicf("return template execute fail. %v", err)
	}

	return code.String()
}

func (t Template) StructTags(tags StructTags) string {
	code := bytes.NewBufferString("")
	err := t.structTagTemplate.Execute(code, tags)
	if err != nil {
		log.Panicf("struct tag template execute fail. %v", err)
	}

	return code.String()
}

func (t Template) StructField(field StructField) string {
	code := bytes.NewBufferString("")
	code.WriteString("  ")
	err := t.structFieldTemplate.Execute(code, field)
	if err != nil {
		log.Panicf("struct field template execute fail. %v", err)
	}

	return code.String()
}

func (t Template) StructFields(fields StructFields) string {
	code := bytes.NewBufferString("")
	for _, field := range fields {
		code.WriteString("  ")
		err := t.structFieldTemplate.Execute(code, field)
		if err != nil {
			log.Panicf("struct field template execute fail. %v", err)
		}
	}

	return code.String()
}

func (t Template) Struct(st Struct) string {
	code := bytes.NewBufferString("")
	err := t.structTemplate.Execute(code, st)
	if err != nil {
		log.Panicf("struct template execute fail. %v", err)
	}

	return code.String()
}

func (t Template) Structs(structs Structs) string {
	code := bytes.NewBufferString("")
	for _, st := range structs {
		err := t.structTemplate.Execute(code, st)
		if err != nil {
			log.Panicf("struct template execute fail. %v", err)
		}
	}

	return code.String()
}

func (t Template) Variable(v Variable) string {
	code := bytes.NewBufferString("")
	err := t.variableTemplate.Execute(code, v)
	if err != nil {
		log.Panicf("variable template execute fail. %v", err)
	}
	code.WriteString("\n")

	return code.String()
}

func (t Template) Variables(vars Variables) string {
	code := bytes.NewBufferString("")
	for idx, v := range vars {
		err := t.variableTemplate.Execute(code, v)
		if err != nil {
			log.Panicf("variable template execute fail. %v", err)
		}
		if idx < len(vars)-1 {
			code.WriteString("\n")
		}
	}

	return code.String()
}

func (t Template) Function(fun Function) string {
	code := bytes.NewBufferString("")
	err := t.functionTemplate.Execute(code, fun)
	if err != nil {
		log.Panicf("function template execute fail. %v", err)
	}

	return code.String()
}

func (t Template) Functions(funs Functions) string {
	code := bytes.NewBufferString("")
	for _, fun := range funs {
		err := t.functionTemplate.Execute(code, fun)
		if err != nil {
			log.Panicf("function template execute fail. %v", err)
		}
	}

	return code.String()
}

func (t Template) Interface(in Interface) string {
	code := bytes.NewBufferString("")
	err := t.interfaceTemplate.Execute(code, in)
	if err != nil {
		log.Panicf("interface template execute fail. %v", err)
	}

	return code.String()
}

func (t Template) Interfaces(ins Interfaces) string {
	code := bytes.NewBufferString("")
	for _, in := range ins {
		err := t.interfaceTemplate.Execute(code, in)
		if err != nil {
			log.Panicf("interface template execute fail. %v", err)
		}
	}

	return code.String()
}

func (t Template) GenericTypes(genericTypes GenericTypes) string {
	code := bytes.NewBufferString("")
	err := t.genericTypeTemplate.Execute(code, genericTypes)
	if err != nil {
		log.Panicf("genericType template execute fail. %v", err)
	}

	return code.String()
}

func (t Template) ParameterizedType(tp ParameterizedType) string {
	code := bytes.NewBufferString("")
	err := t.singleParameterizedTypeTemplate.Execute(code, tp)
	if err != nil {
		log.Panicf("parameterizedType template execute fail. %v", err)
	}

	return code.String()
}

func (t Template) ParameterizedTypes(types ParameterizedTypes) string {
	code := bytes.NewBufferString("")
	var err error
	if len(types) == 1 {
		err = t.singleParameterizedTypeTemplate.Execute(code, types[0])
	} else if len(types) > 1 {
		err = t.multipleParameterizedTypeTemplate.Execute(code, types)
	}
	if err != nil {
		log.Panicf("parameterizedType template execute fail. %v", err)
	}

	return code.String()
}
