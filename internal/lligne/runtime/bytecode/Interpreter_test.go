//
// # Tests of CodeBlock.
//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package bytecode

import (
	"github.com/stretchr/testify/assert"
	"lligne-cli/internal/lligne/runtime/pools"
	"testing"
)

//---------------------------------------------------------------------------------------------------------------------

func TestInterpreter(t *testing.T) {

	t.Run("simple exercise", func(t *testing.T) {
		codeBlock := NewCodeBlock()
		machine := NewMachine()
		interpreter := NewInterpreter(codeBlock, pools.NewStringPool())

		codeBlock.Int64Load(2)
		codeBlock.Int64Load(3)
		codeBlock.Int64LoadZero()
		codeBlock.Int64LoadOne()
		codeBlock.Int64Add()
		codeBlock.Int64Add()
		codeBlock.Int64Add()

		codeBlock.Stop()

		interpreter.Execute(machine)

		actual := machine.Int64GetResult()
		expected := int64(6)

		assert.Equal(t, expected, actual)
	})

}

//---------------------------------------------------------------------------------------------------------------------
