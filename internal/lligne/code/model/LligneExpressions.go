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

//---------------------------------------------------------------------------------------------------------------------

// ILligneExpression is the interface to an expression AST node.
type ILligneExpression interface {
	GetOrigin(tracker scanning.ILligneTokenOriginTracker) scanning.LligneOrigin
	SExpression() string
	TypeCode() LligneExprType
}

//---------------------------------------------------------------------------------------------------------------------

// LligneExprType is an enumeration of Lligne expression types.
type LligneExprType int

const (
	ExprTypeIdentifier LligneExprType = iota
	ExprTypeInfixOperation
	ExprTypeIntegerLiteral
	ExprTypeLeadingDocumentation
	ExprTypeMultilineStringLiteral
	ExprTypeStringLiteral
)

//---------------------------------------------------------------------------------------------------------------------

// LligneInfixOperator is an enumeration of Lligne's binary operators.
type LligneInfixOperator int

const (
	InfixOperatorNone LligneInfixOperator = iota
	InfixOperatorAdd
	InfixOperatorDivide
	InfixOperatorDocument
	InfixOperatorEquality
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
)

//---------------------------------------------------------------------------------------------------------------------

// LlignePrefixOperator is an enumeration of Lligne's prefix operators.
type LlignePrefixOperator int

const (
	PrefixOperatorNone LlignePrefixOperator = iota
	PrefixOperatorLogicalNot
	PrefixOperatorNegation
)

//---------------------------------------------------------------------------------------------------------------------

// LlignePostfixOperator is an enumeration of Lligne's prefix operators.
type LlignePostfixOperator int

const (
	PostfixOperatorNone LlignePostfixOperator = iota
	PostfixOperatorFunctionCall
	PostfixOperatorIndex
	PostfixOperatorOptional
)

// ---------------------------------------------------------------------------------------------------------------------

// TextOfTokenType returns a string describing a Lligne token type.
func (op LligneInfixOperator) String() string {

	switch op {

	case InfixOperatorAdd:
		return " + "
	case InfixOperatorDivide:
		return " / "
	case InfixOperatorDocument:
		return " "
	case InfixOperatorEquality:
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

	}

	panic("Unhandled binary operator: '" + strconv.Itoa(int(op)) + "'.")
}

//---------------------------------------------------------------------------------------------------------------------

// LligneIdentifierExpr represents a single identifier.
type LligneIdentifierExpr struct {
	SourcePos int
	Name      string
}

func (e *LligneIdentifierExpr) GetOrigin(tracker scanning.ILligneTokenOriginTracker) scanning.LligneOrigin {
	return tracker.GetOrigin(e.SourcePos)
}

func (e *LligneIdentifierExpr) SExpression() string {
	return "(identifier " + e.Name + ")"
}

func (e *LligneIdentifierExpr) TypeCode() LligneExprType {
	return ExprTypeIdentifier
}

//---------------------------------------------------------------------------------------------------------------------

// LligneInfixOperationExpr represents an infix operation.
type LligneInfixOperationExpr struct {
	SourcePos int
	Operator  LligneInfixOperator
	Operands  []ILligneExpression
}

func (e *LligneInfixOperationExpr) GetOrigin(tracker scanning.ILligneTokenOriginTracker) scanning.LligneOrigin {
	return tracker.GetOrigin(e.SourcePos)
}

func (e *LligneInfixOperationExpr) SExpression() string {
	result := "(" + strings.TrimSpace(e.Operator.String())

	for _, operand := range e.Operands {
		result += " "
		result += operand.SExpression()
	}

	result += ")"

	return result
}

func (e *LligneInfixOperationExpr) TypeCode() LligneExprType {
	return ExprTypeInfixOperation
}

//---------------------------------------------------------------------------------------------------------------------

// LligneIntegerLiteralExpr represents a single integer literal.
type LligneIntegerLiteralExpr struct {
	SourcePos int
	Text      string
}

func (e *LligneIntegerLiteralExpr) GetOrigin(tracker scanning.ILligneTokenOriginTracker) scanning.LligneOrigin {
	return tracker.GetOrigin(e.SourcePos)
}

func (e *LligneIntegerLiteralExpr) SExpression() string {
	return "(int " + e.Text + ")"
}

func (e *LligneIntegerLiteralExpr) TypeCode() LligneExprType {
	return ExprTypeIntegerLiteral
}

//---------------------------------------------------------------------------------------------------------------------

// LligneLeadingDocumentationExpr represents lines of leading documentation.
type LligneLeadingDocumentationExpr struct {
	SourcePos int
	Text      string
}

func (e *LligneLeadingDocumentationExpr) GetOrigin(tracker scanning.ILligneTokenOriginTracker) scanning.LligneOrigin {
	return tracker.GetOrigin(e.SourcePos)
}

func (e *LligneLeadingDocumentationExpr) SExpression() string {
	return "(leadingdoc\n" + e.Text + ")"
}

func (e *LligneLeadingDocumentationExpr) TypeCode() LligneExprType {
	return ExprTypeLeadingDocumentation
}

//---------------------------------------------------------------------------------------------------------------------

// LligneMultilineStringLiteralExpr represents a single integer literal.
type LligneMultilineStringLiteralExpr struct {
	SourcePos int
	Text      string
}

func (e *LligneMultilineStringLiteralExpr) GetOrigin(tracker scanning.ILligneTokenOriginTracker) scanning.LligneOrigin {
	return tracker.GetOrigin(e.SourcePos)
}

func (e *LligneMultilineStringLiteralExpr) SExpression() string {
	return "(multilinestr\n" + e.Text + ")"
}

func (e *LligneMultilineStringLiteralExpr) TypeCode() LligneExprType {
	return ExprTypeMultilineStringLiteral
}

//---------------------------------------------------------------------------------------------------------------------

// LligneStringLiteralExpr represents a single integer literal.
type LligneStringLiteralExpr struct {
	SourcePos int
	Text      string
}

func (e *LligneStringLiteralExpr) GetOrigin(tracker scanning.ILligneTokenOriginTracker) scanning.LligneOrigin {
	return tracker.GetOrigin(e.SourcePos)
}

func (e *LligneStringLiteralExpr) SExpression() string {
	return "(string " + e.Text + ")"
}

func (e *LligneStringLiteralExpr) TypeCode() LligneExprType {
	return ExprTypeStringLiteral
}

//---------------------------------------------------------------------------------------------------------------------
