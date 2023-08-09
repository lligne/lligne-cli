//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package model

import (
	"lligne-cli/internal/lligne/code/scanning"
	"strconv"
	"strings"
)

//=====================================================================================================================

// ILligneExpression is the interface to an expression AST node.
type ILligneExpression interface {
	ExprType() LligneExprType
	GetOrigin(tracker scanning.ILligneTokenOriginTracker) scanning.LligneOrigin
	SExpression() string
	SetTypeInfo(typeInfo ILligneType)
	SourcePos() int
	TypeInfo() ILligneType
}

//=====================================================================================================================

// LligneExprType is an enumeration of Lligne expression types.
type LligneExprType int

const (
	ExprTypeBooleanLiteral LligneExprType = 1 + iota
	ExprTypeFloatingPointLiteral
	ExprTypeFunctionCall
	ExprTypeIdentifier
	ExprTypeInfixOperation
	ExprTypeIntegerLiteral
	ExprTypeLeadingDocumentation
	ExprTypeMultilineStringLiteral
	ExprTypeOptional
	ExprTypeParenthesized
	ExprTypePrefixOperation
	ExprTypeSequenceLiteral
	ExprTypeStringLiteral
	ExprTypeTrailingDocumentation
)

//=====================================================================================================================

// LligneInfixOperator is an enumeration of Lligne's binary operators.
type LligneInfixOperator int

const (
	InfixOperatorNone LligneInfixOperator = iota
	InfixOperatorAdd
	InfixOperatorDivide
	InfixOperatorDocument
	InfixOperatorEquals
	InfixOperatorFieldReference
	InfixOperatorFunctionCall
	InfixOperatorGreaterThan
	InfixOperatorGreaterThanOrEquals
	InfixOperatorIn
	InfixOperatorIntersect
	InfixOperatorIntersectAssignValue
	InfixOperatorIntersectDefaultValue
	InfixOperatorIntersectLowPrecedence
	InfixOperatorIs
	InfixOperatorLessThan
	InfixOperatorLessThanOrEquals
	InfixOperatorLogicAnd
	InfixOperatorLogicOr
	InfixOperatorMatch
	InfixOperatorMultiply
	InfixOperatorNotMatch
	InfixOperatorQualify
	InfixOperatorRange
	InfixOperatorSubtract
	InfixOperatorUnion
	InfixOperatorWhen
	InfixOperatorWhere
)

// ---------------------------------------------------------------------------------------------------------------------

// String returns a string representing the code of an operator.
func (op LligneInfixOperator) String() string {

	switch op {

	case InfixOperatorAdd:
		return " + "
	case InfixOperatorDivide:
		return " / "
	case InfixOperatorDocument:
		return " "
	case InfixOperatorEquals:
		return " == "
	case InfixOperatorFieldReference:
		return "."
	case InfixOperatorFunctionCall:
		return " -> "
	case InfixOperatorGreaterThan:
		return " > "
	case InfixOperatorGreaterThanOrEquals:
		return " >= "
	case InfixOperatorIn:
		return " in "
	case InfixOperatorIntersect:
		return " & "
	case InfixOperatorIntersectAssignValue:
		return " = "
	case InfixOperatorIntersectDefaultValue:
		return " ?: "
	case InfixOperatorIntersectLowPrecedence:
		return " && "
	case InfixOperatorIs:
		return " is "
	case InfixOperatorLessThan:
		return " < "
	case InfixOperatorLessThanOrEquals:
		return " <= "
	case InfixOperatorLogicAnd:
		return " and "
	case InfixOperatorLogicOr:
		return " or "
	case InfixOperatorMatch:
		return " =~ "
	case InfixOperatorMultiply:
		return " * "
	case InfixOperatorNotMatch:
		return " !~ "
	case InfixOperatorQualify:
		return ": "
	case InfixOperatorRange:
		return ".."
	case InfixOperatorSubtract:
		return " - "
	case InfixOperatorUnion:
		return " | "
	case InfixOperatorWhen:
		return " when "
	case InfixOperatorWhere:
		return " where "

	}

	panic("Unhandled infix operator: '" + strconv.Itoa(int(op)) + "'.")
}

