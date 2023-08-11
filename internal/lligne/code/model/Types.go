//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package model

//=====================================================================================================================

// IType represents the type of expression.
type IType interface {
	BaseType() BaseType
}

//=====================================================================================================================

// BaseType is an enumeration of Lligne data types.
type BaseType int

const (
	BaseTypeBoolean BaseType = 1 + iota
	BaseTypeFloat64
	BaseTypeInt64
)

//=====================================================================================================================

type Type struct {
	baseType BaseType
}

func (t *Type) BaseType() BaseType {
	return t.baseType
}

//=====================================================================================================================

type BoolType struct {
	Type
}

func NewBoolType() IType {
	return &BoolType{Type: Type{baseType: BaseTypeBoolean}}
}

//=====================================================================================================================

type Float64Type struct {
	Type
}

func NewFloat64Type() IType {
	return &Float64Type{Type: Type{baseType: BaseTypeFloat64}}
}

//=====================================================================================================================

type Int64Type struct {
	Type
}

func NewInt64Type() IType {
	return &Int64Type{Type: Type{baseType: BaseTypeInt64}}
}

//=====================================================================================================================
