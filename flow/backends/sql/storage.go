package sql

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/Logiraptor/oak/flow/values"
)

type SQLStorage struct {
	Conn *sql.DB
}

func (s *SQLStorage) PrepareType(t values.Type) error {
	var columns []string

	rt := t.(values.RecordType)
	for i := 0; i < rt.NumFields(); i++ {
		var field = rt.Field(i)
		sqlType, err := typeToSql(field.Type)
		if err != nil {
			return err
		}

		columns = append(columns, field.Name+" "+sqlType)
	}

	var createTable = fmt.Sprintf("create table %s (%s)", t.Name(), strings.Join(columns, ","))
	_, err := s.Conn.Exec(createTable)
	return err
}

func typeToSql(t values.Type) (string, error) {
	switch t.GetKind() {
	case values.Bool:
		return "bit", nil
	case values.String:
		return "text", nil
	case values.Int:
		return "int", nil
	default:
		return "", fmt.Errorf("Cannot prepare sql table with a column of type: %s", t)
	}
}

func (s *SQLStorage) Find(typ values.Type, matcher values.Value) ([]values.Value, error) {
	var columns = columnNames(typ)

	whereClause := ""
	if matcher != nil {
		var whereParts []string

		matchRecord := matcher.(values.RecordValue)
		for i := 0; i < matchRecord.NumFields(); i++ {
			var field = matchRecord.Field(i)
			sqlVal, err := valToSql(field.Value)
			if err != nil {
				return nil, err
			}
			whereParts = append(whereParts, fmt.Sprintf("%s = %s", field.Name, sqlVal))
		}

		if len(whereParts) > 0 {
			whereClause = fmt.Sprintf("where %s", strings.Join(whereParts, " and "))
		}
	}

	var query = fmt.Sprintf("select %s from %s %s",
		strings.Join(columns, ","),
		typ.Name(),
		whereClause)
	rows, err := s.Conn.Query(query)
	if err != nil {
		return nil, err
	}

	scanPtrs, err := createScanPtrs(typ)
	if err != nil {
		return nil, err
	}

	var output []values.Value
	for rows.Next() {
		rows.Scan(scanPtrs...)

		var rec values.MapRecord

		for i := 0; i < len(columns); i++ {
			rec = append(rec, values.Field{
				Name:  columns[i],
				Value: values.NewValue(scanPtrs[i]),
			})
		}

		output = append(output, rec)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return output, nil
}

func createScanPtrs(typ values.Type) ([]interface{}, error) {
	var output []interface{}
	rt := typ.(values.RecordType)
	for i := 0; i < rt.NumFields(); i++ {
		var field = rt.Field(i)
		switch field.Type.GetKind() {
		case values.String:
			output = append(output, new(string))
		case values.Int:
			output = append(output, new(int))
		case values.Bool:
			output = append(output, new(bool))
		default:
			return nil, fmt.Errorf("Cannot scan value of type: %s", values.TypeToString(typ))
		}
	}
	return output, nil
}

func (s *SQLStorage) Put(value values.Value) error {

	var columns = columnNames(value.GetType())
	var data []string

	var recordValue = value.(values.RecordValue)
	for i := 0; i < recordValue.NumFields(); i++ {
		var field = recordValue.Field(i)
		val, err := valToSql(field.Value)
		if err != nil {
			return err
		}
		data = append(data, val)
	}

	var insertRow = fmt.Sprintf("insert into %s (%s) values (%s)",
		value.GetType().Name(),
		strings.Join(columns, ","),
		strings.Join(data, ","))

	_, err := s.Conn.Exec(insertRow)
	return err
}

func columnNames(typ values.Type) []string {
	var columns []string

	var recordValue = typ.(values.RecordType)
	for i := 0; i < recordValue.NumFields(); i++ {
		var field = recordValue.Field(i)
		columns = append(columns, field.Name)
	}
	return columns
}

func valToSql(val values.Value) (string, error) {
	switch val.GetType().GetKind() {
	case values.Int:
		return fmt.Sprint(val.(values.IntValue).IntValue()), nil
	case values.Bool:
		var bVal = val.(values.BoolValue).BoolValue()
		if bVal {
			return "1", nil
		}
		return "0", nil
	case values.String:
		sVal := val.(values.StringValue).StringValue()
		sanitized := strings.Replace(sVal, "'", "\\'", -1)
		return fmt.Sprintf("'%s'", sanitized), nil
	default:
		return "", fmt.Errorf("cannot convert value to sql: %s", values.ValueToString(val))
	}
}
