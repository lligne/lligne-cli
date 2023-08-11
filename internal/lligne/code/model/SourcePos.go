//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package model

import "lligne-cli/internal/lligne/code/scanning"

//=====================================================================================================================

type SourcePos struct {
	startOffset uint32
	endOffset   uint32
}

//---------------------------------------------------------------------------------------------------------------------

func NewSourcePos(token scanning.Token) SourcePos {
	return SourcePos{
		startOffset: token.SourceOffset,
		endOffset:   token.SourceOffset + uint32(token.SourceLength),
	}
}

//---------------------------------------------------------------------------------------------------------------------

// GetText slices the given sourceCode to produce the string demarcated by the source position.
func (s SourcePos) GetText(sourceCode string) string {
	return sourceCode[s.startOffset:s.endOffset]
}

//---------------------------------------------------------------------------------------------------------------------

// Thru creates a new source position extending from the start of one to the end of another.
func (s SourcePos) Thru(s2 SourcePos) SourcePos {

	if s2.endOffset < s.startOffset {
		panic("Source Positions not in correct order.")
	}

	return SourcePos{
		startOffset: s.startOffset,
		endOffset:   s2.endOffset,
	}

}

//=====================================================================================================================
