package core

import (
	"log"
	"time"

	"github.com/Logiraptor/oak/flow/pipeline"
	"github.com/Logiraptor/oak/flow/values"
)

func Repeater(interval time.Duration) pipeline.Component {
	output := values.NewToken("Output")
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
	output := values.NewToken("Output")
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
	input := values.NewToken("Input")
	return pipeline.Component{
		InputPorts: []pipeline.Port{
			{Name: input, Type: values.NewGenericType("a")},
		},
		OutputPorts: []pipeline.Port{},
		Invoke: func(val values.RecordValue, _ pipeline.Emitter) {
			log.Println(val)
		},
	}
}

func Cond() pipeline.Component {
	var (
		cond   = values.NewToken("Condition")
		conseq = values.NewToken("Consequence")
		alt    = values.NewToken("Alternative")
		output = values.NewToken("Output")
		tout   = values.NewGenericType("a")
	)
	return pipeline.Component{
		InputPorts: []pipeline.Port{
			{Name: cond, Type: values.BoolType},
			{Name: conseq, Type: tout},
			{Name: alt, Type: tout},
		},
		OutputPorts: []pipeline.Port{
			{Name: output, Type: tout},
		},
		Invoke: func(val values.RecordValue, emitter pipeline.Emitter) {
			if val.FieldByToken(cond).(values.BoolValue) {
				emitter.Emit(output, val.FieldByToken(conseq))
			} else {
				emitter.Emit(output, val.FieldByToken(alt))
			}
		},
	}
}
