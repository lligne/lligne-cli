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

			rhs := p.parseExprBindingPower(bindingPower.Right)

			lhs = &model.InfixOperationExpr{
				SourcePosition: model.NewSourcePos(opToken),
				Operator:       bindingPower.Operator,
				Lhs:            lhs,
				Rhs:            rhs,
			}

			continue

		}

		break

	}

	return lhs
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
		return p.parsePrefixOperationExpression(token, model.PrefixOperatorNegation)

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
		return p.parsePrefixOperationExpression(token, model.PrefixOperatorLogicalNot)

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

func (p *lligneParser) parsePrefixOperationExpression(
	token scanning.Token,
	operator model.PrefixOperator,
) model.IExpression {
	rightBindingPower := prefixBindingPowers[token.TokenType].Power
	rhs := p.parseExprBindingPower(rightBindingPower)
	return &model.PrefixOperationExpr{
		SourcePosition: model.NewSourcePos(token),
		Operator:       operator,
		Operand:        rhs,
	}
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
	Left     int
	Right    int
	Operator model.InfixOperator
}

//=====================================================================================================================

type prefixBindingPower struct {
	Power    int
	Operator model.PrefixOperator
}

//=====================================================================================================================

type postfixBindingPower struct {
	Power    int
	Operator model.PostfixOperator
}

//=====================================================================================================================

var prefixBindingPowers = make(map[scanning.TokenType]prefixBindingPower)

var infixBindingPowers = make(map[scanning.TokenType]infixBindingPower)

var postfixBindingPowers = make(map[scanning.TokenType]postfixBindingPower)

func init() {

	level := 1

	infixBindingPowers[scanning.TokenTypeColon] = infixBindingPower{level, level + 1, model.InfixOperatorQualify}
	infixBindingPowers[scanning.TokenTypeEquals] = infixBindingPower{level, level + 1, model.InfixOperatorIntersectAssignValue}
	infixBindingPowers[scanning.TokenTypeQuestionMarkColon] = infixBindingPower{level, level + 1, model.InfixOperatorIntersectDefaultValue}

	level += 2

	infixBindingPowers[scanning.TokenTypeAmpersandAmpersand] = infixBindingPower{level, level + 1, model.InfixOperatorIntersectLowPrecedence}

	level += 2

	infixBindingPowers[scanning.TokenTypeVerticalBar] = infixBindingPower{level, level + 1, model.InfixOperatorUnion}

	level += 2

	infixBindingPowers[scanning.TokenTypeAmpersand] = infixBindingPower{level, level + 1, model.InfixOperatorIntersect}

	level += 2

	infixBindingPowers[scanning.TokenTypeWhen] = infixBindingPower{level, level + 1, model.InfixOperatorWhen}
	infixBindingPowers[scanning.TokenTypeWhere] = infixBindingPower{level, level + 1, model.InfixOperatorWhere}

	level += 2

	infixBindingPowers[scanning.TokenTypeSynthDocument] = infixBindingPower{level, level + 1, model.InfixOperatorDocument}

	level += 2

	infixBindingPowers[scanning.TokenTypeOr] = infixBindingPower{level, level + 1, model.InfixOperatorLogicOr}

	level += 2

	infixBindingPowers[scanning.TokenTypeAnd] = infixBindingPower{level, level + 1, model.InfixOperatorLogicAnd}

	level += 2

	prefixBindingPowers[scanning.TokenTypeNot] = prefixBindingPower{level, model.PrefixOperatorLogicalNot}

	level += 2

	infixBindingPowers[scanning.TokenTypeEqualsEquals] = infixBindingPower{level, level + 1, model.InfixOperatorEquals}
	infixBindingPowers[scanning.TokenTypeGreaterThan] = infixBindingPower{level, level + 1, model.InfixOperatorGreaterThan}
	infixBindingPowers[scanning.TokenTypeGreaterThanOrEquals] = infixBindingPower{level, level + 1, model.InfixOperatorGreaterThanOrEquals}
	infixBindingPowers[scanning.TokenTypeLessThan] = infixBindingPower{level, level + 1, model.InfixOperatorLessThan}
	infixBindingPowers[scanning.TokenTypeLessThanOrEquals] = infixBindingPower{level, level + 1, model.InfixOperatorLessThanOrEquals}

	level += 2

	infixBindingPowers[scanning.TokenTypeIn] = infixBindingPower{level, level + 1, model.InfixOperatorIn}
	infixBindingPowers[scanning.TokenTypeIs] = infixBindingPower{level, level + 1, model.InfixOperatorIs}
	infixBindingPowers[scanning.TokenTypeMatches] = infixBindingPower{level, level + 1, model.InfixOperatorMatch}
	infixBindingPowers[scanning.TokenTypeNotMatches] = infixBindingPower{level, level + 1, model.InfixOperatorNotMatch}

	level += 2

	infixBindingPowers[scanning.TokenTypeDotDot] = infixBindingPower{level, level + 1, model.InfixOperatorRange}

	level += 2

	infixBindingPowers[scanning.TokenTypeDash] = infixBindingPower{level, level + 1, model.InfixOperatorSubtract}
	infixBindingPowers[scanning.TokenTypePlus] = infixBindingPower{level, level + 1, model.InfixOperatorAdd}

	level += 2

	infixBindingPowers[scanning.TokenTypeAsterisk] = infixBindingPower{level, level + 1, model.InfixOperatorMultiply}
	infixBindingPowers[scanning.TokenTypeSlash] = infixBindingPower{level, level + 1, model.InfixOperatorDivide}

	level += 2

	prefixBindingPowers[scanning.TokenTypeDash] = prefixBindingPower{level, model.PrefixOperatorNegation}

	level += 2

	infixBindingPowers[scanning.TokenTypeRightArrow] = infixBindingPower{level, level + 1, model.InfixOperatorFunctionCall}

	level += 2

	infixBindingPowers[scanning.TokenTypeDot] = infixBindingPower{level, level + 1, model.InfixOperatorFieldReference}

	level += 2

	postfixBindingPowers[scanning.TokenTypeLeftParenthesis] = postfixBindingPower{level, model.PostfixOperatorFunctionCall}
	postfixBindingPowers[scanning.TokenTypeLeftBracket] = postfixBindingPower{level, model.PostfixOperatorIndex}
	postfixBindingPowers[scanning.TokenTypeQuestionMark] = postfixBindingPower{level, model.PostfixOperatorOptional}

}

//=====================================================================================================================
