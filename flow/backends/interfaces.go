package backends

import "github.com/Logiraptor/oak/flow/values"

type Storage interface {
	PrepareType(values.Type) error
	Put(values.Value) error
	Find(values.Type, values.Value) ([]values.Value, error)
}
