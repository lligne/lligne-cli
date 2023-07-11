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
		scanner := scanning.NewLligneBufferedScanner(
			scanning.NewLligneDocumentationHandlingScanner(
				sourceCode,
				scanning.NewLligneScanner(sourceCode),
			),
		)
		parser := NewLligneParser(scanner)
		model := parser.ParseExpression()

		assert.Equal(t, sExpression, model.SExpression())
	}

	t.Run("identifier literals", func(t *testing.T) {
		check("abc", "(identifier abc)")
		check("\n  d  \n", "(identifier d)")
	})

	t.Run("integer literals", func(t *testing.T) {
		check("123", "(int 123)")
		check("789", "(int 789)")
	})

	t.Run("leading documentation", func(t *testing.T) {
		check("// line one\n // line two\n", "(leadingdoc\n// line one\n// line two\n)")
	})

	t.Run("multiline string literals", func(t *testing.T) {
		check("` line one\n ` line two\n", "(multilinestr\n` line one\n` line two\n)")
	})

	t.Run("string literals", func(t *testing.T) {
		check(`"123"`, `(string "123")`)
		check(`'789'`, `(string '789')`)
	})

}

//---------------------------------------------------------------------------------------------------------------------
