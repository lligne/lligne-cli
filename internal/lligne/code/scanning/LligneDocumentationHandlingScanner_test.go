//
// # Tests of LligneScanner.
//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package scanning

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLligneDocumentationHandlingScanner(t *testing.T) {

	expectToken := func(scanner ILligneScanner, tokenType LligneTokenType, text string, startPos int) {
		expectedToken := LligneToken{
			TokenType:      tokenType,
			Text:           text,
			SourceStartPos: startPos,
		}

		actualToken := scanner.ReadToken()
		assert.Equal(t, expectedToken, actualToken)
	}

	t.Run("documentation to be translated", func(t *testing.T) {
		sourceCode := `
// Leading documentation
  // with two lines
stuff {
    inner, // Trailing documentation 1
    more;  // Trailing documentation 2
    another  // Trailing 3
         // documentation
}
`
		scanner := NewLligneDocumentationHandlingScanner(sourceCode, NewLligneScanner(sourceCode))

		expectToken(scanner, TokenTypeLeadingDocumentation, "// Leading documentation\n// with two lines\n", 1)
		expectToken(scanner, TokenTypeSynthDocument, " ", 1)
		expectToken(scanner, TokenTypeIdentifier, "stuff", 46)
		expectToken(scanner, TokenTypeLeftBrace, "{", 52)
		expectToken(scanner, TokenTypeIdentifier, "inner", 58)
		expectToken(scanner, TokenTypeSynthDocument, " ", 65)
		expectToken(scanner, TokenTypeTrailingDocumentation, "// Trailing documentation 1\n", 65)
		expectToken(scanner, TokenTypeComma, ",", 63)
		expectToken(scanner, TokenTypeIdentifier, "more", 97)
		expectToken(scanner, TokenTypeSynthDocument, " ", 104)
		expectToken(scanner, TokenTypeTrailingDocumentation, "// Trailing documentation 2\n", 104)
		expectToken(scanner, TokenTypeSemicolon, ";", 101)
		expectToken(scanner, TokenTypeIdentifier, "another", 136)
		expectToken(scanner, TokenTypeSynthDocument, " ", 145)
		expectToken(scanner, TokenTypeTrailingDocumentation, "// Trailing 3\n// documentation\n", 145)
		expectToken(scanner, TokenTypeRightBrace, "}", 185)
		expectToken(scanner, TokenTypeEof, "", len(sourceCode))
	})

}
