//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package parsing

import (
	"fmt"
	"lligne-cli/internal/lligne/code/scanning"
	"lligne-cli/internal/lligne/code/util"
	"strconv"
)

//=====================================================================================================================

type Outcome struct {
	SourceCode     string
	NewLineOffsets []uint32
	Model          IExpression
}

//=====================================================================================================================

func ParseExpression(scanResult *scanning.Outcome) *Outcome {
	parser := newParser(scanResult)

	model := parser.parseExprBindingPower(0)

	return &Outcome{
		SourceCode:     scanResult.SourceCode,
		NewLineOffsets: scanResult.NewLineOffsets,
		Model:          model,
	}
}

//---------------------------------------------------------------------------------------------------------------------

// TODO: ParseTopLevel
// ParseParenthesizedItems parses a non-empty sequence of code expected to be the items within a record literal, e.g.
// the top level of a file.
//func ParseParenthesizedItems(sourceCode string, tokens []scanning.Token) IExpression {
//	parser := newParser(sourceCode, tokens)
//
//	return parser.parseParenthesizedExpression(tokens[0], scanning.TokenTypeEof)
//}

//=====================================================================================================================

type lligneParser struct {
	tokens     []scanning.Token
	index      int
	sourceCode string
}

//---------------------------------------------------------------------------------------------------------------------

