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

func TestInterpreter(t *testing.T) {

	t.Run("simple output", func(t *testing.T) {
		codeBlock := &CodeBlock{}
		interpreter := &Interpreter{}

		codeBlock.Int64LoadInt16(2)
		codeBlock.Int64LoadInt16(3)
		codeBlock.Int64LoadZero()
		codeBlock.Int64LoadOne()
		codeBlock.Int64Add()
		codeBlock.Int64Add()
		codeBlock.Int64Add()

		codeBlock.Return()

		codeBlock.Execute(interpreter)

		actual := interpreter.Int64GetResult()
		expected := int64(6)

		assert.Equal(t, expected, actual)
	})

}

//---------------------------------------------------------------------------------------------------------------------
