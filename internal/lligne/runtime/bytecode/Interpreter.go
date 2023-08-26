//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package bytecode

import (
	"fmt"
	"math"
	"unsafe"
)

//=====================================================================================================================

type Interpreter struct {
	// TODO: Nothing needed?
}

//---------------------------------------------------------------------------------------------------------------------

// BoolGetResult returns a boolean result from the top of the value stack.
func (n *Interpreter) BoolGetResult(machine *Machine) bool {
	return machine.Stack[machine.Top] != 0
}

//---------------------------------------------------------------------------------------------------------------------

// Execute runs the op code of the given code block within the given machine.
func (n *Interpreter) Execute(machine *Machine, code *CodeBlock) {

	machine.IP = 0

	for machine.IsRunning {

		opCode := code.OpCodes[machine.IP]
		machine.IP += 1

		dispatch[opCode](machine, code)

	}

}

//---------------------------------------------------------------------------------------------------------------------

// Float64GetResult returns a 64 bit floating point result from the top of the value stack.
func (n *Interpreter) Float64GetResult(machine *Machine) float64 {
	return math.Float64frombits(machine.Stack[machine.Top])
}

//---------------------------------------------------------------------------------------------------------------------

// Int64GetResult returns a 64 bit integer result from the top of the value stack.
func (n *Interpreter) Int64GetResult(machine *Machine) int64 {
	return int64(machine.Stack[machine.Top])
}

//---------------------------------------------------------------------------------------------------------------------

// StringGetResult returns a string result from the top of the value stack.
func (n *Interpreter) StringGetResult(machine *Machine, codeBlock *CodeBlock) string {
	return codeBlock.Strings.Get(machine.Stack[machine.Top])
}

//=====================================================================================================================

const true64 uint64 = 0xFFFFFFFFFFFFFFFF

//---------------------------------------------------------------------------------------------------------------------

// dispatch is a jump table of op code handlers.
var dispatch [OpCode_Count]func(*Machine, *CodeBlock)

//---------------------------------------------------------------------------------------------------------------------

