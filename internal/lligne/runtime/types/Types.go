//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package types

//=====================================================================================================================

// IType represents the type of expression.
type IType interface {
	isType()
}

//=====================================================================================================================

type BoolType struct {
}

func (t *BoolType) isType() {}

var BoolTypeInstance = &BoolType{}

//=====================================================================================================================

type Float64Type struct {
}

func (t *Float64Type) isType() {}

var Float64TypeInstance = &Float64Type{}

//=====================================================================================================================

type Int64Type struct {
}

func (t *Int64Type) isType() {}

var Int64TypeInstance = &Int64Type{}

//=====================================================================================================================

type StringType struct {
}

func (t *StringType) isType() {}

var StringTypeInstance = &StringType{}

//=====================================================================================================================
