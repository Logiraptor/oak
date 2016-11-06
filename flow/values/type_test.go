package values

import "testing"
import "github.com/stretchr/testify/assert"

func TestUnifyType_Primitives(t *testing.T) {
	var a = StringType
	var b = StringType

	out, ok := UnifyType(TypeEnv{}, a, b)
	assert.True(t, ok)
	assert.True(t, EqualTypes(StringType, out))
}

func TestUnifyType_Generics(t *testing.T) {
	{
		var a = StringType
		var b = NewGenericType("a")

		out, ok := UnifyType(TypeEnv{}, a, b)
		assert.True(t, ok)
		assert.True(t, EqualTypes(StringType, out))
	}
	{
		var a = StringType
		var b = NewGenericType("a")

		_, ok := UnifyType(TypeEnv{b.token: IntType}, a, b)
		assert.False(t, ok)
	}
	{
		at := NewGenericType("a")
		var a = RecordType{RecordName: "Record1", Fields: []FieldType{{Name: "Name", Type: StringType}}}
		var b = RecordType{RecordName: "Record2", Fields: []FieldType{{Name: "Name", Type: at}}}
		env := TypeEnv{}
		out, ok := UnifyType(env, a, b)
		assert.True(t, ok)
		assert.True(t, EqualTypes(out, a))
		assert.True(t, EqualTypes(env[at.token], StringType))
	}
}
