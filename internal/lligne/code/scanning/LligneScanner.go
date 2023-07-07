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

	switch {
	case isIdentifierStart(ch):
		return s.scanIdentifierOrKeyword()
	}

	switch ch {
	case '&':
		return s.oneOrTwoRuneToken(TokenTypeAmpersand, '&', TokenTypeAmpersandAmpersand)
	case '*':
		return s.token(TokenTypeAsterisk)
	case ':':
		return s.token(TokenTypeColon)
	case ',':
		return s.token(TokenTypeComma)
	case '-':
		return s.oneOrTwoRuneToken(TokenTypeDash, '>', TokenTypeRightArrow)
	case '.':
		return s.oneToThreeRuneToken(TokenTypeDot, '.', TokenTypeDotDot, '.', TokenTypeDotDotDot)
	case '=':
		return s.scanAfterEquals()
	case '!':
		return s.oneOrTwoRuneToken(TokenTypeExclamationMark, '~', TokenTypeNotMatches)
	case '<':
		return s.oneOrTwoRuneToken(TokenTypeLessThan, '=', TokenTypeLessThanOrEquals)
	case '>':
		return s.oneOrTwoRuneToken(TokenTypeGreaterThan, '=', TokenTypeGreaterThanOrEquals)
	case '{':
		return s.token(TokenTypeLeftBrace)
	case '[':
		return s.token(TokenTypeLeftBracket)
	case '(':
		return s.token(TokenTypeLeftParenthesis)
	case '+':
		return s.token(TokenTypePlus)
	case '?':
		return s.oneOrTwoRuneToken(TokenTypeQuestionMark, ':', TokenTypeQuestionMarkColon)
	case '}':
		return s.token(TokenTypeRightBrace)
	case ']':
		return s.token(TokenTypeRightBracket)
	case ')':
		return s.token(TokenTypeRightParenthesis)
	case ';':
		return s.token(TokenTypeSemicolon)
	case '/':
		return s.token(TokenTypeSlash)
	case '|':
		return s.token(TokenTypeVerticalBar)
	case 0:
		return s.token(TokenTypeEof)
	}

	return s.token(TokenTypeUnrecognizedChar)

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

func (s *LligneScanner) advanceIfIdentifierPart() bool {

	if s.currentPos >= len(s.sourceCode) {
		return false
	}

	found, width := utf8.DecodeRuneInString(s.sourceCode[s.currentPos:])

	if !isIdentifierPart(found) {
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

func isIdentifierPart(ch rune) bool {
	return isIdentifierStart(ch) || '0' <= ch && ch <= '9' || ch >= utf8.RuneSelf && unicode.IsNumber(ch)
}

//---------------------------------------------------------------------------------------------------------------------

func isIdentifierStart(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch >= utf8.RuneSelf && unicode.IsLetter(ch)
}

//---------------------------------------------------------------------------------------------------------------------

func (s *LligneScanner) oneOrTwoRuneToken(
	oneRuneType LligneTokenType,
	secondRune rune,
	twoRuneType LligneTokenType,
) LligneToken {

	if s.advanceIf(secondRune) {
		return s.token(twoRuneType)
	}

	return s.token(oneRuneType)

}

//---------------------------------------------------------------------------------------------------------------------

func (s *LligneScanner) oneToThreeRuneToken(
	oneRuneType LligneTokenType,
	secondRune rune,
	twoRuneType LligneTokenType,
	thirdRune rune,
	threeRuneType LligneTokenType,
) LligneToken {

	if s.advanceIf(secondRune) {

		if s.advanceIf(thirdRune) {
			return s.token(threeRuneType)
		}

		return s.token(twoRuneType)

	}

	return s.token(oneRuneType)

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

func (s *LligneScanner) scanAfterEquals() LligneToken {

	if s.advanceIf('=') {

		if s.advanceIf('=') {
			return s.token(TokenTypeEqualsEqualsEquals)
		}

		return s.token(TokenTypeEqualsEquals)

	}

	if s.advanceIf('~') {
		return s.token(TokenTypeMatches)
	}

	return s.token(TokenTypeEquals)

}

//---------------------------------------------------------------------------------------------------------------------

func (s *LligneScanner) scanIdentifierOrKeyword() LligneToken {

	for s.advanceIfIdentifierPart() {
	}

	text := s.sourceCode[s.markedPos:s.currentPos]

	switch text {
	case "and":
		return LligneToken{TokenTypeAnd, text, s.markedPos}
	case "as":
		return LligneToken{TokenTypeAs, text, s.markedPos}
	case "is":
		return LligneToken{TokenTypeIs, text, s.markedPos}
	case "in":
		return LligneToken{TokenTypeIn, text, s.markedPos}
	case "not":
		return LligneToken{TokenTypeNot, text, s.markedPos}
	case "of":
		return LligneToken{TokenTypeOf, text, s.markedPos}
	case "or":
		return LligneToken{TokenTypeOr, text, s.markedPos}
	case "to":
		return LligneToken{TokenTypeTo, text, s.markedPos}
	default:
		return LligneToken{TokenTypeIdentifier, text, s.markedPos}
	}

}

//---------------------------------------------------------------------------------------------------------------------

func (s *LligneScanner) token(tokenType LligneTokenType) LligneToken {
	return LligneToken{
		TokenType:      tokenType,
		Text:           s.sourceCode[s.markedPos:s.currentPos],
		SourceStartPos: s.markedPos,
	}
}

//---------------------------------------------------------------------------------------------------------------------
