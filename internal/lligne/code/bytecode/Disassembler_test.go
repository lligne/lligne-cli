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

		codeBlock.Int64LoadInt16(2)
		codeBlock.Int64LoadInt16(3)
		codeBlock.Int64LoadZero()
		codeBlock.Int64LoadOne()
		codeBlock.Int64Add()

		codeBlock.Return()

		codeBlock.Execute(disassembler)

		actual := disassembler.GetOutput()

		expected :=
			`
   1  INT64_LOAD_INT16          2
   2  INT64_LOAD_INT16          3
   3  INT64_LOAD_ZERO
   4  INT64_LOAD_ONE
   5  INT64_ADD
   6  RETURN
`

		assert.Equal(t, expected, actual)
	})

}

//---------------------------------------------------------------------------------------------------------------------
