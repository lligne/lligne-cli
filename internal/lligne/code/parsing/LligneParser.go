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

//---------------------------------------------------------------------------------------------------------------------

type ILligneParser interface {
	ParseExpression() model.ILligneExpression
}

//---------------------------------------------------------------------------------------------------------------------

func NewLligneParser(scanner scanning.ILligneBufferedScanner) ILligneParser {
	return &lligneParser{
		scanner: scanner,
	}
}

//---------------------------------------------------------------------------------------------------------------------

type lligneParser struct {
	scanner scanning.ILligneBufferedScanner
}

//---------------------------------------------------------------------------------------------------------------------

func (p *lligneParser) ParseExpression() model.ILligneExpression {
	return p.parseExprBindingPower(0)
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

		//// Handle postfix operators ...
		//pBindingPower := postfixBindingPower.get(opToken.TokenType)
		//
		//if pBindingPower != nil {
		//
		//	if pBindingPower < minBindingPower {
		//		break
		//	}
		//
		//	p.scanner.ReadToken()
		//
		//	lhs = p.ParsePostfixExpression(opToken, lhs)
		//
		//	continue
		//
		//}

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

	//	case LlaceTokenType.DASH:
	//	return this.#parsePrefixOperationExpression(token.origin, token.type, LlaceUnaryOperator.ArithmeticNegation)

	case scanning.TokenTypeDoubleQuotedString:
		return &model.LligneStringLiteralExpr{SourcePos: token.SourceStartPos, Text: token.Text}

	case scanning.TokenTypeIdentifier:
		return &model.LligneIdentifierExpr{SourcePos: token.SourceStartPos, Name: token.Text}

	case scanning.TokenTypeIntegerLiteral:
		return &model.LligneIntegerLiteralExpr{SourcePos: token.SourceStartPos, Text: token.Text}

	case scanning.TokenTypeLeadingDocumentation:
		return &model.LligneLeadingDocumentationExpr{SourcePos: token.SourceStartPos, Text: token.Text}

	case scanning.TokenTypeSingleQuotedString:
		return &model.LligneStringLiteralExpr{SourcePos: token.SourceStartPos, Text: token.Text}

		//	case LlaceTokenType.LEFT_BRACKET:
		//	return this.#parseArrayLiteral(token.origin)
		//
		//	case LlaceTokenType.LEFT_PARENTHESIS:
		//	return this.#parseParenthesizedExpression(token.origin, LlaceTokenType.RIGHT_PARENTHESIS)
		//
		//	case LlaceTokenType.NOT:
		//	return this.#parsePrefixOperationExpression(token.origin, token.type, LlaceUnaryOperator.LogicalNegation)
		//
		//	case LlaceTokenType.TRAILING_DOCUMENTATION: {
		//	const rawLines = token.text.split("\n")
		//	const lines = rawLines.map(line => line.trim()).filter(line => line.length > 0)
		//	return new LlaceTrailingDocumentationExpr(token.origin, lines)
		//	}
		//
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

type infixBindingPower struct {
	Left     int
	Right    int
	Operator model.LligneInfixOperator
}

//---------------------------------------------------------------------------------------------------------------------

type prefixBindingPower struct {
	Power    int
	Operator model.LlignePrefixOperator
}

//---------------------------------------------------------------------------------------------------------------------

type postfixBindingPower struct {
	Power    int
	Operator model.LlignePostfixOperator
}

//---------------------------------------------------------------------------------------------------------------------

var prefixBindingPowers = make(map[scanning.LligneTokenType]prefixBindingPower)

var infixBindingPowers = make(map[scanning.LligneTokenType]infixBindingPower)

var postfixBindingPowers = make(map[scanning.LligneTokenType]postfixBindingPower)

func init() {

	level := 1
	infixBindingPowers[scanning.TokenTypeColon] = infixBindingPower{level, level + 1, model.InfixOperatorQualify}
	infixBindingPowers[scanning.TokenTypeEquals] = infixBindingPower{level, level + 1, model.InfixOperatorIntersectAssignValue}
	infixBindingPowers[scanning.TokenTypeQuestionMark] = infixBindingPower{level, level + 1, model.InfixOperatorIntersectDefaultValue}

	level += 2

	infixBindingPowers[scanning.TokenTypeAmpersandAmpersand] = infixBindingPower{level, level + 1, model.InfixOperatorIntersectLowPrecedence}

	level += 2

	infixBindingPowers[scanning.TokenTypeVerticalBar] = infixBindingPower{level, level + 1, model.InfixOperatorUnion}

	level += 2

	infixBindingPowers[scanning.TokenTypeAmpersand] = infixBindingPower{level, level + 1, model.InfixOperatorIntersect}

	level += 2

	infixBindingPowers[scanning.TokenTypeSynthDocument] = infixBindingPower{level, level + 1, model.InfixOperatorDocument}

	level += 2

	infixBindingPowers[scanning.TokenTypeOr] = infixBindingPower{level, level + 1, model.InfixOperatorLogicOr}

	level += 2

	infixBindingPowers[scanning.TokenTypeAnd] = infixBindingPower{level, level + 1, model.InfixOperatorLogicAnd}

	level += 2

	prefixBindingPowers[scanning.TokenTypeNot] = prefixBindingPower{level, model.PrefixOperatorLogicalNot}

	level += 2

	infixBindingPowers[scanning.TokenTypeEqualsEquals] = infixBindingPower{level, level + 1, model.InfixOperatorEquality}
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

//---------------------------------------------------------------------------------------------------------------------
