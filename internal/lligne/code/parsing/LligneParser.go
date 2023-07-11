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

func (p *lligneParser) parseExprBindingPower(minBindingPower int) model.ILligneExpression {

	lhs := p.parseLeftHandSide()

	// TODO: Pratt expression parsing ...

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
