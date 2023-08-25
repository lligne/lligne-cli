//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package types

//=====================================================================================================================

// TokenType is an enumeration of Lligne token types.
type BuiltInType uint16

const (
	BuiltInTypeType BuiltInType = iota

	BuiltInTypeBool
	BuiltInTypeFloat64
	BuiltInTypeInt64
	BuiltInTypeString

	BuiltInType_Count
)

//---------------------------------------------------------------------------------------------------------------------

var BuiltInTypesByName = make(map[string]BuiltInType)

func init() {
	BuiltInTypesByName["Type"] = BuiltInTypeType
	BuiltInTypesByName["Bool"] = BuiltInTypeBool
	BuiltInTypesByName["Float64"] = BuiltInTypeFloat64
	BuiltInTypesByName["Int64"] = BuiltInTypeInt64
	BuiltInTypesByName["String"] = BuiltInTypeString
}

//=====================================================================================================================

// IType represents the type of expression.
type IType interface {
	isType()
	Name() string
}

//=====================================================================================================================

type BoolType struct {
}

func (t *BoolType) isType()      {}
func (t *BoolType) Name() string { return "Bool" }

var BoolTypeInstance = &BoolType{}

//=====================================================================================================================

type Float64Type struct {
}

func (t *Float64Type) isType()      {}
func (t *Float64Type) Name() string { return "Float64" }

var Float64TypeInstance = &Float64Type{}

//=====================================================================================================================

type Int64Type struct {
}

func (t *Int64Type) isType()      {}
func (t *Int64Type) Name() string { return "Int64" }

var Int64TypeInstance = &Int64Type{}

//=====================================================================================================================

type StringType struct {
}

func (t *StringType) isType()      {}
func (t *StringType) Name() string { return "String" }

var StringTypeInstance = &StringType{}

//=====================================================================================================================

type TypeType struct {
}

func (t *TypeType) isType()      {}
func (t *TypeType) Name() string { return "Type" }

var TypeTypeInstance = &TypeType{}

//=====================================================================================================================
