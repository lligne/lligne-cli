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

// TypedBooleanLiteralExpr represents a single boolean literal.
type TypedBooleanLiteralExpr struct {
	SourcePosition model.SourcePos
	Value          bool
	TypeInfo       IType
}

func (e *TypedBooleanLiteralExpr) isTypeExpression() {}

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

// TypedIdentifierExpr represents a single identifier.
type TypedIdentifierExpr struct {
	SourcePosition model.SourcePos
	Name           string
	TypeInfo       IType
}

func (e *TypedIdentifierExpr) isTypeExpression() {}

//=====================================================================================================================

// TypedInfixOperationExpr represents an infix operation.
type TypedInfixOperationExpr struct {
	SourcePosition model.SourcePos
	Operator       model.InfixOperator
	Lhs            ITypedExpression
	Rhs            ITypedExpression
	TypeInfo       IType
}

func (e *TypedInfixOperationExpr) isTypeExpression() {}

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

// TypedMultilineStringLiteralExpr represents a multiline (back-ticked) string literal.
type TypedMultilineStringLiteralExpr struct {
	SourcePosition model.SourcePos
	Text           string
	TypeInfo       IType
}

func (e *TypedMultilineStringLiteralExpr) isTypeExpression() {}

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

// TypedPrefixOperationExpr represents a prefix operation.
type TypedPrefixOperationExpr struct {
	SourcePosition model.SourcePos
	Operator       model.PrefixOperator
	Operand        ITypedExpression
	TypeInfo       IType
}

func (e *TypedPrefixOperationExpr) isTypeExpression() {}

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

// TypedTrailingDocumentationExpr represents lines of trailing documentation.
type TypedTrailingDocumentationExpr struct {
	SourcePosition model.SourcePos
	Text           string
	TypeInfo       IType
}

func (e *TypedTrailingDocumentationExpr) isTypeExpression() {}

//=====================================================================================================================
