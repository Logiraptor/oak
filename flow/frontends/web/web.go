package web

import (
	"net/http"

	"context"

	"github.com/Logiraptor/oak/flow/pipeline"
)

type key struct{ name string }

var ReqKey = key{name: "req"}
var RWKey = key{name: "rw"}
var DoneKey = key{name: "done"}

func Serve(addr string, ctx context.Context, p pipeline.Pipeline) {
	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		done := make(chan struct{})
		ctx := context.WithValue(ctx, ReqKey, req)
		ctx = context.WithValue(ctx, RWKey, rw)
		ctx = context.WithValue(ctx, DoneKey, done)
		ctx, cancel := context.WithCancel(ctx)
		go p.Run(ctx)
		<-done
		cancel()
	})

	http.ListenAndServe(addr, nil)
}