//=====================================================================================================================

// ParenExprDelimiters is an enumeration of start/stop delimiters for parenthesized expressions.
type ParenExprDelimiters int

const (
	ParenExprDelimitersParentheses ParenExprDelimiters = 1 + iota
	ParenExprDelimitersBraces
	ParenExprDelimitersWholeFile
)

//=====================================================================================================================

// LlignePrefixOperator is an enumeration of Lligne's prefix operators.
type LlignePrefixOperator int

const (
	PrefixOperatorLogicalNot LlignePrefixOperator = 1 + iota
	PrefixOperatorNegation
)

// ---------------------------------------------------------------------------------------------------------------------

// String returns a string representing the code of an operator.
func (op LlignePrefixOperator) String() string {

	switch op {

	case PrefixOperatorLogicalNot:
		return "not "
	case PrefixOperatorNegation:
		return "-"
	}

	panic("Unhandled prefix operator: '" + strconv.Itoa(int(op)) + "'.")
}

//=====================================================================================================================

// LlignePostfixOperator is an enumeration of Lligne's prefix operators.
type LlignePostfixOperator int

const (
	PostfixOperatorNone LlignePostfixOperator = iota
	PostfixOperatorFunctionCall
	PostfixOperatorIndex
	PostfixOperatorOptional
)

//=====================================================================================================================

// LligneExpression represents any expression
type LligneExpression struct {
	exprType  LligneExprType
	sourcePos int
	typeInfo  ILligneType
}

func (e *LligneExpression) ExprType() LligneExprType {
	return e.exprType
}

func (e *LligneExpression) GetOrigin(tracker scanning.ILligneTokenOriginTracker) scanning.LligneOrigin {
	return tracker.GetOrigin(e.sourcePos)
}

func (e *LligneExpression) SetTypeInfo(typeInfo ILligneType) {
	if e.typeInfo != nil {
		panic("Type info is immutable")
	}
	e.typeInfo = typeInfo
}

func (e *LligneExpression) SourcePos() int {
	return e.sourcePos
}

func (e *LligneExpression) TypeInfo() ILligneType {
	return e.typeInfo
}

//=====================================================================================================================

// LligneBooleanLiteralExpr represents a single boolean literal.
type LligneBooleanLiteralExpr struct {
	LligneExpression
	Value bool
}

func NewBooleanLiteralExpr(
	sourcePos int,
	value bool,
) ILligneExpression {
	return &LligneBooleanLiteralExpr{
		LligneExpression: LligneExpression{sourcePos: sourcePos, exprType: ExprTypeBooleanLiteral},
		Value:            value,
	}
}

func (e *LligneBooleanLiteralExpr) SExpression() string {
	if e.Value {
		return "(bool true)"
	}
	return "(bool false)"
}

//=====================================================================================================================

// LligneFloatingPointLiteralExpr represents a single integer literal.
type LligneFloatingPointLiteralExpr struct {
	LligneExpression
	Text string
}

func NewFloatingPointLiteralExpr(
	sourcePos int,
	text string,
) ILligneExpression {
	return &LligneFloatingPointLiteralExpr{
		LligneExpression: LligneExpression{sourcePos: sourcePos, exprType: ExprTypeFloatingPointLiteral},
		Text:             text,
	}
}

func (e *LligneFloatingPointLiteralExpr) SExpression() string {
	return "(float " + e.Text + ")"
}

//=====================================================================================================================

// LligneFunctionCallExpr represents a function call (a function name followed by a parenthesized expression).
type LligneFunctionCallExpr struct {
	LligneExpression
	FunctionReference ILligneExpression
	Argument          ILligneExpression
}

func NewFunctionCallExpr(
	sourcePos int,
	functionReference ILligneExpression,
	argument ILligneExpression,
) ILligneExpression {
	return &LligneFunctionCallExpr{
		LligneExpression:  LligneExpression{sourcePos: sourcePos, exprType: ExprTypeFunctionCall},
		FunctionReference: functionReference,
		Argument:          argument,
	}
}

