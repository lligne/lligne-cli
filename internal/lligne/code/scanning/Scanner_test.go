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
		result := Scan("")

		expectToken(result.Tokens[0], TokenTypeEof, 0, 0)
		assert.Equal(t, 0, len(result.NewLineOffsets))
	})

	t.Run("unrecognized character", func(t *testing.T) {
		result := Scan("‽")

		expectToken(result.Tokens[0], TokenTypeUnrecognizedChar, 0, 3)
		expectToken(result.Tokens[1], TokenTypeEof, 3, 0)
		assert.Equal(t, 0, len(result.NewLineOffsets))
	})

	t.Run("a few punctuation tokens", func(t *testing.T) {
		result := Scan(
			"& &&\n *: , ",
		)

		expectToken(result.Tokens[0], TokenTypeAmpersand, 0, 1)
		expectToken(result.Tokens[1], TokenTypeAmpersandAmpersand, 2, 2)
		expectToken(result.Tokens[2], TokenTypeAsterisk, 6, 1)
		expectToken(result.Tokens[3], TokenTypeColon, 7, 1)
		expectToken(result.Tokens[4], TokenTypeComma, 9, 1)
		expectToken(result.Tokens[5], TokenTypeEof, 11, 0)
		assert.Equal(t, 1, len(result.NewLineOffsets))
	})

	t.Run("a few identifier tokens", func(t *testing.T) {
		result := Scan(
			"a bb c23_f q-code _dfg",
		)

		expectToken(result.Tokens[0], TokenTypeIdentifier, 0, 1)
		expectToken(result.Tokens[1], TokenTypeIdentifier, 2, 2)
		expectToken(result.Tokens[2], TokenTypeIdentifier, 5, 5)
		expectToken(result.Tokens[3], TokenTypeIdentifier, 11, 6)
		expectToken(result.Tokens[4], TokenTypeIdentifier, 18, 4)
		expectToken(result.Tokens[5], TokenTypeEof, 22, 0)
		assert.Equal(t, 0, len(result.NewLineOffsets))
	})

	t.Run("a few integers", func(t *testing.T) {
		result := Scan(
			"123 4\n(99000) 5",
		)

		expectToken(result.Tokens[0], TokenTypeIntegerLiteral, 0, 3)
		expectToken(result.Tokens[1], TokenTypeIntegerLiteral, 4, 1)
		expectToken(result.Tokens[2], TokenTypeLeftParenthesis, 6, 1)
		expectToken(result.Tokens[3], TokenTypeIntegerLiteral, 7, 5)
		expectToken(result.Tokens[4], TokenTypeRightParenthesis, 12, 1)
		expectToken(result.Tokens[5], TokenTypeIntegerLiteral, 14, 1)
		expectToken(result.Tokens[6], TokenTypeEof, 15, 0)
		assert.Equal(t, 1, len(result.NewLineOffsets))
	})

	t.Run("a few numbers", func(t *testing.T) {
		result := Scan(
			"12.3 4\n(990.00) 5.1",
		)

		expectToken(result.Tokens[0], TokenTypeFloatingPointLiteral, 0, 4)
		expectToken(result.Tokens[1], TokenTypeIntegerLiteral, 5, 1)
		expectToken(result.Tokens[2], TokenTypeLeftParenthesis, 7, 1)
		expectToken(result.Tokens[3], TokenTypeFloatingPointLiteral, 8, 6)
		expectToken(result.Tokens[4], TokenTypeRightParenthesis, 14, 1)
		expectToken(result.Tokens[5], TokenTypeFloatingPointLiteral, 16, 3)
		expectToken(result.Tokens[6], TokenTypeEof, 19, 0)
		assert.Equal(t, 1, len(result.NewLineOffsets))
	})

	t.Run("a few double quoted strings", func(t *testing.T) {
		result := Scan(
			`"abc" "xyz" "bad
 "start over"`,
		)

		expectToken(result.Tokens[0], TokenTypeDoubleQuotedString, 0, 5)
		expectToken(result.Tokens[1], TokenTypeDoubleQuotedString, 6, 5)
		expectToken(result.Tokens[2], TokenTypeUnclosedDoubleQuotedString, 12, 4)
		expectToken(result.Tokens[3], TokenTypeDoubleQuotedString, 18, 12)
		expectToken(result.Tokens[4], TokenTypeEof, 30, 0)
		assert.Equal(t, 1, len(result.NewLineOffsets))
	})

	t.Run("a few single quoted strings", func(t *testing.T) {
		result := Scan(
			`'abc' 'xyz' 'bad
 'start over'`,
		)

		expectToken(result.Tokens[0], TokenTypeSingleQuotedString, 0, 5)
		expectToken(result.Tokens[1], TokenTypeSingleQuotedString, 6, 5)
		expectToken(result.Tokens[2], TokenTypeUnclosedSingleQuotedString, 12, 4)
		expectToken(result.Tokens[3], TokenTypeSingleQuotedString, 18, 12)
		expectToken(result.Tokens[4], TokenTypeEof, 30, 0)
		assert.Equal(t, 1, len(result.NewLineOffsets))
	})

	t.Run("all fixed text tokens, one at a time", func(t *testing.T) {
		for tokenType := TokenTypeEof; tokenType < TokenType_Count; tokenType += 1 {
			sourceCode := tokenType.String()

			// Skip tokens that can have different text for the same token type.
			if strings.HasPrefix(sourceCode, "[") && strings.HasSuffix(sourceCode, "]") {
				continue
			}

			result := Scan(
				sourceCode,
			)

			expectToken(result.Tokens[0], tokenType, 0, len(sourceCode))
			expectToken(result.Tokens[1], TokenTypeEof, len(sourceCode), 0)
			assert.Equal(t, 0, len(result.NewLineOffsets))
		}
	})

	t.Run("a few back-ticked string lines", func(t *testing.T) {
		result := Scan(
			"`abc 123\n`  - one\n  `  - two\n\n  `another\n\n  `one more\n `and the end",
		)

		expectToken(result.Tokens[0], TokenTypeBackTickedString, 0, 29)
		expectToken(result.Tokens[1], TokenTypeBackTickedString, 32, 9)
		expectToken(result.Tokens[2], TokenTypeBackTickedString, 44, 23)
		expectToken(result.Tokens[3], TokenTypeEof, 67, 0)
		assert.Equal(t, 7, len(result.NewLineOffsets))
	})

	t.Run("a few documentation lines", func(t *testing.T) {
		result := Scan(
			"// abc 123\n//  - one\n//two\n\n//\n//",
		)

		expectToken(result.Tokens[0], TokenTypeDocumentation, 0, 27)
		expectToken(result.Tokens[1], TokenTypeDocumentation, 28, 5)
		expectToken(result.Tokens[2], TokenTypeEof, 33, 0)
		assert.Equal(t, 5, len(result.NewLineOffsets))
	})

	t.Run("boolean literals", func(t *testing.T) {
		result := Scan(
			"true false",
		)

		expectToken(result.Tokens[0], TokenTypeTrue, 0, 4)
		expectToken(result.Tokens[1], TokenTypeFalse, 5, 5)
		expectToken(result.Tokens[2], TokenTypeEof, 10, 0)
		assert.Equal(t, 0, len(result.NewLineOffsets))
	})

	t.Run("built in types", func(t *testing.T) {
		result := Scan(
			"Bool Float64 Int64 String",
		)

		expectToken(result.Tokens[0], TokenTypeBuiltInType, 0, 4)
		expectToken(result.Tokens[1], TokenTypeBuiltInType, 5, 7)
		expectToken(result.Tokens[2], TokenTypeBuiltInType, 13, 5)
		expectToken(result.Tokens[3], TokenTypeBuiltInType, 19, 6)
		expectToken(result.Tokens[4], TokenTypeEof, 25, 0)
		assert.Equal(t, 0, len(result.NewLineOffsets))
	})

}

//---------------------------------------------------------------------------------------------------------------------
