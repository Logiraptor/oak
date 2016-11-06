package pipeline

import (
	"fmt"

	"context"

	"github.com/Logiraptor/oak/flow/values"
)

const DEBUG = false

func debugPrint(args ...interface{}) {
	if DEBUG {
		fmt.Println(args...)
	}
}

func (p *Pipeline) Run(ctx context.Context) {

	type Update struct {
		ctx  context.Context
		val  values.Value
		port values.Token
	}

	type Handle struct {
		Inputs  chan Update
		Outputs chan Update
	}

	var chans = make(map[values.Token]Handle)
	var handles []Handle

	for componentIndex, component := range p.Components {

		handle := Handle{
			Inputs:  make(chan Update),
			Outputs: make(chan Update),
		}
		handles = append(handles, handle)
		for _, input := range component.InputPorts {
			chans[input.Name] = handle
		}

		go func(componentIndex int, component Component) {
			var currentState values.RecordValue
			currentState.Name = fmt.Sprintf("ComponentInputType$%d", componentIndex)
			for _, port := range component.InputPorts {
				currentState.Fields = append(currentState.Fields, values.Field{
					Name:  port.Name.Name,
					Value: nil,
				})
			}
			if len(component.InputPorts) == 0 {
				component.Invoke(ctx, currentState, FuncEmitter(func(ctx context.Context, name values.Token, value values.Value) {
					handle.Outputs <- Update{
						ctx:  ctx,
						port: name,
						val:  value,
					}
				}))
				return
			}
			debugPrint("Listening for inputs to component", componentIndex)
			for {
				select {
				case <-ctx.Done():
					debugPrint("Shutting down input listener", componentIndex)
					return
				case input := <-handle.Inputs:
					debugPrint("Component", componentIndex, "received input: ", values.ValueToString(input.val))
					for i := range currentState.Fields {
						if currentState.Fields[i].Name == input.port.Name {
							currentState.Fields[i].Value = input.val
						}
					}

					var isValid = true
					for i := range currentState.Fields {
						if currentState.Fields[i].Value == nil {
							debugPrint("Missing field", currentState.Fields[i].Name)
							isValid = false
						}
					}

					if isValid {
						component.Invoke(input.ctx, currentState, FuncEmitter(func(context context.Context, name values.Token, value values.Value) {
							handle.Outputs <- Update{
								ctx:  context,
								port: name,
								val:  value,
							}
						}))
					}
				}
			}
		}(componentIndex, component)
	}

	for i, ch := range handles {
		go func(i int, ch Handle) {
			debugPrint("Listening for outputs from component", i)
			for {
				select {
				case <-ctx.Done():
					debugPrint("Shutting down output listener", i)
					return
				case v := <-ch.Outputs:
					debugPrint("Component", i, "emitted output: ", values.ValueToString(v.val))

					for _, pipe := range p.Pipes {
						debugPrint(pipe, i, v.port)

						if pipe.Source == v.port {
							debugPrint("Forwarding output from ", i, "to", pipe.Dest)

							destCh := chans[pipe.Dest]
							destCh.Inputs <- Update{
								ctx:  v.ctx,
								port: pipe.Dest,
								val:  v.val,
							}
						}
					}
				}
			}
		}(i, ch)
	}

	<-ctx.Done()

}

type FuncEmitter func(context.Context, values.Token, values.Value)

func (f FuncEmitter) Emit(ctx context.Context, name values.Token, value values.Value) {
	f(ctx, name, value)
}
