//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package bytecode

//=====================================================================================================================

type Interpreter struct {
	valueStack     [1000]uint64
	valueStackSize int
	valueStackLast int
}

//---------------------------------------------------------------------------------------------------------------------

const true64 uint64 = 0xFFFFFFFFFFFFFFFF

//---------------------------------------------------------------------------------------------------------------------

func (n *Interpreter) BoolAnd() {
	n.valueStackSize = n.valueStackLast
	n.valueStackLast -= 1
	rhs := n.valueStack[n.valueStackSize] != 0
	lhs := n.valueStack[n.valueStackLast] != 0
	if lhs && rhs {
		n.valueStack[n.valueStackLast] = true64
	} else {
		n.valueStack[n.valueStackLast] = 0
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (n *Interpreter) BoolGetResult() bool {
	return n.valueStack[n.valueStackLast] != 0
}

//---------------------------------------------------------------------------------------------------------------------

func (n *Interpreter) BoolLoadFalse() {
	n.valueStackLast = n.valueStackSize
	n.valueStackSize += 1
	n.valueStack[n.valueStackLast] = 0
}

//---------------------------------------------------------------------------------------------------------------------

func (n *Interpreter) BoolLoadTrue() {
	n.valueStackLast = n.valueStackSize
	n.valueStackSize += 1
	n.valueStack[n.valueStackLast] = true64
}

//---------------------------------------------------------------------------------------------------------------------

func (n *Interpreter) BoolNot() {
	if n.valueStack[n.valueStackLast] == 0 {
		n.valueStack[n.valueStackLast] = true64
	} else {
		n.valueStack[n.valueStackLast] = 0
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (n *Interpreter) BoolOr() {
	n.valueStackSize = n.valueStackLast
	n.valueStackLast -= 1
	rhs := n.valueStack[n.valueStackSize] != 0
	lhs := n.valueStack[n.valueStackLast] != 0
	if lhs || rhs {
		n.valueStack[n.valueStackLast] = true64
	} else {
		n.valueStack[n.valueStackLast] = 0
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (n *Interpreter) Int64Add() {
	n.valueStackSize = n.valueStackLast
	n.valueStackLast -= 1
	rhs := int64(n.valueStack[n.valueStackSize])
	lhs := int64(n.valueStack[n.valueStackLast])
	n.valueStack[n.valueStackLast] = uint64(lhs + rhs)
}

//---------------------------------------------------------------------------------------------------------------------

func (n *Interpreter) Int64Divide() {
	n.valueStackSize = n.valueStackLast
	n.valueStackLast -= 1
	rhs := int64(n.valueStack[n.valueStackSize])
	lhs := int64(n.valueStack[n.valueStackLast])
	n.valueStack[n.valueStackLast] = uint64(lhs / rhs)
}

//---------------------------------------------------------------------------------------------------------------------

func (n *Interpreter) Int64Equals() {
	n.valueStackSize = n.valueStackLast
	n.valueStackLast -= 1
	rhs := int64(n.valueStack[n.valueStackSize])
	lhs := int64(n.valueStack[n.valueStackLast])
	if lhs == rhs {
		n.valueStack[n.valueStackLast] = true64
	} else {
		n.valueStack[n.valueStackLast] = 0
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (n *Interpreter) Int64GreaterThan() {
	n.valueStackSize = n.valueStackLast
	n.valueStackLast -= 1
	rhs := int64(n.valueStack[n.valueStackSize])
	lhs := int64(n.valueStack[n.valueStackLast])
	if lhs > rhs {
		n.valueStack[n.valueStackLast] = true64
	} else {
		n.valueStack[n.valueStackLast] = 0
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (n *Interpreter) Int64GreaterThanOrEquals() {
	n.valueStackSize = n.valueStackLast
	n.valueStackLast -= 1
	rhs := int64(n.valueStack[n.valueStackSize])
	lhs := int64(n.valueStack[n.valueStackLast])
	if lhs >= rhs {
		n.valueStack[n.valueStackLast] = true64
	} else {
		n.valueStack[n.valueStackLast] = 0
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (n *Interpreter) Int64LessThan() {
	n.valueStackSize = n.valueStackLast
	n.valueStackLast -= 1
	rhs := int64(n.valueStack[n.valueStackSize])
	lhs := int64(n.valueStack[n.valueStackLast])
	if lhs < rhs {
		n.valueStack[n.valueStackLast] = true64
	} else {
		n.valueStack[n.valueStackLast] = 0
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (n *Interpreter) Int64LessThanOrEquals() {
	n.valueStackSize = n.valueStackLast
	n.valueStackLast -= 1
	rhs := int64(n.valueStack[n.valueStackSize])
	lhs := int64(n.valueStack[n.valueStackLast])
	if lhs <= rhs {
		n.valueStack[n.valueStackLast] = true64
	} else {
		n.valueStack[n.valueStackLast] = 0
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (n *Interpreter) Int64GetResult() int64 {
	return int64(n.valueStack[n.valueStackLast])
}

//---------------------------------------------------------------------------------------------------------------------

func (n *Interpreter) Int64LoadInt16(operand int16) {
	n.valueStackLast = n.valueStackSize
	n.valueStackSize += 1
	n.valueStack[n.valueStackLast] = uint64(int64(operand))
}

//---------------------------------------------------------------------------------------------------------------------

func (n *Interpreter) Int64LoadOne() {
	n.valueStackLast = n.valueStackSize
	n.valueStackSize += 1
	n.valueStack[n.valueStackLast] = 1
}

//---------------------------------------------------------------------------------------------------------------------

func (n *Interpreter) Int64LoadZero() {
	n.valueStackLast = n.valueStackSize
	n.valueStackSize += 1
	n.valueStack[n.valueStackLast] = 0
}

//---------------------------------------------------------------------------------------------------------------------

func (n *Interpreter) Int64Multiply() {
	n.valueStackSize = n.valueStackLast
	n.valueStackLast -= 1
	rhs := int64(n.valueStack[n.valueStackSize])
	lhs := int64(n.valueStack[n.valueStackLast])
	n.valueStack[n.valueStackLast] = uint64(lhs * rhs)
}

//---------------------------------------------------------------------------------------------------------------------

func (n *Interpreter) Int64Negate() {
	n.valueStack[n.valueStackLast] = uint64(-int64(n.valueStack[n.valueStackLast]))
}

//---------------------------------------------------------------------------------------------------------------------

func (n *Interpreter) Int64Subtract() {
	n.valueStackSize = n.valueStackLast
	n.valueStackLast -= 1
	rhs := int64(n.valueStack[n.valueStackSize])
	lhs := int64(n.valueStack[n.valueStackLast])
	n.valueStack[n.valueStackLast] = uint64(lhs - rhs)
}

//---------------------------------------------------------------------------------------------------------------------

func (n *Interpreter) NoOp() {
	// no operation
}

//---------------------------------------------------------------------------------------------------------------------

func (n *Interpreter) Return() {
}

//=====================================================================================================================
