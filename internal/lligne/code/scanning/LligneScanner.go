//
// # Scanner for Lligne tokens.
//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package scanning

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

//=====================================================================================================================

// ILligneScanner represents the ability to read a sequence of tokens.
type ILligneScanner interface {
	ILligneTokenOriginTracker
	ReadToken() LligneToken
}

//---------------------------------------------------------------------------------------------------------------------

// NewLligneScanner allocates a new scanner for given sourceCode from the given fileName.
func NewLligneScanner(sourceCode string) ILligneScanner {

	// Start out the scan position.
	s := &lligneScanner{
		sourceCode:         sourceCode,
		currentPos:         0,
		markedPos:          0,
		tokenOriginTracker: NewLligneTokenOriginTracker("todo.lligne"),
	}

	// Read the first rune.
	if len(s.sourceCode) == 0 {
		s.nextRune = 0
		s.nextRuneWidth = 0
	} else {
		s.nextRune, s.nextRuneWidth = utf8.DecodeRuneInString(s.sourceCode[s.currentPos:])
	}

	return s

}

//=====================================================================================================================

// LligneScanner converts a string of Lligne source code into tokens.
type lligneScanner struct {
	sourceCode         string
	markedPos          int
	currentPos         int
	nextRune           rune
	nextRuneWidth      int
	tokenOriginTracker LligneTokenOriginTracker
}

//---------------------------------------------------------------------------------------------------------------------

// GetOrigin determines a token origin.
func (s *lligneScanner) GetOrigin(sourcePos int) LligneOrigin {
	return s.tokenOriginTracker.GetOrigin(sourcePos)
}

//---------------------------------------------------------------------------------------------------------------------

