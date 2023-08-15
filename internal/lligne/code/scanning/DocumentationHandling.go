//
// # Scanner for Lligne tokens.
//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package scanning

import "strings"

//=====================================================================================================================

// RemoveDocumentation removes all documentation tokens from a given token array.
func RemoveDocumentation(tokens []Token) []Token {
	result := make([]Token, 0)

	for _, token := range tokens {
		if token.TokenType != TokenTypeDocumentation {
			result = append(result, token)
		}
	}

	return result
}

//=====================================================================================================================

// ProcessLeadingTrailingDocumentation converts multiline documentation tokens to leading or trailing documentation.
func ProcessLeadingTrailingDocumentation(sourceCode string, tokens []Token) []Token {
	result := make([]Token, 0)

	index := 0
	for index < len(tokens)-1 {
		if tokens[index].TokenType == TokenTypeDocumentation {
			result = append(result, Token{
				SourceOffset: tokens[index].SourceOffset,
				SourceLength: tokens[index].SourceLength,
				TokenType:    TokenTypeLeadingDocumentation,
			})
			result = append(result, Token{
				SourceOffset: tokens[index].SourceOffset,
				SourceLength: 0,
				TokenType:    TokenTypeSynthDocument,
			})
			index += 1
		} else if tokens[index+1].TokenType == TokenTypeDocumentation {
			if tokensOnSameLine(sourceCode, tokens[index].SourceOffset, tokens[index+1].SourceOffset) {

				if tokens[index].TokenType == TokenTypeComma || tokens[index].TokenType == TokenTypeSemicolon {
					result = append(result, Token{
						SourceOffset: tokens[index+1].SourceOffset,
						SourceLength: 0,
						TokenType:    TokenTypeSynthDocument,
					})
					result = append(result, Token{
						SourceOffset: tokens[index+1].SourceOffset,
						SourceLength: tokens[index+1].SourceLength,
						TokenType:    TokenTypeTrailingDocumentation,
					})
				}

				result = append(result, tokens[index])

				if tokens[index].TokenType != TokenTypeComma && tokens[index].TokenType != TokenTypeSemicolon {
					result = append(result, Token{
						SourceOffset: tokens[index+1].SourceOffset,
						SourceLength: 0,
						TokenType:    TokenTypeSynthDocument,
					})
					result = append(result, Token{
						SourceOffset: tokens[index+1].SourceOffset,
						SourceLength: tokens[index+1].SourceLength,
						TokenType:    TokenTypeTrailingDocumentation,
					})
				}

				index += 2
			} else {
				result = append(result, tokens[index])

				result = append(result, Token{
					SourceOffset: tokens[index+1].SourceOffset,
					SourceLength: tokens[index+1].SourceLength,
					TokenType:    TokenTypeLeadingDocumentation,
				})
				result = append(result, Token{
					SourceOffset: tokens[index+1].SourceOffset,
					SourceLength: 0,
					TokenType:    TokenTypeSynthDocument,
				})
				index += 2

			}
		} else {
			result = append(result, tokens[index])
			index += 1
		}
	}

	return result
}

//---------------------------------------------------------------------------------------------------------------------

// tokensOnSameLine looks for a line feed in the source code between two tokens.
func tokensOnSameLine(sourceCode string, token1StartPos uint32, token2StartPos uint32) bool {
	return strings.IndexByte(sourceCode[token1StartPos:token2StartPos], '\n') < 0
}

//=====================================================================================================================
