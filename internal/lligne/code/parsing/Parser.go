//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package parsing

import (
	"lligne-cli/internal/lligne/code/scanning"
	"strconv"
)

//=====================================================================================================================

func ParseExpression(sourceCode string, tokens []scanning.Token) (model IExpression) {
	parser := newParser(sourceCode, tokens)

	return parser.parseExprBindingPower(0)
}

//---------------------------------------------------------------------------------------------------------------------

// ParseParenthesizedItems parses a non-empty sequence of code expected to be the items within a record literal, e.g.
// the top level of a file.
func ParseParenthesizedItems(sourceCode string, tokens []scanning.Token) IExpression {
	parser := newParser(sourceCode, tokens)

	return parser.parseParenthesizedExpression(tokens[0], scanning.TokenTypeEof)
}

//=====================================================================================================================

func newParser(sourceCode string, tokens []scanning.Token) *lligneParser {
	return &lligneParser{
		sourceCode: sourceCode,
		tokens:     tokens,
	}
}

//=====================================================================================================================

type lligneParser struct {
	tokens     []scanning.Token
	index      int
	sourceCode string
}

//---------------------------------------------------------------------------------------------------------------------

func (p *lligneParser) parseExprBindingPower(minBindingPower int) IExpression {

	lhs := p.parseLeftHandSide()

	for {

		// Look ahead for an operator continuing the expression
		opToken := p.tokens[p.index]

		// Handle postfix operators ...
		pBindingPower := postfixBindingPowers[opToken.TokenType]

		if pBindingPower.Power != 0 {

			if pBindingPower.Power < minBindingPower {
				break
			}

			p.index += 1

			lhs = p.parsePostfixExpression(opToken, lhs)

			continue

		}

		// Handle infix operators ...
		bindingPower := infixBindingPowers[opToken.TokenType]

		if bindingPower.Left != 0 {

			if bindingPower.Left < minBindingPower {
				break
			}

			p.index += 1

			lhs = p.parseInfixOperation(opToken, bindingPower, lhs)

			continue

		}

		break

	}

	return lhs
}

//---------------------------------------------------------------------------------------------------------------------

