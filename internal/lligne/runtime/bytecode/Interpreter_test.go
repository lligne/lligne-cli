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

	t.Run("simple exercise", func(t *testing.T) {
		codeBlock := NewCodeBlock()
		machine := NewMachine()
		interpreter := &Interpreter{}

		codeBlock.Int64LoadInt16(2)
		codeBlock.Int64LoadInt16(3)
		codeBlock.Int64LoadZero()
		codeBlock.Int64LoadOne()
		codeBlock.Int64Add()
		codeBlock.Int64Add()
		codeBlock.Int64Add()

		codeBlock.Stop()

		interpreter.Execute(machine, codeBlock)

		actual := interpreter.Int64GetResult(machine)
		expected := int64(6)

		assert.Equal(t, expected, actual)
	})

}

//---------------------------------------------------------------------------------------------------------------------
