package main

import (
	"io"
	"net/http"

	"fmt"

	"os/exec"

	"github.com/Logiraptor/oak/flow/pipeline"
)

type debugView struct {
	p pipeline.Pipeline
}

func (d *debugView) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	dotCmd := exec.Command("dot", "-Tpng")
	stdin, err := dotCmd.StdinPipe()
	if err != nil {
		fmt.Fprint(rw, err.Error())
		return
	}
	stdout, err := dotCmd.StdoutPipe()
	if err != nil {
		fmt.Fprint(rw, err.Error())
		return
	}

	go func() {
		d.p.WriteToDot(stdin)
		stdin.Close()
	}()

	rw.Header().Set("Content-Type", "image/png")
	go func() {
		io.Copy(rw, stdout)
		stdout.Close()
	}()

	dotCmd.Run()
}