func (e *LligneFunctionCallExpr) SExpression() string {
	result := "(call "
	result += e.FunctionReference.SExpression()
	result += " "
	result += e.Argument.SExpression()
	result += ")"
	return result
}

//=====================================================================================================================

// LligneIdentifierExpr represents a single identifier.
type LligneIdentifierExpr struct {
	LligneExpression
	Name string
}

func NewIdentifierExpr(
	sourcePos int,
	name string,
) ILligneExpression {
	return &LligneIdentifierExpr{
		LligneExpression: LligneExpression{sourcePos: sourcePos, exprType: ExprTypeIdentifier},
		Name:             name,
	}
}

func (e *LligneIdentifierExpr) SExpression() string {
	return "(id " + e.Name + ")"
}

//=====================================================================================================================

// LligneInfixOperationExpr represents an infix operation.
type LligneInfixOperationExpr struct {
	LligneExpression
	Operator LligneInfixOperator
	Lhs      ILligneExpression
	Rhs      ILligneExpression
}

func NewInfixOperationExpr(
	sourcePos int,
	operator LligneInfixOperator,
	lhs ILligneExpression,
	rhs ILligneExpression,
) ILligneExpression {
	return &LligneInfixOperationExpr{
		LligneExpression: LligneExpression{sourcePos: sourcePos, exprType: ExprTypeInfixOperation},
		Operator:         operator,
		Lhs:              lhs,
		Rhs:              rhs,
	}
}

func (e *LligneInfixOperationExpr) SExpression() string {
	return "(" + strings.TrimSpace(e.Operator.String()) + " " +
		e.Lhs.SExpression() + " " + e.Rhs.SExpression() + ")"
}

//=====================================================================================================================

// LligneIntegerLiteralExpr represents a single integer literal.
type LligneIntegerLiteralExpr struct {
	LligneExpression
	Text string
}

func NewIntegerLiteralExpr(
	sourcePos int,
	text string,
) ILligneExpression {
	return &LligneIntegerLiteralExpr{
		LligneExpression: LligneExpression{sourcePos: sourcePos, exprType: ExprTypeIntegerLiteral},
		Text:             text,
	}
}

func (e *LligneIntegerLiteralExpr) SExpression() string {
	return "(int " + e.Text + ")"
}

//=====================================================================================================================

// LligneLeadingDocumentationExpr represents lines of leading documentation.
type LligneLeadingDocumentationExpr struct {
	LligneExpression
	Text string
}

func NewLeadingDocumentationExpr(
	sourcePos int,
	text string,
) ILligneExpression {
	return &LligneLeadingDocumentationExpr{
		LligneExpression: LligneExpression{sourcePos: sourcePos, exprType: ExprTypeLeadingDocumentation},
		Text:             text,
	}
}

func (e *LligneLeadingDocumentationExpr) SExpression() string {
	return "(leadingdoc\n" + e.Text + ")"
}

//=====================================================================================================================

// LligneMultilineStringLiteralExpr represents a multiline (back-ticked) string literal.
type LligneMultilineStringLiteralExpr struct {
	LligneExpression
	Text string
}

func NewMultilineStringLiteralExpr(
	sourcePos int,
	text string,
) ILligneExpression {
	return &LligneMultilineStringLiteralExpr{
		LligneExpression: LligneExpression{sourcePos: sourcePos, exprType: ExprTypeMultilineStringLiteral},
		Text:             text,
	}
}

func (e *LligneMultilineStringLiteralExpr) SExpression() string {
	return "(multilinestr\n" + e.Text + ")"
}

//=====================================================================================================================

// LligneOptionalExpr represents a parenthesized expression or comma-separated sequence of expressions.
type LligneOptionalExpr struct {
	LligneExpression
	Operand ILligneExpression
}

func NewOptionalExpr(
	sourcePos int,
	operand ILligneExpression,
) ILligneExpression {
	return &LligneOptionalExpr{
		LligneExpression: LligneExpression{sourcePos: sourcePos, exprType: ExprTypeOptional},
		Operand:          operand,
	}
}

