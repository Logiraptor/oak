package core

import (
	"context"

	"github.com/Logiraptor/oak/flow/pipeline"
	"github.com/Logiraptor/oak/flow/values"
)

func Eq() pipeline.Component {
	var (
		lhs    = values.NewToken("LHS")
		rhs    = values.NewToken("RHS")
		output = values.NewToken("Output")
		t      = values.NewGenericType("a")
	)

	return pipeline.Component{
		Name: values.NewToken("=="),
		InputPorts: []pipeline.Port{
			pipeline.Port{Name: lhs, Type: t},
			pipeline.Port{Name: rhs, Type: t},
		},
		OutputPorts: []pipeline.Port{
			pipeline.Port{Name: output, Type: values.BoolType},
		},
		Invoke: func(ctx context.Context, input values.RecordValue, emitter pipeline.Emitter) {
			lhs := input.FieldByToken(lhs)
			rhs := input.FieldByToken(rhs)
			emitter.Emit(ctx, output, values.BoolValue(values.EqualValues(lhs, rhs)))
		},
	}
}
