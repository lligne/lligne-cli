//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package model

import (
	"strconv"
	"strings"
)

//=====================================================================================================================

// IExpression is the interface to an expression AST node.
type IExpression interface {
	SExpression() string
}

//=====================================================================================================================

// InfixOperator is an enumeration of 's binary operators.
type InfixOperator int

const (
	InfixOperatorNone InfixOperator = iota
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
func (op InfixOperator) String() string {

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

// PrefixOperator is an enumeration of 's prefix operators.
type PrefixOperator int

const (
	PrefixOperatorLogicalNot PrefixOperator = 1 + iota
	PrefixOperatorNegation
)

// ---------------------------------------------------------------------------------------------------------------------

// String returns a string representing the code of an operator.
func (op PrefixOperator) String() string {

	switch op {

	case PrefixOperatorLogicalNot:
		return "not "
	case PrefixOperatorNegation:
		return "-"
	}

	panic("Unhandled prefix operator: '" + strconv.Itoa(int(op)) + "'.")
}

//=====================================================================================================================

// PostfixOperator is an enumeration of 's prefix operators.
type PostfixOperator int

const (
	PostfixOperatorNone PostfixOperator = iota
	PostfixOperatorFunctionCall
	PostfixOperatorIndex
	PostfixOperatorOptional
)

//=====================================================================================================================

// TODO: Get rid of this

// Expression represents any expression
type Expression struct {
	SourcePosition SourcePos
}

//=====================================================================================================================

// BooleanLiteralExpr represents a single boolean literal.
type BooleanLiteralExpr struct {
	Expression
	Value bool
}

func NewBooleanLiteralExpr(
	sourcePosition SourcePos,
	value bool,
) IExpression {
	return &BooleanLiteralExpr{
		Expression: Expression{SourcePosition: sourcePosition},
		Value:      value,
	}
}

func (e *BooleanLiteralExpr) SExpression() string {
	if e.Value {
		return "(bool true)"
	}
	return "(bool false)"
}

//=====================================================================================================================

// FloatingPointLiteralExpr represents a single integer literal.
type FloatingPointLiteralExpr struct {
	Expression
	Text string
}

func NewFloatingPointLiteralExpr(
	sourcePosition SourcePos,
	text string,
) IExpression {
	return &FloatingPointLiteralExpr{
		Expression: Expression{SourcePosition: sourcePosition},
		Text:       text,
	}
}

func (e *FloatingPointLiteralExpr) SExpression() string {
	return "(float " + e.Text + ")"
}

//=====================================================================================================================

// FunctionCallExpr represents a function call (a function name followed by a parenthesized expression).
type FunctionCallExpr struct {
	Expression
	FunctionReference IExpression
	Argument          IExpression
}

func NewFunctionCallExpr(
	sourcePosition SourcePos,
	functionReference IExpression,
	argument IExpression,
) IExpression {
	return &FunctionCallExpr{
		Expression:        Expression{SourcePosition: sourcePosition},
		FunctionReference: functionReference,
		Argument:          argument,
	}
}

func (e *FunctionCallExpr) SExpression() string {
	result := "(call "
	result += e.FunctionReference.SExpression()
	result += " "
	result += e.Argument.SExpression()
	result += ")"
	return result
}

//=====================================================================================================================

// IdentifierExpr represents a single identifier.
type IdentifierExpr struct {
	Expression
	Name string
}

func NewIdentifierExpr(
	sourcePosition SourcePos,
	name string,
) IExpression {
	return &IdentifierExpr{
		Expression: Expression{SourcePosition: sourcePosition},
		Name:       name,
	}
}

func (e *IdentifierExpr) SExpression() string {
	return "(id " + e.Name + ")"
}

//=====================================================================================================================

// InfixOperationExpr represents an infix operation.
type InfixOperationExpr struct {
	Expression
	Operator InfixOperator
	Lhs      IExpression
	Rhs      IExpression
}

func NewInfixOperationExpr(
	sourcePosition SourcePos,
	operator InfixOperator,
	lhs IExpression,
	rhs IExpression,
) IExpression {
	return &InfixOperationExpr{
		Expression: Expression{SourcePosition: sourcePosition},
		Operator:   operator,
		Lhs:        lhs,
		Rhs:        rhs,
	}
}

func (e *InfixOperationExpr) SExpression() string {
	return "(" + strings.TrimSpace(e.Operator.String()) + " " +
		e.Lhs.SExpression() + " " + e.Rhs.SExpression() + ")"
}

//=====================================================================================================================

// IntegerLiteralExpr represents a single integer literal.
type IntegerLiteralExpr struct {
	Expression
	Text string
}

func NewIntegerLiteralExpr(
	sourcePosition SourcePos,
	text string,
) IExpression {
	return &IntegerLiteralExpr{
		Expression: Expression{SourcePosition: sourcePosition},
		Text:       text,
	}
}

func (e *IntegerLiteralExpr) SExpression() string {
	return "(int " + e.Text + ")"
}

//=====================================================================================================================

// LeadingDocumentationExpr represents lines of leading documentation.
type LeadingDocumentationExpr struct {
	Expression
	Text string
}

func NewLeadingDocumentationExpr(
	sourcePosition SourcePos,
	text string,
) IExpression {
	return &LeadingDocumentationExpr{
		Expression: Expression{SourcePosition: sourcePosition},
		Text:       text,
	}
}

func (e *LeadingDocumentationExpr) SExpression() string {
	return "(leadingdoc\n" + e.Text + ")"
}

//=====================================================================================================================

// MultilineStringLiteralExpr represents a multiline (back-ticked) string literal.
type MultilineStringLiteralExpr struct {
	Expression
	Text string
}

func NewMultilineStringLiteralExpr(
	sourcePosition SourcePos,
	text string,
) IExpression {
	return &MultilineStringLiteralExpr{
		Expression: Expression{SourcePosition: sourcePosition},
		Text:       text,
	}
}

func (e *MultilineStringLiteralExpr) SExpression() string {
	return "(multilinestr\n" + e.Text + ")"
}

//=====================================================================================================================

// OptionalExpr represents a parenthesized expression or comma-separated sequence of expressions.
type OptionalExpr struct {
	Expression
	Operand IExpression
}

func NewOptionalExpr(
	sourcePosition SourcePos,
	operand IExpression,
) IExpression {
	return &OptionalExpr{
		Expression: Expression{SourcePosition: sourcePosition},
		Operand:    operand,
	}
}

func (e *OptionalExpr) SExpression() string {
	result := "(optional "
	result += e.Operand.SExpression()
	result += ")"
	return result
}

//=====================================================================================================================

// ParenthesizedExpr represents a parenthesized expression or comma-separated sequence of expressions.
type ParenthesizedExpr struct {
	Expression
	Delimiters ParenExprDelimiters
	Items      []IExpression
}

func NewParenthesizedExpr(
	sourcePosition SourcePos,
	delimiters ParenExprDelimiters,
	items []IExpression,
) IExpression {
	return &ParenthesizedExpr{
		Expression: Expression{SourcePosition: sourcePosition},
		Delimiters: delimiters,
		Items:      items,
	}
}

func (e *ParenthesizedExpr) SExpression() string {
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

// PrefixOperationExpr represents a prefix operation.
type PrefixOperationExpr struct {
	Expression
	Operator PrefixOperator
	Operand  IExpression
}

func NewPrefixOperationExpr(
	sourcePosition SourcePos,
	operator PrefixOperator,
	operand IExpression,
) IExpression {
	return &PrefixOperationExpr{
		Expression: Expression{SourcePosition: sourcePosition},
		Operator:   operator,
		Operand:    operand,
	}
}

func (e *PrefixOperationExpr) SExpression() string {
	return "(prefix " + strings.TrimSpace(e.Operator.String()) + " " + e.Operand.SExpression() + ")"
}

//=====================================================================================================================

// SequenceLiteralExpr represents a parenthesized expression or comma-separated sequence of expressions.
type SequenceLiteralExpr struct {
	Expression
	Elements []IExpression
}

func NewSequenceLiteralExpr(
	sourcePosition SourcePos,
	elements []IExpression,
) IExpression {
	return &SequenceLiteralExpr{
		Expression: Expression{SourcePosition: sourcePosition},
		Elements:   elements,
	}
}

func (e *SequenceLiteralExpr) SExpression() string {
	result := "(sequence"

	for _, element := range e.Elements {
		result += " "
		result += element.SExpression()
	}

	result += ")"

	return result
}

//=====================================================================================================================

// StringLiteralExpr represents a single string literal.
type StringLiteralExpr struct {
	Expression
	Text string
}

func NewStringLiteralExpr(
	sourcePosition SourcePos,
	text string,
) IExpression {
	return &StringLiteralExpr{
		Expression: Expression{SourcePosition: sourcePosition},
		Text:       text,
	}
}

func (e *StringLiteralExpr) SExpression() string {
	return "(string " + e.Text + ")"
}

//=====================================================================================================================

// TrailingDocumentationExpr represents lines of trailing documentation.
type TrailingDocumentationExpr struct {
	Expression
	Text string
}

func NewTrailingDocumentationExpr(
	sourcePosition SourcePos,
	text string,
) IExpression {
	return &TrailingDocumentationExpr{
		Expression: Expression{SourcePosition: sourcePosition},
		Text:       text,
	}
}

func (e *TrailingDocumentationExpr) SExpression() string {
	return "(trailingdoc\n" + e.Text + ")"
}

//=====================================================================================================================
