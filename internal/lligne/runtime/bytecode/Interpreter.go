//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package bytecode

//=====================================================================================================================

type Interpreter struct {
	// TODO: Nothing needed?
}

//---------------------------------------------------------------------------------------------------------------------

func (n *Interpreter) BoolGetResult(machine *Machine) bool {
	return machine.Stack[machine.Top] != 0
}

//---------------------------------------------------------------------------------------------------------------------

func (n *Interpreter) Execute(machine *Machine, code *CodeBlock) {

	machine.IP = 0

	for machine.IsRunning {

		opCode := code.OpCodes[machine.IP]
		machine.IP += 1

		dispatch[opCode](machine, code)

	}

}

//---------------------------------------------------------------------------------------------------------------------

func (n *Interpreter) Int64GetResult(machine *Machine) int64 {
	return int64(machine.Stack[machine.Top])
}

//=====================================================================================================================

const true64 uint64 = 0xFFFFFFFFFFFFFFFF

//---------------------------------------------------------------------------------------------------------------------

var dispatch [20]func(*Machine, *CodeBlock)

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

	dispatch[OpCodeInt64Add] = func(m *Machine, c *CodeBlock) {
		rhs := int64(m.Stack[m.Top])
		m.Top -= 1
		lhs := int64(m.Stack[m.Top])
		m.Stack[m.Top] = uint64(lhs + rhs)
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

	dispatch[OpCodeInt64LoadInt16] = func(m *Machine, c *CodeBlock) {
		m.Top += 1
		m.Stack[m.Top] = uint64(c.OpCodes[m.IP])
		m.IP += 1
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
		m.IsRunning = false
	}

}

//=====================================================================================================================
