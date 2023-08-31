//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package bytecode

import (
	"fmt"
	"lligne-cli/internal/lligne/runtime/pools"
	"math"
	"unsafe"
)

//=====================================================================================================================

type Interpreter struct {
	codeBlock  *CodeBlock
	stringPool *pools.StringPool
}

//---------------------------------------------------------------------------------------------------------------------

func NewInterpreter(codeBlock *CodeBlock, stringPool *pools.StringPool) *Interpreter {
	return &Interpreter{
		codeBlock:  codeBlock,
		stringPool: stringPool,
	}
}

//---------------------------------------------------------------------------------------------------------------------

// Execute runs the op code of the given code block within the given machine.
func (n *Interpreter) Execute(machine *Machine) {

	machine.IP = 0

	for machine.IsRunning {

		opCode := n.codeBlock.OpCodes[machine.IP]
		machine.IP += 1

		dispatch[opCode](n, machine)

	}

}

//=====================================================================================================================

const true64 uint64 = 0xFFFFFFFFFFFFFFFF

//---------------------------------------------------------------------------------------------------------------------

// dispatch is a jump table of op code handlers.
var dispatch [OpCode_Count]func(*Interpreter, *Machine)

//---------------------------------------------------------------------------------------------------------------------

func init() {

	dispatch[OpCodeBoolAnd] = func(n *Interpreter, m *Machine) {
		rhs := m.Stack[m.Top] != 0
		m.Top -= 1
		lhs := m.Stack[m.Top] != 0
		if lhs && rhs {
			m.Stack[m.Top] = true64
		} else {
			m.Stack[m.Top] = 0
		}
	}

	dispatch[OpCodeBoolLoadFalse] = func(n *Interpreter, m *Machine) {
		m.Top += 1
		m.Stack[m.Top] = 0
	}

	dispatch[OpCodeBoolLoadTrue] = func(n *Interpreter, m *Machine) {
		m.Top += 1
		m.Stack[m.Top] = true64
	}

	dispatch[OpCodeBoolNot] = func(n *Interpreter, m *Machine) {
		if m.Stack[m.Top] == 0 {
			m.Stack[m.Top] = true64
		} else {
			m.Stack[m.Top] = 0
		}
	}

	dispatch[OpCodeBoolOr] = func(n *Interpreter, m *Machine) {
		rhs := m.Stack[m.Top] != 0
		m.Top -= 1
		lhs := m.Stack[m.Top] != 0
		if lhs || rhs {
			m.Stack[m.Top] = true64
		} else {
			m.Stack[m.Top] = 0
		}
	}

	dispatch[OpCodeFloat64Add] = func(n *Interpreter, m *Machine) {
		rhs := math.Float64frombits(m.Stack[m.Top])
		m.Top -= 1
		lhs := math.Float64frombits(m.Stack[m.Top])
		m.Stack[m.Top] = math.Float64bits(lhs + rhs)
	}

	dispatch[OpCodeFloat64Divide] = func(n *Interpreter, m *Machine) {
		rhs := math.Float64frombits(m.Stack[m.Top])
		m.Top -= 1
		lhs := math.Float64frombits(m.Stack[m.Top])
		m.Stack[m.Top] = math.Float64bits(lhs / rhs)
	}

	dispatch[OpCodeFloat64Equals] = func(n *Interpreter, m *Machine) {
		rhs := math.Float64frombits(m.Stack[m.Top])
		m.Top -= 1
		lhs := math.Float64frombits(m.Stack[m.Top])
		if lhs == rhs {
			m.Stack[m.Top] = true64
		} else {
			m.Stack[m.Top] = 0
		}
	}

	dispatch[OpCodeFloat64GreaterThan] = func(n *Interpreter, m *Machine) {
		rhs := math.Float64frombits(m.Stack[m.Top])
		m.Top -= 1
		lhs := math.Float64frombits(m.Stack[m.Top])
		if lhs > rhs {
			m.Stack[m.Top] = true64
		} else {
			m.Stack[m.Top] = 0
		}
	}

	dispatch[OpCodeFloat64GreaterThanOrEquals] = func(n *Interpreter, m *Machine) {
		rhs := math.Float64frombits(m.Stack[m.Top])
		m.Top -= 1
		lhs := math.Float64frombits(m.Stack[m.Top])
		if lhs >= rhs {
			m.Stack[m.Top] = true64
		} else {
			m.Stack[m.Top] = 0
		}
	}

	dispatch[OpCodeFloat64LessThan] = func(n *Interpreter, m *Machine) {
		rhs := math.Float64frombits(m.Stack[m.Top])
		m.Top -= 1
		lhs := math.Float64frombits(m.Stack[m.Top])
		if lhs < rhs {
			m.Stack[m.Top] = true64
		} else {
			m.Stack[m.Top] = 0
		}
	}

	dispatch[OpCodeFloat64LessThanOrEquals] = func(n *Interpreter, m *Machine) {
		rhs := math.Float64frombits(m.Stack[m.Top])
		m.Top -= 1
		lhs := math.Float64frombits(m.Stack[m.Top])
		if lhs <= rhs {
			m.Stack[m.Top] = true64
		} else {
			m.Stack[m.Top] = 0
		}
	}

	dispatch[OpCodeFloat64Load] = func(n *Interpreter, m *Machine) {
		m.Top += 1
		m.Stack[m.Top] = *(*uint64)(unsafe.Pointer(&n.codeBlock.OpCodes[m.IP]))
		m.IP += 4
	}

	dispatch[OpCodeFloat64LoadOne] = func(n *Interpreter, m *Machine) {
		m.Top += 1
		m.Stack[m.Top] = math.Float64bits(1.0)
	}

	dispatch[OpCodeFloat64LoadZero] = func(n *Interpreter, m *Machine) {
		m.Top += 1
		m.Stack[m.Top] = 0
	}

	dispatch[OpCodeFloat64Multiply] = func(n *Interpreter, m *Machine) {
		rhs := math.Float64frombits(m.Stack[m.Top])
		m.Top -= 1
		lhs := math.Float64frombits(m.Stack[m.Top])
		m.Stack[m.Top] = math.Float64bits(lhs * rhs)
	}

	dispatch[OpCodeFloat64Negate] = func(n *Interpreter, m *Machine) {
		m.Stack[m.Top] = math.Float64bits(-math.Float64frombits(m.Stack[m.Top]))
	}

	dispatch[OpCodeFloat64NotEquals] = func(n *Interpreter, m *Machine) {
		rhs := math.Float64frombits(m.Stack[m.Top])
		m.Top -= 1
		lhs := math.Float64frombits(m.Stack[m.Top])
		if lhs == rhs {
			m.Stack[m.Top] = 0
		} else {
			m.Stack[m.Top] = true64
		}
	}

	dispatch[OpCodeFloat64Subtract] = func(n *Interpreter, m *Machine) {
		rhs := math.Float64frombits(m.Stack[m.Top])
		m.Top -= 1
		lhs := math.Float64frombits(m.Stack[m.Top])
		m.Stack[m.Top] = math.Float64bits(lhs - rhs)
	}

	dispatch[OpCodeInt64Add] = func(n *Interpreter, m *Machine) {
		rhs := int64(m.Stack[m.Top])
		m.Top -= 1
		lhs := int64(m.Stack[m.Top])
		m.Stack[m.Top] = uint64(lhs + rhs)
	}

	dispatch[OpCodeInt64Decrement] = func(n *Interpreter, m *Machine) {
		lhs := int64(m.Stack[m.Top])
		m.Stack[m.Top] = uint64(lhs - 1)
	}

	dispatch[OpCodeInt64Divide] = func(n *Interpreter, m *Machine) {
		rhs := int64(m.Stack[m.Top])
		m.Top -= 1
		lhs := int64(m.Stack[m.Top])
		m.Stack[m.Top] = uint64(lhs / rhs)
	}

	dispatch[OpCodeInt64Equals] = func(n *Interpreter, m *Machine) {
		rhs := int64(m.Stack[m.Top])
		m.Top -= 1
		lhs := int64(m.Stack[m.Top])
		if lhs == rhs {
			m.Stack[m.Top] = true64
		} else {
			m.Stack[m.Top] = 0
		}
	}

	dispatch[OpCodeInt64GreaterThan] = func(n *Interpreter, m *Machine) {
		rhs := int64(m.Stack[m.Top])
		m.Top -= 1
		lhs := int64(m.Stack[m.Top])
		if lhs > rhs {
			m.Stack[m.Top] = true64
		} else {
			m.Stack[m.Top] = 0
		}
	}

	dispatch[OpCodeInt64GreaterThanOrEquals] = func(n *Interpreter, m *Machine) {
		rhs := int64(m.Stack[m.Top])
		m.Top -= 1
		lhs := int64(m.Stack[m.Top])
		if lhs >= rhs {
			m.Stack[m.Top] = true64
		} else {
			m.Stack[m.Top] = 0
		}
	}

	dispatch[OpCodeInt64Increment] = func(n *Interpreter, m *Machine) {
		lhs := int64(m.Stack[m.Top])
		m.Stack[m.Top] = uint64(lhs + 1)
	}

	dispatch[OpCodeInt64LessThan] = func(n *Interpreter, m *Machine) {
		rhs := int64(m.Stack[m.Top])
		m.Top -= 1
		lhs := int64(m.Stack[m.Top])
		if lhs < rhs {
			m.Stack[m.Top] = true64
		} else {
			m.Stack[m.Top] = 0
		}
	}

	dispatch[OpCodeInt64LessThanOrEquals] = func(n *Interpreter, m *Machine) {
		rhs := int64(m.Stack[m.Top])
		m.Top -= 1
		lhs := int64(m.Stack[m.Top])
		if lhs <= rhs {
			m.Stack[m.Top] = true64
		} else {
			m.Stack[m.Top] = 0
		}
	}

	dispatch[OpCodeInt64Load] = func(n *Interpreter, m *Machine) {
		m.Top += 1
		m.Stack[m.Top] = *(*uint64)(unsafe.Pointer(&n.codeBlock.OpCodes[m.IP]))
		m.IP += 4
	}

	dispatch[OpCodeInt64LoadOne] = func(n *Interpreter, m *Machine) {
		m.Top += 1
		m.Stack[m.Top] = 1
	}

	dispatch[OpCodeInt64LoadZero] = func(n *Interpreter, m *Machine) {
		m.Top += 1
		m.Stack[m.Top] = 0
	}

	dispatch[OpCodeInt64Multiply] = func(n *Interpreter, m *Machine) {
		rhs := int64(m.Stack[m.Top])
		m.Top -= 1
		lhs := int64(m.Stack[m.Top])
		m.Stack[m.Top] = uint64(lhs * rhs)
	}

	dispatch[OpCodeInt64Negate] = func(n *Interpreter, m *Machine) {
		m.Stack[m.Top] = uint64(-int64(m.Stack[m.Top]))
	}

	dispatch[OpCodeInt64NotEquals] = func(n *Interpreter, m *Machine) {
		rhs := int64(m.Stack[m.Top])
		m.Top -= 1
		lhs := int64(m.Stack[m.Top])
		if lhs == rhs {
			m.Stack[m.Top] = 0
		} else {
			m.Stack[m.Top] = true64
		}
	}

	dispatch[OpCodeInt64Subtract] = func(n *Interpreter, m *Machine) {
		rhs := int64(m.Stack[m.Top])
		m.Top -= 1
		lhs := int64(m.Stack[m.Top])
		m.Stack[m.Top] = uint64(lhs - rhs)
	}

	dispatch[OpCodeNoOp] = func(n *Interpreter, m *Machine) {
		// do nothing
	}

	dispatch[OpCodeReturn] = func(n *Interpreter, m *Machine) {
		// TO DO
	}

	dispatch[OpCodeStop] = func(n *Interpreter, m *Machine) {
		m.IsRunning = false
	}

	dispatch[OpCodeStringConcatenate] = func(n *Interpreter, m *Machine) {
		rhs := n.stringPool.Get(m.Stack[m.Top])
		m.Top -= 1
		lhs := n.stringPool.Get(m.Stack[m.Top])
		m.Stack[m.Top] = n.stringPool.Put(lhs + rhs)
	}

	dispatch[OpCodeStringEquals] = func(n *Interpreter, m *Machine) {
		rhs := n.stringPool.Get(m.Stack[m.Top])
		m.Top -= 1
		lhs := n.stringPool.Get(m.Stack[m.Top])
		if lhs == rhs {
			m.Stack[m.Top] = true64
		} else {
			m.Stack[m.Top] = 0
		}
	}

	dispatch[OpCodeStringLoad] = func(n *Interpreter, m *Machine) {
		m.Top += 1
		m.Stack[m.Top] = *(*uint64)(unsafe.Pointer(&n.codeBlock.OpCodes[m.IP]))
		m.IP += 4
	}

	dispatch[OpCodeStringNotEquals] = func(n *Interpreter, m *Machine) {
		rhs := n.stringPool.Get(m.Stack[m.Top])
		m.Top -= 1
		lhs := n.stringPool.Get(m.Stack[m.Top])
		if lhs == rhs {
			m.Stack[m.Top] = 0
		} else {
			m.Stack[m.Top] = true64
		}
	}

	dispatch[OpCodeTypeEquals] = func(n *Interpreter, m *Machine) {
		rhs := m.Stack[m.Top]
		m.Top -= 1
		lhs := m.Stack[m.Top]
		if lhs == rhs {
			m.Stack[m.Top] = true64
		} else {
			m.Stack[m.Top] = 0
		}
	}

	dispatch[OpCodeTypeLoad] = func(n *Interpreter, m *Machine) {
		m.Top += 1
		m.Stack[m.Top] = uint64(n.codeBlock.OpCodes[m.IP])
		m.IP += 1
	}

	dispatch[OpCodeTypeNotEquals] = func(n *Interpreter, m *Machine) {
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
