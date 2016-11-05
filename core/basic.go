package core

import (
	"log"
	"time"

	"github.com/Logiraptor/oak/flow/pipeline"
	"github.com/Logiraptor/oak/flow/values"
)

func Repeater(interval time.Duration) pipeline.Component {
	output := pipeline.NewToken("Output")
	return pipeline.Component{
		InputPorts: []pipeline.Port{},
		OutputPorts: []pipeline.Port{
			{
				Name: output,
				Type: values.BoolType,
			},
		},
		Invoke: func(_ values.RecordValue, emitter pipeline.Emitter) {
			for range time.Tick(interval) {
				emitter.Emit(output, values.BoolValue(true))
			}
		},
	}
}

func Constant(val values.Value) pipeline.Component {
	output := pipeline.NewToken("Output")
	return pipeline.Component{
		InputPorts: []pipeline.Port{},
		OutputPorts: []pipeline.Port{
			{
				Name: output,
				Type: val.GetType(),
			},
		},
		Invoke: func(_ values.RecordValue, emitter pipeline.Emitter) {
			emitter.Emit(output, val)
		},
	}
}

func Logger() pipeline.Component {
	input := pipeline.NewToken("Input")
	return pipeline.Component{
		InputPorts: []pipeline.Port{
			{Name: input, Type: values.StringType},
		},
		OutputPorts: []pipeline.Port{},
		Invoke: func(val values.RecordValue, _ pipeline.Emitter) {
			log.Println(val)
		},
	}
}

func Cond() pipeline.Component {
	var (
		cond   = pipeline.NewToken("Condition")
		conseq = pipeline.NewToken("Consequence")
		alt    = pipeline.NewToken("Alternative")
		output = pipeline.NewToken("Output")
	)
	return pipeline.Component{
		InputPorts: []pipeline.Port{
			{Name: cond, Type: values.BoolType},
			{Name: conseq, Type: values.StringType},
			{Name: alt, Type: values.StringType},
		},
		OutputPorts: []pipeline.Port{
			{Name: output, Type: values.StringType},
		},
		Invoke: func(val values.RecordValue, emitter pipeline.Emitter) {
			var input = new(struct {
				Condition                bool
				Consequence, Alternative string
			})
			values.FillNative(val, input)
			if input.Condition {
				emitter.Emit(output, values.StringValue(input.Consequence))
			} else {
				emitter.Emit(output, values.StringValue(input.Alternative))
			}
		},
	}
}
