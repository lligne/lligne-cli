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

func TestLligneScanner(t *testing.T) {

	expectToken := func(scanner *LligneScanner, tokenType LligneTokenType, text string, line int, column int) {
		actualToken := scanner.ReadToken()

		expectedToken := LligneToken{
			TokenType: tokenType,
			Text:      text,
			Origin: &LligneOrigin{
				FileName: "sample.lligne",
				Line:     line,
				Column:   column,
			},
		}

		assert.Equal(t, expectedToken, actualToken)
	}

	expectTokenType := func(scanner *LligneScanner, tokenType LligneTokenType, line int, column int) {
		expectToken(scanner, tokenType, tokenType.String(), line, column)
	}

	t.Run("a few punctuation tokens", func(t *testing.T) {
		scanner := NewLligneScanner(
			"sample.lligne",
			"& &&\n *: , ",
		)

		expectTokenType(&scanner, TokenTypeAmpersand, 1, 1)
		expectTokenType(&scanner, TokenTypeAmpersandAmpersand, 1, 3)
		expectTokenType(&scanner, TokenTypeAsterisk, 2, 2)
		expectTokenType(&scanner, TokenTypeColon, 2, 3)
		expectTokenType(&scanner, TokenTypeComma, 2, 5)
		expectToken(&scanner, TokenTypeEof, "", 2, 7)
	})

}
