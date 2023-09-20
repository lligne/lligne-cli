//
// # Tests of Disassembler.
//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package bytecode

import (
	"github.com/stretchr/testify/assert"
	"lligne-cli/internal/lligne/runtime/pools"
	"lligne-cli/internal/lligne/runtime/types"
	"testing"
)

//---------------------------------------------------------------------------------------------------------------------

func TestCodeBlockDisassembly(t *testing.T) {

	t.Run("simple output", func(t *testing.T) {
		typePool := types.NewTypePool().Freeze()

		codeBlock := NewCodeBlock()

		codeBlock.BoolAnd()
		codeBlock.BoolLoadFalse()
		codeBlock.BoolLoadTrue()
		codeBlock.BoolNot()
		codeBlock.BoolOr()

		codeBlock.Float64Add()
		codeBlock.Float64Divide()
		codeBlock.Float64Equals()
		codeBlock.Float64GreaterThan()
		codeBlock.Float64GreaterThanOrEquals()
		codeBlock.Float64LessThan()
		codeBlock.Float64LessThanOrEquals()
		codeBlock.Float64Load(3)
		codeBlock.Float64LoadOne()
		codeBlock.Float64LoadZero()
		codeBlock.Float64Multiply()
		codeBlock.Float64Negate()
		codeBlock.Float64Subtract()

		codeBlock.Int64Add()
		codeBlock.Int64Divide()
		codeBlock.Int64Equals()
		codeBlock.Int64GreaterThan()
		codeBlock.Int64GreaterThanOrEquals()
		codeBlock.Int64LessThan()
		codeBlock.Int64LessThanOrEquals()
		codeBlock.Int64Load(3)
		codeBlock.Int64LoadOne()
		codeBlock.Int64LoadZero()
		codeBlock.Int64Multiply()
		codeBlock.Int64Negate()
		codeBlock.Int64Subtract()

		codeBlock.StringConcatenate()
		codeBlock.StringEquals()
		codeBlock.StringLoad(0)
		codeBlock.StringLoad(1)

		codeBlock.TypeLoad(1)
		codeBlock.TypeLoad(2)
		codeBlock.TypeLoad(3)
		codeBlock.TypeLoad(4)
		codeBlock.TypeEquals()
		codeBlock.TypeNotEquals()

		codeBlock.RecordStore(5)
		codeBlock.RecordEquals()
		codeBlock.RecordFieldIndexLoad(17)
		codeBlock.RecordNotEquals()

		codeBlock.StackPop()
		codeBlock.StackPopSecond()
		codeBlock.StackSwapTopTwo()

		codeBlock.Return()
		codeBlock.Stop()

		stringPool := pools.NewStringPool()
		stringPool.Put("String0")
		stringPool.Put("String1")

		actual := codeBlock.Disassemble(stringPool, typePool)

		expected :=
			`
   1  BOOL_AND
   2  BOOL_LOAD_FALSE
   3  BOOL_LOAD_TRUE
   4  BOOL_NOT
   5  BOOL_OR
   6  FLOAT64_ADD
   7  FLOAT64_DIVIDE
   8  FLOAT64_EQUALS
   9  FLOAT64_GREATER
  10  FLOAT64_NOT_LESS
  11  FLOAT64_LESS
  12  FLOAT64_NOT_GREATER
  13  FLOAT64_LOAD              3.000
  18  FLOAT64_LOAD_ONE
  19  FLOAT64_LOAD_ZERO
  20  FLOAT64_MULTIPLY
  21  FLOAT64_NEGATE
  22  FLOAT64_SUBTRACT
  23  INT64_ADD
  24  INT64_DIVIDE
  25  INT64_EQUALS
  26  INT64_GREATER
  27  INT64_NOT_LESS
  28  INT64_LESS
  29  INT64_NOT_GREATER
  30  INT64_LOAD                3
  35  INT64_LOAD_ONE
  36  INT64_LOAD_ZERO
  37  INT64_MULTIPLY
  38  INT64_NEGATE
  39  INT64_SUBTRACT
  40  STRING_CONCATENATE
  41  STRING_EQUALS
  42  STRING_LOAD          'String0'
  47  STRING_LOAD          'String1'
  52  TYPE_LOAD            Bool
  57  TYPE_LOAD            Float64
  62  TYPE_LOAD            Int64
  67  TYPE_LOAD            String
  72  TYPE_EQUALS
  73  TYPE_NOT_EQUALS
  74  RECORD_STORE              5
  79  RECORD_EQUALS
  80  RECORD_FLD_IDX_LOAD      17
  85  RECORD_NOT_EQUALS
  86  STACK_POP
  87  STACK_POP_SECOND
  88  STACK_SWAP_TOP_TWO
  89  RETURN
  90  STOP
`

		assert.Equal(t, expected, actual)
	})

}

//---------------------------------------------------------------------------------------------------------------------
