//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package bytecode

import (
	"fmt"
	"lligne-cli/internal/lligne/runtime/pools"
	"lligne-cli/internal/lligne/runtime/types"
	"math"
	"strings"
	"unsafe"
)

//=====================================================================================================================

// CodeBlock consists of a sequence of op codes plus a string constant pool.
type CodeBlock struct {
	OpCodes []uint16
}

//---------------------------------------------------------------------------------------------------------------------

// NewCodeBlock constructs a new empty code block.
func NewCodeBlock() *CodeBlock {
	result := &CodeBlock{
		OpCodes: nil,
	}

	return result
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

func (cb *CodeBlock) Float64Load(operand float64) {
	cb.OpCodes = append(cb.OpCodes, OpCodeFloat64Load)
	bits := math.Float64bits(operand)
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

func (cb *CodeBlock) Float64NotEquals() {
	cb.OpCodes = append(cb.OpCodes, OpCodeFloat64NotEquals)
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

func (cb *CodeBlock) Int64Load(operand int64) {
	cb.OpCodes = append(cb.OpCodes, OpCodeInt64Load)
	cb.append64BitOperand(uint64(operand))
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

func (cb *CodeBlock) Int64NotEquals() {
	cb.OpCodes = append(cb.OpCodes, OpCodeInt64NotEquals)
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

func (cb *CodeBlock) RecordEquals() {
	cb.OpCodes = append(cb.OpCodes, OpCodeRecordEquals)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) RecordFieldIndexLoad(fieldIndex uint64) {
	cb.OpCodes = append(cb.OpCodes, OpCodeRecordFieldIndexLoad)
	cb.append64BitOperand(fieldIndex)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) RecordFieldReference() {
	cb.OpCodes = append(cb.OpCodes, OpCodeRecordFieldReference)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) RecordNotEquals() {
	cb.OpCodes = append(cb.OpCodes, OpCodeRecordNotEquals)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) RecordStore(fieldCount int) {
	cb.OpCodes = append(cb.OpCodes, OpCodeRecordStore)
	cb.append64BitOperand(uint64(fieldCount))
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) Return() {
	cb.OpCodes = append(cb.OpCodes, OpCodeReturn)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) StackPop() {
	cb.OpCodes = append(cb.OpCodes, OpCodeStackPop)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) StackPopSecond() {
	cb.OpCodes = append(cb.OpCodes, OpCodeStackPopSecond)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) StackSwapTopTwo() {
	cb.OpCodes = append(cb.OpCodes, OpCodeStackSwapTopTwo)
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

func (cb *CodeBlock) StringEquals() {
	cb.OpCodes = append(cb.OpCodes, OpCodeStringEquals)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) StringLoad(valueIndex pools.StringIndex) {
	cb.OpCodes = append(cb.OpCodes, OpCodeStringLoad)
	cb.append64BitOperand(uint64(valueIndex))
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) StringNotEquals() {
	cb.OpCodes = append(cb.OpCodes, OpCodeStringNotEquals)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) TypeEquals() {
	cb.OpCodes = append(cb.OpCodes, OpCodeTypeEquals)
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) TypeLoad(valueIndex types.TypeIndex) {
	cb.OpCodes = append(cb.OpCodes, OpCodeTypeLoad)
	cb.append64BitOperand(uint64(valueIndex))
}

//---------------------------------------------------------------------------------------------------------------------

func (cb *CodeBlock) TypeNotEquals() {
	cb.OpCodes = append(cb.OpCodes, OpCodeTypeNotEquals)
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
func (cb *CodeBlock) Disassemble(
	stringPool *pools.StringPool,
	typePool *types.TypeConstantPool,
) string {

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
		case OpCodeFloat64Load:
			value := *(*float64)(unsafe.Pointer(&cb.OpCodes[ip]))
			writeFloat64(output, ip, "FLOAT64_LOAD", value)
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
		case OpCodeInt64Load:
			value := *(*int64)(unsafe.Pointer(&cb.OpCodes[ip]))
			writeInt64(output, ip, "INT64_LOAD", value)
			ip += 4
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

		case OpCodeRecordEquals:
			write(output, ip, "RECORD_EQUALS")
		case OpCodeRecordFieldIndexLoad:
			fieldIndex := *(*uint64)(unsafe.Pointer(&cb.OpCodes[ip]))
			writeUInt64(output, ip, "RECORD_FLD_IDX_LOAD", fieldIndex)
			ip += 4
		case OpCodeRecordNotEquals:
			write(output, ip, "RECORD_NOT_EQUALS")
		case OpCodeRecordStore:
			writeUInt64(output, ip, "RECORD_STORE", uint64(cb.OpCodes[ip]))
			ip += 4

		case OpCodeReturn:
			write(output, ip, "RETURN")

		case OpCodeStackPop:
			write(output, ip, "STACK_POP")
		case OpCodeStackPopSecond:
			write(output, ip, "STACK_POP_SECOND")
		case OpCodeStackSwapTopTwo:
			write(output, ip, "STACK_SWAP_TOP_TWO")

		case OpCodeStop:
			write(output, ip, "STOP")
			return output.String() + "\n"

		case OpCodeStringConcatenate:
			write(output, ip, "STRING_CONCATENATE")
		case OpCodeStringEquals:
			write(output, ip, "STRING_EQUALS")
		case OpCodeStringLoad:
			valueIndex := *(*pools.StringIndex)(unsafe.Pointer(&cb.OpCodes[ip]))
			writeString(output, ip, "STRING_LOAD", stringPool.Get(valueIndex))
			ip += 4

		case OpCodeTypeEquals:
			write(output, ip, "TYPE_EQUALS")
		case OpCodeTypeLoad:
			writeType(output, ip, "TYPE_LOAD", typePool.Get(types.TypeIndex(cb.OpCodes[ip])))
			ip += 4
		case OpCodeTypeNotEquals:
			write(output, ip, "TYPE_NOT_EQUALS")

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

func writeInt64(output *strings.Builder, line int, opCode string, operand int64) {
	output.WriteString("\n")
	output.WriteString(fmt.Sprintf("%4d  %-20s %6d", line, opCode, operand))
}

//---------------------------------------------------------------------------------------------------------------------

func writeString(output *strings.Builder, line int, opCode string, operand string) {
	output.WriteString("\n")
	output.WriteString(fmt.Sprintf("%4d  %-20s '%s'", line, opCode, operand))
}

//---------------------------------------------------------------------------------------------------------------------

func writeType(output *strings.Builder, line int, opCode string, operand types.IType) {
	output.WriteString("\n")
	output.WriteString(fmt.Sprintf("%4d  %-20s %s", line, opCode, operand.Name()))
}

//---------------------------------------------------------------------------------------------------------------------

func writeUInt64(output *strings.Builder, line int, opCode string, operand uint64) {
	output.WriteString("\n")
	output.WriteString(fmt.Sprintf("%4d  %-20s %6d", line, opCode, operand))
}

//=====================================================================================================================
