//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package types

//=====================================================================================================================

type TypeCategory uint16

const (
	TypeCategoryUnit TypeCategory = iota
	TypeCategoryBool
	TypeCategoryFloat64
	TypeCategoryInt64
	TypeCategoryString
	TypeCategoryType

	TypeCategoryOptional
	TypeCategoryRecord
)

//=====================================================================================================================

// IType represents the type of expression.
type IType interface {
	isType()
	Category() TypeCategory
	Name() string
}

//=====================================================================================================================

type BoolType struct {
}

func (t *BoolType) isType()                {}
func (t *BoolType) Category() TypeCategory { return TypeCategoryBool }
func (t *BoolType) Name() string           { return "Bool" }

var BoolTypeInstance = &BoolType{}

//=====================================================================================================================

type Float64Type struct {
}

func (t *Float64Type) isType()                {}
func (t *Float64Type) Category() TypeCategory { return TypeCategoryFloat64 }
func (t *Float64Type) Name() string           { return "Float64" }

var Float64TypeInstance = &Float64Type{}

//=====================================================================================================================

type Int64Type struct {
}

func (t *Int64Type) isType()                {}
func (t *Int64Type) Category() TypeCategory { return TypeCategoryInt64 }
func (t *Int64Type) Name() string           { return "Int64" }

var Int64TypeInstance = &Int64Type{}

//=====================================================================================================================

type RecordType struct {
	FieldNameIndexes []uint64
	FieldTypeIndexes []uint64
}

func (t *RecordType) isType()                {}
func (t *RecordType) Category() TypeCategory { return TypeCategoryRecord }
func (t *RecordType) Name() string           { return "Record-TBD" }

//=====================================================================================================================

type StringType struct {
}

func (t *StringType) isType()                {}
func (t *StringType) Category() TypeCategory { return TypeCategoryString }
func (t *StringType) Name() string           { return "String" }

var StringTypeInstance = &StringType{}

//=====================================================================================================================

type TypeType struct {
}

func (t *TypeType) isType()                {}
func (t *TypeType) Category() TypeCategory { return TypeCategoryType }
func (t *TypeType) Name() string           { return "Type" }

var TypeTypeInstance = &TypeType{}

//=====================================================================================================================

type UnitType struct {
}

func (t *UnitType) isType()                {}
func (t *UnitType) Category() TypeCategory { return TypeCategoryUnit }
func (t *UnitType) Name() string           { return "Unit" }

var UnitTypeInstance = &UnitType{}

//=====================================================================================================================
