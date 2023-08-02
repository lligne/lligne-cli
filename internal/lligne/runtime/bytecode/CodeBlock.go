//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package bytecode

//=====================================================================================================================

type ICodeBlockExecutor interface {
	BoolAnd()
	BoolLoadFalse()
	BoolLoadTrue()
	BoolNot()
	BoolOr()
	Int64Add()
	Int64Divide()
	Int64Equals()
	Int64GreaterThan()
	Int64GreaterThanOrEquals()
	Int64LessThan()
	Int64LessThanOrEquals()
	Int64LoadInt16(operand int16)
	Int64LoadOne()
	Int64LoadZero()
	Int64Multiply()
	Int64Negate()
	Int64Subtract()
	NoOp()
	Return()
}

//=====================================================================================================================

// InterpretResult is the status after executing a bytecode operation.
type InterpretResult uint8

const (
	InterpretResultOk InterpretResult = 1 + iota

	InterpretResultError
)

//=====================================================================================================================

type CodeBlock struct {
	codes []LligneOpCode
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) Execute(executor ICodeBlockExecutor) InterpretResult {

	ip := 0

	for {

		opCode := cb.codes[ip]
		ip += 1

		switch opCode {

		case OpCodeBoolAnd:
			executor.BoolAnd()
		case OpCodeBoolLoadFalse:
			executor.BoolLoadFalse()
		case OpCodeBoolLoadTrue:
			executor.BoolLoadTrue()
		case OpCodeBoolNot:
			executor.BoolNot()
		case OpCodeBoolOr:
			executor.BoolOr()

		case OpCodeInt64Add:
			executor.Int64Add()
		case OpCodeInt64Divide:
			executor.Int64Divide()
		case OpCodeInt64Equals:
			executor.Int64Equals()
		case OpCodeInt64GreaterThan:
			executor.Int64GreaterThan()
		case OpCodeInt64GreaterThanOrEquals:
			executor.Int64GreaterThanOrEquals()
		case OpCodeInt64LessThan:
			executor.Int64LessThan()
		case OpCodeInt64LessThanOrEquals:
			executor.Int64LessThanOrEquals()
		case OpCodeInt64LoadInt16:
			value := int16(cb.codes[ip])
			ip += 1
			executor.Int64LoadInt16(value)
		case OpCodeInt64LoadOne:
			executor.Int64LoadOne()
		case OpCodeInt64LoadZero:
			executor.Int64LoadZero()
		case OpCodeInt64Multiply:
			executor.Int64Multiply()
		case OpCodeInt64Negate:
			executor.Int64Negate()
		case OpCodeInt64Subtract:
			executor.Int64Subtract()

		case OpCodeNoOp:
			executor.NoOp()

		case OpCodeReturn:
			executor.Return()
			return InterpretResultOk
		}

	}

}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) BoolAnd() {
	cb.codes = append(cb.codes, OpCodeBoolAnd)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) BoolLoadFalse() {
	cb.codes = append(cb.codes, OpCodeBoolLoadFalse)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) BoolLoadTrue() {
	cb.codes = append(cb.codes, OpCodeBoolLoadTrue)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) BoolNot() {
	cb.codes = append(cb.codes, OpCodeBoolNot)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) BoolOr() {
	cb.codes = append(cb.codes, OpCodeBoolOr)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) Int64Add() {
	cb.codes = append(cb.codes, OpCodeInt64Add)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) Int64Divide() {
	cb.codes = append(cb.codes, OpCodeInt64Divide)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) Int64Equals() {
	cb.codes = append(cb.codes, OpCodeInt64Equals)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) Int64GreaterThan() {
	cb.codes = append(cb.codes, OpCodeInt64GreaterThan)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) Int64GreaterThanOrEquals() {
	cb.codes = append(cb.codes, OpCodeInt64GreaterThanOrEquals)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) Int64LessThan() {
	cb.codes = append(cb.codes, OpCodeInt64LessThan)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) Int64LessThanOrEquals() {
	cb.codes = append(cb.codes, OpCodeInt64LessThanOrEquals)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) Int64LoadInt16(operand int16) {
	cb.codes = append(cb.codes, OpCodeInt64LoadInt16)
	cb.codes = append(cb.codes, LligneOpCode(operand))
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) Int64LoadOne() {
	cb.codes = append(cb.codes, OpCodeInt64LoadOne)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) Int64LoadZero() {
	cb.codes = append(cb.codes, OpCodeInt64LoadZero)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) Int64Multiply() {
	cb.codes = append(cb.codes, OpCodeInt64Multiply)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) Int64Negate() {
	cb.codes = append(cb.codes, OpCodeInt64Negate)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) Int64Subtract() {
	cb.codes = append(cb.codes, OpCodeInt64Subtract)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) NoOp() {
	cb.codes = append(cb.codes, OpCodeNoOp)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) Return() {
	cb.codes = append(cb.codes, OpCodeReturn)
}

//=====================================================================================================================
