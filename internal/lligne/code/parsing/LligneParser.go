//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package parsing

import (
	"lligne-cli/internal/lligne/code/model"
	"lligne-cli/internal/lligne/code/scanning"
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

func (p *lligneParser) parseExprBindingPower(minBindingPower int) model.ILligneExpression {

	lhs := p.parseLeftHandSide()

	// TODO: Pratt expression parsing ...

	return lhs
}

//---------------------------------------------------------------------------------------------------------------------

func (p *lligneParser) parseLeftHandSide() model.ILligneExpression {

	token := p.scanner.ReadToken()

	switch token.TokenType {

	//case LlaceTokenType.CHAR_LITERAL:
	//	return new LlaceCharLiteralExpr(token.origin, token.text)
	//
	//	case LlaceTokenType.DASH:
	//	return this.#parsePrefixOperationExpression(token.origin, token.type, LlaceUnaryOperator.ArithmeticNegation)
	//
	//	case LlaceTokenType.DATE_LITERAL:
	//	return new LlaceDateLiteralExpr(token.origin, token.text)
	//
	//	case LlaceTokenType.DATETIME_LITERAL:
	//	return new LlaceDateTimeLiteralExpr(token.origin, token.text)
	//
	case scanning.TokenTypeIdentifier:
		return model.NewLligneIdentifierExpr(token.SourceStartPos, token.Text)

		//	case LlaceTokenType.INTEGER_LITERAL:
		//	return new LlaceIntegerLiteralExpr(token.origin, token.text)
		//
		//	case LlaceTokenType.LEADING_DOCUMENTATION: {
		//	const rawLines = token.text.split("\n")
		//	const lines = rawLines.map(line => line.trim()).filter(line => line.length > 0)
		//	return new LlaceLeadingDocumentationExpr(token.origin, lines)
		//	}
		//
		//	case LlaceTokenType.LEFT_BRACKET:
		//	return this.#parseArrayLiteral(token.origin)
		//
		//	case LlaceTokenType.LEFT_PARENTHESIS:
		//	return this.#parseParenthesizedExpression(token.origin, LlaceTokenType.RIGHT_PARENTHESIS)
		//
		//	case LlaceTokenType.MULTILINE_STRING:
		//	const rawLines = token.text.split("\n")
		//	const lines = rawLines.map(line => line.trim()).filter(line => line.length > 0)
		//	return new LlaceMultilineStringLiteralExpr(token.origin, lines)
		//
		//	case LlaceTokenType.NOT:
		//	return this.#parsePrefixOperationExpression(token.origin, token.type, LlaceUnaryOperator.LogicalNegation)
		//
		//	case LlaceTokenType.STRING_LITERAL:
		//	return new LlaceStringLiteralExpr(token.origin, token.text)
		//
		//	case LlaceTokenType.TRAILING_DOCUMENTATION: {
		//	const rawLines = token.text.split("\n")
		//	const lines = rawLines.map(line => line.trim()).filter(line => line.length > 0)
		//	return new LlaceTrailingDocumentationExpr(token.origin, lines)
		//	}
		//
		//	case LlaceTokenType.UUID_LITERAL:
		//	return new LlaceUuidLiteralExpr(token.origin, token.text)
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

	panic("Unfinished parsing code")

}

//---------------------------------------------------------------------------------------------------------------------
