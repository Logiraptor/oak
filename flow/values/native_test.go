package values

import (
	"reflect"
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/assert"
)

type Person struct {
	Name   string
	Age    int
	Active bool
	Tags   []string
	Pets   []Pet
}

type Pet struct {
	Breed string
}

type IncompletePerson struct {
	Name   string
	Age    int
	Active bool
	Tags   []string
	Pets   []Value
}

func TestFillNative(t *testing.T) {
	err := quick.Check(func(p Person) bool {

		v := NewValue(p)
		newP := Person{}
		FillNative(v, &newP)

		return reflect.DeepEqual(p, newP)
	}, nil)
	assert.NoError(t, err)
}

func TestToNative(t *testing.T) {
	err := quick.Check(func(p Person) bool {

		v := NewValue(p)
		native := ToNative(v)

		return EqualValues(v, NewValue(native))
	}, nil)
	assert.NoError(t, err)
}