// parseInfixOperation parses an infix expression after the left hand side and the operator token have been consumed
func (p *lligneParser) parseInfixOperation(
	opToken scanning.Token,
	bindingPower infixBindingPower,
	lhs IExpression,
) IExpression {
	rhs := p.parseExprBindingPower(bindingPower.Right)

	switch opToken.TokenType {

	case scanning.TokenTypeAmpersand:
		return &IntersectExpr{
			SourcePosition: lhs.GetSourcePosition().Thru(rhs.GetSourcePosition()),
			Lhs:            lhs,
			Rhs:            rhs,
		}

	case scanning.TokenTypeAmpersandAmpersand:
		return &IntersectLowPrecedenceExpr{
			SourcePosition: lhs.GetSourcePosition().Thru(rhs.GetSourcePosition()),
			Lhs:            lhs,
			Rhs:            rhs,
		}

	case scanning.TokenTypeAnd:
		return &LogicalAndExpr{
			SourcePosition: lhs.GetSourcePosition().Thru(rhs.GetSourcePosition()),
			Lhs:            lhs,
			Rhs:            rhs,
		}

	case scanning.TokenTypeAsterisk:
		return &MultiplicationExpr{
			SourcePosition: lhs.GetSourcePosition().Thru(rhs.GetSourcePosition()),
			Lhs:            lhs,
			Rhs:            rhs,
		}

	case scanning.TokenTypeColon:
		return &QualifyExpr{
			SourcePosition: lhs.GetSourcePosition().Thru(rhs.GetSourcePosition()),
			Lhs:            lhs,
			Rhs:            rhs,
		}

	case scanning.TokenTypeDash:
		return &SubtractionExpr{
			SourcePosition: lhs.GetSourcePosition().Thru(rhs.GetSourcePosition()),
			Lhs:            lhs,
			Rhs:            rhs,
		}

	case scanning.TokenTypeDot:
		return &FieldReferenceExpr{
			SourcePosition: lhs.GetSourcePosition().Thru(rhs.GetSourcePosition()),
			Parent:         lhs,
			Child:          rhs,
		}

	case scanning.TokenTypeDotDot:
		return &RangeExpr{
			SourcePosition: lhs.GetSourcePosition().Thru(rhs.GetSourcePosition()),
			First:          lhs,
			Last:           rhs,
		}

	case scanning.TokenTypeEquals:
		return &IntersectAssignValueExpr{
			SourcePosition: lhs.GetSourcePosition().Thru(rhs.GetSourcePosition()),
			Lhs:            lhs,
			Rhs:            rhs,
		}

	case scanning.TokenTypeEqualsEquals:
		return &EqualsExpr{
			SourcePosition: lhs.GetSourcePosition().Thru(rhs.GetSourcePosition()),
			Lhs:            lhs,
			Rhs:            rhs,
		}

	case scanning.TokenTypeGreaterThan:
		return &GreaterThanExpr{
			SourcePosition: lhs.GetSourcePosition().Thru(rhs.GetSourcePosition()),
			Lhs:            lhs,
			Rhs:            rhs,
		}

	case scanning.TokenTypeGreaterThanOrEquals:
		return &GreaterThanOrEqualsExpr{
			SourcePosition: lhs.GetSourcePosition().Thru(rhs.GetSourcePosition()),
			Lhs:            lhs,
			Rhs:            rhs,
		}

	case scanning.TokenTypeIn:
		return &InExpr{
			SourcePosition: lhs.GetSourcePosition().Thru(rhs.GetSourcePosition()),
			Lhs:            lhs,
			Rhs:            rhs,
		}

	case scanning.TokenTypeIs:
		return &IsExpr{
			SourcePosition: lhs.GetSourcePosition().Thru(rhs.GetSourcePosition()),
			Lhs:            lhs,
			Rhs:            rhs,
		}

	case scanning.TokenTypeLessThan:
		return &LessThanExpr{
			SourcePosition: lhs.GetSourcePosition().Thru(rhs.GetSourcePosition()),
			Lhs:            lhs,
			Rhs:            rhs,
		}

	case scanning.TokenTypeLessThanOrEquals:
		return &LessThanOrEqualsExpr{
			SourcePosition: lhs.GetSourcePosition().Thru(rhs.GetSourcePosition()),
			Lhs:            lhs,
			Rhs:            rhs,
		}

	case scanning.TokenTypeMatches:
		return &MatchExpr{
			SourcePosition: lhs.GetSourcePosition().Thru(rhs.GetSourcePosition()),
			Lhs:            lhs,
			Rhs:            rhs,
		}

	case scanning.TokenTypeNotMatches:
		return &NotMatchExpr{
			SourcePosition: lhs.GetSourcePosition().Thru(rhs.GetSourcePosition()),
			Lhs:            lhs,
			Rhs:            rhs,
		}

	case scanning.TokenTypeOr:
		return &LogicalOrExpr{
			SourcePosition: lhs.GetSourcePosition().Thru(rhs.GetSourcePosition()),
			Lhs:            lhs,
			Rhs:            rhs,
		}

	case scanning.TokenTypePlus:
		return &AdditionExpr{
			SourcePosition: lhs.GetSourcePosition().Thru(rhs.GetSourcePosition()),
			Lhs:            lhs,
			Rhs:            rhs,
		}

	case scanning.TokenTypeQuestionMarkColon:
		return &IntersectDefaultValueExpr{
			SourcePosition: lhs.GetSourcePosition().Thru(rhs.GetSourcePosition()),
			Lhs:            lhs,
			Rhs:            rhs,
		}

	case scanning.TokenTypeRightArrow:
		return &FunctionArrowExpr{
			SourcePosition: lhs.GetSourcePosition().Thru(rhs.GetSourcePosition()),
			Argument:       lhs,
			Result:         rhs,
		}

	case scanning.TokenTypeSlash:
		return &DivisionExpr{
			SourcePosition: lhs.GetSourcePosition().Thru(rhs.GetSourcePosition()),
			Lhs:            lhs,
			Rhs:            rhs,
		}

	case scanning.TokenTypeSynthDocument:
		return &DocumentExpr{
			SourcePosition: lhs.GetSourcePosition().Thru(rhs.GetSourcePosition()),
			Lhs:            lhs,
			Rhs:            rhs,
		}

	case scanning.TokenTypeVerticalBar:
		return &UnionExpr{
			SourcePosition: lhs.GetSourcePosition().Thru(rhs.GetSourcePosition()),
			Lhs:            lhs,
			Rhs:            rhs,
		}

	case scanning.TokenTypeWhen:
		return &WhenExpr{
			SourcePosition: lhs.GetSourcePosition().Thru(rhs.GetSourcePosition()),
			Lhs:            lhs,
			Rhs:            rhs,
		}

	case scanning.TokenTypeWhere:
		return &WhereExpr{
			SourcePosition: lhs.GetSourcePosition().Thru(rhs.GetSourcePosition()),
			Lhs:            lhs,
			Rhs:            rhs,
		}

	}

	panic(opToken.TokenType)

}

