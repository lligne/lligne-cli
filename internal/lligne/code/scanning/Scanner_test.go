//
// # Tests of LligneScanner.
//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package scanning

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

//---------------------------------------------------------------------------------------------------------------------

func TestLligneScanner(t *testing.T) {

	expectToken := func(token Token, expectedTokenType TokenType, expectedSourceOffset int, expectedLength int) {
		assert.Equal(t, expectedTokenType, token.TokenType, "Wrong token type")
		assert.Equal(t, uint32(expectedSourceOffset), token.SourceOffset, "Wrong source offset")
		assert.Equal(t, uint16(expectedLength), token.SourceLength, "Wrong source length")
	}

	t.Run("empty string", func(t *testing.T) {
		tokens, newLineOffsets := Scan("")

		expectToken(tokens[0], TokenTypeEof, 0, 0)
		assert.Equal(t, 0, len(newLineOffsets))
	})

	t.Run("unrecognized character", func(t *testing.T) {
		tokens, newLineOffsets := Scan("‽")

		expectToken(tokens[0], TokenTypeUnrecognizedChar, 0, 3)
		expectToken(tokens[1], TokenTypeEof, 3, 0)
		assert.Equal(t, 0, len(newLineOffsets))
	})

	t.Run("a few punctuation tokens", func(t *testing.T) {
		tokens, newLineOffsets := Scan(
			"& &&\n *: , ",
		)

		expectToken(tokens[0], TokenTypeAmpersand, 0, 1)
		expectToken(tokens[1], TokenTypeAmpersandAmpersand, 2, 2)
		expectToken(tokens[2], TokenTypeAsterisk, 6, 1)
		expectToken(tokens[3], TokenTypeColon, 7, 1)
		expectToken(tokens[4], TokenTypeComma, 9, 1)
		expectToken(tokens[5], TokenTypeEof, 11, 0)
		assert.Equal(t, 1, len(newLineOffsets))
	})

	t.Run("a few identifier tokens", func(t *testing.T) {
		tokens, newLineOffsets := Scan(
			"a bb c23_f _dfg",
		)

		expectToken(tokens[0], TokenTypeIdentifier, 0, 1)
		expectToken(tokens[1], TokenTypeIdentifier, 2, 2)
		expectToken(tokens[2], TokenTypeIdentifier, 5, 5)
		expectToken(tokens[3], TokenTypeIdentifier, 11, 4)
		expectToken(tokens[4], TokenTypeEof, 15, 0)
		assert.Equal(t, 0, len(newLineOffsets))
	})

	t.Run("a few integers", func(t *testing.T) {
		tokens, newLineOffsets := Scan(
			"123 4\n(99000) 5",
		)

		expectToken(tokens[0], TokenTypeIntegerLiteral, 0, 3)
		expectToken(tokens[1], TokenTypeIntegerLiteral, 4, 1)
		expectToken(tokens[2], TokenTypeLeftParenthesis, 6, 1)
		expectToken(tokens[3], TokenTypeIntegerLiteral, 7, 5)
		expectToken(tokens[4], TokenTypeRightParenthesis, 12, 1)
		expectToken(tokens[5], TokenTypeIntegerLiteral, 14, 1)
		expectToken(tokens[6], TokenTypeEof, 15, 0)
		assert.Equal(t, 1, len(newLineOffsets))
	})

	t.Run("a few numbers", func(t *testing.T) {
		tokens, newLineOffsets := Scan(
			"12.3 4\n(990.00) 5.1",
		)

		expectToken(tokens[0], TokenTypeFloatingPointLiteral, 0, 4)
		expectToken(tokens[1], TokenTypeIntegerLiteral, 5, 1)
		expectToken(tokens[2], TokenTypeLeftParenthesis, 7, 1)
		expectToken(tokens[3], TokenTypeFloatingPointLiteral, 8, 6)
		expectToken(tokens[4], TokenTypeRightParenthesis, 14, 1)
		expectToken(tokens[5], TokenTypeFloatingPointLiteral, 16, 3)
		expectToken(tokens[6], TokenTypeEof, 19, 0)
		assert.Equal(t, 1, len(newLineOffsets))
	})

	t.Run("a few double quoted strings", func(t *testing.T) {
		tokens, newLineOffsets := Scan(
			`"abc" "xyz" "bad
 "start over"`,
		)

		expectToken(tokens[0], TokenTypeDoubleQuotedString, 0, 5)
		expectToken(tokens[1], TokenTypeDoubleQuotedString, 6, 5)
		expectToken(tokens[2], TokenTypeUnclosedDoubleQuotedString, 12, 4)
		expectToken(tokens[3], TokenTypeDoubleQuotedString, 18, 12)
		expectToken(tokens[4], TokenTypeEof, 30, 0)
		assert.Equal(t, 1, len(newLineOffsets))
	})

	t.Run("a few single quoted strings", func(t *testing.T) {
		tokens, newLineOffsets := Scan(
			`'abc' 'xyz' 'bad
 'start over'`,
		)

		expectToken(tokens[0], TokenTypeSingleQuotedString, 0, 5)
		expectToken(tokens[1], TokenTypeSingleQuotedString, 6, 5)
		expectToken(tokens[2], TokenTypeUnclosedSingleQuotedString, 12, 4)
		expectToken(tokens[3], TokenTypeSingleQuotedString, 18, 12)
		expectToken(tokens[4], TokenTypeEof, 30, 0)
		assert.Equal(t, 1, len(newLineOffsets))
	})

	t.Run("all fixed text tokens, one at a time", func(t *testing.T) {
		for tokenType := TokenTypeEof; tokenType < TokenType_Count; tokenType += 1 {
			sourceCode := tokenType.String()

			// Skip tokens that can have different text for the same token type.
			if strings.HasPrefix(sourceCode, "[") && strings.HasSuffix(sourceCode, "]") {
				continue
			}

			tokens, newLineOffsets := Scan(
				sourceCode,
			)

			expectToken(tokens[0], tokenType, 0, len(sourceCode))
			expectToken(tokens[1], TokenTypeEof, len(sourceCode), 0)
			assert.Equal(t, 0, len(newLineOffsets))
		}
	})

	t.Run("a few back-ticked string lines", func(t *testing.T) {
		tokens, newLineOffsets := Scan(
			"`abc 123\n`  - one\n  `  - two\n\n  `another\n\n  `one more\n `and the end",
		)

		expectToken(tokens[0], TokenTypeBackTickedString, 0, 29)
		expectToken(tokens[1], TokenTypeBackTickedString, 32, 9)
		expectToken(tokens[2], TokenTypeBackTickedString, 44, 23)
		expectToken(tokens[3], TokenTypeEof, 67, 0)
		assert.Equal(t, 7, len(newLineOffsets))
	})

	t.Run("a few documentation lines", func(t *testing.T) {
		tokens, newLineOffsets := Scan(
			"// abc 123\n//  - one\n//two\n\n//\n//",
		)

		expectToken(tokens[0], TokenTypeDocumentation, 0, 27)
		expectToken(tokens[1], TokenTypeDocumentation, 28, 5)
		expectToken(tokens[2], TokenTypeEof, 33, 0)
		assert.Equal(t, 5, len(newLineOffsets))
	})

	t.Run("boolean literals", func(t *testing.T) {
		tokens, newLineOffsets := Scan(
			"true false",
		)

		expectToken(tokens[0], TokenTypeTrue, 0, 4)
		expectToken(tokens[1], TokenTypeFalse, 5, 5)
		expectToken(tokens[2], TokenTypeEof, 10, 0)
		assert.Equal(t, 0, len(newLineOffsets))
	})

}

//---------------------------------------------------------------------------------------------------------------------
