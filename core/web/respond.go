package web

import (
	"context"

	"net/http"

	"io"

	"github.com/Logiraptor/oak/flow/frontends/web"
	"github.com/Logiraptor/oak/flow/pipeline"
	"github.com/Logiraptor/oak/flow/values"
)

func HTTPResponder() pipeline.Component {
	var body = values.NewToken("Body")
	var status = values.NewToken("Status")

	return pipeline.Component{
		InputPorts: []pipeline.Port{
			{Name: body, Type: values.StringType},
			{Name: status, Type: values.IntType},
		},
		OutputPorts: []pipeline.Port{},
		Invoke: func(ctx context.Context, input values.RecordValue, emitter pipeline.Emitter) {
			rw := ctx.Value(web.RWKey).(http.ResponseWriter)
			rw.WriteHeader(int(input.FieldByToken(status).(values.IntValue)))
			io.WriteString(rw, string(input.FieldByToken(body).(values.StringValue)))

			close(ctx.Value(web.DoneKey).(chan struct{}))
		},
	}
}
