//
// # Tests of CodeBlock.
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

func TestDisassembler(t *testing.T) {

	t.Run("simple output", func(t *testing.T) {
		codeBlock := &CodeBlock{}
		disassembler := &Disassembler{}

		codeBlock.BoolAnd()
		codeBlock.BoolLoadFalse()
		codeBlock.BoolLoadTrue()
		codeBlock.BoolOr()

		codeBlock.Int64Add()
		codeBlock.Int64Divide()
		codeBlock.Int64LoadInt16(3)
		codeBlock.Int64LoadOne()
		codeBlock.Int64LoadZero()
		codeBlock.Int64Multiply()
		codeBlock.Int64Negate()
		codeBlock.Int64Subtract()

		codeBlock.Return()

		codeBlock.Execute(disassembler)

		actual := disassembler.GetOutput()

		expected :=
			`
   1  BOOL_AND
   2  BOOL_LOAD_FALSE
   3  BOOL_LOAD_TRUE
   4  BOOL_OR
   5  INT64_ADD
   6  INT64_DIVIDE
   7  INT64_LOAD_INT16          3
   8  INT64_LOAD_ONE
   9  INT64_LOAD_ZERO
  10  INT64_MULTIPLY
  11  INT64_NEGATE
  12  INT64_SUBTRACT
  13  RETURN
`

		assert.Equal(t, expected, actual)
	})

}

//---------------------------------------------------------------------------------------------------------------------
