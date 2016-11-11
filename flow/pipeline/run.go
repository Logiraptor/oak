package pipeline

import (
	"fmt"

	"context"

	"github.com/Logiraptor/oak/flow/values"
)

type key struct{ Name string }

var InspectorKey = &key{"Inspector"}

const DEBUG = true

func debugPrint(args ...interface{}) {
	if DEBUG {
		fmt.Println(args...)
	}
}

type Update struct {
	Ctx  context.Context
	Val  values.Value
	Port values.Token
}

type Inspector interface {
	MessageSent(from, to values.Token, val values.Value)
}

type NilInspector struct{}

func (NilInspector) MessageSent(values.Token, values.Token, values.Value) {}

type pipelineExecutionContext struct {
	inspector               Inspector
	pipeline                *Pipeline
	componentContexts       []*componentExecutionContext
	componentContextsByPort map[values.Token]*componentExecutionContext
	stack                   []values.Token
}

func (p *pipelineExecutionContext) triggerUpdate(ctx context.Context, c *componentExecutionContext) {
	if c.canRun() {
		c.component.Invoke(ctx, c.inputState, FuncEmitter(func(ctx context.Context, name values.Token, value values.Value) {
			for _, pipe := range p.pipeline.Pipes {
				if pipe.Source == name {
					p.inspector.MessageSent(name, pipe.Dest, value)
					destComponentContext := p.componentContextsByPort[pipe.Dest]
					destComponentContext.applyUpdate(Update{
						Ctx: ctx, Port: pipe.Dest, Val: value,
					})
					p.triggerUpdate(ctx, destComponentContext)
				}
			}
		}))
	}
}

type componentExecutionContext struct {
	inputState values.RecordValue
	component  Component
}

func (c *componentExecutionContext) canRun() bool {
	for i := range c.inputState.Fields {
		if c.inputState.Fields[i].Value == nil {
			return false
		}
	}
	return true
}

func (c *componentExecutionContext) applyUpdate(u Update) {
	for i := range c.inputState.Fields {
		if c.inputState.Fields[i].Name == u.Port.Name {
			c.inputState.Fields[i].Value = u.Val
		}
	}
}

func (p *Pipeline) Run(ctx context.Context) {

	inspector, ok := ctx.Value(InspectorKey).(Inspector)
	if !ok {
		inspector = NilInspector{}
	}

	var pipelineExecutionContext = pipelineExecutionContext{
		pipeline:                p,
		inspector:               inspector,
		componentContextsByPort: make(map[values.Token]*componentExecutionContext),
	}

	for componentIndex, component := range p.Components {
		componentContext := &componentExecutionContext{
			inputState: values.RecordValue{},
			component:  component,
		}
		componentContext.inputState.Name = fmt.Sprintf("ComponentInputType$%d", componentIndex)
		for _, port := range component.InputPorts {
			componentContext.inputState.Fields = append(componentContext.inputState.Fields, values.Field{
				Name:  port.Name.Name,
				Value: nil,
			})
		}
		pipelineExecutionContext.componentContexts = append(pipelineExecutionContext.componentContexts, componentContext)

		for _, input := range component.InputPorts {
			pipelineExecutionContext.componentContextsByPort[input.Name] = componentContext
		}
	}

	for _, component := range pipelineExecutionContext.componentContexts {
		if len(component.component.InputPorts) == 0 {
			pipelineExecutionContext.triggerUpdate(ctx, component)
		}
	}

	<-ctx.Done()
}

type FuncEmitter func(context.Context, values.Token, values.Value)

func (f FuncEmitter) Emit(ctx context.Context, name values.Token, value values.Value) {
	f(ctx, name, value)
}
