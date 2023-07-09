//
// # Scanner adapter to handle Lligne documentation, splitting it into leading and trailing documentation
//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package scanning

import "strings"

//---------------------------------------------------------------------------------------------------------------------

type lligneDocumentationHandlingScanner struct {
	sourceCode      string
	scanner         ILligneScanner
	tokenAhead      [3]LligneToken
	tokenAheadCount int
}

//---------------------------------------------------------------------------------------------------------------------

func NewLligneDocumentationHandlingScanner(sourceCode string, scanner ILligneScanner) ILligneScanner {

	// Allocate the handling scanner, delegating to the given scanner.
	s := &lligneDocumentationHandlingScanner{
		sourceCode: sourceCode,
		scanner:    scanner,
	}

	// Buffer the first token
	s.tokenAhead[0] = scanner.ReadToken()
	s.tokenAheadCount = 1

	if s.tokenAhead[0].TokenType == TokenTypeDocumentation {
		// Documentation right off the bat is necessarily leading documentation; change the token type.
		s.tokenAhead[0] = LligneToken{TokenTypeLeadingDocumentation, s.tokenAhead[0].Text, s.tokenAhead[0].SourceStartPos}

		// Add a synthetic documentation operator.
		s.tokenAhead[1] = LligneToken{TokenTypeSynthDocument, " ", s.tokenAhead[0].SourceStartPos}
		s.tokenAheadCount = 2
	}

	return s
}

//---------------------------------------------------------------------------------------------------------------------

func (s *lligneDocumentationHandlingScanner) ReadToken() LligneToken {

	// If needed, read the next token from the inner scanner.
	if s.tokenAheadCount == 1 {
		s.tokenAhead[1] = s.scanner.ReadToken()
		s.tokenAheadCount = 2
	}

	// Convert documentation tokens to leading or trailing.
	if s.tokenAhead[1].TokenType == TokenTypeDocumentation {

		if s.tokenAhead[0].TokenType != TokenTypeVerticalBar &&
			s.tokensOnSameLine(s.tokenAhead[0].SourceStartPos, s.tokenAhead[1].SourceStartPos) {

			// Add a synthetic documentation operator before the trailing documentation.
			s.tokenAhead[1] = LligneToken{
				TokenTypeSynthDocument,
				" ",
				s.tokenAhead[0].SourceStartPos,
			}

			// Convert to trailing documentation after other tokens on the same line except '|'.
			s.tokenAhead[2] = LligneToken{
				TokenTypeTrailingDocumentation,
				s.tokenAhead[1].Text,
				s.tokenAhead[1].SourceStartPos,
			}

			// Reorder trailing documentation as if it came before a comma or semicolon.
			if s.tokenAhead[0].TokenType == TokenTypeComma || s.tokenAhead[0].TokenType == TokenTypeSemicolon {
				temp := s.tokenAhead[0]
				s.tokenAhead[0] = s.tokenAhead[1]
				s.tokenAhead[1] = s.tokenAhead[2]
				s.tokenAhead[2] = temp
			}

		} else {

			// Otherwise convert to leading documentation.
			s.tokenAhead[1] = LligneToken{
				TokenTypeLeadingDocumentation,
				s.tokenAhead[1].Text,
				s.tokenAhead[1].SourceStartPos,
			}

			// Add a synthetic documentation operator.
			s.tokenAhead[2] = LligneToken{TokenTypeSynthDocument, " ", s.tokenAhead[0].SourceStartPos}

		}

		s.tokenAheadCount = 3

	}

	// Advance the three level token buffer.
	result := s.tokenAhead[0]
	s.tokenAhead[0] = s.tokenAhead[1]
	s.tokenAhead[1] = s.tokenAhead[2]
	s.tokenAheadCount -= 1

	return result

}

//---------------------------------------------------------------------------------------------------------------------

// tokensOnSameLine looks for a line feed in the source code between two tokens.
func (s *lligneDocumentationHandlingScanner) tokensOnSameLine(token1StartPos int, token2StartPos int) bool {
	return strings.IndexByte(s.sourceCode[token1StartPos:token2StartPos], '\n') < 0
}

//---------------------------------------------------------------------------------------------------------------------
