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

type ILligneParser interface {
	ParseExpression() model.ILligneExpression
	ParseParenthesizedItems() model.ILligneExpression
}

//---------------------------------------------------------------------------------------------------------------------

func NewLligneParser(scanner scanning.ILligneBufferedScanner) ILligneParser {
	return &lligneParser{
		scanner: scanner,
	}
}

//=====================================================================================================================

type lligneParser struct {
	scanner scanning.ILligneBufferedScanner
}

//---------------------------------------------------------------------------------------------------------------------

func (p *lligneParser) ParseExpression() model.ILligneExpression {
	return p.parseExprBindingPower(0)
}

//---------------------------------------------------------------------------------------------------------------------

// ParseParenthesizedItems parses a non-empty sequence of code expected to be the items within a record literal, e.g.
// the top level of a file.
func (p *lligneParser) ParseParenthesizedItems() model.ILligneExpression {
	return p.parseParenthesizedExpression(p.scanner.PeekToken(), scanning.TokenTypeEof)
}

//---------------------------------------------------------------------------------------------------------------------

/**
 * Constructs an infix operator expression from given operator token [opToken], left-hand side [lhs],
 * and right-hand side [rhs].
 */
func (p *lligneParser) makeInfixExpression(
	sourceCodePos int,
	operator model.LligneInfixOperator,
	lhs model.ILligneExpression,
	rhs model.ILligneExpression,
) model.ILligneExpression {

	if lhs.TypeCode() == model.ExprTypeInfixOperation {
		lhsOrig := lhs.(*model.LligneInfixOperationExpr)
		if lhsOrig.Operator == operator {
			return &model.LligneInfixOperationExpr{SourcePos: lhsOrig.SourcePos, Operator: operator, Operands: append(lhsOrig.Operands, rhs)}
		}
	}

	return &model.LligneInfixOperationExpr{SourcePos: sourceCodePos, Operator: operator, Operands: []model.ILligneExpression{lhs, rhs}}

}

//---------------------------------------------------------------------------------------------------------------------

func (p *lligneParser) parseExprBindingPower(minBindingPower int) model.ILligneExpression {

	lhs := p.parseLeftHandSide()

	for {

		// Look ahead for an operator continuing the expression
		opToken := p.scanner.PeekToken()

		// Handle postfix operators ...
		pBindingPower := postfixBindingPowers[opToken.TokenType]

		if pBindingPower.Operator != model.PostfixOperatorNone {

			if pBindingPower.Power < minBindingPower {
				break
			}

			p.scanner.ReadToken()

			lhs = p.parsePostfixExpression(opToken, lhs)

			continue

		}

		// Handle infix operators ...
		bindingPower := infixBindingPowers[opToken.TokenType]

		if bindingPower.Operator != model.InfixOperatorNone {

			if bindingPower.Left < minBindingPower {
				break
			}

			p.scanner.ReadToken()

			rhs := p.parseExprBindingPower(bindingPower.Right)

			lhs = p.makeInfixExpression(opToken.SourceStartPos, bindingPower.Operator, lhs, rhs)

			continue

		}

		break

	}

	return lhs
}

//---------------------------------------------------------------------------------------------------------------------

