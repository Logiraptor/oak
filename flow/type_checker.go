package flow

import (
	"bytes"
	"go/types"

	"golang.org/x/tools/go/loader"
)

func Load(conf Config) (App, []error) {
	pkg, err := load(writeProgram(conf))
	if err != nil {
		return App{}, []error{err}
	}

	var app App
	app.Entry = conf.Entry
	app.Imports = conf.Imports
	for name, comp := range conf.Components {
		typ := lookupDef(pkg, string(name)).Type().(*types.Signature)
		app.Components = append(app.Components, Component{
			Label:   name,
			Func:    comp,
			Inputs:  typ.Params(),
			Outputs: typ.Results(),
		})
	}
	app.Flow = NewGraph(app, conf.Flow)

	errs := typeCheck(pkg, conf)
	return app, errs
}

func lookupDef(pkg *loader.PackageInfo, name string) types.Object {
	for id, obj := range pkg.Defs {
		if name == id.String() {
			return obj
		}
	}
	panic("undefined definition: " + name)
}

func typeCheck(pkg *loader.PackageInfo, conf Config) []error {
	var errs []error
	for start, end := range conf.Flow {
		startSig := lookupDef(pkg, string(start)).Type().(*types.Signature).Results()
		endSig := lookupDef(pkg, string(end)).Type().(*types.Signature).Params()

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
	return errs
}

func writeProgram(conf Config) string {
	buf := new(bytes.Buffer)
	tmpl.ExecuteTemplate(buf, "typeChecker", conf)
	return buf.String()
}

func load(src string) (*loader.PackageInfo, error) {
	conf := loader.Config{}
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