//---------------------------------------------------------------------------------------------------------------------

func (p *lligneParser) parseLeftHandSide() IExpression {

	token := p.tokens[p.index]
	p.index += 1

	switch token.TokenType {

	case scanning.TokenTypeBackTickedString:
		return &StringLiteralExpr{
			SourcePosition: NewSourcePos(token),
			Delimiters:     StringDelimitersBackTicksMultiline,
		}

	case scanning.TokenTypeDash:
		return p.parseNegationOperationExpression(token)

	case scanning.TokenTypeDoubleQuotedString:
		return &StringLiteralExpr{
			SourcePosition: NewSourcePos(token),
			Delimiters:     StringDelimitersDoubleQuotes,
		}

	case scanning.TokenTypeFalse:
		return &BooleanLiteralExpr{
			SourcePosition: NewSourcePos(token),
			Value:          false,
		}

	case scanning.TokenTypeFloatingPointLiteral:
		return &FloatingPointLiteralExpr{
			SourcePosition: NewSourcePos(token),
		}

	case scanning.TokenTypeIdentifier:
		return &IdentifierExpr{
			SourcePosition: NewSourcePos(token),
		}

	case scanning.TokenTypeIntegerLiteral:
		return &IntegerLiteralExpr{
			SourcePosition: NewSourcePos(token),
		}

	case scanning.TokenTypeLeadingDocumentation:
		return &LeadingDocumentationExpr{
			SourcePosition: NewSourcePos(token),
		}

	case scanning.TokenTypeLeftBrace:
		return p.parseParenthesizedExpression(token, scanning.TokenTypeRightBrace)

	case scanning.TokenTypeLeftBracket:
		return p.parseSequenceLiteral(token)

	case scanning.TokenTypeLeftParenthesis:
		return p.parseParenthesizedExpression(token, scanning.TokenTypeRightParenthesis)

	case scanning.TokenTypeNot:
		return p.parseLogicalNotOperationExpression(token)

	case scanning.TokenTypeSingleQuotedString:
		return &StringLiteralExpr{
			SourcePosition: NewSourcePos(token),
			Delimiters:     StringDelimitersSingleQuotes,
		}

	case scanning.TokenTypeTrailingDocumentation:
		return &TrailingDocumentationExpr{
			SourcePosition: NewSourcePos(token),
		}

	case scanning.TokenTypeTrue:
		return &BooleanLiteralExpr{
			SourcePosition: NewSourcePos(token),
			Value:          true,
		}

		//	default:
		//	this.expectedType(
		//	LlaceTokenType.CHAR_LITERAL,
		//	LlaceTokenType.DASH,
		//	LlaceTokenType.IDENTIFIER,
		//	LlaceTokenType.INTEGER_LITERAL,
		//	LlaceTokenType.STRING_LITERAL
		//	)

	}

	panic("Unfinished parsing code: '" + strconv.Itoa(int(token.TokenType)) + "'.")

}

//---------------------------------------------------------------------------------------------------------------------

