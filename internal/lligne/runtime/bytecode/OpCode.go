//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package bytecode

//=====================================================================================================================

// LligneOpCode is an enumeration of Lligne operation codes.
type LligneOpCode uint16

const (
	OpCodeNoOp LligneOpCode = iota

	OpCodeReturn

	// Booleans
	OpCodeBoolAnd
	OpCodeBoolLoadFalse
	OpCodeBoolLoadTrue
	OpCodeBoolOr

	// 64 Bit Integers
	OpCodeInt64Add
	OpCodeInt64Divide
	OpCodeInt64LoadInt16
	OpCodeInt64LoadOne
	OpCodeInt64LoadZero
	OpCodeInt64Multiply
	OpCodeInt64Negate
	OpCodeInt64Subtract
)

//=====================================================================================================================
