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

//=====================================================================================================================

// Scan converts the given source code to an array of tokens plus an array of new line character offsets.
func Scan(sourceCode string) (tokens []Token, newLineOffsets []uint32) {

	// Create a scanner.
	scanner := newScanner(sourceCode)

	// Scan the entire source code.
	scanner.scan()

	// Extract the results.
	tokens = scanner.tokens
	newLineOffsets = scanner.newLineOffsets

	return

}

//=====================================================================================================================

// newScanner allocates a new scanner for given sourceCode from the given fileName.
func newScanner(sourceCode string) *scanner {

	// Create a scanner
	s := &scanner{
		sourceCode:     sourceCode,
		markedPos:      0,
		currentPos:     0,
		newLineOffsets: make([]uint32, 0),
		tokens:         make([]Token, 0),
	}

	// Read the first rune.
	if len(s.sourceCode) > 0 {
		s.runeAhead1, s.runeAhead1Width = utf8.DecodeRuneInString(s.sourceCode[0:])
	}

	if len(s.sourceCode) > s.runeAhead1Width {
		s.runeAhead2, s.runeAhead2Width = utf8.DecodeRuneInString(s.sourceCode[s.runeAhead1Width:])
	}

	return s

}

//=====================================================================================================================

// LligneScanner converts a string of Lligne source code into tokens.
type scanner struct {
	sourceCode      string
	markedPos       int
	currentPos      int
	runeAhead1      rune
	runeAhead2      rune
	runeAhead1Width int
	runeAhead2Width int
	tokens          []Token
	newLineOffsets  []uint32
}

//---------------------------------------------------------------------------------------------------------------------

// scan converts the source code to an array of tokens
func (s *scanner) scan() {
	for {
		token := s.readToken()
		s.tokens = append(s.tokens, token)

		if token.TokenType == TokenTypeEof {
			s.tokens = append(s.tokens, token)
			s.tokens = append(s.tokens, token)
			break
		}
	}
}

//---------------------------------------------------------------------------------------------------------------------

// readToken returns the next token from the scanner.
func (s *scanner) readToken() Token {

	// Ignore whitespace
	for unicode.IsSpace(s.runeAhead1) {
		s.advance()
	}

	// Mark the start of the token
	s.markedPos = s.currentPos

	// Consume the next character.
	ch := s.runeAhead1
	s.advance()

	// Handle character ranges.
	switch {
	case isIdentifierStart(ch):
		return s.scanIdentifierOrKeyword()
	case isDigit(ch):
		return s.scanNumber()
	}

	// Handle individual characters.
	switch ch {
	case '&':
		return s.oneOrTwoRuneToken(TokenTypeAmpersand, '&', TokenTypeAmpersandAmpersand)
	case '*':
		return s.token(TokenTypeAsterisk)
	case '`':
		return s.scanBackTickedString()
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
		return s.scanAfterSlash()
	case '"':
		return s.scanDoubleQuotedString()
	case '\'':
		return s.scanSingleQuotedString()
	case '|':
		return s.token(TokenTypeVerticalBar)
	case 0:
		return s.token(TokenTypeEof)
	}

	return s.token(TokenTypeUnrecognizedChar)

}

//---------------------------------------------------------------------------------------------------------------------

// advance consumes one rune and stages the next one in the scanner.
func (s *scanner) advance() {

	if s.runeAhead1 == '\n' {
		s.newLineOffsets = append(s.newLineOffsets, uint32(s.currentPos))
	}
	s.currentPos += s.runeAhead1Width
	s.runeAhead1 = s.runeAhead2
	s.runeAhead1Width = s.runeAhead2Width

	if s.currentPos+1 >= len(s.sourceCode) {
		s.runeAhead2 = 0
		s.runeAhead2Width = 0
	} else {
		s.runeAhead2, s.runeAhead2Width = utf8.DecodeRuneInString(s.sourceCode[s.currentPos+1:])
	}

}

//---------------------------------------------------------------------------------------------------------------------

// isDigit determines whether a rune is a number.
func isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9' || ch >= utf8.RuneSelf && unicode.IsNumber(ch)
}

//---------------------------------------------------------------------------------------------------------------------

