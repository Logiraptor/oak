package values

import (
	"fmt"
	"strings"
)

type Value interface {
	GetType() Type
}

var _ = Value(IntValue(0))
var _ = Value(StringValue(""))
var _ = Value(BoolValue(false))
var _ = Value(RecordValue{})
var _ = Value(ListValue{})

type IntValue int
type StringValue string
type BoolValue bool
type RecordValue struct {
	Name   string
	Fields []Field
}
type ListValue struct {
	Elements []Value
	Type     Type
}

type Field struct {
	Value
	Name string
}

func (IntValue) GetType() Type {
	return IntType
}

func (StringValue) GetType() Type {
	return StringType
}

func (BoolValue) GetType() Type {
	return BoolType
}

func (r RecordValue) GetType() Type {
	var output = RecordType{
		RecordName: r.Name,
	}

	for _, field := range r.Fields {
		output.Fields = append(output.Fields, FieldType{
			Name: field.Name,
			Type: field.Value.GetType(),
		})
	}

	return output
}

func (r RecordValue) FieldByToken(tok Token) Value {
	for _, field := range r.Fields {
		if field.Name == tok.Name {
			return field.Value
		}
	}
	panic(fmt.Sprintf("No such field %s on record type %s", tok.Name, r.Name))
}

func (l ListValue) GetType() Type {
	return ListType{
		ElementType: l.Type,
	}
}

func EqualValues(a, b Value) bool {
	if !EqualTypes(a.GetType(), b.GetType()) {
		return false
	}
	switch v := a.(type) {
	case StringValue:
		return v == b.(StringValue)
	case IntValue:
		return v == b.(IntValue)
	case BoolValue:
		return v == b.(BoolValue)
	case RecordValue:
		arec := v
		brec := b.(RecordValue)
		for i, aField := range arec.Fields {
			bField := brec.Fields[i]
			if aField.Name != bField.Name {
				return false
			}
			if !EqualValues(aField.Value, bField.Value) {
				return false
			}
		}
		return true
	case ListValue:
		alist := a.(ListValue)
		blist := b.(ListValue)
		if len(alist.Elements) != len(blist.Elements) {
			return false
		}

		for i, aElem := range alist.Elements {
			var bElem = blist.Elements[i]
			if !EqualValues(aElem, bElem) {
				return false
			}
		}

		return true
	}
	return false
}

func ValueToString(val Value) string {
	switch v := val.(type) {
	case RecordValue:
		return recordValueToString(v)
	case IntValue, StringValue, BoolValue:
		return fmt.Sprint(v)
	case ListValue:
		return listValueToString(v)
	}
	panic(fmt.Sprintf("Cannot convert value type: %s to string", val))
}

func recordValueToString(v RecordValue) string {
	var parts = []string{}
	for _, field := range v.Fields {
		parts = append(parts, fmt.Sprintf("%s= %s", field.Name, ValueToString(field.Value)))
	}
	return fmt.Sprintf("{%s}", strings.Join(parts, ", "))
}

func listValueToString(v ListValue) string {
	var parts = []string{}
	for _, elem := range v.Elements {
		parts = append(parts, ValueToString(elem))
	}
	return fmt.Sprintf("[%s]", strings.Join(parts, ", "))
}
