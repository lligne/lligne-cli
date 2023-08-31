//
// # Scanner for Lligne tokens.
//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package tokenfilters

import (
	"lligne-cli/internal/lligne/code/scanning"
	"strings"
)

//=====================================================================================================================

// ProcessLeadingTrailingDocumentation converts multiline documentation tokens to leading or trailing documentation.
func ProcessLeadingTrailingDocumentation(scanResult *scanning.Outcome) *scanning.Outcome {
	tokens := scanResult.Tokens
	result := make([]scanning.Token, 0)

	index := 0
	for index < len(tokens)-1 {
		if tokens[index].TokenType == scanning.TokenTypeDocumentation {
			result = append(result, scanning.Token{
				SourceOffset: tokens[index].SourceOffset,
				SourceLength: tokens[index].SourceLength,
				TokenType:    scanning.TokenTypeLeadingDocumentation,
			})
			result = append(result, scanning.Token{
				SourceOffset: tokens[index].SourceOffset,
				SourceLength: 0,
				TokenType:    scanning.TokenTypeSynthDocument,
			})
			index += 1
		} else if tokens[index+1].TokenType == scanning.TokenTypeDocumentation {
			if tokensOnSameLine(scanResult.SourceCode, tokens[index].SourceOffset, tokens[index+1].SourceOffset) {

				if tokens[index].TokenType == scanning.TokenTypeComma || tokens[index].TokenType == scanning.TokenTypeSemicolon {
					result = append(result, scanning.Token{
						SourceOffset: tokens[index+1].SourceOffset,
						SourceLength: 0,
						TokenType:    scanning.TokenTypeSynthDocument,
					})
					result = append(result, scanning.Token{
						SourceOffset: tokens[index+1].SourceOffset,
						SourceLength: tokens[index+1].SourceLength,
						TokenType:    scanning.TokenTypeTrailingDocumentation,
					})
				}

				result = append(result, tokens[index])

				if tokens[index].TokenType != scanning.TokenTypeComma && tokens[index].TokenType != scanning.TokenTypeSemicolon {
					result = append(result, scanning.Token{
						SourceOffset: tokens[index+1].SourceOffset,
						SourceLength: 0,
						TokenType:    scanning.TokenTypeSynthDocument,
					})
					result = append(result, scanning.Token{
						SourceOffset: tokens[index+1].SourceOffset,
						SourceLength: tokens[index+1].SourceLength,
						TokenType:    scanning.TokenTypeTrailingDocumentation,
					})
				}

				index += 2
			} else {
				result = append(result, tokens[index])

				result = append(result, scanning.Token{
					SourceOffset: tokens[index+1].SourceOffset,
					SourceLength: tokens[index+1].SourceLength,
					TokenType:    scanning.TokenTypeLeadingDocumentation,
				})
				result = append(result, scanning.Token{
					SourceOffset: tokens[index+1].SourceOffset,
					SourceLength: 0,
					TokenType:    scanning.TokenTypeSynthDocument,
				})
				index += 2

			}
		} else {
			result = append(result, tokens[index])
			index += 1
		}
	}

	return &scanning.Outcome{
		SourceCode:     scanResult.SourceCode,
		Tokens:         result,
		NewLineOffsets: scanResult.NewLineOffsets,
	}
}

//---------------------------------------------------------------------------------------------------------------------

// tokensOnSameLine looks for a line feed in the source code between two tokens.
func tokensOnSameLine(sourceCode string, token1StartPos uint32, token2StartPos uint32) bool {
	return strings.IndexByte(sourceCode[token1StartPos:token2StartPos], '\n') < 0
}

//=====================================================================================================================
