//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package bytecode

import (
	"fmt"
	"math"
	"strings"
	"unsafe"
)

//=====================================================================================================================

// CodeBlock consists of a sequence of op codes plus a string constant pool.
type CodeBlock struct {
	OpCodes []uint16
	Strings StringConstantPool
}

//---------------------------------------------------------------------------------------------------------------------

// NewCodeBlock constructs a new empty code block.
func NewCodeBlock() *CodeBlock {
	return &CodeBlock{
		OpCodes: nil,
		Strings: NewStringConstantPool(),
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) BoolAnd() {
	cb.OpCodes = append(cb.OpCodes, OpCodeBoolAnd)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) BoolLoadFalse() {
	cb.OpCodes = append(cb.OpCodes, OpCodeBoolLoadFalse)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) BoolLoadTrue() {
	cb.OpCodes = append(cb.OpCodes, OpCodeBoolLoadTrue)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) BoolNot() {
	cb.OpCodes = append(cb.OpCodes, OpCodeBoolNot)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) BoolOr() {
	cb.OpCodes = append(cb.OpCodes, OpCodeBoolOr)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) Float64Add() {
	cb.OpCodes = append(cb.OpCodes, OpCodeFloat64Add)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) Float64Divide() {
	cb.OpCodes = append(cb.OpCodes, OpCodeFloat64Divide)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) Float64Equals() {
	cb.OpCodes = append(cb.OpCodes, OpCodeFloat64Equals)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) Float64GreaterThan() {
	cb.OpCodes = append(cb.OpCodes, OpCodeFloat64GreaterThan)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) Float64GreaterThanOrEquals() {
	cb.OpCodes = append(cb.OpCodes, OpCodeFloat64GreaterThanOrEquals)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) Float64LessThan() {
	cb.OpCodes = append(cb.OpCodes, OpCodeFloat64LessThan)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) Float64LessThanOrEquals() {
	cb.OpCodes = append(cb.OpCodes, OpCodeFloat64LessThanOrEquals)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) Float64LoadFloat64(operand float64) {
	bits := math.Float64bits(operand)
	cb.OpCodes = append(cb.OpCodes, OpCodeFloat64LoadFloat64)
	cb.append64BitOperand(bits)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) Float64LoadOne() {
	cb.OpCodes = append(cb.OpCodes, OpCodeFloat64LoadOne)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) Float64LoadZero() {
	cb.OpCodes = append(cb.OpCodes, OpCodeFloat64LoadZero)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) Float64Multiply() {
	cb.OpCodes = append(cb.OpCodes, OpCodeFloat64Multiply)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) Float64Negate() {
	cb.OpCodes = append(cb.OpCodes, OpCodeFloat64Negate)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) Float64Subtract() {
	cb.OpCodes = append(cb.OpCodes, OpCodeFloat64Subtract)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) Int64Add() {
	cb.OpCodes = append(cb.OpCodes, OpCodeInt64Add)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) Int64Decrement() {
	cb.OpCodes = append(cb.OpCodes, OpCodeInt64Decrement)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) Int64Divide() {
	cb.OpCodes = append(cb.OpCodes, OpCodeInt64Divide)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) Int64Equals() {
	cb.OpCodes = append(cb.OpCodes, OpCodeInt64Equals)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) Int64GreaterThan() {
	cb.OpCodes = append(cb.OpCodes, OpCodeInt64GreaterThan)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) Int64GreaterThanOrEquals() {
	cb.OpCodes = append(cb.OpCodes, OpCodeInt64GreaterThanOrEquals)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) Int64Increment() {
	cb.OpCodes = append(cb.OpCodes, OpCodeInt64Increment)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) Int64LessThan() {
	cb.OpCodes = append(cb.OpCodes, OpCodeInt64LessThan)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) Int64LessThanOrEquals() {
	cb.OpCodes = append(cb.OpCodes, OpCodeInt64LessThanOrEquals)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) Int64LoadInt16(operand int16) {
	cb.OpCodes = append(cb.OpCodes, OpCodeInt64LoadInt16)
	cb.OpCodes = append(cb.OpCodes, uint16(operand))
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) Int64LoadOne() {
	cb.OpCodes = append(cb.OpCodes, OpCodeInt64LoadOne)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) Int64LoadZero() {
	cb.OpCodes = append(cb.OpCodes, OpCodeInt64LoadZero)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) Int64Multiply() {
	cb.OpCodes = append(cb.OpCodes, OpCodeInt64Multiply)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) Int64Negate() {
	cb.OpCodes = append(cb.OpCodes, OpCodeInt64Negate)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) Int64Subtract() {
	cb.OpCodes = append(cb.OpCodes, OpCodeInt64Subtract)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) NoOp() {
	cb.OpCodes = append(cb.OpCodes, OpCodeNoOp)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) Return() {
	cb.OpCodes = append(cb.OpCodes, OpCodeReturn)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) Stop() {
	cb.OpCodes = append(cb.OpCodes, OpCodeStop)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) StringConcatenate() {
	cb.OpCodes = append(cb.OpCodes, OpCodeStringConcatenate)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) StringLoad(value string) {
	cb.OpCodes = append(cb.OpCodes, OpCodeStringLoad)
	cb.OpCodes = append(cb.OpCodes, cb.Strings.Put(value))
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) append64BitOperand(bits uint64) {
	cb.OpCodes = append(cb.OpCodes, uint16(bits))
	cb.OpCodes = append(cb.OpCodes, uint16(bits>>16))
	cb.OpCodes = append(cb.OpCodes, uint16(bits>>32))
	cb.OpCodes = append(cb.OpCodes, uint16(bits>>48))
}

//=====================================================================================================================

// Disassemble dumps out the code block op codes.
func (cb *CodeBlock) Disassemble() string {

	output := &strings.Builder{}

	ip := 0

	for {

		opCode := cb.OpCodes[ip]
		ip += 1

		switch opCode {

		case OpCodeBoolAnd:
			write(output, ip, "BOOL_AND")
		case OpCodeBoolLoadFalse:
			write(output, ip, "BOOL_LOAD_FALSE")
		case OpCodeBoolLoadTrue:
			write(output, ip, "BOOL_LOAD_TRUE")
		case OpCodeBoolNot:
			write(output, ip, "BOOL_NOT")
		case OpCodeBoolOr:
			write(output, ip, "BOOL_OR")

		case OpCodeFloat64Add:
			write(output, ip, "FLOAT64_ADD")
		case OpCodeFloat64Divide:
			write(output, ip, "FLOAT64_DIVIDE")
		case OpCodeFloat64Equals:
			write(output, ip, "FLOAT64_EQUALS")
		case OpCodeFloat64GreaterThan:
			write(output, ip, "FLOAT64_GREATER")
		case OpCodeFloat64GreaterThanOrEquals:
			write(output, ip, "FLOAT64_NOT_LESS")
		case OpCodeFloat64LessThan:
			write(output, ip, "FLOAT64_LESS")
		case OpCodeFloat64LessThanOrEquals:
			write(output, ip, "FLOAT64_NOT_GREATER")
		case OpCodeFloat64LoadFloat64:
			value := *(*float64)(unsafe.Pointer(&cb.OpCodes[ip]))
			writeFloat64(output, ip, "FLOAT64_LOAD_FLOAT64", value)
			ip += 4
		case OpCodeFloat64LoadOne:
			write(output, ip, "FLOAT64_LOAD_ONE")
		case OpCodeFloat64LoadZero:
			write(output, ip, "FLOAT64_LOAD_ZERO")
		case OpCodeFloat64Multiply:
			write(output, ip, "FLOAT64_MULTIPLY")
		case OpCodeFloat64Negate:
			write(output, ip, "FLOAT64_NEGATE")
		case OpCodeFloat64Subtract:
			write(output, ip, "FLOAT64_SUBTRACT")

		case OpCodeInt64Add:
			write(output, ip, "INT64_ADD")
		case OpCodeInt64Decrement:
			write(output, ip, "INT64_DECREMENT")
		case OpCodeInt64Divide:
			write(output, ip, "INT64_DIVIDE")
		case OpCodeInt64Equals:
			write(output, ip, "INT64_EQUALS")
		case OpCodeInt64GreaterThan:
			write(output, ip, "INT64_GREATER")
		case OpCodeInt64GreaterThanOrEquals:
			write(output, ip, "INT64_NOT_LESS")
		case OpCodeInt64Increment:
			write(output, ip, "INT64_INCREMENT")
		case OpCodeInt64LessThan:
			write(output, ip, "INT64_LESS")
		case OpCodeInt64LessThanOrEquals:
			write(output, ip, "INT64_NOT_GREATER")
		case OpCodeInt64LoadInt16:
			value := int16(cb.OpCodes[ip])
			writeInt16(output, ip, "INT64_LOAD_INT16", value)
			ip += 1
		case OpCodeInt64LoadOne:
			write(output, ip, "INT64_LOAD_ONE")
		case OpCodeInt64LoadZero:
			write(output, ip, "INT64_LOAD_ZERO")
		case OpCodeInt64Multiply:
			write(output, ip, "INT64_MULTIPLY")
		case OpCodeInt64Negate:
			write(output, ip, "INT64_NEGATE")
		case OpCodeInt64Subtract:
			write(output, ip, "INT64_SUBTRACT")

		case OpCodeNoOp:
			write(output, ip, "NO_OP")

		case OpCodeReturn:
			write(output, ip, "RETURN")
		case OpCodeStop:
			write(output, ip, "STOP")
			return output.String() + "\n"

		case OpCodeStringConcatenate:
			write(output, ip, "STRING_CONCATENATE")
		case OpCodeStringLoad:
			value := cb.Strings.Get(cb.OpCodes[ip])
			writeString(output, ip, "STRING_LOAD", value)
			ip += 1
		}

	}

}

// ---------------------------------------------------------------------------------------------------------------------

func write(output *strings.Builder, line int, opCode string) {
	output.WriteString("\n")
	output.WriteString(fmt.Sprintf("%4d  %s", line, opCode))
}

//---------------------------------------------------------------------------------------------------------------------

func writeFloat64(output *strings.Builder, line int, opCode string, operand float64) {
	output.WriteString("\n")
	output.WriteString(fmt.Sprintf("%4d  %-20s %10.3f", line, opCode, operand))
}

//---------------------------------------------------------------------------------------------------------------------

func writeInt16(output *strings.Builder, line int, opCode string, operand int16) {
	output.WriteString("\n")
	output.WriteString(fmt.Sprintf("%4d  %-20s %6d", line, opCode, operand))
}

//---------------------------------------------------------------------------------------------------------------------

func writeString(output *strings.Builder, line int, opCode string, operand string) {
	output.WriteString("\n")
	output.WriteString(fmt.Sprintf("%4d  %-20s '%s'", line, opCode, operand))
}

//=====================================================================================================================
