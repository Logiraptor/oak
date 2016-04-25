package base

import (
	"fmt"
	"os"
	"sync"

	"golang.org/x/net/context"
)

var wgKey = new(struct{})

func StringCLI() string {
	return os.Args[1]
}

func Finish(ctx context.Context) {
	wg := ctx.Value(wgKey).(*sync.WaitGroup)
	wg.Done()
}

func PrintString(val string) {
	fmt.Println(val)
}
