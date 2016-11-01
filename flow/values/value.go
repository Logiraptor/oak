package values

import (
	"fmt"
	"reflect"
	"strings"
)

type Value interface {
	GetType() Type
}

type IntValue interface {
	Value
	IntValue() int
}

type StringValue interface {
	Value
	StringValue() string
}

type BoolValue interface {
	Value
	BoolValue() bool
}

type RecordValue interface {
	Value
	NumFields() int
	Field(int) Field
}

type Field struct {
	Value
	Name string
}

type ListValue interface {
	Value
	Len() int
	Index(int) Value
}

func NewValue(i interface{}) Value {
	v := reflect.ValueOf(i)
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return &reflectValue{
		value: v,
	}
}

func EqualValues(a, b Value) bool {
	if !EqualTypes(a.GetType(), b.GetType()) {
		return false
	}
	switch a.GetType().GetKind() {
	case String:
		return a.(StringValue).StringValue() == b.(StringValue).StringValue()
	case Int:
		return a.(IntValue).IntValue() == b.(IntValue).IntValue()
	case Bool:
		return a.(BoolValue).BoolValue() == b.(BoolValue).BoolValue()
	case Record:
		arec := a.(RecordValue)
		brec := b.(RecordValue)
		for i := 0; i < arec.NumFields(); i++ {
			aField := arec.Field(i)
			bField := brec.Field(i)
			if aField.Name != bField.Name {
				return false
			}
			if !EqualValues(aField.Value, bField.Value) {
				return false
			}
		}
		return true
	case List:
		alist := a.(ListValue)
		blist := b.(ListValue)
		if alist.Len() != blist.Len() {
			return false
		}

		for i := 0; i < alist.Len(); i++ {
			var aElem = alist.Index(i)
			var bElem = blist.Index(i)
			if !EqualValues(aElem, bElem) {
				return false
			}
		}

		return true
	}
	return false
}

func ValueToString(v Value) string {
	switch v.GetType().GetKind() {
	case Record:
		return recordValueToString(v.(RecordValue))
	case Int:
		return fmt.Sprint(v.(IntValue).IntValue())
	case String:
		return fmt.Sprint(v.(StringValue).StringValue())
	case Bool:
		return fmt.Sprint(v.(BoolValue).BoolValue())
	case List:
		return listValueToString(v.(ListValue))
	}
	panic(fmt.Sprintf("Cannot convert value type: %s to string", v.GetType().GetKind()))
}

func recordValueToString(v RecordValue) string {
	var parts = []string{}
	for i := 0; i < v.NumFields(); i++ {
		var field = v.Field(i)
		parts = append(parts, fmt.Sprintf("%s= %s", field.Name, ValueToString(field.Value)))
	}
	return fmt.Sprintf("{%s}", strings.Join(parts, ", "))
}

func listValueToString(v ListValue) string {
	var parts = []string{}
	for i := 0; i < v.Len(); i++ {
		elem := v.Index(i)
		parts = append(parts, ValueToString(elem))
	}
	return fmt.Sprintf("[%s]", strings.Join(parts, ", "))
}
