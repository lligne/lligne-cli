//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package bytecode

//=====================================================================================================================

// LligneOpCode is an enumeration of Lligne byte codes.
type LligneOpCode uint8

const (
	OpCodeNoOp LligneOpCode = iota

	OpCodeReturn

	// Integers
	OpCodeInt64Add
	OpCodeInt64LoadInt16
	OpCodeInt64LoadOne
	OpCodeInt64LoadZero
)

//=====================================================================================================================