// ReadToken returns the next token from the scanner.
func (s *lligneScanner) ReadToken() LligneToken {

	// Ignore whitespace
	for unicode.IsSpace(s.nextRune) {
		s.advance()
	}

	// Mark the start of the token
	s.markedPos = s.currentPos

	// Consume the next character.
	ch := s.nextRune
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
//---------------------------------------------------------------------------------------------------------------------

// advance consumes one rune and stages the next one in the scanner.
func (s *lligneScanner) advance() {

	if s.nextRune == '\n' {
		s.tokenOriginTracker.AppendNewLinePosition(s.currentPos)
	}
	s.currentPos += s.nextRuneWidth

	if s.currentPos >= len(s.sourceCode) {
		s.nextRune = 0
		s.nextRuneWidth = 0
	} else {
		s.nextRune, s.nextRuneWidth = utf8.DecodeRuneInString(s.sourceCode[s.currentPos:])
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
func (s *lligneScanner) oneOrTwoRuneToken(
	oneRuneType LligneTokenType,
	secondRune rune,
	twoRuneType LligneTokenType,
) LligneToken {

	if s.nextRune == secondRune {
		s.advance()
		return s.token(twoRuneType)
	}

	return s.token(oneRuneType)

}

//---------------------------------------------------------------------------------------------------------------------

// oneToThreeRuneToken scans a sequence of runes that could be one, two, or three characters in length.
func (s *lligneScanner) oneToThreeRuneToken(
	oneRuneType LligneTokenType,
	secondRune rune,
	twoRuneType LligneTokenType,
	thirdRune rune,
	threeRuneType LligneTokenType,
) LligneToken {

	if s.nextRune == secondRune {

		s.advance()

		if s.nextRune == thirdRune {
			s.advance()
			return s.token(threeRuneType)
		}

		return s.token(twoRuneType)

	}

	return s.token(oneRuneType)

}

//---------------------------------------------------------------------------------------------------------------------

// scanAfterEquals scans one of: '=', '==', '===', '=~'.
func (s *lligneScanner) scanAfterEquals() LligneToken {

	if s.nextRune == '=' {

		s.advance()

		if s.nextRune == '=' {
			s.advance()
			return s.token(TokenTypeEqualsEqualsEquals)
		}

		return s.token(TokenTypeEqualsEquals)

	}

	if s.nextRune == '~' {
		s.advance()
		return s.token(TokenTypeMatches)
	}

	return s.token(TokenTypeEquals)

}

//---------------------------------------------------------------------------------------------------------------------

// scanAfterSlash scans either just the slash or else a comment extending to the end of the line.
func (s *lligneScanner) scanAfterSlash() LligneToken {

	if s.nextRune == '/' {
		s.advance()
		return s.scanDocumentation()
	}

	return s.token(TokenTypeSlash)

}

//---------------------------------------------------------------------------------------------------------------------

// scanBackTickedString consumes a multiline back-ticked string.
func (s *lligneScanner) scanBackTickedString() LligneToken {

	text := strings.Builder{}
	mark := s.markedPos

	for {

		// Consume to the end of the line.
		for s.nextRune != '\n' && s.nextRune != 0 {
			s.advance()
		}

		// Accumulate this line and always a new line character.
		text.WriteString(s.sourceCode[s.markedPos:s.currentPos])
		text.WriteRune('\n')

		// Quit if hit the end of input.
		if s.nextRune == 0 {
			break
		}

		s.advance()

		// Ignore whitespace
		for s.nextRune != '\n' && unicode.IsSpace(s.nextRune) {
			s.advance()
		}

		// Quit after seeing something other than another back-ticked string on the subsequent line.
		if s.nextRune != '`' {
			break
		}

		// Mark the start of the next line and consume the back tick
		s.markedPos = s.currentPos
		s.advance()

	}

	return LligneToken{TokenTypeBackTickedString, text.String(), mark}

}

//---------------------------------------------------------------------------------------------------------------------

// scanDocumentation consumes a multiline comment.
func (s *lligneScanner) scanDocumentation() LligneToken {

	text := strings.Builder{}
	mark := s.markedPos

	for {

		// Consume to the end of the line.
		for s.nextRune != '\n' && s.nextRune != 0 {
			s.advance()
		}

		// Accumulate this line and always a new line character.
		text.WriteString(s.sourceCode[s.markedPos:s.currentPos])
		text.WriteRune('\n')

		// Quit if hit the end of input.
		if s.nextRune == 0 {
			break
		}

		s.advance()

		// Ignore whitespace
		for s.nextRune != '\n' && unicode.IsSpace(s.nextRune) {
			s.advance()
		}

		// Quit after seeing something other than another back-ticked string on the subsequent line.
		if s.nextRune != '/' || s.sourceCode[s.currentPos+1] != '/' {
			break
		}

		// Mark the start of the next line and consume the "//"
		s.markedPos = s.currentPos
		s.advance()
		s.advance()

	}

	return LligneToken{TokenTypeDocumentation, text.String(), mark}

}

//---------------------------------------------------------------------------------------------------------------------

// scanDoubleQuotedString scans the remainder of a string literal after the initial double quote character has been consumed.
func (s *lligneScanner) scanDoubleQuotedString() LligneToken {

	for {
		switch s.nextRune {
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
func (s *lligneScanner) scanIdentifierOrKeyword() LligneToken {

	for isIdentifierPart(s.nextRune) {
		s.advance()
	}

	text := s.sourceCode[s.markedPos:s.currentPos]

	tokenType, isKeyword := keywords[text]
	if isKeyword {
		return LligneToken{tokenType, text, s.markedPos}
	}

	return LligneToken{TokenTypeIdentifier, text, s.markedPos}

}

//---------------------------------------------------------------------------------------------------------------------

// scanNumber scans a numeric literal after the opening digit has been consumed.
func (s *lligneScanner) scanNumber() LligneToken {

	for isDigit(s.nextRune) {
		s.advance()
	}

	// TODO: also floating point literals

	return s.token(TokenTypeIntegerLiteral)

}

//---------------------------------------------------------------------------------------------------------------------

// scanSingleQuotedString scans the remainder of a string literal after the initial single quote character has been consumed.
func (s *lligneScanner) scanSingleQuotedString() LligneToken {

	for {
		switch s.nextRune {
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
func (s *lligneScanner) token(tokenType LligneTokenType) LligneToken {
	return LligneToken{
		TokenType:      tokenType,
		Text:           s.sourceCode[s.markedPos:s.currentPos],
		SourceStartPos: s.markedPos,
	}
}

//=====================================================================================================================

var keywords = map[string]LligneTokenType{
	TokenTypeAnd.String():   TokenTypeAnd,
	TokenTypeAs.String():    TokenTypeAs,
	TokenTypeIs.String():    TokenTypeIs,
	TokenTypeIn.String():    TokenTypeIn,
	TokenTypeNot.String():   TokenTypeNot,
	TokenTypeOf.String():    TokenTypeOf,
	TokenTypeOr.String():    TokenTypeOr,
	TokenTypeTo.String():    TokenTypeTo,
	TokenTypeWhen.String():  TokenTypeWhen,
	TokenTypeWhere.String(): TokenTypeWhere,
}

//=====================================================================================================================
