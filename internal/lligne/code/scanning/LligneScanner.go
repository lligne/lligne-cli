//
// # Scanner for Lligne tokens.
//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package scanning

import (
	"unicode"
	"unicode/utf8"
)

//---------------------------------------------------------------------------------------------------------------------

// LligneScanner converts a string of Lligne source code into tokens.
type LligneScanner struct {
	fileName   string
	sourceCode string
	currentPos int
	markedPos  int
}

//---------------------------------------------------------------------------------------------------------------------

// NewLligneScanner allocates a new scanner for given sourceCode from the given fileName.
func NewLligneScanner(fileName string, sourceCode string) LligneScanner {
	return LligneScanner{
		fileName:   fileName,
		sourceCode: sourceCode,
		currentPos: 0,
		markedPos:  0,
	}
}

//---------------------------------------------------------------------------------------------------------------------

// ReadToken returns the next token from the scanner.
func (s *LligneScanner) ReadToken() LligneToken {

	// Ignore whitespace
	for s.advanceIfWhitespace() {
	}

	// Mark the start of the token
	s.markedPos = s.currentPos

	ch := s.readRune()

	switch ch {
	case '&':
		return s.oneOrTwoRuneToken(TokenTypeAmpersand, '&', TokenTypeAmpersandAmpersand)
	case '*':
		return s.newToken(TokenTypeAsterisk)
	case ':':
		return s.newToken(TokenTypeColon)
	case ',':
		return s.newToken(TokenTypeComma)
	case 0:
		return s.newToken(TokenTypeEof)
	}

	return s.newToken(TokenTypeUnrecognizedChar)

}

//---------------------------------------------------------------------------------------------------------------------
//---------------------------------------------------------------------------------------------------------------------

func (s *LligneScanner) advance(width int, result rune) {

	if result == '\n' {
		// TODO: track line break positions
	}

	s.currentPos += width

}

//---------------------------------------------------------------------------------------------------------------------

func (s *LligneScanner) advanceIf(needed rune) bool {

	if s.currentPos >= len(s.sourceCode) {
		return false
	}

	found, width := utf8.DecodeRuneInString(s.sourceCode[s.currentPos:])

	if found != needed {
		return false
	}

	s.advance(width, found)

	return true

}

//---------------------------------------------------------------------------------------------------------------------

func (s *LligneScanner) advanceIfWhitespace() bool {

	if s.currentPos >= len(s.sourceCode) {
		return false
	}

	found, width := utf8.DecodeRuneInString(s.sourceCode[s.currentPos:])

	if !unicode.IsSpace(found) {
		return false
	}

	s.advance(width, found)

	return true

}

//---------------------------------------------------------------------------------------------------------------------

func (s *LligneScanner) newToken(tokenType LligneTokenType) LligneToken {
	return LligneToken{
		TokenType:      tokenType,
		Text:           s.sourceCode[s.markedPos:s.currentPos],
		SourceStartPos: s.markedPos,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *LligneScanner) oneOrTwoRuneToken(
	oneRuneType LligneTokenType,
	secondRune rune,
	twoRuneType LligneTokenType,
) LligneToken {

	if s.advanceIf(secondRune) {
		return s.newToken(twoRuneType)
	}

	return s.newToken(oneRuneType)

}

//---------------------------------------------------------------------------------------------------------------------

func (s *LligneScanner) peekRune() rune {

	if s.currentPos >= len(s.sourceCode) {
		return 0
	}

	result, _ := utf8.DecodeRuneInString(s.sourceCode[s.currentPos:])

	return result

}

//---------------------------------------------------------------------------------------------------------------------

func (s *LligneScanner) readRune() rune {

	if s.currentPos >= len(s.sourceCode) {
		return 0
	}

	result, width := utf8.DecodeRuneInString(s.sourceCode[s.currentPos:])

	s.advance(width, result)

	return result

}

//---------------------------------------------------------------------------------------------------------------------
