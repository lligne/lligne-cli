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

	expectToken := func(scanner ILligneScanner, tokenType LligneTokenType, text string, startPos int) {
		actualToken := scanner.ReadToken()

		expectedToken := LligneToken{
			TokenType:      tokenType,
			Text:           text,
			SourceStartPos: startPos,
		}

		assert.Equal(t, expectedToken, actualToken)
	}

	expectTokenType := func(scanner ILligneScanner, tokenType LligneTokenType, startPos int) {
		expectToken(scanner, tokenType, tokenType.String(), startPos)
	}

	t.Run("empty string", func(t *testing.T) {
		scanner := NewLligneScanner(
			"",
		)

		expectToken(scanner, TokenTypeEof, "", 0)
	})

	t.Run("a few punctuation tokens", func(t *testing.T) {
		scanner := NewLligneScanner(
			"& &&\n *: , ",
		)

		expectTokenType(scanner, TokenTypeAmpersand, 0)
		expectTokenType(scanner, TokenTypeAmpersandAmpersand, 2)
		expectTokenType(scanner, TokenTypeAsterisk, 6)
		expectTokenType(scanner, TokenTypeColon, 7)
		expectTokenType(scanner, TokenTypeComma, 9)
		expectToken(scanner, TokenTypeEof, "", 11)
	})

	t.Run("a few identifier tokens", func(t *testing.T) {
		scanner := NewLligneScanner(
			"a bb c23_f _dfg",
		)

		expectToken(scanner, TokenTypeIdentifier, "a", 0)
		expectToken(scanner, TokenTypeIdentifier, "bb", 2)
		expectToken(scanner, TokenTypeIdentifier, "c23_f", 5)
		expectToken(scanner, TokenTypeIdentifier, "_dfg", 11)
		expectToken(scanner, TokenTypeEof, "", 15)
	})

	t.Run("a few integers", func(t *testing.T) {
		scanner := NewLligneScanner(
			"123 4\n(99000) 5",
		)

		expectToken(scanner, TokenTypeIntegerLiteral, "123", 0)
		expectToken(scanner, TokenTypeIntegerLiteral, "4", 4)
		expectToken(scanner, TokenTypeLeftParenthesis, "(", 6)
		expectToken(scanner, TokenTypeIntegerLiteral, "99000", 7)
		expectToken(scanner, TokenTypeRightParenthesis, ")", 12)
		expectToken(scanner, TokenTypeIntegerLiteral, "5", 14)
		expectToken(scanner, TokenTypeEof, "", 15)
	})

	t.Run("a few numbers", func(t *testing.T) {
		scanner := NewLligneScanner(
			"12.3 4\n(990.00) 5.1",
		)

		expectToken(scanner, TokenTypeFloatingPointLiteral, "12.3", 0)
		expectToken(scanner, TokenTypeIntegerLiteral, "4", 5)
		expectToken(scanner, TokenTypeLeftParenthesis, "(", 7)
		expectToken(scanner, TokenTypeFloatingPointLiteral, "990.00", 8)
		expectToken(scanner, TokenTypeRightParenthesis, ")", 14)
		expectToken(scanner, TokenTypeFloatingPointLiteral, "5.1", 16)
		expectToken(scanner, TokenTypeEof, "", 19)
	})

	t.Run("a few double quoted strings", func(t *testing.T) {
		scanner := NewLligneScanner(
			`"abc" "xyz" "bad
 "start over"`,
		)

		expectToken(scanner, TokenTypeDoubleQuotedString, `"abc"`, 0)
		expectToken(scanner, TokenTypeDoubleQuotedString, `"xyz"`, 6)
		expectToken(scanner, TokenTypeUnclosedDoubleQuotedString, `"bad`, 12)
		expectToken(scanner, TokenTypeDoubleQuotedString, `"start over"`, 18)
		expectToken(scanner, TokenTypeEof, "", 30)
	})

	t.Run("a few single quoted strings", func(t *testing.T) {
		scanner := NewLligneScanner(
			`'abc' 'xyz' 'bad
 'start over'`,
		)

		expectToken(scanner, TokenTypeSingleQuotedString, `'abc'`, 0)
		expectToken(scanner, TokenTypeSingleQuotedString, `'xyz'`, 6)
		expectToken(scanner, TokenTypeUnclosedSingleQuotedString, `'bad`, 12)
		expectToken(scanner, TokenTypeSingleQuotedString, `'start over'`, 18)
		expectToken(scanner, TokenTypeEof, "", 30)
	})

	t.Run("all fixed text tokens, one at a time", func(t *testing.T) {
		for tokenType := TokenTypeEof; tokenType < TokenType_Count; tokenType += 1 {
			sourceCode := tokenType.String()

			// Skip tokens that can have different text for the same token type.
			if strings.HasPrefix(sourceCode, "[") && strings.HasSuffix(sourceCode, "]") {
				continue
			}

			scanner := NewLligneScanner(
				sourceCode,
			)

			expectTokenType(scanner, tokenType, 0)
			expectToken(scanner, TokenTypeEof, "", len(sourceCode))
		}
	})

	t.Run("a few back-ticked string lines", func(t *testing.T) {
		scanner := NewLligneScanner(
			"`abc 123\n`  - one\n  `  - two\n\n  `another\n\n  `one more\n `and the end",
		)

		expectToken(scanner, TokenTypeBackTickedString, "`abc 123\n`  - one\n`  - two\n", 0)
		expectToken(scanner, TokenTypeBackTickedString, "`another\n", 32)
		expectToken(scanner, TokenTypeBackTickedString, "`one more\n`and the end\n", 44)
		expectToken(scanner, TokenTypeEof, "", 67)
	})

	t.Run("a few documentation lines", func(t *testing.T) {
		scanner := NewLligneScanner(
			"// abc 123\n//  - one\n//two\n\n//\n//",
		)

		expectToken(scanner, TokenTypeDocumentation, "// abc 123\n//  - one\n//two\n", 0)
		expectToken(scanner, TokenTypeDocumentation, "//\n//\n", 28)
		expectToken(scanner, TokenTypeEof, "", 33)
	})

	t.Run("boolean literals", func(t *testing.T) {
		scanner := NewLligneScanner(
			"true false",
		)

		expectToken(scanner, TokenTypeTrue, "true", 0)
		expectToken(scanner, TokenTypeFalse, "false", 5)
		expectToken(scanner, TokenTypeEof, "", 10)
	})

}

//---------------------------------------------------------------------------------------------------------------------
