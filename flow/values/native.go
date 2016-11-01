package values

import (
	"fmt"
	"reflect"
)

func FillNative(val Value, out interface{}) {
	outVal := reflect.ValueOf(out)
	for outVal.Kind() == reflect.Ptr {
		outVal = outVal.Elem()
	}
	reflectFillNative(val, outVal)
}

func reflectFillNative(val Value, outVal reflect.Value) {
	switch v := val.(type) {
	case StringValue:
		outVal.Set(reflect.ValueOf(string(v)))
	case IntValue:
		outVal.Set(reflect.ValueOf(int(v)))
	case BoolValue:
		outVal.Set(reflect.ValueOf(bool(v)))
	case RecordValue:
		for _, field := range v.Fields {
			outField := outVal.FieldByName(field.Name)
			reflectFillNative(field.Value, outField)
		}
	case ListValue:
		var output = reflect.MakeSlice(outVal.Type(), len(v.Elements), len(v.Elements))
		for i, value := range v.Elements {
			reflectFillNative(value, output.Index(i))
		}
		outVal.Set(output)
	default:
		panic(fmt.Sprintf("Cannot fill native type: %T", v))
	}
}

func ToNative(val Value) interface{} {
	ptr := createNativePtr(val.GetType())
	fmt.Printf("%T", ptr)
	FillNative(val, ptr)
	return ptr
}

func createNativePtr(typ Type) interface{} {
	return reflect.New(reflectCreateNativePtr(typ)).Interface()
}

func reflectCreateNativePtr(typ Type) reflect.Type {
	switch v := typ.(type) {
	case PrimitiveType:
		switch v {
		case StringType:
			return reflect.TypeOf("")
		case IntType:
			return reflect.TypeOf(0)
		case BoolType:
			return reflect.TypeOf(false)
		}
	case RecordType:
		var fields []reflect.StructField
		for _, field := range v.Fields {
			fields = append(fields, reflect.StructField{
				Name: field.Name,
				Type: reflectCreateNativePtr(field.Type),
			})
		}
		return reflect.StructOf(fields)
	case ListType:
		return reflect.SliceOf(reflectCreateNativePtr(v.ElementType))
	}
	panic(fmt.Sprintf("Cannot create native type for Type: %T", typ))
}
