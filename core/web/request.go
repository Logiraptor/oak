package web

import (
	"context"
	"net/http"

	"github.com/Logiraptor/oak/flow/frontends/web"
	"github.com/Logiraptor/oak/flow/pipeline"
	"github.com/Logiraptor/oak/flow/values"
)

func HTTPRequest() pipeline.Component {
	output := values.NewToken("Output")
	return pipeline.Component{
		Name:       values.NewToken("HTTPRequest"),
		InputPorts: []pipeline.Port{},
		OutputPorts: []pipeline.Port{
			pipeline.Port{Name: output, Type: values.RecordType{
				RecordName: "HttpRequest",
				Fields: []values.FieldType{
					{Name: "Method", Type: values.StringType},
					{Name: "Path", Type: values.StringType},
				},
			}},
		},
		Invoke: func(ctx context.Context, _ values.RecordValue, emitter pipeline.Emitter) {
			req := ctx.Value(web.ReqKey).(*http.Request)
			emitter.Emit(ctx, output, values.RecordValue{
				Name: "HttpRequest",
				Fields: []values.Field{
					{Name: "Method", Value: values.StringValue(req.Method)},
					{Name: "Path", Value: values.StringValue(req.URL.Path)},
				},
			})
		},
	}
}
