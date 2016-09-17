package main

import (
	"fmt"
	"reflect"
)

type Type int

const (
	Int Type = iota
	String
	Bool
)

type Field struct {
	Name string
	Type Type
}

func structToFields(s interface{}) []Field {
	typ := reflect.TypeOf(s)
	out := []Field{}

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		out = append(out, Field{
			Name: field.Name,
			Type: typeFromReflect(field.Type),
		})
	}

	return out
}

func typeFromReflect(typ reflect.Type) Type {
	switch typ.Kind() {
	case reflect.Int:
		return Int
	case reflect.String:
		return String
	case reflect.Bool:
		return Bool
	default:
		panic(fmt.Sprintf("Cannot convert reflect Type %s to Type", typ))
	}
}
