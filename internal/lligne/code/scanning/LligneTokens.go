//
// # Data types related to Lligne token scanning.
//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package scanning

import (
	"strconv"
)

//---------------------------------------------------------------------------------------------------------------------

// LligneTokenType is an enumeration of Lligne token types.
type LligneTokenType int

const (
	TokenTypeEof LligneTokenType = iota
	TokenTypeAmpersand
	TokenTypeAmpersandAmpersand
	TokenTypeAnd
	TokenTypeAs
	TokenTypeAsterisk
	TokenTypeBackTickedString
	TokenTypeCharLiteral
	TokenTypeColon
	TokenTypeComma
	TokenTypeCompare
	TokenTypeCompareReversed
	TokenTypeDash
	TokenTypeDateLiteral
	TokenTypeDateTimeLiteral
	TokenTypeDocumentation
	TokenTypeDot
	TokenTypeDotDot
	TokenTypeDotDotDot
	TokenTypeEquals
	TokenTypeEqualsEquals
	TokenTypeEqualsEqualsEquals
	TokenTypeExclamationMark
	TokenTypeGreaterThan
	TokenTypeGreaterThanOrEquals
	TokenTypeIdentifier
	TokenTypeIntegerLiteral
	TokenTypeIn
	TokenTypeIpAddressV4
	TokenTypeIpAddressV6
	TokenTypeIs
	TokenTypeLeadingDocumentation
	TokenTypeLeftBrace
	TokenTypeLeftBracket
	TokenTypeLeftParenthesis
	TokenTypeLessThan
	TokenTypeLessThanOrEquals
	TokenTypeMatches
	TokenTypeMultilineString
	TokenTypeNot
	TokenTypeNotMatches
	TokenTypeOf
	TokenTypeOr
	TokenTypePlus
	TokenTypeQuestionMark
	TokenTypeQuestionMarkColon
	TokenTypeRightArrow
	TokenTypeRightBrace
	TokenTypeRightBracket
	TokenTypeRightParenthesis
	TokenTypeSemicolon
	TokenTypeSlash
	TokenTypeStringLiteral
	TokenTypeSynthDocument
	TokenTypeTo
	TokenTypeTrailingDocumentation
	TokenTypeUnclosedLiteral
	TokenTypeUnclosedString
	TokenTypeUnrecognizedChar
	TokenTypeUnrecognizedLiteral
	TokenTypeUuidLiteral
	TokenTypeVerticalBar
)

// ---------------------------------------------------------------------------------------------------------------------

// TextOfTokenType returns a string describing a Lligne token type.
func (tt LligneTokenType) String() string {

	switch tt {

	case TokenTypeEof:
		return "[end of file]"

	// Punctuation
	case TokenTypeAmpersand:
		return "&"
	case TokenTypeAmpersandAmpersand:
		return "&&"
	case TokenTypeAsterisk:
		return "*"
	case TokenTypeColon:
		return ":"
	case TokenTypeComma:
		return ","
	case TokenTypeCompare:
		return "<=>"
	case TokenTypeCompareReversed:
		return ">=<"
	case TokenTypeDash:
		return "-"
	case TokenTypeDot:
		return "."
	case TokenTypeDotDot:
		return ".."
	case TokenTypeDotDotDot:
		return "..."
	case TokenTypeEquals:
		return "="
	case TokenTypeEqualsEquals:
		return "=="
	case TokenTypeEqualsEqualsEquals:
		return "==="
	case TokenTypeExclamationMark:
		return "!"
	case TokenTypeGreaterThan:
		return "<"
	case TokenTypeGreaterThanOrEquals:
		return "<="
	case TokenTypeLeftBrace:
		return "{"
	case TokenTypeLeftBracket:
		return "["
	case TokenTypeLeftParenthesis:
		return "("
	case TokenTypeLessThan:
		return "<"
	case TokenTypeLessThanOrEquals:
		return "<="
	case TokenTypeMatches:
		return "=~"
	case TokenTypeNotMatches:
		return "!~"
	case TokenTypePlus:
		return "+"
	case TokenTypeQuestionMark:
		return "?"
	case TokenTypeQuestionMarkColon:
		return "?:"
	case TokenTypeRightArrow:
		return "->"
	case TokenTypeRightBrace:
		return "}"
	case TokenTypeRightBracket:
		return "]"
	case TokenTypeRightParenthesis:
		return ")"
	case TokenTypeSemicolon:
		return ";"
	case TokenTypeSlash:
		return "/"
	case TokenTypeVerticalBar:
		return "|"

	// Keywords
	case TokenTypeAnd:
		return "and"
	case TokenTypeAs:
		return "as"
	case TokenTypeIn:
		return "in"
	case TokenTypeIs:
		return "is"
	case TokenTypeNot:
		return "not"
	case TokenTypeOf:
		return "of"
	case TokenTypeOr:
		return "or"
	case TokenTypeTo:
		return "to"

	// Identifiers
	case TokenTypeIdentifier:
		return "[identifier]"

	// Numeric Literals
	case TokenTypeIntegerLiteral:
		return "[integer literal]"

	// String Literals
	case TokenTypeBackTickedString:
		return "[back-ticked string]"
	case TokenTypeMultilineString:
		return "[multiline string]"
	case TokenTypeStringLiteral:
		return "[string literal]"

	// Other Literals
	case TokenTypeCharLiteral:
		return "[character literal]"
	case TokenTypeDateLiteral:
		return "[date literal]"
	case TokenTypeDateTimeLiteral:
		return "[date-time literal]"
	case TokenTypeIpAddressV4:
		return "[IP address V4 literal]"
	case TokenTypeIpAddressV6:
		return "[IP address V6 literal]"
	case TokenTypeUuidLiteral:
		return "[UUID literal]"

	// Documentation
	case TokenTypeDocumentation:
		return "[documentation]"
	case TokenTypeLeadingDocumentation:
		return "[leading documentation]"
	case TokenTypeSynthDocument:
		return "[synthetic documentation operator]"
	case TokenTypeTrailingDocumentation:
		return "[trailing documentation]"

	// Errors
	case TokenTypeUnclosedLiteral:
		return "[error - literal extends past end of line]"
	case TokenTypeUnclosedString:
		return "[error - string extends past end of line]"
	case TokenTypeUnrecognizedChar:
		return "[error - unrecognized character]"
	case TokenTypeUnrecognizedLiteral:
		return "[error - unrecognized literal]"

	}

	panic("Unhandled token type: '" + strconv.Itoa(int(tt)) + "'.")
}

//---------------------------------------------------------------------------------------------------------------------

// LligneToken is an abstract token occurring at line [line] and column [column] (both 1-based) in its source file.
type LligneToken struct {
	TokenType LligneTokenType
	Text      string
	Origin    *LligneOrigin
}

//---------------------------------------------------------------------------------------------------------------------