func init() {

	dispatch[OpCodeBoolAnd] = func(m *Machine, c *CodeBlock) {
		rhs := m.Stack[m.Top] != 0
		m.Top -= 1
		lhs := m.Stack[m.Top] != 0
		if lhs && rhs {
			m.Stack[m.Top] = true64
		} else {
			m.Stack[m.Top] = 0
		}
	}

	dispatch[OpCodeBoolLoadFalse] = func(m *Machine, c *CodeBlock) {
		m.Top += 1
		m.Stack[m.Top] = 0
	}

	dispatch[OpCodeBoolLoadTrue] = func(m *Machine, c *CodeBlock) {
		m.Top += 1
		m.Stack[m.Top] = true64
	}

	dispatch[OpCodeBoolNot] = func(m *Machine, c *CodeBlock) {
		if m.Stack[m.Top] == 0 {
			m.Stack[m.Top] = true64
		} else {
			m.Stack[m.Top] = 0
		}
	}

	dispatch[OpCodeBoolOr] = func(m *Machine, c *CodeBlock) {
		rhs := m.Stack[m.Top] != 0
		m.Top -= 1
		lhs := m.Stack[m.Top] != 0
		if lhs || rhs {
			m.Stack[m.Top] = true64
		} else {
			m.Stack[m.Top] = 0
		}
	}

	dispatch[OpCodeFloat64Add] = func(m *Machine, c *CodeBlock) {
		rhs := math.Float64frombits(m.Stack[m.Top])
		m.Top -= 1
		lhs := math.Float64frombits(m.Stack[m.Top])
		m.Stack[m.Top] = math.Float64bits(lhs + rhs)
	}

	dispatch[OpCodeFloat64Divide] = func(m *Machine, c *CodeBlock) {
		rhs := math.Float64frombits(m.Stack[m.Top])
		m.Top -= 1
		lhs := math.Float64frombits(m.Stack[m.Top])
		m.Stack[m.Top] = math.Float64bits(lhs / rhs)
	}

	dispatch[OpCodeFloat64Equals] = func(m *Machine, c *CodeBlock) {
		rhs := math.Float64frombits(m.Stack[m.Top])
		m.Top -= 1
		lhs := math.Float64frombits(m.Stack[m.Top])
		if lhs == rhs {
			m.Stack[m.Top] = true64
		} else {
			m.Stack[m.Top] = 0
		}
	}

	dispatch[OpCodeFloat64GreaterThan] = func(m *Machine, c *CodeBlock) {
		rhs := math.Float64frombits(m.Stack[m.Top])
		m.Top -= 1
		lhs := math.Float64frombits(m.Stack[m.Top])
		if lhs > rhs {
			m.Stack[m.Top] = true64
		} else {
			m.Stack[m.Top] = 0
		}
	}

	dispatch[OpCodeFloat64GreaterThanOrEquals] = func(m *Machine, c *CodeBlock) {
		rhs := math.Float64frombits(m.Stack[m.Top])
		m.Top -= 1
		lhs := math.Float64frombits(m.Stack[m.Top])
		if lhs >= rhs {
			m.Stack[m.Top] = true64
		} else {
			m.Stack[m.Top] = 0
		}
	}

	dispatch[OpCodeFloat64LessThan] = func(m *Machine, c *CodeBlock) {
		rhs := math.Float64frombits(m.Stack[m.Top])
		m.Top -= 1
		lhs := math.Float64frombits(m.Stack[m.Top])
		if lhs < rhs {
			m.Stack[m.Top] = true64
		} else {
			m.Stack[m.Top] = 0
		}
	}

	dispatch[OpCodeFloat64LessThanOrEquals] = func(m *Machine, c *CodeBlock) {
		rhs := math.Float64frombits(m.Stack[m.Top])
		m.Top -= 1
		lhs := math.Float64frombits(m.Stack[m.Top])
		if lhs <= rhs {
			m.Stack[m.Top] = true64
		} else {
			m.Stack[m.Top] = 0
		}
	}

	dispatch[OpCodeFloat64Load] = func(m *Machine, c *CodeBlock) {
		m.Top += 1
		m.Stack[m.Top] = *(*uint64)(unsafe.Pointer(&c.OpCodes[m.IP]))
		m.IP += 4
	}

	dispatch[OpCodeFloat64LoadOne] = func(m *Machine, c *CodeBlock) {
		m.Top += 1
		m.Stack[m.Top] = math.Float64bits(1.0)
	}

	dispatch[OpCodeFloat64LoadZero] = func(m *Machine, c *CodeBlock) {
		m.Top += 1
		m.Stack[m.Top] = 0
	}

	dispatch[OpCodeFloat64Multiply] = func(m *Machine, c *CodeBlock) {
		rhs := math.Float64frombits(m.Stack[m.Top])
		m.Top -= 1
		lhs := math.Float64frombits(m.Stack[m.Top])
		m.Stack[m.Top] = math.Float64bits(lhs * rhs)
	}

	dispatch[OpCodeFloat64Negate] = func(m *Machine, c *CodeBlock) {
		m.Stack[m.Top] = math.Float64bits(-math.Float64frombits(m.Stack[m.Top]))
	}

	dispatch[OpCodeFloat64NotEquals] = func(m *Machine, c *CodeBlock) {
		rhs := math.Float64frombits(m.Stack[m.Top])
		m.Top -= 1
		lhs := math.Float64frombits(m.Stack[m.Top])
		if lhs == rhs {
			m.Stack[m.Top] = 0
		} else {
			m.Stack[m.Top] = true64
		}
	}

	dispatch[OpCodeFloat64Subtract] = func(m *Machine, c *CodeBlock) {
		rhs := math.Float64frombits(m.Stack[m.Top])
		m.Top -= 1
		lhs := math.Float64frombits(m.Stack[m.Top])
		m.Stack[m.Top] = math.Float64bits(lhs - rhs)
	}

	dispatch[OpCodeInt64Add] = func(m *Machine, c *CodeBlock) {
		rhs := int64(m.Stack[m.Top])
		m.Top -= 1
		lhs := int64(m.Stack[m.Top])
		m.Stack[m.Top] = uint64(lhs + rhs)
	}

	dispatch[OpCodeInt64Decrement] = func(m *Machine, c *CodeBlock) {
		lhs := int64(m.Stack[m.Top])
		m.Stack[m.Top] = uint64(lhs - 1)
	}

	dispatch[OpCodeInt64Divide] = func(m *Machine, c *CodeBlock) {
		rhs := int64(m.Stack[m.Top])
		m.Top -= 1
		lhs := int64(m.Stack[m.Top])
		m.Stack[m.Top] = uint64(lhs / rhs)
	}

	dispatch[OpCodeInt64Equals] = func(m *Machine, c *CodeBlock) {
		rhs := int64(m.Stack[m.Top])
		m.Top -= 1
		lhs := int64(m.Stack[m.Top])
		if lhs == rhs {
			m.Stack[m.Top] = true64
		} else {
			m.Stack[m.Top] = 0
		}
	}

	dispatch[OpCodeInt64GreaterThan] = func(m *Machine, c *CodeBlock) {
		rhs := int64(m.Stack[m.Top])
		m.Top -= 1
		lhs := int64(m.Stack[m.Top])
		if lhs > rhs {
			m.Stack[m.Top] = true64
		} else {
			m.Stack[m.Top] = 0
		}
	}

	dispatch[OpCodeInt64GreaterThanOrEquals] = func(m *Machine, c *CodeBlock) {
		rhs := int64(m.Stack[m.Top])
		m.Top -= 1
		lhs := int64(m.Stack[m.Top])
		if lhs >= rhs {
			m.Stack[m.Top] = true64
		} else {
			m.Stack[m.Top] = 0
		}
	}

	dispatch[OpCodeInt64Increment] = func(m *Machine, c *CodeBlock) {
		lhs := int64(m.Stack[m.Top])
		m.Stack[m.Top] = uint64(lhs + 1)
	}

	dispatch[OpCodeInt64LessThan] = func(m *Machine, c *CodeBlock) {
		rhs := int64(m.Stack[m.Top])
		m.Top -= 1
		lhs := int64(m.Stack[m.Top])
		if lhs < rhs {
			m.Stack[m.Top] = true64
		} else {
			m.Stack[m.Top] = 0
		}
	}

	dispatch[OpCodeInt64LessThanOrEquals] = func(m *Machine, c *CodeBlock) {
		rhs := int64(m.Stack[m.Top])
		m.Top -= 1
		lhs := int64(m.Stack[m.Top])
		if lhs <= rhs {
			m.Stack[m.Top] = true64
		} else {
			m.Stack[m.Top] = 0
		}
	}

	dispatch[OpCodeInt64Load] = func(m *Machine, c *CodeBlock) {
		m.Top += 1
		m.Stack[m.Top] = *(*uint64)(unsafe.Pointer(&c.OpCodes[m.IP]))
		m.IP += 4
	}

	dispatch[OpCodeInt64LoadOne] = func(m *Machine, c *CodeBlock) {
		m.Top += 1
		m.Stack[m.Top] = 1
	}

	dispatch[OpCodeInt64LoadZero] = func(m *Machine, c *CodeBlock) {
		m.Top += 1
		m.Stack[m.Top] = 0
	}

	dispatch[OpCodeInt64Multiply] = func(m *Machine, c *CodeBlock) {
		rhs := int64(m.Stack[m.Top])
		m.Top -= 1
		lhs := int64(m.Stack[m.Top])
		m.Stack[m.Top] = uint64(lhs * rhs)
	}

	dispatch[OpCodeInt64Negate] = func(m *Machine, c *CodeBlock) {
		m.Stack[m.Top] = uint64(-int64(m.Stack[m.Top]))
	}

	dispatch[OpCodeInt64NotEquals] = func(m *Machine, c *CodeBlock) {
		rhs := int64(m.Stack[m.Top])
		m.Top -= 1
		lhs := int64(m.Stack[m.Top])
		if lhs == rhs {
			m.Stack[m.Top] = 0
		} else {
			m.Stack[m.Top] = true64
		}
	}

	dispatch[OpCodeInt64Subtract] = func(m *Machine, c *CodeBlock) {
		rhs := int64(m.Stack[m.Top])
		m.Top -= 1
		lhs := int64(m.Stack[m.Top])
		m.Stack[m.Top] = uint64(lhs - rhs)
	}

	dispatch[OpCodeNoOp] = func(m *Machine, c *CodeBlock) {
		// do nothing
	}

	dispatch[OpCodeReturn] = func(m *Machine, c *CodeBlock) {
		// TO DO
	}

	dispatch[OpCodeStop] = func(m *Machine, c *CodeBlock) {
		m.IsRunning = false
	}

	dispatch[OpCodeStringConcatenate] = func(m *Machine, c *CodeBlock) {
		rhs := c.Strings.Get(m.Stack[m.Top])
		m.Top -= 1
		lhs := c.Strings.Get(m.Stack[m.Top])
		m.Stack[m.Top] = c.Strings.Put(lhs + rhs)
	}

	dispatch[OpCodeStringEquals] = func(m *Machine, c *CodeBlock) {
		rhs := c.Strings.Get(m.Stack[m.Top])
		m.Top -= 1
		lhs := c.Strings.Get(m.Stack[m.Top])
		if lhs == rhs {
			m.Stack[m.Top] = true64
		} else {
			m.Stack[m.Top] = 0
		}
	}

	dispatch[OpCodeStringLoad] = func(m *Machine, c *CodeBlock) {
		m.Top += 1
		m.Stack[m.Top] = *(*uint64)(unsafe.Pointer(&c.OpCodes[m.IP]))
		m.IP += 4
	}

	dispatch[OpCodeStringNotEquals] = func(m *Machine, c *CodeBlock) {
		rhs := c.Strings.Get(m.Stack[m.Top])
		m.Top -= 1
		lhs := c.Strings.Get(m.Stack[m.Top])
		if lhs == rhs {
			m.Stack[m.Top] = 0
		} else {
			m.Stack[m.Top] = true64
		}
	}

	dispatch[OpCodeTypeEquals] = func(m *Machine, c *CodeBlock) {
		rhs := m.Stack[m.Top]
		m.Top -= 1
		lhs := m.Stack[m.Top]
		if lhs == rhs {
			m.Stack[m.Top] = true64
		} else {
			m.Stack[m.Top] = 0
		}
	}

	dispatch[OpCodeTypeLoad] = func(m *Machine, c *CodeBlock) {
		m.Top += 1
		m.Stack[m.Top] = uint64(c.OpCodes[m.IP])
		m.IP += 1
	}

	dispatch[OpCodeTypeNotEquals] = func(m *Machine, c *CodeBlock) {
		rhs := m.Stack[m.Top]
		m.Top -= 1
		lhs := m.Stack[m.Top]
		if lhs == rhs {
			m.Stack[m.Top] = 0
		} else {
			m.Stack[m.Top] = true64
		}
	}

	for i := uint16(0); i < OpCode_Count; i += 1 {
		if dispatch[i] == nil {
			panic(fmt.Sprintf("Missing dispatch function %d", i))
		}
	}

}

//=====================================================================================================================
