package values

import (
	"fmt"
	"strings"
)

type TypeEnv map[Token]Type

type Type interface {
	Name() string
}

func NewGenericType(name string) *GenericType {
	return &GenericType{
		token: NewToken(name),
	}
}

func UnifyType(env TypeEnv, def, use Type) (Type, bool) {
	fmt.Println("Attempting to unify types", TypeToString(def), TypeToString(use))
	if EqualTypes(def, use) {
		return def, true
	}

	switch vuse := use.(type) {
	case *GenericType:
		if derived, ok := env[vuse.token]; ok {
			fmt.Println("Generic type previously derived as", derived)
			return UnifyType(env, def, derived)
		}
		fmt.Println("Generic type resolves to", def)
		env[vuse.token] = def
		return def, true
	case RecordType:
		switch vdef := def.(type) {
		case RecordType:
			var result RecordType
			result.RecordName = "UnifiedType"
			fmt.Println("Unifying record types", TypeToString(vuse), TypeToString(vdef))
		useloop:
			for _, useField := range vuse.Fields {
				for _, defField := range vdef.Fields {
					if useField.Name == defField.Name {
						fieldType, ok := UnifyType(env, defField.Type, useField.Type)
						if !ok {
							return nil, false
						}
						result.Fields = append(result.Fields, FieldType{
							Name: useField.Name,
							Type: fieldType,
						})
						continue useloop
					}
				}
				fmt.Println("Missing field ", useField.Name)
				return nil, false
			}
			return result, true
		case *GenericType:
			if derived, ok := env[vdef.token]; ok {
				return UnifyType(env, derived, use)
			}
			env[vdef.token] = vuse
			return vuse, true
		}
	case PrimitiveType:
		switch vdef := def.(type) {
		case PrimitiveType:
			if vdef == vuse {
				return vdef, true
			}
			return nil, false
		case *GenericType:
			if derived, ok := env[vdef.token]; ok {
				return UnifyType(env, derived, use)
			}
			env[vdef.token] = vuse
			return vuse, true
		}
	default:
		fmt.Printf("Unhandled type in unification: %T\n", use)
	}

	return nil, false
}

type GenericType struct {
	token Token
}

func (g *GenericType) Name() string {
	return fmt.Sprint(g.token)
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
		bprim, ok := b.(PrimitiveType)
		if !ok {
			return false
		}
		return v == bprim
	case RecordType:
		brec, ok := b.(RecordType)
		if !ok {
			return false
		}
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
	case *GenericType:
		return fmt.Sprintf("'%s", v.token.Name)
	default:
		return fmt.Sprintf("<Cannot convert type to string: %T>", typ)
	}
}

func recordTypeToString(typ RecordType) string {
	var parts []string
	for _, field := range typ.Fields {
		parts = append(parts, fmt.Sprintf("%s: %s", field.Name, TypeToString(field.Type)))
	}
	return fmt.Sprintf("record{%s}", strings.Join(parts, ", "))
}
