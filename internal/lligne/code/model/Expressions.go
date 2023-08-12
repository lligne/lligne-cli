//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package model

import (
	"strconv"
)

//=====================================================================================================================

// IExpression is the interface to an expression AST node.
type IExpression interface {
	isExpression()
}

//=====================================================================================================================

// TODO: get rid of this

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

// TODO: get rid of this

// ParenExprDelimiters is an enumeration of start/stop delimiters for parenthesized expressions.
type ParenExprDelimiters int

const (
	ParenExprDelimitersParentheses ParenExprDelimiters = 1 + iota
	ParenExprDelimitersBraces
	ParenExprDelimitersWholeFile
)

//=====================================================================================================================

// TODO: get rid of this

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

// BooleanLiteralExpr represents a single boolean literal.
type BooleanLiteralExpr struct {
	SourcePosition SourcePos
	Value          bool
}

func (e *BooleanLiteralExpr) isExpression() {}

//=====================================================================================================================

// FloatingPointLiteralExpr represents a single integer literal.
type FloatingPointLiteralExpr struct {
	SourcePosition SourcePos
	Text           string
}

func (e *FloatingPointLiteralExpr) isExpression() {}

//=====================================================================================================================

// FunctionCallExpr represents a function call (a function name followed by a parenthesized expression).
type FunctionCallExpr struct {
	SourcePosition    SourcePos
	FunctionReference IExpression
	Argument          IExpression
}

func (e *FunctionCallExpr) isExpression() {}

//=====================================================================================================================

// IdentifierExpr represents a single identifier.
type IdentifierExpr struct {
	SourcePosition SourcePos
	Name           string
}

func (e *IdentifierExpr) isExpression() {}

//=====================================================================================================================

// InfixOperationExpr represents an infix operation.
type InfixOperationExpr struct {
	SourcePosition SourcePos
	Operator       InfixOperator
	Lhs            IExpression
	Rhs            IExpression
}

func (e *InfixOperationExpr) isExpression() {}

//=====================================================================================================================

// IntegerLiteralExpr represents a single integer literal.
type IntegerLiteralExpr struct {
	SourcePosition SourcePos
	Text           string
}

func (e *IntegerLiteralExpr) isExpression() {}

//=====================================================================================================================

// LeadingDocumentationExpr represents lines of leading documentation.
type LeadingDocumentationExpr struct {
	SourcePosition SourcePos
	Text           string
}

func (e *LeadingDocumentationExpr) isExpression() {}

//=====================================================================================================================

// MultilineStringLiteralExpr represents a multiline (back-ticked) string literal.
type MultilineStringLiteralExpr struct {
	SourcePosition SourcePos
	Text           string
}

func (e *MultilineStringLiteralExpr) isExpression() {}

//=====================================================================================================================

// OptionalExpr represents a parenthesized expression or comma-separated sequence of expressions.
type OptionalExpr struct {
	SourcePosition SourcePos
	Operand        IExpression
}

func (e *OptionalExpr) isExpression() {}

//=====================================================================================================================

// ParenthesizedExpr represents a parenthesized expression or comma-separated sequence of expressions.
type ParenthesizedExpr struct {
	SourcePosition SourcePos
	Delimiters     ParenExprDelimiters
	Items          []IExpression
}

func (e *ParenthesizedExpr) isExpression() {}

//=====================================================================================================================

// PrefixOperationExpr represents a prefix operation.
type PrefixOperationExpr struct {
	SourcePosition SourcePos
	Operator       PrefixOperator
	Operand        IExpression
}

func (e *PrefixOperationExpr) isExpression() {}

//=====================================================================================================================

// SequenceLiteralExpr represents a parenthesized expression or comma-separated sequence of expressions.
type SequenceLiteralExpr struct {
	SourcePosition SourcePos
	Elements       []IExpression
}

func (e *SequenceLiteralExpr) isExpression() {}

//=====================================================================================================================

// StringLiteralExpr represents a single string literal.
type StringLiteralExpr struct {
	SourcePosition SourcePos
	Text           string
}

func (e *StringLiteralExpr) isExpression() {}

//=====================================================================================================================

// TrailingDocumentationExpr represents lines of trailing documentation.
type TrailingDocumentationExpr struct {
	SourcePosition SourcePos
	Text           string
}

func (e *TrailingDocumentationExpr) isExpression() {}

//=====================================================================================================================
