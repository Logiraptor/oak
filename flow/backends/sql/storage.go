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
	for _, field := range rt.Fields {
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
	switch v := t.(type) {
	case values.PrimitiveType:
		switch v {
		case values.BoolType:
			return "bit", nil
		case values.StringType:
			return "text", nil
		case values.IntType:
			return "int", nil
		}
	}
	return "", fmt.Errorf("Cannot prepare sql table with a column of type: %s", values.TypeToString(t))
}

func (s *SQLStorage) Find(typ values.Type, matcher values.Value) ([]values.Value, error) {
	var columns = columnNames(typ)

	whereClause := ""
	if matcher != nil {
		var whereParts []string

		matchRecord := matcher.(values.RecordValue)
		for _, field := range matchRecord.Fields {
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

		var rec values.RecordValue
		rec.Name = typ.Name()

		for i := 0; i < len(columns); i++ {
			rec.Fields = append(rec.Fields, values.Field{
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
	for _, field := range rt.Fields {
		switch p := field.Type.(type) {
		case values.PrimitiveType:
			switch p {
			case values.StringType:
				output = append(output, new(string))
			case values.IntType:
				output = append(output, new(int))
			case values.BoolType:
				output = append(output, new(bool))
			default:
				return nil, fmt.Errorf("Cannot scan value of type: %s", values.TypeToString(typ))
			}
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
	for _, field := range recordValue.Fields {
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
	for _, field := range recordValue.Fields {
		columns = append(columns, field.Name)
	}
	return columns
}

func valToSql(val values.Value) (string, error) {
	switch v := val.(type) {
	case values.IntValue:
		return fmt.Sprint(v), nil
	case values.BoolValue:
		if v {
			return "1", nil
		}
		return "0", nil
	case values.StringValue:
		sanitized := strings.Replace(string(v), "'", "\\'", -1)
		return fmt.Sprintf("'%s'", sanitized), nil
	default:
		return "", fmt.Errorf("cannot convert value to sql: %s", values.ValueToString(val))
	}
}