func (p *lligneParser) parseLogicalNotOperationExpression(
	token scanning.Token,
) IExpression {
	rightBindingPower := prefixBindingPowers[token.TokenType].Power
	rhs := p.parseExprBindingPower(rightBindingPower)
	return &LogicalNotOperationExpr{
		SourcePosition: NewSourcePos(token),
		Operand:        rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (p *lligneParser) parseNegationOperationExpression(
	token scanning.Token,
) IExpression {
	rightBindingPower := prefixBindingPowers[token.TokenType].Power
	rhs := p.parseExprBindingPower(rightBindingPower)
	return &NegationOperationExpr{
		SourcePosition: NewSourcePos(token).Thru(rhs.GetSourcePosition()),
		Operand:        rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (p *lligneParser) parseParenthesizedExpression(
	token scanning.Token,
	endingTokenType scanning.TokenType,
) IExpression {

	var items []IExpression

	for p.tokens[p.index].TokenType != endingTokenType {
		// Parse one expression.
		items = append(items, p.parseExprBindingPower(0))

		if p.tokens[p.index].TokenType != scanning.TokenTypeComma {
			break
		}
		p.index += 1
	}

	if p.tokens[p.index].TokenType != endingTokenType {
		panic("Expected " + endingTokenType.String())
	}
	endSourcePos := NewSourcePos(p.tokens[p.index])
	p.index += 1

	var delimiters ParenExprDelimiters
	switch endingTokenType {
	case scanning.TokenTypeEof:
		delimiters = ParenExprDelimitersWholeFile
	case scanning.TokenTypeRightBrace:
		delimiters = ParenExprDelimitersBraces
	case scanning.TokenTypeRightParenthesis:
		delimiters = ParenExprDelimitersParentheses
	}

	return &ParenthesizedExpr{
		SourcePosition: NewSourcePos(token).Thru(endSourcePos),
		Delimiters:     delimiters,
		Items:          items,
	}

}

//---------------------------------------------------------------------------------------------------------------------

func (p *lligneParser) parsePostfixExpression(opToken scanning.Token, lhs IExpression) IExpression {

	switch opToken.TokenType {

	case scanning.TokenTypeLeftParenthesis:
		args := p.parseParenthesizedExpression(opToken, scanning.TokenTypeRightParenthesis)
		return &FunctionCallExpr{
			SourcePosition:    lhs.GetSourcePosition().Thru(args.GetSourcePosition()),
			FunctionReference: lhs,
			Argument:          args,
		}

	case scanning.TokenTypeQuestionMark:
		return &OptionalExpr{
			SourcePosition: lhs.GetSourcePosition(),
			Operand:        lhs,
		}

	}

	panic("Unfinished parsing code: '" + strconv.Itoa(int(opToken.TokenType)) + "'.")

}

//---------------------------------------------------------------------------------------------------------------------

func (p *lligneParser) parseSequenceLiteral(token scanning.Token) IExpression {

	startSourcePos := NewSourcePos(token)
	var items []IExpression

	if p.tokens[p.index].TokenType == scanning.TokenTypeRightBracket {
		endSourcePos := NewSourcePos(p.tokens[p.index])
		p.index += 1
		return &SequenceLiteralExpr{
			SourcePosition: startSourcePos.Thru(endSourcePos),
			Elements:       items,
		}
	}

	for p.tokens[p.index].TokenType != scanning.TokenTypeRightBracket {
		// Parse one expression.
		items = append(items, p.parseExprBindingPower(0))

		if p.tokens[p.index].TokenType != scanning.TokenTypeComma {
			break
		}
		p.index += 1
	}

	if p.tokens[p.index].TokenType != scanning.TokenTypeRightBracket {
		panic("Expected " + scanning.TokenTypeRightBracket.String())
	}
	endSourcePos := NewSourcePos(p.tokens[p.index])
	p.index += 1

	return &SequenceLiteralExpr{
		SourcePosition: startSourcePos.Thru(endSourcePos),
		Elements:       items,
	}

}

//=====================================================================================================================

type infixBindingPower struct {
	Left  int
	Right int
}

//=====================================================================================================================

type prefixBindingPower struct {
	Power int
}

//=====================================================================================================================

type postfixBindingPower struct {
	Power int
}

//=====================================================================================================================

var prefixBindingPowers = make(map[scanning.TokenType]prefixBindingPower)

var infixBindingPowers = make(map[scanning.TokenType]infixBindingPower)

var postfixBindingPowers = make(map[scanning.TokenType]postfixBindingPower)

func init() {

	level := 1

	infixBindingPowers[scanning.TokenTypeColon] = infixBindingPower{level, level + 1}
	infixBindingPowers[scanning.TokenTypeEquals] = infixBindingPower{level, level + 1}
	infixBindingPowers[scanning.TokenTypeQuestionMarkColon] = infixBindingPower{level, level + 1}

	level += 2

	infixBindingPowers[scanning.TokenTypeAmpersandAmpersand] = infixBindingPower{level, level + 1}

	level += 2

	infixBindingPowers[scanning.TokenTypeVerticalBar] = infixBindingPower{level, level + 1}

	level += 2

	infixBindingPowers[scanning.TokenTypeAmpersand] = infixBindingPower{level, level + 1}

	level += 2

	infixBindingPowers[scanning.TokenTypeWhen] = infixBindingPower{level, level + 1}
	infixBindingPowers[scanning.TokenTypeWhere] = infixBindingPower{level, level + 1}

	level += 2

	infixBindingPowers[scanning.TokenTypeSynthDocument] = infixBindingPower{level, level + 1}

	level += 2

	infixBindingPowers[scanning.TokenTypeOr] = infixBindingPower{level, level + 1}

	level += 2

	infixBindingPowers[scanning.TokenTypeAnd] = infixBindingPower{level, level + 1}

	level += 2

	prefixBindingPowers[scanning.TokenTypeNot] = prefixBindingPower{level}

	level += 2

	infixBindingPowers[scanning.TokenTypeEqualsEquals] = infixBindingPower{level, level + 1}
	infixBindingPowers[scanning.TokenTypeGreaterThan] = infixBindingPower{level, level + 1}
	infixBindingPowers[scanning.TokenTypeGreaterThanOrEquals] = infixBindingPower{level, level + 1}
	infixBindingPowers[scanning.TokenTypeLessThan] = infixBindingPower{level, level + 1}
	infixBindingPowers[scanning.TokenTypeLessThanOrEquals] = infixBindingPower{level, level + 1}

	level += 2

	infixBindingPowers[scanning.TokenTypeIn] = infixBindingPower{level, level + 1}
	infixBindingPowers[scanning.TokenTypeIs] = infixBindingPower{level, level + 1}
	infixBindingPowers[scanning.TokenTypeMatches] = infixBindingPower{level, level + 1}
	infixBindingPowers[scanning.TokenTypeNotMatches] = infixBindingPower{level, level + 1}

	level += 2

	infixBindingPowers[scanning.TokenTypeDotDot] = infixBindingPower{level, level + 1}

	level += 2

	infixBindingPowers[scanning.TokenTypeDash] = infixBindingPower{level, level + 1}
	infixBindingPowers[scanning.TokenTypePlus] = infixBindingPower{level, level + 1}

	level += 2

	infixBindingPowers[scanning.TokenTypeAsterisk] = infixBindingPower{level, level + 1}
	infixBindingPowers[scanning.TokenTypeSlash] = infixBindingPower{level, level + 1}

	level += 2

	prefixBindingPowers[scanning.TokenTypeDash] = prefixBindingPower{level}

	level += 2

	infixBindingPowers[scanning.TokenTypeRightArrow] = infixBindingPower{level, level + 1}

	level += 2

	infixBindingPowers[scanning.TokenTypeDot] = infixBindingPower{level, level + 1}

	level += 2

	postfixBindingPowers[scanning.TokenTypeLeftParenthesis] = postfixBindingPower{level}
	postfixBindingPowers[scanning.TokenTypeLeftBracket] = postfixBindingPower{level}
	postfixBindingPowers[scanning.TokenTypeQuestionMark] = postfixBindingPower{level}

}

//=====================================================================================================================
