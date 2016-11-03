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

func TestPipeline_Verify(t *testing.T) {
	type fields struct {
		Components []Component
		Pipes      []Pipe
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "Empty pipeline",
			wantErr: false,
		},
		{
			name:    "Linear pipeline",
			wantErr: false,
			fields: fields{

				Components: []Component{
					0: Component{
						InputPorts:  []Port{},
						OutputPorts: []Port{{Name: "Output", Type: values.NewValue(Struct1{}).GetType()}},
					},
					1: Component{
						InputPorts:  []Port{{Name: "Input", Type: values.NewValue(Struct1{}).GetType()}},
						OutputPorts: []Port{},
					},
				},

				Pipes: []Pipe{
					{
						SourcePort:      "Output",
						DestPort:        "Input",
						SourceComponent: 0,
						DestComponent:   1,
					},
				},
			},
		},

		{
			name:    "Linear pipeline with incompatible types",
			wantErr: true,
			fields: fields{

				Components: []Component{
					0: Component{
						InputPorts:  []Port{},
						OutputPorts: []Port{{Name: "Output", Type: values.NewValue(Struct1{}).GetType()}},
					},
					1: Component{
						InputPorts:  []Port{{Name: "Input", Type: values.NewValue(Struct2{}).GetType()}},
						OutputPorts: []Port{},
					},
				},

				Pipes: []Pipe{
					{
						SourcePort:      "Output",
						DestPort:        "Input",
						SourceComponent: 0,
						DestComponent:   1,
					},
				},
			},
		},

		{
			name:    "Linear pipeline with unspecified inputs",
			wantErr: true,
			fields: fields{

				Components: []Component{
					0: Component{
						InputPorts:  []Port{},
						OutputPorts: []Port{{Name: "Output", Type: values.NewValue(Struct1{}).GetType()}},
					},
					1: Component{
						InputPorts: []Port{
							{Name: "Input1", Type: values.NewValue(Struct1{}).GetType()},
							{Name: "Input2", Type: values.NewValue(Struct1{}).GetType()},
						},
						OutputPorts: []Port{},
					},
				},

				Pipes: []Pipe{
					{
						SourcePort:      "Output",
						DestPort:        "Input1",
						SourceComponent: 0,
						DestComponent:   1,
					},
				},
			},
		},

		{
			name:    "Alpha conversion prevents aliasing",
			wantErr: true,
			fields: fields{

				Components: []Component{
					0: Component{
						InputPorts:  []Port{},
						OutputPorts: []Port{{Name: "Output", Type: values.NewValue(Struct1{}).GetType()}},
					},
					1: Component{
						InputPorts: []Port{
							{Name: "Input", Type: values.NewValue(Struct1{}).GetType()},
						},
						OutputPorts: []Port{},
					},
					2: Component{
						InputPorts: []Port{
							{Name: "Input", Type: values.NewValue(Struct1{}).GetType()},
						},
						OutputPorts: []Port{},
					},
				},

				Pipes: []Pipe{
					{
						SourcePort:      "Output",
						DestPort:        "Input",
						SourceComponent: 0,
						DestComponent:   1,
					},
				},
			},
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
