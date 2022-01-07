package poet

import (
	"bytes"
	"log"
	"text/template"
)

const packageTemplate = `package {{.Name}}`
const singleImportTemplate = `import {{.Name}}`
const multipleImportTemplate = `import (
  {{range $index, $import := . -}}
    {{- $import.Alias}} "{{$import.Package}}"
  {{end -}}
)
`

const singleConstantTemplate = `const {{.Name}} = {{.Value}}`
const multipleConstantTemplate = `const (
  {{range $index, $con := . -}}
    {{- $con.Name}} = {{$con.Value}}
  {{end -}}
)`

const singleArgumentTemplate = `{{.Name}} {{.Type}}`
const multipleArgumentTemplate = `{{$args := .}}{{- range $index, $ret := $args -}}{{if $index}}, {{end}}{{.Name}} {{- if $args.NeedType $index }} {{.Type}}{{end}}{{end}}`

const singleReturnTemplate = `{{if .Name}}({{.Name}}{{end}} {{.Type}}{{if .Name}}){{end}}`
const multipleReturnTemplate = `{{$returns := .}}({{- range $index, $ret := $returns -}}{{if $index}}, {{end}}{{.Name}} {{- if $returns.NeedType $index }} {{.Type}}{{end}}{{end}})`

const variableTemplate = `
{{- if not .StructField}}var {{end -}}
{{- range $index, $name := .Names -}}
  {{- if $index}} ,{{end}}{{$name -}}
{{- end -}}
{{- if .Type}} {{.Type}}{{end -}}
{{- if not .StructField}}
  {{- if .Value}} = {{.Value}}{{end -}}
{{- end -}}`

const structTemplate = `type {{.Name}} struct {
  {{.Variables}}
}`

const functionTemplate = `
{{- if not .Interface }}func {{end -}}
{{- if .Receiver }}({{.Receiver.Name}} {{.Receiver.Type}}){{end -}}
{{.Name}}({{.Arguments}}) {{.Returns -}}
{{- if .Statements}} {
  {{.Statements}}
}
{{end -}}
`

const interfaceTemplate = `
type {{.Name}} interface {
  {{ .Functions }}
}`

type Template struct {
	packageTemplate          *template.Template
	singleImportTemplate     *template.Template
	multipleImportTemplate   *template.Template
	singleConstantTemplate   *template.Template
	multipleConstantTemplate *template.Template
	singleArgumentTemplate   *template.Template
	multipleArgumentTemplate *template.Template
	singleReturnTemplate     *template.Template
	multipleReturnTemplate   *template.Template
	structTemplate           *template.Template
	variableTemplate         *template.Template
	functionTemplate         *template.Template
	interfaceTemplate        *template.Template
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

	instance = &Template{
		packageTemplate:          packageTmpl,
		singleImportTemplate:     singleImportTmpl,
		multipleImportTemplate:   multipleImportTmpl,
		singleConstantTemplate:   singleConstantTmpl,
		multipleConstantTemplate: multipleConstantTmpl,
		singleArgumentTemplate:   singleArgumentTmpl,
		multipleArgumentTemplate: multipleArgumentTmpl,
		singleReturnTemplate:     singleReturnTmpl,
		multipleReturnTemplate:   multipleReturnTmpl,
		structTemplate:           structTmpl,
		variableTemplate:         variableTmpl,
		functionTemplate:         functionTmpl,
		interfaceTemplate:        interfaceTmpl,
	}
}

func GetTemplate() *Template {
	return instance
}

func (t Template) Package(pkg Package) string {
	code := bytes.NewBufferString("")
	err := t.packageTemplate.Execute(code, pkg)
	if err != nil {
		log.Panicf("package template execute fail. %v", err)
	}

	return code.String()
}

func (t Template) Import(imp Imports) string {
	code := bytes.NewBufferString("")
	var err error
	if len(imp) == 1 {
		err = t.singleImportTemplate.Execute(code, imp[0])
	} else if len(imp) > 1 {
		err = t.multipleImportTemplate.Execute(code, imp)
	}
	if err != nil {
		log.Panicf("import template execute fail. %v", err)
	}

	return code.String()
}

func (t Template) Constant(cons Constants) string {
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

func (t Template) Argument(args Arguments) string {
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

func (t Template) Return(returns Returns) string {
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

func (t Template) Struct(structs Structs) string {
	code := bytes.NewBufferString("")
	for _, st := range structs {
		err := t.structTemplate.Execute(code, st)
		if err != nil {
			log.Panicf("struct template execute fail. %v", err)
		}
	}

	return code.String()
}

func (t Template) Variable(vars Variables) string {
	code := bytes.NewBufferString("")
	for _, v := range vars {
		err := t.variableTemplate.Execute(code, v)
		if err != nil {
			log.Panicf("variable template execute fail. %v", err)
		}
		code.WriteString("\n")
	}

	return code.String()
}

func (t Template) Function(funs Functions) string {
	code := bytes.NewBufferString("")
	for _, fun := range funs {
		err := t.functionTemplate.Execute(code, fun)
		if err != nil {
			log.Panicf("function template execute fail. %v", err)
		}
	}

	return code.String()
}

func (t Template) Interface(ins Interfaces) string {
	code := bytes.NewBufferString("")
	for _, in := range ins {
		err := t.interfaceTemplate.Execute(code, in)
		if err != nil {
			log.Panicf("interface template execute fail. %v", err)
		}
	}

	return code.String()
}
