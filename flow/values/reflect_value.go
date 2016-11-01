package values

import "reflect"

type reflectValue struct {
	value reflect.Value
}

func (r *reflectValue) GetType() Type {
	return &reflectType{
		typ: r.value.Type(),
	}
}

func (r *reflectValue) NumFields() int {
	return r.value.NumField()
}

func (r *reflectValue) Field(i int) Field {
	return Field{
		Name: r.value.Type().Field(i).Name,
		Value: &reflectValue{
			r.value.Field(i),
		},
	}
}

func (r *reflectValue) StringValue() string {
	return r.value.String()
}

func (r *reflectValue) IntValue() int {
	return int(r.value.Int())
}

func (r *reflectValue) BoolValue() bool {
	return bool(r.value.Bool())
}
