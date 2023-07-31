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
   1  INT64_ADD
   2  INT64_DIVIDE
   3  INT64_LOAD_INT16          3
   4  INT64_LOAD_ONE
   5  INT64_LOAD_ZERO
   6  INT64_MULTIPLY
   7  INT64_NEGATE
   8  INT64_SUBTRACT
   9  RETURN
`

		assert.Equal(t, expected, actual)
	})

}

//---------------------------------------------------------------------------------------------------------------------
