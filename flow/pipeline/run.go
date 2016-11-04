package pipeline

import (
	"fmt"

	"github.com/Logiraptor/oak/flow/values"
)

func (p *Pipeline) Run() {

	type Update struct {
		val  values.Value
		port Token
	}

	type Handle struct {
		Inputs  chan Update
		Outputs chan Update
	}

	var chans = make(map[Token]Handle)
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
			currentState.Name = fmt.Sprintf("input$%d", componentIndex)
			for portIndex := range component.InputPorts {
				currentState.Fields = append(currentState.Fields, values.Field{
					Name:  fmt.Sprintf("input$%d$%d", componentIndex, portIndex),
					Value: nil,
				})
			}
			if len(component.InputPorts) == 0 {
				component.Invoke(currentState, FuncEmitter(func(name Token, value values.Value) {
					handle.Outputs <- Update{
						port: name,
						val:  value,
					}
				}))
				return
			}
			fmt.Println("Listening for inputs to component", componentIndex)
			for {
				select {
				case input := <-handle.Inputs:
					fmt.Println("Component", componentIndex, "received input: ", values.ValueToString(input.val))
					for i := range currentState.Fields {
						if currentState.Fields[i].Name == input.port.Name {
							currentState.Fields[i].Value = input.val
						}
					}

					var isValid = true
					for i := range currentState.Fields {
						if currentState.Fields[i].Value == nil {
							isValid = false
						}
					}

					if isValid {
						component.Invoke(currentState, FuncEmitter(func(name Token, value values.Value) {
							handle.Outputs <- Update{
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
			for {
				select {
				case v := <-ch.Outputs:
					fmt.Println("Component", i, "emitted output: ", values.ValueToString(v.val))

					for _, pipe := range p.Pipes {
						fmt.Println(pipe, i, v.port)

						if pipe.Source == v.port {
							fmt.Println("Forwarding output from ", i, "to", pipe.Dest)

							destCh := chans[pipe.Dest]
							destCh.Inputs <- Update{
								port: pipe.Dest,
								val:  v.val,
							}
						}
					}
				}
			}
		}(i, ch)
	}

	select {}

}

type FuncEmitter func(Token, values.Value)

func (f FuncEmitter) Emit(name Token, value values.Value) {
	f(name, value)
}
