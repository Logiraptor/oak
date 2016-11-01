package values

import (
	"fmt"
	"reflect"
)

func NewValue(i interface{}) Value {
	v := reflect.ValueOf(i)

	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Int:
		return IntValue(v.Int())
	case reflect.String:
		return StringValue(v.String())
	case reflect.Bool:
		return BoolValue(v.Bool())
	case reflect.Struct:
		var output RecordValue
		output.Name = v.Type().Name()
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			if v.Type().Field(i).PkgPath == "" {
				output.Fields = append(output.Fields, Field{
					Name:  v.Type().Field(i).Name,
					Value: NewValue(field.Interface()),
				})
			}
		}
		return output
	case reflect.Slice:
		var output ListValue
		var inner = v.Type().Elem()
		var innerVal = reflect.New(inner).Elem()
		var valOfInner = NewValue(innerVal.Interface())
		var typeOfInner = valOfInner.GetType()

		output.Type = typeOfInner
		for i := 0; i < v.Len(); i++ {
			output.Elements = append(output.Elements, NewValue(v.Index(i).Interface()))
		}
		return output
	default:
		panic(fmt.Sprintf("Cannot create a Value from %s", v.Type().String()))
	}
}