// isIdentifierPart determines whether a given rune could be the second or later character of an identifier.
func isIdentifierPart(ch rune) bool {
	return isIdentifierStart(ch) || isDigit(ch)
}

//---------------------------------------------------------------------------------------------------------------------

// isIdentifierStart determines whether a given rune could be the opening character of an identifier.
func isIdentifierStart(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch >= utf8.RuneSelf && unicode.IsLetter(ch)
}

//---------------------------------------------------------------------------------------------------------------------

// oneOrTwoRuneToken scans a sequence of runes that could be one or two characters in length.
func (s *scanner) oneOrTwoRuneToken(
	oneRuneType TokenType,
	secondRune rune,
	twoRuneType TokenType,
) Token {

	if s.runeAhead1 == secondRune {
		s.advance()
		return s.token(twoRuneType)
	}

	return s.token(oneRuneType)

}

//---------------------------------------------------------------------------------------------------------------------

// oneToThreeRuneToken scans a sequence of runes that could be one, two, or three characters in length.
func (s *scanner) oneToThreeRuneToken(
	oneRuneType TokenType,
	secondRune rune,
	twoRuneType TokenType,
	thirdRune rune,
	threeRuneType TokenType,
) Token {

	if s.runeAhead1 == secondRune {

		s.advance()

		if s.runeAhead1 == thirdRune {
			s.advance()
			return s.token(threeRuneType)
		}

		return s.token(twoRuneType)

	}

	return s.token(oneRuneType)

}

//---------------------------------------------------------------------------------------------------------------------

// scanAfterEquals scans one of: '=', '==', '===', '=~'.
func (s *scanner) scanAfterEquals() Token {

	if s.runeAhead1 == '=' {

		s.advance()

		if s.runeAhead1 == '=' {
			s.advance()
			return s.token(TokenTypeEqualsEqualsEquals)
		}

		return s.token(TokenTypeEqualsEquals)

	}

	if s.runeAhead1 == '~' {
		s.advance()
		return s.token(TokenTypeMatches)
	}

	return s.token(TokenTypeEquals)

}

//---------------------------------------------------------------------------------------------------------------------

// scanAfterSlash scans either just the slash or else a comment extending to the end of the line.
func (s *scanner) scanAfterSlash() Token {

	if s.runeAhead1 == '/' {
		s.advance()
		return s.scanDocumentation()
	}

	return s.token(TokenTypeSlash)

}

//---------------------------------------------------------------------------------------------------------------------

// scanBackTickedString consumes a multiline back-ticked string.
func (s *scanner) scanBackTickedString() Token {

	mark := s.markedPos

	for {

		// Consume to the end of the line.
		for s.runeAhead1 != '\n' && s.runeAhead1 != 0 {
			s.advance()
		}

		// Quit if hit the end of input.
		if s.runeAhead1 == 0 {
			break
		}

		s.advance()

		// Ignore whitespace
		for s.runeAhead1 != '\n' && unicode.IsSpace(s.runeAhead1) {
			s.advance()
		}

		// Quit after seeing something other than another back-ticked string on the subsequent line.
		if s.runeAhead1 != '`' {
			break
		}

		// Mark the start of the next line and consume the back tick
		s.markedPos = s.currentPos
		s.advance()

	}

	return Token{
		SourceOffset: uint32(mark),
		SourceLength: uint16(s.currentPos - mark),
		TokenType:    TokenTypeBackTickedString,
	}

}

//---------------------------------------------------------------------------------------------------------------------

// scanDocumentation consumes a multiline comment.
func (s *scanner) scanDocumentation() Token {

	mark := s.markedPos

	for {

		// Consume to the end of the line.
		for s.runeAhead1 != '\n' && s.runeAhead1 != 0 {
			s.advance()
		}

		// Quit if hit the end of input.
		if s.runeAhead1 == 0 {
			break
		}

		s.advance()

		// Ignore whitespace
		for s.runeAhead1 != '\n' && unicode.IsSpace(s.runeAhead1) {
			s.advance()
		}

		// Quit after seeing something other than another back-ticked string on the subsequent line.
		if s.runeAhead1 != '/' || s.sourceCode[s.currentPos+1] != '/' {
			break
		}

		// Mark the start of the next line and consume the "//"
		s.markedPos = s.currentPos
		s.advance()
		s.advance()

	}

	return Token{
		SourceOffset: uint32(mark),
		SourceLength: uint16(s.currentPos - mark),
		TokenType:    TokenTypeDocumentation,
	}

}

