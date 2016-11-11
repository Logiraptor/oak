package main

import (
	"io"
	"net/http"

	"fmt"

	"os/exec"

	"github.com/Logiraptor/oak/flow/pipeline"
	"github.com/Logiraptor/oak/flow/values"
	"golang.org/x/net/websocket"
)

type message struct {
	from, to values.Token
	value    values.Value
}

type debugView struct {
	p pipeline.Pipeline

	messages  chan message
	listeners chan chan message
}

func (d *debugView) Start() {
	var listeners []chan message

	for {
		select {
		case l := <-d.listeners:
			listeners = append(listeners, l)
		case m := <-d.messages:
			for _, l := range listeners {
				l <- m
			}
		}
	}
}

func (d *debugView) MessageSent(from, to values.Token, value values.Value) {
	d.messages <- message{
		from: from, to: to, value: value,
	}
}

func (d *debugView) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		d.homePage(rw, req)
	case "/graph.png":
		d.graphImage(rw, req)
	case "/messages":
		websocket.Handler(d.messagePipe).ServeHTTP(rw, req)
	default:
		http.NotFound(rw, req)
	}
}

func (d *debugView) messagePipe(ws *websocket.Conn) {
	in := make(chan message)
	d.listeners <- in
	for m := range in {
		fmt.Fprintf(ws, "%v -> %v (%s)", m.from, m.to, values.ValueToString(m.value))
	}
}

func (d *debugView) homePage(rw http.ResponseWriter, req *http.Request) {
	http.ServeFile(rw, req, "index.html")
}

func (d *debugView) graphImage(rw http.ResponseWriter, req *http.Request) {
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