func (p *lligneParser) parseLeftHandSide() model.ILligneExpression {

	token := p.scanner.ReadToken()

	switch token.TokenType {

	case scanning.TokenTypeBackTickedString:
		return &model.LligneMultilineStringLiteralExpr{SourcePos: token.SourceStartPos, Text: token.Text}

	case scanning.TokenTypeDash:
		return p.parsePrefixOperationExpression(token, model.PrefixOperatorNegation)

	case scanning.TokenTypeDoubleQuotedString:
		return &model.LligneStringLiteralExpr{SourcePos: token.SourceStartPos, Text: token.Text}

	case scanning.TokenTypeFalse:
		return &model.LligneBooleanLiteralExpr{SourcePos: token.SourceStartPos, Value: false}

	case scanning.TokenTypeIdentifier:
		return &model.LligneIdentifierExpr{SourcePos: token.SourceStartPos, Name: token.Text}

	case scanning.TokenTypeIntegerLiteral:
		return &model.LligneIntegerLiteralExpr{SourcePos: token.SourceStartPos, Text: token.Text}

	case scanning.TokenTypeLeadingDocumentation:
		return &model.LligneLeadingDocumentationExpr{SourcePos: token.SourceStartPos, Text: token.Text}

	case scanning.TokenTypeLeftBrace:
		return p.parseParenthesizedExpression(token, scanning.TokenTypeRightBrace)

	case scanning.TokenTypeLeftBracket:
		return p.parseSequenceLiteral(token)

	case scanning.TokenTypeLeftParenthesis:
		return p.parseParenthesizedExpression(token, scanning.TokenTypeRightParenthesis)

	case scanning.TokenTypeNot:
		return p.parsePrefixOperationExpression(token, model.PrefixOperatorLogicalNot)

	case scanning.TokenTypeSingleQuotedString:
		return &model.LligneStringLiteralExpr{SourcePos: token.SourceStartPos, Text: token.Text}

	case scanning.TokenTypeTrailingDocumentation:
		return &model.LligneTrailingDocumentationExpr{SourcePos: token.SourceStartPos, Text: token.Text}

	case scanning.TokenTypeTrue:
		return &model.LligneBooleanLiteralExpr{SourcePos: token.SourceStartPos, Value: true}

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
	token scanning.LligneToken,
	endingTokenType scanning.LligneTokenType,
) model.ILligneExpression {

	var items []model.ILligneExpression

	for !p.scanner.PeekTokenIsType(endingTokenType) {
		// Parse one expression.
		items = append(items, p.parseExprBindingPower(0))

		if !p.scanner.AdvanceTokenIfType(scanning.TokenTypeComma) {
			break
		}
	}

	if !p.scanner.AdvanceTokenIfType(endingTokenType) {
		panic("Expected " + endingTokenType.String())
	}

	var delimiters model.ParenExprDelimiters
	switch endingTokenType {
	case scanning.TokenTypeEof:
		delimiters = model.ParenExprDelimitersWholeFile
	case scanning.TokenTypeRightBrace:
		delimiters = model.ParenExprDelimitersBraces
	case scanning.TokenTypeRightParenthesis:
		delimiters = model.ParenExprDelimitersParentheses
	}

	return &model.LligneParenthesizedExpr{
		SourcePos:  token.SourceStartPos,
		Delimiters: delimiters,
		Items:      items,
	}

}

//---------------------------------------------------------------------------------------------------------------------

func (p *lligneParser) parsePostfixExpression(opToken scanning.LligneToken, lhs model.ILligneExpression) model.ILligneExpression {

	switch opToken.TokenType {

	case scanning.TokenTypeLeftParenthesis:
		args := p.parseParenthesizedExpression(opToken, scanning.TokenTypeRightParenthesis)
		return &model.LligneFunctionCallExpr{
			SourcePos:         opToken.SourceStartPos,
			FunctionReference: lhs,
			Argument:          args,
		}

	case scanning.TokenTypeQuestionMark:
		return &model.LligneOptionalExpr{
			SourcePos: opToken.SourceStartPos,
			Operand:   lhs,
		}

	}

	panic("Unfinished parsing code: '" + strconv.Itoa(int(opToken.TokenType)) + "'.")

}

//---------------------------------------------------------------------------------------------------------------------

func (p *lligneParser) parsePrefixOperationExpression(
	token scanning.LligneToken,
	operator model.LlignePrefixOperator,
) model.ILligneExpression {
	rightBindingPower := prefixBindingPowers[token.TokenType].Power
	rhs := p.parseExprBindingPower(rightBindingPower)
	return &model.LlignePrefixOperationExpr{
		SourcePos: token.SourceStartPos,
		Operator:  operator,
		Operand:   rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (p *lligneParser) parseSequenceLiteral(token scanning.LligneToken) model.ILligneExpression {

	var items []model.ILligneExpression

	if p.scanner.AdvanceTokenIfType(scanning.TokenTypeRightBracket) {
		return &model.LligneSequenceLiteralExpr{SourcePos: token.SourceStartPos, Elements: items}
	}

	for !p.scanner.PeekTokenIsType(scanning.TokenTypeRightBracket) {
		// Parse one expression.
		items = append(items, p.parseExprBindingPower(0))

		if !p.scanner.AdvanceTokenIfType(scanning.TokenTypeComma) {
			break
		}
	}

	if !p.scanner.AdvanceTokenIfType(scanning.TokenTypeRightBracket) {
		panic("Expected " + scanning.TokenTypeRightBracket.String())
	}

	return &model.LligneSequenceLiteralExpr{
		SourcePos: token.SourceStartPos,
		Elements:  items,
	}

	return &model.LligneSequenceLiteralExpr{}

}

//=====================================================================================================================

type infixBindingPower struct {
	Left     int
	Right    int
	Operator model.LligneInfixOperator
}

//=====================================================================================================================

type prefixBindingPower struct {
	Power    int
	Operator model.LlignePrefixOperator
}

//=====================================================================================================================

type postfixBindingPower struct {
	Power    int
	Operator model.LlignePostfixOperator
}

//=====================================================================================================================

var prefixBindingPowers = make(map[scanning.LligneTokenType]prefixBindingPower)

var infixBindingPowers = make(map[scanning.LligneTokenType]infixBindingPower)

var postfixBindingPowers = make(map[scanning.LligneTokenType]postfixBindingPower)

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
