package loader

import (
	"bytes"
	"go/types"

	"github.com/Logiraptor/oak/flow/internal/templates"
	"github.com/Logiraptor/oak/flow/parser"
	goloader "golang.org/x/tools/go/loader"
)

// Load loads the components of a flow app and
// reports type errors found in the graph.
func Load(prog parser.Program) (App, []error) {
	pkg, err := load(writeProgram(prog))
	if err != nil {
		return App{}, []error{err}
	}

	var app App
	app.Entry = prog.Entry
	app.Imports = prog.Imports
	for name, comp := range prog.Components {
		typ := lookupDef(pkg, string(name)).Type().(*types.Signature)
		app.Components = append(app.Components, Component{
			Label:   name,
			Func:    comp,
			Inputs:  typ.Params(),
			Outputs: typ.Results(),
		})
	}
	app.Flow = newGraph(app, prog.Flow)

	errs := typeCheck(pkg, prog)
	return app, errs
}

func lookupDef(pkg *goloader.PackageInfo, name string) types.Object {
	for id, obj := range pkg.Defs {
		if name == id.String() {
			return obj
		}
	}
	panic("undefined definition: " + name)
}

func typeCheck(pkg *goloader.PackageInfo, conf parser.Program) []error {
	var errs []error
	for start, end := range conf.Flow {
		startSig := lookupDef(pkg, string(start)).Type().(*types.Signature)
		endSig := lookupDef(pkg, string(end)).Type().(*types.Signature)

		startResults := startSig.Results()
		endParams := endSig.Params()

		if !endSig.Variadic() && startResults.Len() != endParams.Len() {
			errs = append(errs, cardinalityMismatchError(start, end, startResults, endParams))
			continue
		}

	outer:
		for i := 0; i < startResults.Len(); i++ {
			source := startResults.At(i)
			dest := endParams.At(i)

			// drain the rest of startResults if the dest is variadic
			if endSig.Variadic() && i == endParams.Len()-1 {
				destType := dest.Type().(*types.Slice).Elem()
				for j := i; j < startResults.Len(); j++ {
					source = startResults.At(j)

					if !types.AssignableTo(source.Type(), destType) {
						errs = append(errs, typeMismatchError(start, end, j, source.Type(), dest.Type()))
					}
				}
				break outer
			}

			if !types.AssignableTo(source.Type(), dest.Type()) {
				errs = append(errs, typeMismatchError(start, end, i, source.Type(), dest.Type()))
			}
		}
	}
	return errs
}

func writeProgram(conf parser.Program) string {
	buf := new(bytes.Buffer)
	templates.TypeChecker(buf, conf)
	return buf.String()
}

func load(src string) (*goloader.PackageInfo, error) {
	conf := goloader.Config{}
	// This line prevents the loader package from printing errors
	// to stderr.
	conf.TypeChecker.Error = func(err error) {}
	file, err := conf.ParseFile("__loader.go", src)
	if err != nil {
		return nil, err
	}

	conf.CreateFromFiles("flow", file)

	prog, err := conf.Load()
	if err != nil {
		return nil, err
	}
	return prog.Package("flow"), err
}
