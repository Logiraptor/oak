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

	assert.Equal(t, Record, typ.GetKind())
	assert.Implements(t, (*RecordType)(nil), typ)

	rt := typ.(RecordType)

	assert.Equal(t, 2, rt.NumFields())

	assert.Equal(t, String, rt.Field(0).GetKind())
	assert.Equal(t, "Name", rt.Field(0).Name)
	assert.Equal(t, Int, rt.Field(1).GetKind())
	assert.Equal(t, "Age", rt.Field(1).Name)

	assert.Implements(t, (*RecordValue)(nil), v)

	rv := v.(RecordValue)

	assert.Equal(t, 2, rv.NumFields())
	assert.Equal(t, "Name", rv.Field(0).Name)
	assert.Equal(t, String, rv.Field(0).Value.GetType().GetKind())
	assert.Equal(t, "Age", rv.Field(1).Name)
	assert.Equal(t, Int, rv.Field(1).Value.GetType().GetKind())

	nameValue := rv.Field(0).Value.(StringValue)
	ageValue := rv.Field(1).Value.(IntValue)

	assert.Equal(t, "Foo", nameValue.StringValue())
	assert.Equal(t, 23, ageValue.IntValue())
}
