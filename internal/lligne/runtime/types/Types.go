//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package types

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

type RecordType struct {
	FieldNameIndexes []uint64
	FieldTypeIndexes []uint64
}

func (t *RecordType) isType()      {}
func (t *RecordType) Name() string { return "Record-TBD" }

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
