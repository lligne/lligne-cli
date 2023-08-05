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
		codeBlock := &CodeBlock{}

		codeBlock.BoolAnd()
		codeBlock.BoolLoadFalse()
		codeBlock.BoolLoadTrue()
		codeBlock.BoolNot()
		codeBlock.BoolOr()

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
   6  INT64_ADD
   7  INT64_DIVIDE
   8  INT64_EQUALS
   9  INT64_GREATER
  10  INT64_NOT_LESS
  11  INT64_LESS
  12  INT64_NOT_GREATER
  13  INT64_LOAD_INT16          3
  14  INT64_LOAD_ONE
  15  INT64_LOAD_ZERO
  16  INT64_MULTIPLY
  17  INT64_NEGATE
  18  INT64_SUBTRACT
  19  RETURN
  20  STOP
`

		assert.Equal(t, expected, actual)
	})

}

//---------------------------------------------------------------------------------------------------------------------
