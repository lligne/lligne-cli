//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package bytecode

//=====================================================================================================================

const (
	OpCodeNoOp uint16 = iota
	OpCodeStop
	OpCodeReturn

	// Booleans
	OpCodeBoolAnd
	OpCodeBoolLoadFalse
	OpCodeBoolLoadTrue
	OpCodeBoolNot
	OpCodeBoolOr

	// 64 Bit Floating Point
	OpCodeFloat64Add
	OpCodeFloat64Divide
	OpCodeFloat64Equals
	OpCodeFloat64GreaterThan
	OpCodeFloat64GreaterThanOrEquals
	OpCodeFloat64LessThan
	OpCodeFloat64LessThanOrEquals
	OpCodeFloat64Load
	OpCodeFloat64LoadOne
	OpCodeFloat64LoadZero
	OpCodeFloat64Multiply
	OpCodeFloat64Negate
	OpCodeFloat64NotEquals
	OpCodeFloat64Subtract

	// 64 Bit Integers
	OpCodeInt64Add
	OpCodeInt64Decrement
	OpCodeInt64Divide
	OpCodeInt64Equals
	OpCodeInt64GreaterThan
	OpCodeInt64GreaterThanOrEquals
	OpCodeInt64Increment
	OpCodeInt64LessThan
	OpCodeInt64LessThanOrEquals
	OpCodeInt64Load
	OpCodeInt64LoadOne
	OpCodeInt64LoadZero
	OpCodeInt64Multiply
	OpCodeInt64Negate
	OpCodeInt64NotEquals
	OpCodeInt64Subtract

	// Strings
	OpCodeStringConcatenate
	OpCodeStringEquals
	OpCodeStringLoad
	OpCodeStringNotEquals

	// Types
	OpCodeTypeEquals
	OpCodeTypeLoad
	OpCodeTypeNotEquals

	// Records
	OpCodeRecordEquals
	OpCodeRecordFieldIndexLoad
	OpCodeRecordFieldReference
	OpCodeRecordNotEquals
	OpCodeRecordStore

	OpCode_Count
)

//=====================================================================================================================
