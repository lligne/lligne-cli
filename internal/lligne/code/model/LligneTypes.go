//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package model

//=====================================================================================================================

// ILligneType represents the type of expression.
type ILligneType interface {
	BaseType() LligneBaseType
}

//=====================================================================================================================

// LligneType is an enumeration of Lligne data types.
type LligneBaseType int

const (
	BaseTypeBoolean LligneBaseType = 1 + iota
	BaseTypeFloat64
	BaseTypeInt64
)

//=====================================================================================================================

type LligneType struct {
	baseType LligneBaseType
}

func (t *LligneType) BaseType() LligneBaseType {
	return t.baseType
}

//=====================================================================================================================

type LligneBoolType struct {
	LligneType
}

func NewBoolType() ILligneType {
	return &LligneBoolType{LligneType: LligneType{baseType: BaseTypeBoolean}}
}

//=====================================================================================================================

type LligneFloat64Type struct {
	LligneType
}

func NewFloat64Type() ILligneType {
	return &LligneFloat64Type{LligneType: LligneType{baseType: BaseTypeFloat64}}
}

//=====================================================================================================================

type LligneInt64Type struct {
	LligneType
}

func NewInt64Type() ILligneType {
	return &LligneInt64Type{LligneType: LligneType{baseType: BaseTypeInt64}}
}

//=====================================================================================================================
