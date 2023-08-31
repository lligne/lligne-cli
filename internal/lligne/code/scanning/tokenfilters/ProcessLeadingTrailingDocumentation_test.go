//
// # Tests of LligneScanner.
//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package tokenfilters

import (
	"github.com/stretchr/testify/assert"
	"lligne-cli/internal/lligne/code/scanning"
	"testing"
)

//---------------------------------------------------------------------------------------------------------------------

func TestLligneDocumentationHandlingScanner(t *testing.T) {

	expectToken := func(token scanning.Token, expectedTokenType scanning.TokenType, expectedSourceOffset int, expectedLength int) {
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

    // Leading documentation after trailing
    onemorevar

	gadget

	// Leading after non-doc
	junk
}
`
		scanOutcome := scanning.Scan(sourceCode)
		scanOutcome = ProcessLeadingTrailingDocumentation(scanOutcome)
		tokens := scanOutcome.Tokens

		expectToken(tokens[0], scanning.TokenTypeLeadingDocumentation, 1, 45)
		expectToken(tokens[1], scanning.TokenTypeSynthDocument, 1, 0)
		expectToken(tokens[2], scanning.TokenTypeIdentifier, 46, 5)
		expectToken(tokens[3], scanning.TokenTypeLeftBrace, 52, 1)
		expectToken(tokens[4], scanning.TokenTypeIdentifier, 58, 5)
		expectToken(tokens[5], scanning.TokenTypeSynthDocument, 65, 0)
		expectToken(tokens[6], scanning.TokenTypeTrailingDocumentation, 65, 32)
		expectToken(tokens[7], scanning.TokenTypeComma, 63, 1)
		expectToken(tokens[8], scanning.TokenTypeIdentifier, 97, 4)
		expectToken(tokens[9], scanning.TokenTypeSynthDocument, 104, 0)
		expectToken(tokens[10], scanning.TokenTypeTrailingDocumentation, 104, 32)
		expectToken(tokens[11], scanning.TokenTypeSemicolon, 101, 1)
		expectToken(tokens[12], scanning.TokenTypeIdentifier, 136, 7)
		expectToken(tokens[13], scanning.TokenTypeSynthDocument, 145, 0)
		expectToken(tokens[14], scanning.TokenTypeTrailingDocumentation, 145, 40)
		expectToken(tokens[15], scanning.TokenTypeLeadingDocumentation, 190, 44)
		expectToken(tokens[16], scanning.TokenTypeSynthDocument, 190, 0)
		expectToken(tokens[17], scanning.TokenTypeIdentifier, 234, 10)
		expectToken(tokens[18], scanning.TokenTypeIdentifier, 247, 6)
		expectToken(tokens[19], scanning.TokenTypeLeadingDocumentation, 256, 26)
		expectToken(tokens[20], scanning.TokenTypeSynthDocument, 256, 0)
		expectToken(tokens[21], scanning.TokenTypeIdentifier, 282, 4)
		expectToken(tokens[22], scanning.TokenTypeRightBrace, 287, 1)
		expectToken(tokens[23], scanning.TokenTypeEof, len(sourceCode), 0)

	})

}

//---------------------------------------------------------------------------------------------------------------------
