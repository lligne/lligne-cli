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
	currentPos sourcePos
	markedPos  sourcePos
}

//---------------------------------------------------------------------------------------------------------------------

// NewLligneScanner allocates a new scanner for given sourceCode from the given fileName.
func NewLligneScanner(fileName string, sourceCode string) LligneScanner {
	return LligneScanner{
		fileName:   fileName,
		sourceCode: sourceCode,
		currentPos: sourcePos{
			charPos: 0,
			line:    1,
			column:  1,
		},
		markedPos: sourcePos{
			charPos: 0,
			line:    1,
			column:  1,
		},
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
		if s.advanceIf('&') {
			return s.newToken(TokenTypeAmpersandAmpersand)
		}
		return s.newToken(TokenTypeAmpersand)
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

type sourcePos struct {
	charPos int
	line    int
	column  int
}

//---------------------------------------------------------------------------------------------------------------------

func (s *LligneScanner) advance(width int, result rune) {

	s.currentPos.charPos += width

	s.currentPos.column += 1
	if result == '\n' {
		s.currentPos.line += 1
		s.currentPos.column = 1
	}

}

//---------------------------------------------------------------------------------------------------------------------

func (s *LligneScanner) advanceIf(needed rune) bool {

	if s.currentPos.charPos >= len(s.sourceCode) {
		return false
	}

	found, width := utf8.DecodeRuneInString(s.sourceCode[s.currentPos.charPos:])

	if found != needed {
		return false
	}

	s.advance(width, found)

	return true

}

//---------------------------------------------------------------------------------------------------------------------

func (s *LligneScanner) advanceIfWhitespace() bool {

	if s.currentPos.charPos >= len(s.sourceCode) {
		return false
	}

	found, width := utf8.DecodeRuneInString(s.sourceCode[s.currentPos.charPos:])

	if !unicode.IsSpace(found) {
		return false
	}

	s.advance(width, found)

	return true

}

//---------------------------------------------------------------------------------------------------------------------

func (s *LligneScanner) newToken(tokenType LligneTokenType) LligneToken {
	return LligneToken{
		TokenType: tokenType,
		Text:      s.sourceCode[s.markedPos.charPos:s.currentPos.charPos],
		Origin:    s.originFromMark(),
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *LligneScanner) originFromMark() *LligneOrigin {
	return &LligneOrigin{
		FileName: s.fileName,
		Line:     s.markedPos.line,
		Column:   s.markedPos.column,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *LligneScanner) readRune() rune {

	if s.currentPos.charPos >= len(s.sourceCode) {
		return 0
	}

	result, width := utf8.DecodeRuneInString(s.sourceCode[s.currentPos.charPos:])

	s.advance(width, result)

	return result

}

//---------------------------------------------------------------------------------------------------------------------
