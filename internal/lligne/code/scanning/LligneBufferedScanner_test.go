//
// # Tests of LligneBufferedScanner.
//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package scanning

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLligneBufferedScanner(t *testing.T) {

	expectToken := func(scanner ILligneBufferedScanner, tokenType LligneTokenType, text string, startPos int, line int, column int) {
		expectedToken := LligneToken{
			TokenType:      tokenType,
			Text:           text,
			SourceStartPos: startPos,
		}

		assert.True(t, scanner.PeekTokenIsType(tokenType))

		actualToken := scanner.PeekToken()
		assert.Equal(t, expectedToken, actualToken)

		origin := scanner.GetOrigin(startPos)
		assert.Equal(t, line, origin.Line)
		assert.Equal(t, column, origin.Column)

		assert.True(t, scanner.AdvanceTokenIfType(tokenType))
	}

	expectTokenType := func(scanner ILligneBufferedScanner, tokenType LligneTokenType, startPos int, line int, column int) {
		expectToken(scanner, tokenType, tokenType.String(), startPos, line, column)
	}

	t.Run("a few punctuation tokens", func(t *testing.T) {
		scanner := NewLligneBufferedScanner(NewLligneScanner(
			"          & &&\n           * : , ",
		))

		expectTokenType(scanner, TokenTypeAmpersand, 10, 1, 11)
		expectTokenType(scanner, TokenTypeAmpersandAmpersand, 12, 1, 13)
		expectTokenType(scanner, TokenTypeAsterisk, 26, 2, 12)
		expectTokenType(scanner, TokenTypeColon, 28, 2, 14)
		expectTokenType(scanner, TokenTypeComma, 30, 2, 16)
		expectToken(scanner, TokenTypeEof, "", 32, 2, 18)
	})

	t.Run("a few identifier tokens", func(t *testing.T) {
		scanner := NewLligneBufferedScanner(NewLligneScanner(
			"a bb c23_f _dfg",
		))

		expectToken(scanner, TokenTypeIdentifier, "a", 0, 1, 1)
		expectToken(scanner, TokenTypeIdentifier, "bb", 2, 1, 3)
		expectToken(scanner, TokenTypeIdentifier, "c23_f", 5, 1, 6)
		expectToken(scanner, TokenTypeIdentifier, "_dfg", 11, 1, 12)
		expectToken(scanner, TokenTypeEof, "", 15, 1, 16)
	})

	t.Run("a few numbers", func(t *testing.T) {
		scanner := NewLligneBufferedScanner(NewLligneScanner(
			"123 4\n(99000) 5",
		))

		expectToken(scanner, TokenTypeIntegerLiteral, "123", 0, 1, 1)
		expectToken(scanner, TokenTypeIntegerLiteral, "4", 4, 1, 5)
		expectToken(scanner, TokenTypeLeftParenthesis, "(", 6, 2, 1)
		expectToken(scanner, TokenTypeIntegerLiteral, "99000", 7, 2, 2)
		expectToken(scanner, TokenTypeRightParenthesis, ")", 12, 2, 7)
		expectToken(scanner, TokenTypeIntegerLiteral, "5", 14, 2, 9)
		expectToken(scanner, TokenTypeEof, "", 15, 2, 10)
	})

	t.Run("a few double quoted strings", func(t *testing.T) {
		scanner := NewLligneBufferedScanner(NewLligneScanner(
			`"abc" "xyz" "bad
"start over"`,
		))

		expectToken(scanner, TokenTypeDoubleQuotedString, `"abc"`, 0, 1, 1)
		expectToken(scanner, TokenTypeDoubleQuotedString, `"xyz"`, 6, 1, 7)
		expectToken(scanner, TokenTypeUnclosedDoubleQuotedString, `"bad`, 12, 1, 13)
		expectToken(scanner, TokenTypeDoubleQuotedString, `"start over"`, 17, 2, 1)
		expectToken(scanner, TokenTypeEof, "", 29, 2, 13)
	})

	t.Run("a few single quoted strings", func(t *testing.T) {
		scanner := NewLligneBufferedScanner(NewLligneScanner(
			`'abc' 'xyz' 'bad
'start over'`,
		))

		expectToken(scanner, TokenTypeSingleQuotedString, `'abc'`, 0, 1, 1)
		expectToken(scanner, TokenTypeSingleQuotedString, `'xyz'`, 6, 1, 7)
		expectToken(scanner, TokenTypeUnclosedSingleQuotedString, `'bad`, 12, 1, 13)
		expectToken(scanner, TokenTypeSingleQuotedString, `'start over'`, 17, 2, 1)
		expectToken(scanner, TokenTypeEof, "", 29, 2, 13)
	})

	t.Run("a few back-ticked string lines", func(t *testing.T) {
		scanner := NewLligneBufferedScanner(NewLligneScanner(
			"`abc 123\n`  - one\n  `  - two\n\n  `another\n\n  `one more\n `and the end",
		))

		expectToken(scanner, TokenTypeBackTickedString, "`abc 123\n`  - one\n`  - two\n", 0, 1, 1)
		expectToken(scanner, TokenTypeBackTickedString, "`another\n", 32, 5, 3)
		expectToken(scanner, TokenTypeBackTickedString, "`one more\n`and the end\n", 44, 7, 3)
		expectToken(scanner, TokenTypeEof, "", 67, 8, 14)
	})

	t.Run("a few documentation lines", func(t *testing.T) {
		scanner := NewLligneBufferedScanner(NewLligneScanner(
			"// abc 123\n//  - one\n//two\n\n//\n//",
		))

		expectToken(scanner, TokenTypeDocumentation, "// abc 123\n//  - one\n//two\n", 0, 1, 1)
		expectToken(scanner, TokenTypeDocumentation, "//\n//\n", 28, 5, 1)
		expectToken(scanner, TokenTypeEof, "", 33, 6, 3)
	})

}