func (e *LligneOptionalExpr) SExpression() string {
	result := "(optional "
	result += e.Operand.SExpression()
	result += ")"
	return result
}

//=====================================================================================================================

// LligneParenthesizedExpr represents a parenthesized expression or comma-separated sequence of expressions.
type LligneParenthesizedExpr struct {
	LligneExpression
	Delimiters ParenExprDelimiters
	Items      []ILligneExpression
}

func NewParenthesizedExpr(
	sourcePos int,
	delimiters ParenExprDelimiters,
	items []ILligneExpression,
) ILligneExpression {
	return &LligneParenthesizedExpr{
		LligneExpression: LligneExpression{sourcePos: sourcePos, exprType: ExprTypeParenthesized},
		Delimiters:       delimiters,
		Items:            items,
	}
}

func (e *LligneParenthesizedExpr) SExpression() string {
	var result string

	switch e.Delimiters {
	case ParenExprDelimitersParentheses:
		result = "(parenthesized \"()\""
	case ParenExprDelimitersBraces:
		result = "(parenthesized \"{}\""
	case ParenExprDelimitersWholeFile:
		result = "(parenthesized \"\""
	}

	for _, item := range e.Items {
		result += " "
		result += item.SExpression()
	}

	result += ")"

	return result
}

//=====================================================================================================================

// LlignePrefixOperationExpr represents a prefix operation.
type LlignePrefixOperationExpr struct {
	LligneExpression
	Operator LlignePrefixOperator
	Operand  ILligneExpression
}

func NewPrefixOperationExpr(
	sourcePos int,
	operator LlignePrefixOperator,
	operand ILligneExpression,
) ILligneExpression {
	return &LlignePrefixOperationExpr{
		LligneExpression: LligneExpression{sourcePos: sourcePos, exprType: ExprTypePrefixOperation},
		Operator:         operator,
		Operand:          operand,
	}
}

func (e *LlignePrefixOperationExpr) SExpression() string {
	return "(prefix " + strings.TrimSpace(e.Operator.String()) + " " + e.Operand.SExpression() + ")"
}

//=====================================================================================================================

// LligneSequenceLiteralExpr represents a parenthesized expression or comma-separated sequence of expressions.
type LligneSequenceLiteralExpr struct {
	LligneExpression
	Elements []ILligneExpression
}

func NewSequenceLiteralExpr(
	sourcePos int,
	elements []ILligneExpression,
) ILligneExpression {
	return &LligneSequenceLiteralExpr{
		LligneExpression: LligneExpression{sourcePos: sourcePos, exprType: ExprTypeSequenceLiteral},
		Elements:         elements,
	}
}

func (e *LligneSequenceLiteralExpr) SExpression() string {
	result := "(sequence"

	for _, element := range e.Elements {
		result += " "
		result += element.SExpression()
	}

	result += ")"

	return result
}

//=====================================================================================================================

// LligneStringLiteralExpr represents a single string literal.
type LligneStringLiteralExpr struct {
	LligneExpression
	Text string
}

func NewStringLiteralExpr(
	sourcePos int,
	text string,
) ILligneExpression {
	return &LligneStringLiteralExpr{
		LligneExpression: LligneExpression{sourcePos: sourcePos, exprType: ExprTypeStringLiteral},
		Text:             text,
	}
}

func (e *LligneStringLiteralExpr) SExpression() string {
	return "(string " + e.Text + ")"
}

//=====================================================================================================================

// LligneTrailingDocumentationExpr represents lines of trailing documentation.
type LligneTrailingDocumentationExpr struct {
	LligneExpression
	Text string
}

func NewTrailingDocumentationExpr(
	sourcePos int,
	text string,
) ILligneExpression {
	return &LligneTrailingDocumentationExpr{
		LligneExpression: LligneExpression{sourcePos: sourcePos, exprType: ExprTypeTrailingDocumentation},
		Text:             text,
	}
}

func (e *LligneTrailingDocumentationExpr) SExpression() string {
	return "(trailingdoc\n" + e.Text + ")"
}

//=====================================================================================================================
