package test

import (
	"testing"

	"github.com/Logiraptor/oak/flow/backends"
	"github.com/Logiraptor/oak/flow/values"
	"github.com/stretchr/testify/assert"
)

type ExampleModel struct {
	Name string
	Age  int
}

type NameFilter struct {
	Name string
}

func StorageContract(t *testing.T, cut backends.Storage) {
	userValue := values.NewValue(ExampleModel{})
	userType := userValue.GetType()

	err := cut.PrepareType(userType)
	assert.NoError(t, err)

	err = cut.Put(values.NewValue(ExampleModel{
		Name: "Alan Turing",
		Age:  23,
	}))
	assert.NoError(t, err)

	err = cut.Put(values.NewValue(ExampleModel{
		Name: "Donald Knuth",
		Age:  45,
	}))
	assert.NoError(t, err)

	results, err := cut.Find(userType, values.NewValue(NameFilter{
		Name: "Alan Turing",
	}))

	assert.Len(t, results, 1)
	assert.True(t, values.EqualValues(results[0], values.NewValue(ExampleModel{
		Name: "Alan Turing",
		Age:  23,
	})))

	results, err = cut.Find(userType, nil)

	assert.Len(t, results, 2)
}