//---------------------------------------------------------------------------------------------------------------------

// scanDoubleQuotedString scans the remainder of a string literal after the initial double quote character has been consumed.
func (s *scanner) scanDoubleQuotedString() Token {

	for {
		switch s.runeAhead1 {
		case '"':
			s.advance()
			return s.token(TokenTypeDoubleQuotedString)
		case '\\':
			s.advance()
			// TODO: handle escape sequences properly
			s.advance()
		case '\n':
			return s.token(TokenTypeUnclosedDoubleQuotedString)
		default:
			s.advance()
		}
	}

}

//---------------------------------------------------------------------------------------------------------------------

// scanIdentifierOrKeyword scans the remainder of an identifier after the opening letter has been consumed.
func (s *scanner) scanIdentifierOrKeyword() Token {

	for isIdentifierPart(s.runeAhead1) {
		s.advance()
	}

	text := s.sourceCode[s.markedPos:s.currentPos]

	tokenType, isKeyword := keywords[text]
	if isKeyword {
		return Token{
			SourceOffset: uint32(s.markedPos),
			SourceLength: uint16(s.currentPos - s.markedPos),
			TokenType:    tokenType,
		}
	}

	return Token{
		SourceOffset: uint32(s.markedPos),
		SourceLength: uint16(s.currentPos - s.markedPos),
		TokenType:    TokenTypeIdentifier,
	}

}

//---------------------------------------------------------------------------------------------------------------------

// scanNumber scans a numeric literal after the opening digit has been consumed.
func (s *scanner) scanNumber() Token {

	for isDigit(s.runeAhead1) {
		s.advance()
	}

	if s.runeAhead1 == '.' && isDigit(s.runeAhead2) {
		s.advance()
		return s.scanNumberFloatingPoint()
	}

	return s.token(TokenTypeIntegerLiteral)

}

//---------------------------------------------------------------------------------------------------------------------

// scanNumberFloatingPoint scans a floating point literal after the decimal point has been consumed.
func (s *scanner) scanNumberFloatingPoint() Token {

	for isDigit(s.runeAhead1) {
		s.advance()
	}

	// TODO: exponents

	return s.token(TokenTypeFloatingPointLiteral)

}

//---------------------------------------------------------------------------------------------------------------------

// scanSingleQuotedString scans the remainder of a string literal after the initial single quote character has been consumed.
func (s *scanner) scanSingleQuotedString() Token {

	for {
		switch s.runeAhead1 {
		case '\'':
			s.advance()
			return s.token(TokenTypeSingleQuotedString)
		case '\\':
			s.advance()
			// TODO: handle escape sequences properly
			s.advance()
		case '\n':
			return s.token(TokenTypeUnclosedSingleQuotedString)
		default:
			s.advance()
		}
	}

}

//---------------------------------------------------------------------------------------------------------------------

// Function token builds a new token of given type with text from the marked position to the current position.
func (s *scanner) token(tokenType TokenType) Token {
	return Token{
		SourceOffset: uint32(s.markedPos),
		SourceLength: uint16(s.currentPos - s.markedPos),
		TokenType:    tokenType,
	}
}

//=====================================================================================================================

var keywords = map[string]TokenType{
	TokenTypeAnd.String():   TokenTypeAnd,
	TokenTypeAs.String():    TokenTypeAs,
	TokenTypeFalse.String(): TokenTypeFalse,
	TokenTypeIs.String():    TokenTypeIs,
	TokenTypeIn.String():    TokenTypeIn,
	TokenTypeNot.String():   TokenTypeNot,
	TokenTypeOr.String():    TokenTypeOr,
	TokenTypeTrue.String():  TokenTypeTrue,
	TokenTypeWhen.String():  TokenTypeWhen,
	TokenTypeWhere.String(): TokenTypeWhere,
}

//=====================================================================================================================
