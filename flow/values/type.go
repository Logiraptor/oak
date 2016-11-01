package values

import (
	"fmt"
	"strings"
)

type Type interface {
	Name() string
}

type RecordType struct {
	RecordName string
	Fields     []FieldType
}

func (r RecordType) Name() string {
	return r.RecordName
}

type FieldType struct {
	Name string
	Type
}

type ListType struct {
	ElementType Type
}

func (l ListType) Name() string {
	return fmt.Sprintf("[%s]", l.ElementType.Name())
}

type PrimitiveType Kind

func (p PrimitiveType) GetKind() Kind {
	return Kind(p)
}

func (p PrimitiveType) Name() string {
	switch p {
	case IntType:
		return "int"
	case StringType:
		return "string"
	case BoolType:
		return "bool"
	}
	panic(fmt.Sprintf("Cannot name primitiveType: %d", p))
}

var _ = Type(PrimitiveType(0))

const (
	IntType    PrimitiveType = PrimitiveType(Int)
	StringType PrimitiveType = PrimitiveType(String)
	BoolType   PrimitiveType = PrimitiveType(Bool)
)

type Kind int

const (
	String Kind = iota
	Int
	Bool
	Record
	List
)

func EqualTypes(a, b Type) bool {
	switch v := a.(type) {
	case PrimitiveType:
		return v == b.(PrimitiveType)
	case RecordType:
		brec := b.(RecordType)
		if len(v.Fields) != len(brec.Fields) {
			return false
		}
		for i, aField := range v.Fields {
			bField := brec.Fields[i]
			if aField.Name != bField.Name {
				return false
			}
			if !EqualTypes(aField.Type, bField.Type) {
				return false
			}
		}
		return true
	case ListType:
		return EqualTypes(v.ElementType, b.(ListType).ElementType)
	}
	return false
}

func TypeToString(typ Type) string {
	switch v := typ.(type) {
	case PrimitiveType:
		return v.Name()
	case RecordType:
		return recordTypeToString(v)
	case ListType:
		return fmt.Sprintf("[%s]", TypeToString(v.ElementType))
	default:
		panic(fmt.Sprintf("Cannot convert type to string: %s", typ))
	}
}

func recordTypeToString(typ RecordType) string {
	var parts []string
	for _, field := range typ.Fields {
		parts = append(parts, fmt.Sprintf("%s: %s", field.Name, TypeToString(field.Type)))
	}
	return fmt.Sprintf("record{%s}", strings.Join(parts, ", "))
}
