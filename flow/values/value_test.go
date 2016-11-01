package values

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type SimpleRecord struct {
	Name string
	Age  int
}

func TestNewValue(t *testing.T) {
	r := SimpleRecord{
		Name: "Foo",
		Age:  23,
	}
	v := NewValue(r)

	typ := v.GetType()

	assert.IsType(t, RecordType{}, typ)

	rt := typ.(RecordType)

	assert.Equal(t, 2, len(rt.Fields))

	assert.Equal(t, StringType, rt.Fields[0].Type)
	assert.Equal(t, "Name", rt.Fields[0].Name)
	assert.Equal(t, IntType, rt.Fields[1].Type)
	assert.Equal(t, "Age", rt.Fields[1].Name)

	assert.IsType(t, RecordValue{}, v)

	rv := v.(RecordValue)

	assert.Equal(t, 2, len(rv.Fields))
	assert.Equal(t, "Name", rv.Fields[0].Name)
	assert.Equal(t, StringType, rv.Fields[0].Value.GetType())
	assert.Equal(t, "Age", rv.Fields[1].Name)
	assert.Equal(t, IntType, rv.Fields[1].Value.GetType())

	nameValue := rv.Fields[0].Value.(StringValue)
	ageValue := rv.Fields[1].Value.(IntValue)

	assert.Equal(t, "Foo", string(nameValue))
	assert.Equal(t, 23, int(ageValue))
}
