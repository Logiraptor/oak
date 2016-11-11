package core

import (
	"context"

	"github.com/Logiraptor/oak/flow/pipeline"
	"github.com/Logiraptor/oak/flow/values"
)

func FieldAccessor(name string) pipeline.Component {
	var (
		outputType = values.NewGenericType("a")
		inputType  = values.RecordType{
			RecordName: "RecordWith_" + name + "_Field",
			Fields: []values.FieldType{
				{Name: name, Type: outputType},
			},
		}
		input  = values.NewToken("Input")
		output = values.NewToken("Output")
	)

	return pipeline.Component{
		Name: values.NewToken("Field Accessor: " + name),
		InputPorts: []pipeline.Port{
			pipeline.Port{Name: input, Type: inputType},
		},
		OutputPorts: []pipeline.Port{
			pipeline.Port{Name: output, Type: outputType},
		},
		Invoke: func(ctx context.Context, inputRecord values.RecordValue, emitter pipeline.Emitter) {
			inputRecord = inputRecord.FieldByToken(input).(values.RecordValue)
			for _, field := range inputRecord.Fields {
				if field.Name == name {
					emitter.Emit(ctx, output, field.Value)
				}
			}
		},
	}
}
