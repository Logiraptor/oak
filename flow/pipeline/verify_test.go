package pipeline

import "testing"
import "github.com/Logiraptor/oak/flow/values"

type Struct1 struct {
	Name string
}

type Struct2 struct {
	Name string
	Age  int
}

var struct1Type = values.NewValue(Struct1{}).GetType()
var struct2Type = values.NewValue(Struct2{}).GetType()

func TestPipeline_Verify(t *testing.T) {
	tests := []struct {
		name    string
		fields  Pipeline
		wantErr bool
	}{
		{
			name:    "Empty pipeline",
			wantErr: false,
		},
		{
			name:    "Linear pipeline",
			wantErr: false,
			fields: pipeline([]Component{
				component(nil, Out{"Output": struct1Type}),
				component(In{"Input": struct1Type}, nil),
			}, []Pipe{
				pipe("Output", "Input"),
			}),
		},

		{
			name:    "Linear pipeline with incompatible types",
			wantErr: true,
			fields: pipeline([]Component{
				component(nil, Out{"Output": struct1Type}),
				component(In{"Input": struct2Type}, nil),
			}, []Pipe{
				pipe("Output", "Input"),
			}),
		},

		{
			name:    "Linear pipeline with unspecified inputs",
			wantErr: true,
			fields: pipeline([]Component{
				component(nil, Out{"Output": struct1Type}),
				component(In{"Input1": struct1Type, "Input2": struct1Type}, nil),
			}, []Pipe{
				pipe("Output", "Input1"),
			}),
		},

		{
			name:    "Generic inputs",
			wantErr: false,
			fields: pipeline([]Component{
				component(nil, Out{"Output": struct1Type}),
				component(In{"Input": values.NewGenericType("a")}, nil),
			}, []Pipe{
				pipe("Output", "Input"),
			}),
		},

		{
			name:    "Generic inputs and outputs (passthrough) -- correctly typed",
			wantErr: false,
			fields: pipeline([]Component{
				component(nil, Out{"Output1": struct1Type}),
				func() Component {
					t := values.NewGenericType("a")
					return component(In{"Input1": t}, Out{"Output2": t})
				}(),
				component(In{"Input2": struct1Type}, nil),
			}, []Pipe{
				pipe("Output1", "Input1"),
				pipe("Output2", "Input2"),
			}),
		},

		{
			name:    "Generic inputs and outputs (passthrough) -- incorrectly typed",
			wantErr: true,
			fields: pipeline([]Component{
				component(nil, Out{"Output1": struct1Type}),
				func() Component {
					t := values.NewGenericType("a")
					return component(In{"Input1": t}, Out{"Output2": t})
				}(),
				component(In{"Input2": struct2Type}, nil),
			}, []Pipe{
				pipe("Output1", "Input1"),
				pipe("Output2", "Input2"),
			}),
		},
	}
	for _, tt := range tests {
		p := &Pipeline{
			Components: tt.fields.Components,
			Pipes:      tt.fields.Pipes,
		}
		if err := p.Verify(); (err != nil) != tt.wantErr {
			t.Errorf("%q. Pipeline.Verify() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func port(name string, typ values.Type) Port {
	return Port{
		Name: values.Token{name, 1},
		Type: typ,
	}
}

type In map[string]values.Type
type Out map[string]values.Type

func component(input In, output Out) Component {

	var inputs []Port
	for name, typ := range input {
		inputs = append(inputs, port(name, typ))
	}

	var outputs []Port
	for name, typ := range output {
		outputs = append(outputs, port(name, typ))
	}

	return Component{
		InputPorts:  inputs,
		OutputPorts: outputs,
	}
}

func pipeline(components []Component, pipes []Pipe) Pipeline {
	return Pipeline{
		Components: components,
		Pipes:      pipes,
	}
}

func pipe(from, to string) Pipe {
	return Pipe{
		Source: values.Token{from, 1},
		Dest:   values.Token{to, 1},
	}
}

func pipelines() {
	p := pipeline([]Component{
		component(nil, Out{"Output": struct1Type}),
		component(In{"Input": struct2Type}, nil),
	}, []Pipe{
		pipe("Output", "Input"),
	})
	p.Verify()
}
