package values

import "reflect"

type reflectType struct {
	typ reflect.Type
}

func (r *reflectType) GetKind() Kind {
	switch r.typ.Kind() {
	case reflect.String:
		return String
	case reflect.Int:
		return Int
	case reflect.Bool:
		return Bool
	case reflect.Struct:
		return Record
	case reflect.Slice:
		return List
	default:
		panic("unsupported kind in Value: " + r.typ.Kind().String())
	}
}

func (r *reflectType) Name() string {
	return r.typ.Name()
}

func (r *reflectType) NumFields() int {
	return r.typ.NumField()
}

func (r *reflectType) Field(i int) FieldType {
	f := r.typ.Field(i)
	return FieldType{
		Name: f.Name,
		Type: &reflectType{
			typ: f.Type,
		},
	}
}
