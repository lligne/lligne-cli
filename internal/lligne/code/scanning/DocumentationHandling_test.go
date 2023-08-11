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

//---------------------------------------------------------------------------------------------------------------------

func TestLligneDocumentationHandlingScanner(t *testing.T) {

	expectToken := func(token Token, expectedTokenType TokenType, expectedSourceOffset int, expectedLength int) {
		assert.Equal(t, expectedTokenType, token.TokenType, "Wrong token type")
		assert.Equal(t, uint32(expectedSourceOffset), token.SourceOffset, "Wrong source offset")
		assert.Equal(t, uint16(expectedLength), token.SourceLength, "Wrong source length")
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
		tokens, _ := Scan(sourceCode)

		tokens = ProcessLeadingTrailingDocumentation(sourceCode, tokens)

		expectToken(tokens[0], TokenTypeLeadingDocumentation, 1, 45)
		expectToken(tokens[1], TokenTypeSynthDocument, 1, 0)
		expectToken(tokens[2], TokenTypeIdentifier, 46, 5)
		expectToken(tokens[3], TokenTypeLeftBrace, 52, 1)
		expectToken(tokens[4], TokenTypeIdentifier, 58, 5)
		expectToken(tokens[5], TokenTypeSynthDocument, 65, 0)
		expectToken(tokens[6], TokenTypeTrailingDocumentation, 65, 32)
		expectToken(tokens[7], TokenTypeComma, 63, 1)
		expectToken(tokens[8], TokenTypeIdentifier, 97, 4)
		expectToken(tokens[9], TokenTypeSynthDocument, 104, 0)
		expectToken(tokens[10], TokenTypeTrailingDocumentation, 104, 32)
		expectToken(tokens[11], TokenTypeSemicolon, 101, 1)
		expectToken(tokens[12], TokenTypeIdentifier, 136, 7)
		expectToken(tokens[13], TokenTypeSynthDocument, 145, 0)
		expectToken(tokens[14], TokenTypeTrailingDocumentation, 145, 40)
		expectToken(tokens[15], TokenTypeRightBrace, 185, 1)
		expectToken(tokens[16], TokenTypeEof, len(sourceCode), 0)
	})

}

//---------------------------------------------------------------------------------------------------------------------
