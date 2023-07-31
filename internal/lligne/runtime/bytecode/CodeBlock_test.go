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

func TestLligneCodeBlock(t *testing.T) {

	t.Run("just return", func(t *testing.T) {
		codeBlock := &CodeBlock{}
		codeBlock.Return()

		assert.Equal(t, "TBD", "TBD")
	})

}

//---------------------------------------------------------------------------------------------------------------------