func newParser(scanResult *scanning.Outcome) *lligneParser {
	return &lligneParser{
		sourceCode: scanResult.SourceCode,
		tokens:     scanResult.Tokens,
	}
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

func (p *lligneParser) parseFunctionArgumentsExpression(
	token scanning.Token,
) IExpression {

	var items []IExpression

	for p.tokens[p.index].TokenType != scanning.TokenTypeRightParenthesis {
		// Parse one expression.
		items = append(items, p.parseExprBindingPower(0))

		if p.tokens[p.index].TokenType != scanning.TokenTypeComma {
			break
		}
		p.index += 1
	}

	if p.tokens[p.index].TokenType != scanning.TokenTypeRightParenthesis {
		panic("Expected " + scanning.TokenTypeRightParenthesis.String())
	}
	endSourcePos := util.NewSourcePos(p.tokens[p.index])
	p.index += 1

	return &FunctionArgumentsExpr{
		SourcePosition: util.NewSourcePos(token).Thru(endSourcePos),
		Items:          items,
	}

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

	case scanning.TokenTypeEqualsTilde:
		return &MatchExpr{
			SourcePosition: lhs.GetSourcePosition().Thru(rhs.GetSourcePosition()),
			Lhs:            lhs,
			Rhs:            rhs,
		}

	case scanning.TokenTypeExclamationEquals:
		return &NotEqualsExpr{
			SourcePosition: lhs.GetSourcePosition().Thru(rhs.GetSourcePosition()),
			Lhs:            lhs,
			Rhs:            rhs,
		}

	case scanning.TokenTypeExclamationTilde:
		return &NotMatchExpr{
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

	case scanning.TokenTypeQuestionColon:
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

	default:
		panic(fmt.Sprintf("Missing case in parseInfixOperation: " + strconv.Itoa(int(opToken.TokenType)) + "'."))

	}

}

//---------------------------------------------------------------------------------------------------------------------

func (p *lligneParser) parseLeftHandSide() IExpression {

	token := p.tokens[p.index]
	p.index += 1

	switch token.TokenType {

	case scanning.TokenTypeBackTickedString:
		return &StringLiteralExpr{
			SourcePosition: util.NewSourcePos(token),
			Delimiters:     StringDelimitersBackTicksMultiline,
		}

	case scanning.TokenTypeBuiltInType:
		return &BuiltInTypeExpr{
			SourcePosition: util.NewSourcePos(token),
		}

	case scanning.TokenTypeDash:
		return p.parseNegationOperationExpression(token)

	case scanning.TokenTypeDoubleQuotedString:
		return &StringLiteralExpr{
			SourcePosition: util.NewSourcePos(token),
			Delimiters:     StringDelimitersDoubleQuotes,
		}

	case scanning.TokenTypeFalse:
		return &BooleanLiteralExpr{
			SourcePosition: util.NewSourcePos(token),
			Value:          false,
		}

	case scanning.TokenTypeFloatingPointLiteral:
		sourcePosition := util.NewSourcePos(token)
		valueStr := sourcePosition.GetText(p.sourceCode)
		value, _ := strconv.ParseFloat(valueStr, 64)
		return &Float64LiteralExpr{
			SourcePosition: sourcePosition,
			Value:          value,
		}

	case scanning.TokenTypeIdentifier:
		return &IdentifierExpr{
			SourcePosition: util.NewSourcePos(token),
		}

	case scanning.TokenTypeIntegerLiteral:
		sourcePosition := util.NewSourcePos(token)
		valueStr := sourcePosition.GetText(p.sourceCode)
		value, _ := strconv.ParseInt(valueStr, 10, 64)
		return &Int64LiteralExpr{
			SourcePosition: util.NewSourcePos(token),
			Value:          value,
		}

	case scanning.TokenTypeLeadingDocumentation:
		return &LeadingDocumentationExpr{
			SourcePosition: util.NewSourcePos(token),
		}

	case scanning.TokenTypeLeftBrace:
		return p.parseRecordExpression(token)

	case scanning.TokenTypeLeftBracket:
		return p.parseSequenceLiteral(token)

	case scanning.TokenTypeLeftParenthesis:
		return p.parseParenthesizedExpression(token)

	case scanning.TokenTypeNot:
		return p.parseLogicalNotOperationExpression(token)

	case scanning.TokenTypeSingleQuotedString:
		return &StringLiteralExpr{
			SourcePosition: util.NewSourcePos(token),
			Delimiters:     StringDelimitersSingleQuotes,
		}

	case scanning.TokenTypeTrailingDocumentation:
		return &TrailingDocumentationExpr{
			SourcePosition: util.NewSourcePos(token),
		}

	case scanning.TokenTypeTrue:
		return &BooleanLiteralExpr{
			SourcePosition: util.NewSourcePos(token),
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
		SourcePosition: util.NewSourcePos(token),
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
		SourcePosition: util.NewSourcePos(token).Thru(rhs.GetSourcePosition()),
		Operand:        rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (p *lligneParser) parseParenthesizedExpression(
	token scanning.Token,
) IExpression {

	// Handle empty parentheses specially.
	if p.tokens[p.index].TokenType == scanning.TokenTypeRightParenthesis {
		endSourcePos := util.NewSourcePos(p.tokens[p.index])
		p.index += 1

		return &UnitExpr{
			SourcePosition: util.NewSourcePos(token).Thru(endSourcePos),
		}
	}

	// Parse one expression.
	inner := p.parseExprBindingPower(0)

	// Comma means function parameters
	if p.tokens[p.index].TokenType == scanning.TokenTypeComma {

		p.index += 1

		var items []IExpression
		items = append(items, inner)

		for p.tokens[p.index].TokenType != scanning.TokenTypeRightParenthesis {
			// Parse one expression.
			items = append(items, p.parseExprBindingPower(0))

			if p.tokens[p.index].TokenType != scanning.TokenTypeComma {
				break
			}
			p.index += 1
		}

		if p.tokens[p.index].TokenType != scanning.TokenTypeRightParenthesis {
			panic("Expected " + scanning.TokenTypeRightParenthesis.String())
		}
		endSourcePos := util.NewSourcePos(p.tokens[p.index])
		p.index += 1

		return &FunctionArgumentsExpr{
			SourcePosition: util.NewSourcePos(token).Thru(endSourcePos),
			Items:          items,
		}

	}

	if p.tokens[p.index].TokenType != scanning.TokenTypeRightParenthesis {
		panic("Expected " + scanning.TokenTypeRightParenthesis.String())
	}

	endSourcePos := util.NewSourcePos(p.tokens[p.index])
	p.index += 1

	return &ParenthesizedExpr{
		SourcePosition: util.NewSourcePos(token).Thru(endSourcePos),
		InnerExpr:      inner,
	}

}

//---------------------------------------------------------------------------------------------------------------------

func (p *lligneParser) parsePostfixExpression(opToken scanning.Token, lhs IExpression) IExpression {

	switch opToken.TokenType {

	case scanning.TokenTypeLeftParenthesis:
		args := p.parseFunctionArgumentsExpression(opToken)
		return &FunctionCallExpr{
			SourcePosition:    lhs.GetSourcePosition().Thru(args.GetSourcePosition()),
			FunctionReference: lhs,
			Argument:          args,
		}

	case scanning.TokenTypeQuestion:
		return &OptionalExpr{
			SourcePosition: lhs.GetSourcePosition(),
			Operand:        lhs,
		}

	}

	panic("Unfinished postfix parsing code: '" + strconv.Itoa(int(opToken.TokenType)) + "'.")

}

//---------------------------------------------------------------------------------------------------------------------

func (p *lligneParser) parseRecordExpression(
	token scanning.Token,
) IExpression {

	var items []IExpression

	for p.tokens[p.index].TokenType != scanning.TokenTypeRightBrace {
		// Parse one expression.
		items = append(items, p.parseExprBindingPower(0))

		if p.tokens[p.index].TokenType != scanning.TokenTypeComma {
			break
		}
		p.index += 1
	}

	if p.tokens[p.index].TokenType != scanning.TokenTypeRightBrace {
		panic("Expected " + scanning.TokenTypeRightBrace.String())
	}
	endSourcePos := util.NewSourcePos(p.tokens[p.index])
	p.index += 1

	return &RecordExpr{
		SourcePosition: util.NewSourcePos(token).Thru(endSourcePos),
		Items:          items,
	}

}

//---------------------------------------------------------------------------------------------------------------------

func (p *lligneParser) parseSequenceLiteral(token scanning.Token) IExpression {

	startSourcePos := util.NewSourcePos(token)
	var items []IExpression

	if p.tokens[p.index].TokenType == scanning.TokenTypeRightBracket {
		endSourcePos := util.NewSourcePos(p.tokens[p.index])
		p.index += 1
		return &ArrayLiteralExpr{
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
	endSourcePos := util.NewSourcePos(p.tokens[p.index])
	p.index += 1

	return &ArrayLiteralExpr{
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
	infixBindingPowers[scanning.TokenTypeQuestionColon] = infixBindingPower{level, level + 1}

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
	infixBindingPowers[scanning.TokenTypeExclamationEquals] = infixBindingPower{level, level + 1}
	infixBindingPowers[scanning.TokenTypeGreaterThan] = infixBindingPower{level, level + 1}
	infixBindingPowers[scanning.TokenTypeGreaterThanOrEquals] = infixBindingPower{level, level + 1}
	infixBindingPowers[scanning.TokenTypeLessThan] = infixBindingPower{level, level + 1}
	infixBindingPowers[scanning.TokenTypeLessThanOrEquals] = infixBindingPower{level, level + 1}

	level += 2

	infixBindingPowers[scanning.TokenTypeIn] = infixBindingPower{level, level + 1}
	infixBindingPowers[scanning.TokenTypeIs] = infixBindingPower{level, level + 1}
	infixBindingPowers[scanning.TokenTypeEqualsTilde] = infixBindingPower{level, level + 1}
	infixBindingPowers[scanning.TokenTypeExclamationTilde] = infixBindingPower{level, level + 1}

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
	postfixBindingPowers[scanning.TokenTypeQuestion] = postfixBindingPower{level}

}

//=====================================================================================================================
