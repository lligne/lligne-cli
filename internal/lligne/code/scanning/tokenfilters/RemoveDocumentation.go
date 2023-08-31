//
// # Scanner for Lligne tokens.
//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package tokenfilters

import (
	"lligne-cli/internal/lligne/code/scanning"
)

//=====================================================================================================================

// RemoveDocumentation removes all documentation tokens from a given scan result.
func RemoveDocumentation(scanResult *scanning.Outcome) *scanning.Outcome {
	filteredTokens := make([]scanning.Token, 0)

	for _, token := range scanResult.Tokens {
		if token.TokenType != scanning.TokenTypeDocumentation {
			filteredTokens = append(filteredTokens, token)
		}
	}

	return &scanning.Outcome{
		SourceCode:     scanResult.SourceCode,
		Tokens:         filteredTokens,
		NewLineOffsets: scanResult.NewLineOffsets,
	}
}

//=====================================================================================================================
