package flow

import (
	"bytes"
	"go/types"
	"text/template"

	"golang.org/x/tools/go/loader"
)

type Component struct {
	Func    string
	Inputs  *types.Tuple
	Outputs *types.Tuple
}

type App struct {
	Imports    map[string]string
	Entry      ComponentID
	Components map[ComponentID]Component
	Flow       Graph
}

var outputTempl = `
package flow

import (
{{range $k, $v := .Imports}}
	{{$k}} "{{$v}}"
{{- end}}
)

{{range $name, $comp := .Components}}
var {{$name}} = {{$comp}}
{{- end}}

`

type TypeChecker struct {
	Conf *Config
}

func (t TypeChecker) writeProgram() string {
	tmpl := template.Must(template.New("root").Funcs(template.FuncMap{
	// "process": makeProcess,
	}).Parse(outputTempl))

	buf := new(bytes.Buffer)
	tmpl.Execute(buf, t.Conf)
	return buf.String()
}

func load(src string) (*loader.Program, error) {
	conf := loader.Config{}
	file, err := conf.ParseFile("__loader.go", src)
	if err != nil {
		return nil, err
	}

	conf.CreateFromFiles("flow", file)

	return conf.Load()
}

func (t TypeChecker) loadProgram() (*loader.Program, error) {
	return load(t.writeProgram())
}

func (t TypeChecker) Check() (*App, []error) {
	var app = new(App)
	app.Components = make(map[ComponentID]Component)
	app.Entry = t.Conf.Entry
	app.Imports = t.Conf.Imports

	prog, err := t.loadProgram()
	if err != nil {
		return nil, []error{err}
	}

	var lookup = make(map[string]types.Object)
	for id, obj := range prog.Package("flow").Defs {
		if _, ok := t.Conf.Components[ComponentID(id.String())]; ok {
			lookup[id.String()] = obj
		}
	}

	for name, comp := range t.Conf.Components {
		typ := lookup[string(name)].Type().(*types.Signature)
		app.Components[name] = Component{
			Func:    comp,
			Inputs:  typ.Params(),
			Outputs: typ.Results(),
		}
	}

	app.Flow = NewFlowGraph(app.Components, t.Conf.Flow)

	var errs []error
	for start, end := range t.Conf.Flow {
		startSig := lookup[string(start)].Type().(*types.Signature).Results()
		endSig := lookup[string(end)].Type().(*types.Signature).Params()

		if startSig.Len() != endSig.Len() {
			errs = append(errs, cardinalityMismatchError(start, end, startSig, endSig))
			continue
		}

		for i := 0; i < startSig.Len(); i++ {
			source := startSig.At(i)
			dest := endSig.At(i)
			if !types.Identical(source.Type(), dest.Type()) {
				errs = append(errs, typeMismatchError(start, end, i, source.Type(), dest.Type()))
			}
		}
	}
	return app, errs
}
