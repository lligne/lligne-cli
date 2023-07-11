//
// # Tests of the parser for Lligne code
//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package parsing

import (
	"github.com/stretchr/testify/assert"
	"lligne-cli/internal/lligne/code/scanning"
	"testing"
)

//---------------------------------------------------------------------------------------------------------------------

func TestLligneBufferedScanner(t *testing.T) {

	check := func(sourceCode string, sExpression string) {
		scanner := scanning.NewLligneBufferedScanner(scanning.NewLligneScanner(sourceCode))
		parser := NewLligneParser(scanner)
		model := parser.ParseExpression()

		assert.Equal(t, sExpression, model.SExpression())
	}

	t.Run("single identifiers", func(t *testing.T) {
		check("abc", "(identifier abc)")
		check("\n  d  \n", "(identifier d)")
	})

}

//---------------------------------------------------------------------------------------------------------------------
