//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package bytecode

//=====================================================================================================================

type ICodeBlockExecutor interface {
	Int64Add()
	Int64LoadInt16(operand int16)
	Int64LoadOne()
	Int64LoadZero()
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

		case OpCodeInt64Add:
			executor.Int64Add()
		case OpCodeInt64LoadInt16:
			value := int16(cb.codes[ip])
			ip += 1
			executor.Int64LoadInt16(value)
		case OpCodeInt64LoadOne:
			executor.Int64LoadOne()
		case OpCodeInt64LoadZero:
			executor.Int64LoadZero()

		case OpCodeNoOp:
			executor.NoOp()

		case OpCodeReturn:
			executor.Return()
			print("Returning")
			return InterpretResultOk
		}

	}

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

func (cb *CodeBlock) Int64Add() {
	cb.codes = append(cb.codes, OpCodeInt64Add)
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
