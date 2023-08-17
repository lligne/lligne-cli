//
// # Tests of Disassembler.
//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package bytecode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//---------------------------------------------------------------------------------------------------------------------

func TestCodeBlockDisassembly(t *testing.T) {

	t.Run("simple output", func(t *testing.T) {
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
		codeBlock.Float64LoadFloat64(3)
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
		codeBlock.Int64LoadInt16(3)
		codeBlock.Int64LoadOne()
		codeBlock.Int64LoadZero()
		codeBlock.Int64Multiply()
		codeBlock.Int64Negate()
		codeBlock.Int64Subtract()

		codeBlock.StringLoad("Example")

		codeBlock.Return()
		codeBlock.Stop()

		actual := codeBlock.Disassemble()

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
  13  FLOAT64_LOAD_FLOAT64      3.000
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
  30  INT64_LOAD_INT16          3
  32  INT64_LOAD_ONE
  33  INT64_LOAD_ZERO
  34  INT64_MULTIPLY
  35  INT64_NEGATE
  36  INT64_SUBTRACT
  37  STRING_LOAD          'Example'
  39  RETURN
  40  STOP
`

		assert.Equal(t, expected, actual)
	})

}

//---------------------------------------------------------------------------------------------------------------------
