package main

import (
	"io"
	"net/http"

	"fmt"

	"os/exec"

	"encoding/json"

	"github.com/Logiraptor/oak/flow/pipeline"
	"github.com/Logiraptor/oak/flow/values"
	"golang.org/x/net/websocket"
)

type message struct {
	from, to values.Token
	value    values.Value
	color    string
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
	var pipeIndex = 0
	for i, p := range d.p.Pipes {
		if p.Source == from && p.Dest == to {
			pipeIndex = i
		}
	}
	d.messages <- message{
		from: from, to: to, value: value, color: colors[pipeIndex],
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
	type WSMessage struct {
		Message string
		Color   string
	}
	in := make(chan message)
	d.listeners <- in
	encoder := json.NewEncoder(ws)
	for m := range in {
		encoder.Encode(WSMessage{
			Message: fmt.Sprintf("%v -> %v (%s)", m.from.Name, m.to.Name, values.ValueToString(m.value)),
			Color:   m.color,
		})
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
		WriteToDot(stdin, d.p)
		stdin.Close()
	}()

	rw.Header().Set("Content-Type", "image/png")
	go func() {
		io.Copy(rw, stdout)
		stdout.Close()
	}()

	dotCmd.Run()
}

type errWriter struct {
	err error
	w   io.Writer
}

func (e *errWriter) Write(buf []byte) (int, error) {
	if e.err != nil {
		return 0, e.err
	}
	n, err := e.w.Write(buf)
	e.err = err
	return n, err
}

func WriteToDot(w io.Writer, p pipeline.Pipeline) error {
	we := &errWriter{w: w}
	io.WriteString(we, "digraph {")
	for pipeIndex, pipe := range p.Pipes {
		var sourceIndex int
		var sourceType values.Type
		var destIndex int
		var destType values.Type
		for i, component := range p.Components {
			for _, port := range component.InputPorts {
				if port.Name == pipe.Source {
					sourceIndex = i
					sourceType = port.Type
				}
				if port.Name == pipe.Dest {
					destIndex = i
					destType = port.Type
				}
			}
			for _, port := range component.OutputPorts {
				if port.Name == pipe.Source {
					sourceIndex = i
					sourceType = port.Type
				}
				if port.Name == pipe.Dest {
					destIndex = i
					destType = port.Type
				}
			}
		}

		fmt.Fprintf(we, "%q -> %q [headlabel = %q taillabel = %q  color = %q];",
			p.Components[sourceIndex].Name.Name,
			p.Components[destIndex].Name.Name,

			values.TypeToString(destType),
			values.TypeToString(sourceType),
			colors[pipeIndex])
	}
	io.WriteString(we, "}")
	return we.err
}

var colors = []string{
	"#000000", "#FFFF00", "#1CE6FF", "#FF34FF", "#FF4A46", "#008941", "#006FA6", "#A30059",
	"#FFDBE5", "#7A4900", "#0000A6", "#63FFAC", "#B79762", "#004D43", "#8FB0FF", "#997D87",
	"#5A0007", "#809693", "#FEFFE6", "#1B4400", "#4FC601", "#3B5DFF", "#4A3B53", "#FF2F80",
	"#61615A", "#BA0900", "#6B7900", "#00C2A0", "#FFAA92", "#FF90C9", "#B903AA", "#D16100",
	"#DDEFFF", "#000035", "#7B4F4B", "#A1C299", "#300018", "#0AA6D8", "#013349", "#00846F",
	"#372101", "#FFB500", "#C2FFED", "#A079BF", "#CC0744", "#C0B9B2", "#C2FF99", "#001E09",
	"#00489C", "#6F0062", "#0CBD66", "#EEC3FF", "#456D75", "#B77B68", "#7A87A1", "#788D66",
	"#885578", "#FAD09F", "#FF8A9A", "#D157A0", "#BEC459", "#456648", "#0086ED", "#886F4C",

	"#34362D", "#B4A8BD", "#00A6AA", "#452C2C", "#636375", "#A3C8C9", "#FF913F", "#938A81",
	"#575329", "#00FECF", "#B05B6F", "#8CD0FF", "#3B9700", "#04F757", "#C8A1A1", "#1E6E00",
	"#7900D7", "#A77500", "#6367A9", "#A05837", "#6B002C", "#772600", "#D790FF", "#9B9700",
	"#549E79", "#FFF69F", "#201625", "#72418F", "#BC23FF", "#99ADC0", "#3A2465", "#922329",
	"#5B4534", "#FDE8DC", "#404E55", "#0089A3", "#CB7E98", "#A4E804", "#324E72", "#6A3A4C",
	"#83AB58", "#001C1E", "#D1F7CE", "#004B28", "#C8D0F6", "#A3A489", "#806C66", "#222800",
	"#BF5650", "#E83000", "#66796D", "#DA007C", "#FF1A59", "#8ADBB4", "#1E0200", "#5B4E51",
	"#C895C5", "#320033", "#FF6832", "#66E1D3", "#CFCDAC", "#D0AC94", "#7ED379", "#012C58",
}
