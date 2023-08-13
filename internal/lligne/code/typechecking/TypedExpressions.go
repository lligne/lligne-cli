//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package typechecking

import (
	"lligne-cli/internal/lligne/code/model"
)

//=====================================================================================================================

// ITypedExpression is the interface to an expression AST node with types added.
type ITypedExpression interface {
	isTypeExpression()
}

//=====================================================================================================================

// TypedAdditionExpr represents an addition operation.
type TypedAdditionExpr struct {
	SourcePosition model.SourcePos
	Lhs            ITypedExpression
	Rhs            ITypedExpression
	TypeInfo       IType
}

func (e *TypedAdditionExpr) isTypeExpression() {}

//=====================================================================================================================

// TypedBooleanLiteralExpr represents a single boolean literal.
type TypedBooleanLiteralExpr struct {
	SourcePosition model.SourcePos
	Value          bool
	TypeInfo       IType
}

func (e *TypedBooleanLiteralExpr) isTypeExpression() {}

//=====================================================================================================================

// TypedDivisionExpr represents a division operation.
type TypedDivisionExpr struct {
	SourcePosition model.SourcePos
	Lhs            ITypedExpression
	Rhs            ITypedExpression
	TypeInfo       IType
}

func (e *TypedDivisionExpr) isTypeExpression() {}

//=====================================================================================================================

// TypedEqualsExpr represents a equals operation.
type TypedEqualsExpr struct {
	SourcePosition model.SourcePos
	Lhs            ITypedExpression
	Rhs            ITypedExpression
	TypeInfo       IType
}

func (e *TypedEqualsExpr) isTypeExpression() {}

//=====================================================================================================================

// TypedFloatingPointLiteralExpr represents a single integer literal.
type TypedFloatingPointLiteralExpr struct {
	SourcePosition model.SourcePos
	Text           string
	TypeInfo       IType
}

func (e *TypedFloatingPointLiteralExpr) isTypeExpression() {}

//=====================================================================================================================

// TypedFunctionCallExpr represents a function call (a function name followed by a parenthesized expression).
type TypedFunctionCallExpr struct {
	SourcePosition    model.SourcePos
	FunctionReference ITypedExpression
	Argument          ITypedExpression
	TypeInfo          IType
}

func (e *TypedFunctionCallExpr) isTypeExpression() {}

//=====================================================================================================================

// TypedGreaterThanExpr represents a greater than operation.
type TypedGreaterThanExpr struct {
	SourcePosition model.SourcePos
	Lhs            ITypedExpression
	Rhs            ITypedExpression
	TypeInfo       IType
}

func (e *TypedGreaterThanExpr) isTypeExpression() {}

//=====================================================================================================================

// TypedGreaterThanOrEqualsExpr represents a greater than operation.
type TypedGreaterThanOrEqualsExpr struct {
	SourcePosition model.SourcePos
	Lhs            ITypedExpression
	Rhs            ITypedExpression
	TypeInfo       IType
}

func (e *TypedGreaterThanOrEqualsExpr) isTypeExpression() {}

//=====================================================================================================================

// TypedIdentifierExpr represents a single identifier.
type TypedIdentifierExpr struct {
	SourcePosition model.SourcePos
	Name           string
	TypeInfo       IType
}

func (e *TypedIdentifierExpr) isTypeExpression() {}

//=====================================================================================================================

// TypedIntegerLiteralExpr represents a single integer literal.
type TypedIntegerLiteralExpr struct {
	SourcePosition model.SourcePos
	Text           string
	TypeInfo       IType
}

func (e *TypedIntegerLiteralExpr) isTypeExpression() {}

//=====================================================================================================================

// TypedLeadingDocumentationExpr represents lines of leading documentation.
type TypedLeadingDocumentationExpr struct {
	SourcePosition model.SourcePos
	Text           string
	TypeInfo       IType
}

func (e *TypedLeadingDocumentationExpr) isTypeExpression() {}

//=====================================================================================================================

// TypedLessThanExpr represents a less than operation.
type TypedLessThanExpr struct {
	SourcePosition model.SourcePos
	Lhs            ITypedExpression
	Rhs            ITypedExpression
	TypeInfo       IType
}

func (e *TypedLessThanExpr) isTypeExpression() {}

//=====================================================================================================================

