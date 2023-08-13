//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package parsing

import (
	"lligne-cli/internal/lligne/code/model"
	"lligne-cli/internal/lligne/code/scanning"
	"strconv"
)

//=====================================================================================================================

func ParseExpression(sourceCode string, tokens []scanning.Token) (model model.IExpression) {
	parser := newParser(sourceCode, tokens)

	return parser.parseExprBindingPower(0)
}

//---------------------------------------------------------------------------------------------------------------------

// ParseParenthesizedItems parses a non-empty sequence of code expected to be the items within a record literal, e.g.
// the top level of a file.
func ParseParenthesizedItems(sourceCode string, tokens []scanning.Token) model.IExpression {
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

func (p *lligneParser) parseExprBindingPower(minBindingPower int) model.IExpression {

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
	lhs model.IExpression,
) model.IExpression {
	rhs := p.parseExprBindingPower(bindingPower.Right)

	switch opToken.TokenType {

	case scanning.TokenTypeAmpersand:
		return &model.IntersectExpr{
			SourcePosition: model.NewSourcePos(opToken),
			Lhs:            lhs,
			Rhs:            rhs,
		}

	case scanning.TokenTypeAmpersandAmpersand:
		return &model.IntersectLowPrecedenceExpr{
			SourcePosition: model.NewSourcePos(opToken),
			Lhs:            lhs,
			Rhs:            rhs,
		}

	case scanning.TokenTypeAnd:
		return &model.LogicalAndExpr{
			SourcePosition: model.NewSourcePos(opToken),
			Lhs:            lhs,
			Rhs:            rhs,
		}

	case scanning.TokenTypeAsterisk:
		return &model.MultiplicationExpr{
			SourcePosition: model.NewSourcePos(opToken),
			Lhs:            lhs,
			Rhs:            rhs,
		}

	case scanning.TokenTypeColon:
		return &model.QualifyExpr{
			SourcePosition: model.NewSourcePos(opToken),
			Lhs:            lhs,
			Rhs:            rhs,
		}

	case scanning.TokenTypeDash:
		return &model.SubtractionExpr{
			SourcePosition: model.NewSourcePos(opToken),
			Lhs:            lhs,
			Rhs:            rhs,
		}

	case scanning.TokenTypeDot:
		return &model.FieldReferenceExpr{
			SourcePosition: model.NewSourcePos(opToken),
			Parent:         lhs,
			Child:          rhs,
		}

	case scanning.TokenTypeDotDot:
		return &model.RangeExpr{
			SourcePosition: model.NewSourcePos(opToken),
			First:          lhs,
			Last:           rhs,
		}

	case scanning.TokenTypeEquals:
		return &model.IntersectAssignValueExpr{
			SourcePosition: model.NewSourcePos(opToken),
			Lhs:            lhs,
			Rhs:            rhs,
		}

	case scanning.TokenTypeEqualsEquals:
		return &model.EqualsExpr{
			SourcePosition: model.NewSourcePos(opToken),
			Lhs:            lhs,
			Rhs:            rhs,
		}

	case scanning.TokenTypeGreaterThan:
		return &model.GreaterThanExpr{
			SourcePosition: model.NewSourcePos(opToken),
			Lhs:            lhs,
			Rhs:            rhs,
		}

	case scanning.TokenTypeGreaterThanOrEquals:
		return &model.GreaterThanOrEqualsExpr{
			SourcePosition: model.NewSourcePos(opToken),
			Lhs:            lhs,
			Rhs:            rhs,
		}

	case scanning.TokenTypeIn:
		return &model.InExpr{
			SourcePosition: model.NewSourcePos(opToken),
			Lhs:            lhs,
			Rhs:            rhs,
		}

	case scanning.TokenTypeIs:
		return &model.IsExpr{
			SourcePosition: model.NewSourcePos(opToken),
			Lhs:            lhs,
			Rhs:            rhs,
		}

	case scanning.TokenTypeLessThan:
		return &model.LessThanExpr{
			SourcePosition: model.NewSourcePos(opToken),
			Lhs:            lhs,
			Rhs:            rhs,
		}

	case scanning.TokenTypeLessThanOrEquals:
		return &model.LessThanOrEqualsExpr{
			SourcePosition: model.NewSourcePos(opToken),
			Lhs:            lhs,
			Rhs:            rhs,
		}

	case scanning.TokenTypeMatches:
		return &model.MatchExpr{
			SourcePosition: model.NewSourcePos(opToken),
			Lhs:            lhs,
			Rhs:            rhs,
		}

	case scanning.TokenTypeNotMatches:
		return &model.NotMatchExpr{
			SourcePosition: model.NewSourcePos(opToken),
			Lhs:            lhs,
			Rhs:            rhs,
		}

	case scanning.TokenTypeOr:
		return &model.LogicalOrExpr{
			SourcePosition: model.NewSourcePos(opToken),
			Lhs:            lhs,
			Rhs:            rhs,
		}

	case scanning.TokenTypePlus:
		return &model.AdditionExpr{
			SourcePosition: model.NewSourcePos(opToken),
			Lhs:            lhs,
			Rhs:            rhs,
		}

	case scanning.TokenTypeQuestionMarkColon:
		return &model.IntersectDefaultValueExpr{
			SourcePosition: model.NewSourcePos(opToken),
			Lhs:            lhs,
			Rhs:            rhs,
		}

	case scanning.TokenTypeRightArrow:
		return &model.FunctionArrowExpr{
			SourcePosition: model.NewSourcePos(opToken),
			Argument:       lhs,
			Result:         rhs,
		}

	case scanning.TokenTypeSlash:
		return &model.DivisionExpr{
			SourcePosition: model.NewSourcePos(opToken),
			Lhs:            lhs,
			Rhs:            rhs,
		}

	case scanning.TokenTypeSynthDocument:
		return &model.DocumentExpr{
			SourcePosition: model.NewSourcePos(opToken),
			Lhs:            lhs,
			Rhs:            rhs,
		}

	case scanning.TokenTypeVerticalBar:
		return &model.UnionExpr{
			SourcePosition: model.NewSourcePos(opToken),
			Lhs:            lhs,
			Rhs:            rhs,
		}

	case scanning.TokenTypeWhen:
		return &model.WhenExpr{
			SourcePosition: model.NewSourcePos(opToken),
			Lhs:            lhs,
			Rhs:            rhs,
		}

	case scanning.TokenTypeWhere:
		return &model.WhereExpr{
			SourcePosition: model.NewSourcePos(opToken),
			Lhs:            lhs,
			Rhs:            rhs,
		}

	}

	panic(opToken.TokenType)

}

//---------------------------------------------------------------------------------------------------------------------

func (p *lligneParser) parseLeftHandSide() model.IExpression {

	token := p.tokens[p.index]
	p.index += 1

	switch token.TokenType {

	case scanning.TokenTypeBackTickedString:
		return &model.MultilineStringLiteralExpr{
			SourcePosition: model.NewSourcePos(token),
			Text:           p.sourceCode[token.SourceOffset : token.SourceOffset+uint32(token.SourceLength)],
		}

	case scanning.TokenTypeDash:
		return p.parseNegationOperationExpression(token)

	case scanning.TokenTypeDoubleQuotedString:
		return &model.StringLiteralExpr{
			SourcePosition: model.NewSourcePos(token),
			Text:           p.sourceCode[token.SourceOffset : token.SourceOffset+uint32(token.SourceLength)],
		}

	case scanning.TokenTypeFalse:
		return &model.BooleanLiteralExpr{
			SourcePosition: model.NewSourcePos(token),
			Value:          false,
		}

	case scanning.TokenTypeFloatingPointLiteral:
		return &model.FloatingPointLiteralExpr{
			SourcePosition: model.NewSourcePos(token),
			Text:           p.sourceCode[token.SourceOffset : token.SourceOffset+uint32(token.SourceLength)],
		}

	case scanning.TokenTypeIdentifier:
		return &model.IdentifierExpr{
			SourcePosition: model.NewSourcePos(token),
			Name:           p.sourceCode[token.SourceOffset : token.SourceOffset+uint32(token.SourceLength)],
		}

	case scanning.TokenTypeIntegerLiteral:
		return &model.IntegerLiteralExpr{
			SourcePosition: model.NewSourcePos(token),
			Text:           p.sourceCode[token.SourceOffset : token.SourceOffset+uint32(token.SourceLength)],
		}

	case scanning.TokenTypeLeadingDocumentation:
		return &model.LeadingDocumentationExpr{
			SourcePosition: model.NewSourcePos(token),
			Text:           p.sourceCode[token.SourceOffset : token.SourceOffset+uint32(token.SourceLength)],
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
		return &model.StringLiteralExpr{
			SourcePosition: model.NewSourcePos(token),
			Text:           p.sourceCode[token.SourceOffset : token.SourceOffset+uint32(token.SourceLength)],
		}

	case scanning.TokenTypeTrailingDocumentation:
		return &model.TrailingDocumentationExpr{
			SourcePosition: model.NewSourcePos(token),
			Text:           p.sourceCode[token.SourceOffset : token.SourceOffset+uint32(token.SourceLength)],
		}

	case scanning.TokenTypeTrue:
		return &model.BooleanLiteralExpr{
			SourcePosition: model.NewSourcePos(token),
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
) model.IExpression {
	rightBindingPower := prefixBindingPowers[token.TokenType].Power
	rhs := p.parseExprBindingPower(rightBindingPower)
	return &model.LogicalNotOperationExpr{
		SourcePosition: model.NewSourcePos(token),
		Operand:        rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (p *lligneParser) parseNegationOperationExpression(
	token scanning.Token,
) model.IExpression {
	rightBindingPower := prefixBindingPowers[token.TokenType].Power
	rhs := p.parseExprBindingPower(rightBindingPower)
	return &model.NegationOperationExpr{
		SourcePosition: model.NewSourcePos(token),
		Operand:        rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (p *lligneParser) parseParenthesizedExpression(
	token scanning.Token,
	endingTokenType scanning.TokenType,
) model.IExpression {

	var items []model.IExpression

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
	p.index += 1

	var delimiters model.ParenExprDelimiters
	switch endingTokenType {
	case scanning.TokenTypeEof:
		delimiters = model.ParenExprDelimitersWholeFile
	case scanning.TokenTypeRightBrace:
		delimiters = model.ParenExprDelimitersBraces
	case scanning.TokenTypeRightParenthesis:
		delimiters = model.ParenExprDelimitersParentheses
	}

	return &model.ParenthesizedExpr{
		SourcePosition: model.NewSourcePos(token),
		Delimiters:     delimiters,
		Items:          items,
	}

}

//---------------------------------------------------------------------------------------------------------------------

func (p *lligneParser) parsePostfixExpression(opToken scanning.Token, lhs model.IExpression) model.IExpression {

	switch opToken.TokenType {

	case scanning.TokenTypeLeftParenthesis:
		args := p.parseParenthesizedExpression(opToken, scanning.TokenTypeRightParenthesis)
		return &model.FunctionCallExpr{
			SourcePosition:    model.NewSourcePos(opToken),
			FunctionReference: lhs,
			Argument:          args,
		}

	case scanning.TokenTypeQuestionMark:
		return &model.OptionalExpr{
			SourcePosition: model.NewSourcePos(opToken),
			Operand:        lhs,
		}

	}

	panic("Unfinished parsing code: '" + strconv.Itoa(int(opToken.TokenType)) + "'.")

}

//---------------------------------------------------------------------------------------------------------------------

func (p *lligneParser) parseSequenceLiteral(token scanning.Token) model.IExpression {

	var items []model.IExpression

	if p.tokens[p.index].TokenType == scanning.TokenTypeRightBracket {
		p.index += 1
		return &model.SequenceLiteralExpr{
			SourcePosition: model.NewSourcePos(token),
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
	p.index += 1

	return &model.SequenceLiteralExpr{
		SourcePosition: model.NewSourcePos(token),
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
