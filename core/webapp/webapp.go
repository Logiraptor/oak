package webapp

import (
	"fmt"
	"net/http"

	"golang.org/x/net/context"
)

func QueryParam(name string) func(context.Context, *http.Request) (context.Context, string) {
	return func(ctx context.Context, req *http.Request) (context.Context, string) {
		return ctx, req.FormValue(name)
	}
}

func WriteString(ctx context.Context, val string) context.Context {
	fmt.Println(val)
	return ctx
}