// TypedLessThanOrEqualsExpr represents a less than operation.
type TypedLessThanOrEqualsExpr struct {
	SourcePosition model.SourcePos
	Lhs            ITypedExpression
	Rhs            ITypedExpression
	TypeInfo       IType
}

func (e *TypedLessThanOrEqualsExpr) isTypeExpression() {}

//=====================================================================================================================

// TypedLogicalAndExpr represents a logical "and" operation.
type TypedLogicalAndExpr struct {
	SourcePosition model.SourcePos
	Lhs            ITypedExpression
	Rhs            ITypedExpression
	TypeInfo       IType
}

func (e *TypedLogicalAndExpr) isTypeExpression() {}

//=====================================================================================================================

// TypedLogicalNotOperationExpr represents a logical "not" operation.
type TypedLogicalNotOperationExpr struct {
	SourcePosition model.SourcePos
	Operand        ITypedExpression
	TypeInfo       IType
}

func (e *TypedLogicalNotOperationExpr) isTypeExpression() {}

//=====================================================================================================================

// TypedLogicalOrExpr represents a logical "or" operation.
type TypedLogicalOrExpr struct {
	SourcePosition model.SourcePos
	Lhs            ITypedExpression
	Rhs            ITypedExpression
	TypeInfo       IType
}

func (e *TypedLogicalOrExpr) isTypeExpression() {}

//=====================================================================================================================

// TypedMultilineStringLiteralExpr represents a multiline (back-ticked) string literal.
type TypedMultilineStringLiteralExpr struct {
	SourcePosition model.SourcePos
	Text           string
	TypeInfo       IType
}

func (e *TypedMultilineStringLiteralExpr) isTypeExpression() {}

//=====================================================================================================================

// TypedMultiplicationExpr represents a multiplication operation.
type TypedMultiplicationExpr struct {
	SourcePosition model.SourcePos
	Lhs            ITypedExpression
	Rhs            ITypedExpression
	TypeInfo       IType
}

func (e *TypedMultiplicationExpr) isTypeExpression() {}

//=====================================================================================================================

// TypedNegationOperationExpr represents an arithmetic negation operation.
type TypedNegationOperationExpr struct {
	SourcePosition model.SourcePos
	Operand        ITypedExpression
	TypeInfo       IType
}

func (e *TypedNegationOperationExpr) isTypeExpression() {}

//=====================================================================================================================

// TypedOptionalExpr represents a parenthesized expression or comma-separated sequence of expressions.
type TypedOptionalExpr struct {
	SourcePosition model.SourcePos
	Operand        ITypedExpression
	TypeInfo       IType
}

func (e *TypedOptionalExpr) isTypeExpression() {}

//=====================================================================================================================

// TypedParenthesizedExpr represents a parenthesized expression or comma-separated sequence of expressions.
type TypedParenthesizedExpr struct {
	SourcePosition model.SourcePos
	Delimiters     model.ParenExprDelimiters
	Items          []ITypedExpression
	TypeInfo       IType
}

func (e *TypedParenthesizedExpr) isTypeExpression() {}

//=====================================================================================================================

// TypedSequenceLiteralExpr represents a parenthesized expression or comma-separated sequence of expressions.
type TypedSequenceLiteralExpr struct {
	SourcePosition model.SourcePos
	Elements       []ITypedExpression
	TypeInfo       IType
}

func (e *TypedSequenceLiteralExpr) isTypeExpression() {}

//=====================================================================================================================

// TypedStringLiteralExpr represents a single string literal.
type TypedStringLiteralExpr struct {
	SourcePosition model.SourcePos
	Text           string
	TypeInfo       IType
}

func (e *TypedStringLiteralExpr) isTypeExpression() {}

//=====================================================================================================================

// TypedSubtractionExpr represents a subtraction operation.
type TypedSubtractionExpr struct {
	SourcePosition model.SourcePos
	Lhs            ITypedExpression
	Rhs            ITypedExpression
	TypeInfo       IType
}

func (e *TypedSubtractionExpr) isTypeExpression() {}

//=====================================================================================================================

// TypedTrailingDocumentationExpr represents lines of trailing documentation.
type TypedTrailingDocumentationExpr struct {
	SourcePosition model.SourcePos
	Text           string
	TypeInfo       IType
}

func (e *TypedTrailingDocumentationExpr) isTypeExpression() {}

//=====================================================================================================================
