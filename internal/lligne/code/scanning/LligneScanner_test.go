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

func TestLligneScanner(t *testing.T) {

	expectToken := func(scanner *LligneScanner, tokenType LligneTokenType, text string, startPos int) {
		actualToken := scanner.ReadToken()

		expectedToken := LligneToken{
			TokenType:      tokenType,
			Text:           text,
			SourceStartPos: startPos,
		}

		assert.Equal(t, expectedToken, actualToken)
	}

	expectTokenType := func(scanner *LligneScanner, tokenType LligneTokenType, startPos int) {
		expectToken(scanner, tokenType, tokenType.String(), startPos)
	}

	t.Run("a few punctuation tokens", func(t *testing.T) {
		scanner := NewLligneScanner(
			"sample.lligne",
			"& &&\n *: , ",
		)

		expectTokenType(&scanner, TokenTypeAmpersand, 0)
		expectTokenType(&scanner, TokenTypeAmpersandAmpersand, 2)
		expectTokenType(&scanner, TokenTypeAsterisk, 6)
		expectTokenType(&scanner, TokenTypeColon, 7)
		expectTokenType(&scanner, TokenTypeComma, 9)
		expectToken(&scanner, TokenTypeEof, "", 11)
	})

	t.Run("a few identifier tokens", func(t *testing.T) {
		scanner := NewLligneScanner(
			"sample.lligne",
			"a bb c23_f _dfg",
		)

		expectToken(&scanner, TokenTypeIdentifier, "a", 0)
		expectToken(&scanner, TokenTypeIdentifier, "bb", 2)
		expectToken(&scanner, TokenTypeIdentifier, "c23_f", 5)
		expectToken(&scanner, TokenTypeIdentifier, "_dfg", 11)
		expectToken(&scanner, TokenTypeEof, "", 15)
	})

	t.Run("a few numbers", func(t *testing.T) {
		scanner := NewLligneScanner(
			"sample.lligne",
			"123 4\n(99000) 5",
		)

		expectToken(&scanner, TokenTypeIntegerLiteral, "123", 0)
		expectToken(&scanner, TokenTypeIntegerLiteral, "4", 4)
		expectToken(&scanner, TokenTypeLeftParenthesis, "(", 6)
		expectToken(&scanner, TokenTypeIntegerLiteral, "99000", 7)
		expectToken(&scanner, TokenTypeRightParenthesis, ")", 12)
		expectToken(&scanner, TokenTypeIntegerLiteral, "5", 14)
		expectToken(&scanner, TokenTypeEof, "", 15)
	})

	t.Run("a few strings", func(t *testing.T) {
		scanner := NewLligneScanner(
			"sample.lligne",
			`"abc" "xyz" "bad
 "start over"`,
		)

		expectToken(&scanner, TokenTypeStringLiteral, `"abc"`, 0)
		expectToken(&scanner, TokenTypeStringLiteral, `"xyz"`, 6)
		expectToken(&scanner, TokenTypeUnclosedString, `"bad`, 12)
		expectToken(&scanner, TokenTypeStringLiteral, `"start over"`, 18)
		expectToken(&scanner, TokenTypeEof, "", 30)
	})

	t.Run("all fixed text tokens, one at a time", func(t *testing.T) {
		for tokenType := TokenTypeEof; tokenType < TokenType_Count; tokenType += 1 {
			sourceCode := tokenType.String()

			// Skip tokens that can have different text for the same token type.
			if strings.HasPrefix(sourceCode, "[") && strings.HasSuffix(sourceCode, "]") {
				continue
			}

			scanner := NewLligneScanner(
				"sample.lligne",
				sourceCode,
			)

			expectTokenType(&scanner, tokenType, 0)
			expectToken(&scanner, TokenTypeEof, "", len(sourceCode))
		}
	})

	t.Run("a few back-ticked string lines", func(t *testing.T) {
		scanner := NewLligneScanner(
			"sample.lligne",
			"`abc 123\n`  - one\n`  - two\n`",
		)

		expectToken(&scanner, TokenTypeBackTickedString, "`abc 123", 0)
		expectToken(&scanner, TokenTypeBackTickedString, "`  - one", 9)
		expectToken(&scanner, TokenTypeBackTickedString, "`  - two", 18)
		expectToken(&scanner, TokenTypeBackTickedString, "`", 27)
		expectToken(&scanner, TokenTypeEof, "", 28)
	})

	t.Run("a few documentation lines", func(t *testing.T) {
		scanner := NewLligneScanner(
			"sample.lligne",
			"// abc 123\n//  - one\n//two\n//\n//",
		)

		expectToken(&scanner, TokenTypeDocumentation, "// abc 123", 0)
		expectToken(&scanner, TokenTypeDocumentation, "//  - one", 11)
		expectToken(&scanner, TokenTypeDocumentation, "//two", 21)
		expectToken(&scanner, TokenTypeDocumentation, "//", 27)
		expectToken(&scanner, TokenTypeDocumentation, "//", 30)
		expectToken(&scanner, TokenTypeEof, "", 32)
	})

}
