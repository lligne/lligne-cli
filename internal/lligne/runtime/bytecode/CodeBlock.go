//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package bytecode

import (
	"fmt"
	"math"
	"strings"
)

//=====================================================================================================================

type CodeBlock struct {
	OpCodes []uint16
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
	cb.OpCodes = append(cb.OpCodes, uint16(bits))
	cb.OpCodes = append(cb.OpCodes, uint16(bits>>16))
	cb.OpCodes = append(cb.OpCodes, uint16(bits>>32))
	cb.OpCodes = append(cb.OpCodes, uint16(bits>>48))
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

//=====================================================================================================================

func (cb *CodeBlock) Disassemble() string {

	output := &strings.Builder{}
	var line int

	ip := 0

	for {

		opCode := cb.OpCodes[ip]
		ip += 1
		line += 1

		switch opCode {

		case OpCodeBoolAnd:
			write(output, line, "BOOL_AND")
		case OpCodeBoolLoadFalse:
			write(output, line, "BOOL_LOAD_FALSE")
		case OpCodeBoolLoadTrue:
			write(output, line, "BOOL_LOAD_TRUE")
		case OpCodeBoolNot:
			write(output, line, "BOOL_NOT")
		case OpCodeBoolOr:
			write(output, line, "BOOL_OR")

		case OpCodeFloat64Add:
			write(output, line, "FLOAT64_ADD")
		case OpCodeFloat64Divide:
			write(output, line, "FLOAT64_DIVIDE")
		case OpCodeFloat64Equals:
			write(output, line, "FLOAT64_EQUALS")
		case OpCodeFloat64GreaterThan:
			write(output, line, "FLOAT64_GREATER")
		case OpCodeFloat64GreaterThanOrEquals:
			write(output, line, "FLOAT64_NOT_LESS")
		case OpCodeFloat64LessThan:
			write(output, line, "FLOAT64_LESS")
		case OpCodeFloat64LessThanOrEquals:
			write(output, line, "FLOAT64_NOT_GREATER")
		case OpCodeFloat64LoadFloat64:
			bits := uint64(cb.OpCodes[ip])
			ip += 1
			bits |= uint64(cb.OpCodes[ip]) << 16
			ip += 1
			bits |= uint64(cb.OpCodes[ip]) << 32
			ip += 1
			bits |= uint64(cb.OpCodes[ip]) << 48
			ip += 1
			value := math.Float64frombits(bits)
			writeFloat64(output, line, "FLOAT64_LOAD_FLOAT64", value)
		case OpCodeFloat64LoadOne:
			write(output, line, "FLOAT64_LOAD_ONE")
		case OpCodeFloat64LoadZero:
			write(output, line, "FLOAT64_LOAD_ZERO")
		case OpCodeFloat64Multiply:
			write(output, line, "FLOAT64_MULTIPLY")
		case OpCodeFloat64Negate:
			write(output, line, "FLOAT64_NEGATE")
		case OpCodeFloat64Subtract:
			write(output, line, "FLOAT64_SUBTRACT")

		case OpCodeInt64Add:
			write(output, line, "INT64_ADD")
		case OpCodeInt64Divide:
			write(output, line, "INT64_DIVIDE")
		case OpCodeInt64Equals:
			write(output, line, "INT64_EQUALS")
		case OpCodeInt64GreaterThan:
			write(output, line, "INT64_GREATER")
		case OpCodeInt64GreaterThanOrEquals:
			write(output, line, "INT64_NOT_LESS")
		case OpCodeInt64LessThan:
			write(output, line, "INT64_LESS")
		case OpCodeInt64LessThanOrEquals:
			write(output, line, "INT64_NOT_GREATER")
		case OpCodeInt64LoadInt16:
			value := int16(cb.OpCodes[ip])
			ip += 1
			writeInt16(output, line, "INT64_LOAD_INT16", value)
		case OpCodeInt64LoadOne:
			write(output, line, "INT64_LOAD_ONE")
		case OpCodeInt64LoadZero:
			write(output, line, "INT64_LOAD_ZERO")
		case OpCodeInt64Multiply:
			write(output, line, "INT64_MULTIPLY")
		case OpCodeInt64Negate:
			write(output, line, "INT64_NEGATE")
		case OpCodeInt64Subtract:
			write(output, line, "INT64_SUBTRACT")

		case OpCodeNoOp:
			write(output, line, "NO_OP")

		case OpCodeReturn:
			write(output, line, "RETURN")
		case OpCodeStop:
			write(output, line, "STOP")
			return output.String() + "\n"
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

//=====================================================================================================================
