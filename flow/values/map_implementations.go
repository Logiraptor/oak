package values

import "fmt"

type MapRecord []Field

type MapType []FieldType

func (m MapRecord) GetType() Type {
	var t MapType
	for _, f := range m {
		t = append(t, FieldType{
			Name: f.Name,
			Type: f.GetType(),
		})
	}
	return t
}

func (m MapRecord) NumFields() int {
	return len(m)
}

func (m MapRecord) Field(i int) Field {
	return m[i]
}

func (m MapType) Name() string {
	return fmt.Sprint("MapRecord")
}

func (m MapType) GetKind() Kind {
	return Record
}

func (m MapType) Field(i int) FieldType {
	return m[i]
}

func (m MapType) NumFields() int {
	return len(m)
}

var _ = RecordValue(MapRecord{})
var _ = RecordType(MapType{})
