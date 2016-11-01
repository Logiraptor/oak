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
	case reflect.Struct:
		var output RecordValue
		output.Name = v.Type().Name()
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			output.Fields = append(output.Fields, Field{
				Name:  v.Type().Field(i).Name,
				Value: NewValue(field.Interface()),
			})
		}
		return output
	default:
		panic(fmt.Sprintf("Cannot create a Value from %s", v.Type().String()))
	}
}
