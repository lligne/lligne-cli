//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package bytecode

import (
	"lligne-cli/internal/lligne/runtime/pools"
	"math"
)

//=====================================================================================================================

// Machine is a stack of operands for bytecode operations.
type Machine struct {
	Stack     [1000]uint64
	Top       int
	IP        int
	IsRunning bool
}

//---------------------------------------------------------------------------------------------------------------------

func NewMachine() *Machine {
	return &Machine{Top: -1, IsRunning: true}
}

//---------------------------------------------------------------------------------------------------------------------

// BoolGetResult returns a boolean result from the top of the value stack.
func (m *Machine) BoolGetResult() bool {
	return m.Stack[m.Top] != 0
}

//---------------------------------------------------------------------------------------------------------------------

// Float64GetResult returns a 64-bit floating point result from the top of the value stack.
func (m *Machine) Float64GetResult() float64 {
	return math.Float64frombits(m.Stack[m.Top])
}

//---------------------------------------------------------------------------------------------------------------------

// Int64GetResult returns a 64-bit integer result from the top of the value stack.
func (m *Machine) Int64GetResult() int64 {
	return int64(m.Stack[m.Top])
}

//---------------------------------------------------------------------------------------------------------------------

// StringGetResult returns a string result from the top of the value stack.
func (m *Machine) StringGetResult(stringPool *pools.StringPool) string {
	return stringPool.Get(m.Stack[m.Top])
}

//=====================================================================================================================
