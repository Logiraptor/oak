package flow

import (
	"go/types"
	"os"
	"strings"
	"text/template"
)

var tmpl = template.Must(template.New("root").Funcs(template.FuncMap{
	"join": func(a []string) string {
		return strings.Join(a, ",")
	},
}).Parse(flowAppSrc))

var flowAppSrc = `
// THIS FILE WAS GENERATED BY THE OAK FLOW TOOL
package main

import (
{{range $k, $v := .App.Imports}}
	{{$k}} "{{$v}}"
{{- end}}
)

func main() {
	{{range $name, $value := .App.Components}}
	var {{$name}} = {{$value.Func}}
	{{- end}}

	{{$root := .App}}

	{{range .Steps}}
	{{with .LHS}}{{. | join}} := {{end}}{{.Func}}({{.Args | join}})
	{{end}}
}

`

type TemplateStep struct {
	LHS  []string
	Func ComponentID
	Args []string
}

type TemplateParams struct {
	App   App
	Steps []TemplateStep
}

func WriteFlowApp(app App) {

	order := app.Flow.TopologicalSort(app.Entry)

	var steps = make([]TemplateStep, len(order))

	steps[0] = TemplateStep{
		LHS:  names(app.Components[order[1].Label].Inputs),
		Func: order[0].Label,
		Args: nil,
	}

	for i, node := range order[1 : len(order)-1] {
		steps[i+1].Func = node.Label
		steps[i+1].Args = steps[i].LHS
		steps[i+1].LHS = names(app.Components[order[i+2].Label].Inputs)
	}

	steps[len(steps)-1] = TemplateStep{
		LHS:  nil,
		Func: order[len(order)-1].Label,
		Args: steps[len(steps)-2].LHS,
	}

	tmpl.Execute(os.Stdout, TemplateParams{
		App:   app,
		Steps: steps,
	})
}

func names(in *types.Tuple) []string {
	out := make([]string, in.Len())
	for i := 0; i < in.Len(); i++ {
		out[i] = in.At(i).Name()
	}
	return out
}
