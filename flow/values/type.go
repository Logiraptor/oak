package values

import (
	"fmt"
	"strings"
)

type Type interface {
	Name() string
	GetKind() Kind
}

type RecordType interface {
	Type
	NumFields() int
	Field(int) FieldType
}

type FieldType struct {
	Type
	Name string
}

type ListType interface {
	Type
	ElementType() Type
}

type Kind int

const (
	String Kind = iota
	Int
	Bool
	Record
	List
)

func EqualTypes(a, b Type) bool {
	if a.GetKind() != b.GetKind() {
		return false
	}
	switch a.GetKind() {
	case String, Int, Bool:
		return true
	case Record:
		arec := a.(RecordType)
		brec := b.(RecordType)
		if arec.NumFields() != brec.NumFields() {
			return false
		}
		for i := 0; i < arec.NumFields(); i++ {
			aField := arec.Field(i)
			bField := brec.Field(i)
			if aField.Name != bField.Name {
				return false
			}
			if !EqualTypes(aField.Type, bField.Type) {
				return false
			}
		}
		return true
	case List:
		return EqualTypes(a.(ListType).ElementType(), b.(ListType).ElementType())
	}
	return false
}

func TypeToString(typ Type) string {
	switch typ.GetKind() {
	case Record:
		return recordTypeToString(typ.(RecordType))
	case List:
		return fmt.Sprintf("[%s]", TypeToString(typ.(ListType).ElementType()))
	case Int:
		return "int"
	case String:
		return "string"
	case Bool:
		return "bool"
	default:
		panic(fmt.Sprintf("Cannot convert type to string: %s", typ))
	}
}

func recordTypeToString(typ RecordType) string {
	var parts []string
	for i := 0; i < typ.NumFields(); i++ {
		var field = typ.Field(i)
		parts = append(parts, fmt.Sprintf("%s: %s", field.Name, TypeToString(field.Type)))
	}
	return fmt.Sprintf("record{%s}", strings.Join(parts, ", "))
}
