//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package typechecking

import (
	"lligne-cli/internal/lligne/code/parsing"
	"lligne-cli/internal/lligne/runtime/types"
)

//=====================================================================================================================

// ITypedExpression is the interface to an expression AST node with types added.
type ITypedExpression interface {
	GetSourcePosition() parsing.SourcePos
	GetTypeInfo() types.IType
	isTypeExpression()
}

//=====================================================================================================================

// TypedAdditionExpr represents an addition operation.
type TypedAdditionExpr struct {
	SourcePosition parsing.SourcePos
	Lhs            ITypedExpression
	Rhs            ITypedExpression
	TypeInfo       types.IType
}

func (e *TypedAdditionExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *TypedAdditionExpr) GetTypeInfo() types.IType             { return e.TypeInfo }
func (e *TypedAdditionExpr) isTypeExpression()                    {}

//=====================================================================================================================

// TypedBooleanLiteralExpr represents a single boolean literal.
type TypedBooleanLiteralExpr struct {
	SourcePosition parsing.SourcePos
	Value          bool
}

func (e *TypedBooleanLiteralExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *TypedBooleanLiteralExpr) GetTypeInfo() types.IType             { return types.BoolTypeInstance }
func (e *TypedBooleanLiteralExpr) isTypeExpression()                    {}

//=====================================================================================================================

// TypedDivisionExpr represents a division operation.
type TypedDivisionExpr struct {
	SourcePosition parsing.SourcePos
	Lhs            ITypedExpression
	Rhs            ITypedExpression
	TypeInfo       types.IType
}

func (e *TypedDivisionExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *TypedDivisionExpr) GetTypeInfo() types.IType             { return e.TypeInfo }
func (e *TypedDivisionExpr) isTypeExpression()                    {}

//=====================================================================================================================

// TypedEqualsExpr represents a equals operation.
type TypedEqualsExpr struct {
	SourcePosition parsing.SourcePos
	Lhs            ITypedExpression
	Rhs            ITypedExpression
}

func (e *TypedEqualsExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *TypedEqualsExpr) GetTypeInfo() types.IType             { return types.BoolTypeInstance }
func (e *TypedEqualsExpr) isTypeExpression()                    {}

//=====================================================================================================================

// TypedFloat64LiteralExpr represents a single 64-bit floating point literal.
type TypedFloat64LiteralExpr struct {
	SourcePosition parsing.SourcePos
	Value          float64
}

func (e *TypedFloat64LiteralExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *TypedFloat64LiteralExpr) GetTypeInfo() types.IType             { return types.Float64TypeInstance }
func (e *TypedFloat64LiteralExpr) isTypeExpression()                    {}

//=====================================================================================================================

// TypedFunctionCallExpr represents a function call (a function name followed by a parenthesized expression).
type TypedFunctionCallExpr struct {
	SourcePosition    parsing.SourcePos
	FunctionReference ITypedExpression
	Argument          ITypedExpression
	TypeInfo          types.IType
}

func (e *TypedFunctionCallExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *TypedFunctionCallExpr) GetTypeInfo() types.IType             { return e.TypeInfo }
func (e *TypedFunctionCallExpr) isTypeExpression()                    {}

//=====================================================================================================================

// TypedGreaterThanExpr represents a greater than operation.
type TypedGreaterThanExpr struct {
	SourcePosition parsing.SourcePos
	Lhs            ITypedExpression
	Rhs            ITypedExpression
}

func (e *TypedGreaterThanExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *TypedGreaterThanExpr) GetTypeInfo() types.IType             { return types.BoolTypeInstance }
func (e *TypedGreaterThanExpr) isTypeExpression()                    {}

//=====================================================================================================================

// TypedGreaterThanOrEqualsExpr represents a greater than operation.
type TypedGreaterThanOrEqualsExpr struct {
	SourcePosition parsing.SourcePos
	Lhs            ITypedExpression
	Rhs            ITypedExpression
}

func (e *TypedGreaterThanOrEqualsExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *TypedGreaterThanOrEqualsExpr) GetTypeInfo() types.IType             { return types.BoolTypeInstance }
func (e *TypedGreaterThanOrEqualsExpr) isTypeExpression()                    {}

//=====================================================================================================================

// TypedIdentifierExpr represents a single identifier.
type TypedIdentifierExpr struct {
	SourcePosition parsing.SourcePos
	Name           string
	TypeInfo       types.IType
}

func (e *TypedIdentifierExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *TypedIdentifierExpr) GetTypeInfo() types.IType             { return e.TypeInfo }
func (e *TypedIdentifierExpr) isTypeExpression()                    {}

//=====================================================================================================================

// TypedInt64LiteralExpr represents a single 64-bit integer literal.
type TypedInt64LiteralExpr struct {
	SourcePosition parsing.SourcePos
	Value          int64
}

func (e *TypedInt64LiteralExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *TypedInt64LiteralExpr) GetTypeInfo() types.IType             { return types.Int64TypeInstance }
func (e *TypedInt64LiteralExpr) isTypeExpression()                    {}

//=====================================================================================================================

// TypedLeadingDocumentationExpr represents lines of leading documentation.
type TypedLeadingDocumentationExpr struct {
	SourcePosition parsing.SourcePos
	Text           string
	TypeInfo       types.IType
}

func (e *TypedLeadingDocumentationExpr) GetSourcePosition() parsing.SourcePos {
	return e.SourcePosition
}
func (e *TypedLeadingDocumentationExpr) GetTypeInfo() types.IType { return e.TypeInfo }
func (e *TypedLeadingDocumentationExpr) isTypeExpression()        {}

//=====================================================================================================================

// TypedLessThanExpr represents a less than operation.
type TypedLessThanExpr struct {
	SourcePosition parsing.SourcePos
	Lhs            ITypedExpression
	Rhs            ITypedExpression
}

func (e *TypedLessThanExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *TypedLessThanExpr) GetTypeInfo() types.IType             { return types.BoolTypeInstance }
func (e *TypedLessThanExpr) isTypeExpression()                    {}

//=====================================================================================================================

// TypedLessThanOrEqualsExpr represents a less than operation.
type TypedLessThanOrEqualsExpr struct {
	SourcePosition parsing.SourcePos
	Lhs            ITypedExpression
	Rhs            ITypedExpression
}

func (e *TypedLessThanOrEqualsExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *TypedLessThanOrEqualsExpr) GetTypeInfo() types.IType             { return types.BoolTypeInstance }
func (e *TypedLessThanOrEqualsExpr) isTypeExpression()                    {}

//=====================================================================================================================

// TypedLogicalAndExpr represents a logical "and" operation.
type TypedLogicalAndExpr struct {
	SourcePosition parsing.SourcePos
	Lhs            ITypedExpression
	Rhs            ITypedExpression
}

func (e *TypedLogicalAndExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *TypedLogicalAndExpr) GetTypeInfo() types.IType             { return types.BoolTypeInstance }
func (e *TypedLogicalAndExpr) isTypeExpression()                    {}

//=====================================================================================================================

// TypedLogicalNotOperationExpr represents a logical "not" operation.
type TypedLogicalNotOperationExpr struct {
	SourcePosition parsing.SourcePos
	Operand        ITypedExpression
}

func (e *TypedLogicalNotOperationExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *TypedLogicalNotOperationExpr) GetTypeInfo() types.IType             { return types.BoolTypeInstance }
func (e *TypedLogicalNotOperationExpr) isTypeExpression()                    {}

//=====================================================================================================================

// TypedLogicalOrExpr represents a logical "or" operation.
type TypedLogicalOrExpr struct {
	SourcePosition parsing.SourcePos
	Lhs            ITypedExpression
	Rhs            ITypedExpression
}

func (e *TypedLogicalOrExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *TypedLogicalOrExpr) GetTypeInfo() types.IType             { return types.BoolTypeInstance }
func (e *TypedLogicalOrExpr) isTypeExpression()                    {}

//=====================================================================================================================

// TypedMultilineStringLiteralExpr represents a multiline (back-ticked) string literal.
type TypedMultilineStringLiteralExpr struct {
	SourcePosition parsing.SourcePos
	Text           string
	TypeInfo       types.IType
}

func (e *TypedMultilineStringLiteralExpr) GetSourcePosition() parsing.SourcePos {
	return e.SourcePosition
}
func (e *TypedMultilineStringLiteralExpr) GetTypeInfo() types.IType { return e.TypeInfo }
func (e *TypedMultilineStringLiteralExpr) isTypeExpression()        {}

//=====================================================================================================================

// TypedMultiplicationExpr represents a multiplication operation.
type TypedMultiplicationExpr struct {
	SourcePosition parsing.SourcePos
	Lhs            ITypedExpression
	Rhs            ITypedExpression
	TypeInfo       types.IType
}

func (e *TypedMultiplicationExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *TypedMultiplicationExpr) GetTypeInfo() types.IType             { return e.TypeInfo }
func (e *TypedMultiplicationExpr) isTypeExpression()                    {}

//=====================================================================================================================

// TypedNegationOperationExpr represents an arithmetic negation operation.
type TypedNegationOperationExpr struct {
	SourcePosition parsing.SourcePos
	Operand        ITypedExpression
	TypeInfo       types.IType
}

func (e *TypedNegationOperationExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *TypedNegationOperationExpr) GetTypeInfo() types.IType             { return e.TypeInfo }
func (e *TypedNegationOperationExpr) isTypeExpression()                    {}

//=====================================================================================================================

// TypedOptionalExpr represents a parenthesized expression or comma-separated sequence of expressions.
type TypedOptionalExpr struct {
	SourcePosition parsing.SourcePos
	Operand        ITypedExpression
	TypeInfo       types.IType
}

func (e *TypedOptionalExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *TypedOptionalExpr) GetTypeInfo() types.IType             { return e.TypeInfo }
func (e *TypedOptionalExpr) isTypeExpression()                    {}

//=====================================================================================================================

// TypedParenthesizedExpr represents a parenthesized expression or comma-separated sequence of expressions.
type TypedParenthesizedExpr struct {
	SourcePosition parsing.SourcePos
	Delimiters     parsing.ParenExprDelimiters
	Items          []ITypedExpression
	TypeInfo       types.IType
}

func (e *TypedParenthesizedExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *TypedParenthesizedExpr) GetTypeInfo() types.IType             { return e.TypeInfo }
func (e *TypedParenthesizedExpr) isTypeExpression()                    {}

//=====================================================================================================================

// TypedSequenceLiteralExpr represents a parenthesized expression or comma-separated sequence of expressions.
type TypedSequenceLiteralExpr struct {
	SourcePosition parsing.SourcePos
	Elements       []ITypedExpression
	TypeInfo       types.IType
}

func (e *TypedSequenceLiteralExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *TypedSequenceLiteralExpr) GetTypeInfo() types.IType             { return e.TypeInfo }
func (e *TypedSequenceLiteralExpr) isTypeExpression()                    {}

//=====================================================================================================================

// TypedStringConcatenationExpr represents concatenation of two strings.
type TypedStringConcatenationExpr struct {
	SourcePosition parsing.SourcePos
	Lhs            ITypedExpression
	Rhs            ITypedExpression
}

func (e *TypedStringConcatenationExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *TypedStringConcatenationExpr) GetTypeInfo() types.IType             { return types.StringTypeInstance }
func (e *TypedStringConcatenationExpr) isTypeExpression()                    {}

//=====================================================================================================================

// TypedStringLiteralExpr represents a single string literal.
type TypedStringLiteralExpr struct {
	SourcePosition parsing.SourcePos
	Value          string
}

func (e *TypedStringLiteralExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *TypedStringLiteralExpr) GetTypeInfo() types.IType             { return types.StringTypeInstance }
func (e *TypedStringLiteralExpr) isTypeExpression()                    {}

//=====================================================================================================================

// TypedSubtractionExpr represents a subtraction operation.
type TypedSubtractionExpr struct {
	SourcePosition parsing.SourcePos
	Lhs            ITypedExpression
	Rhs            ITypedExpression
	TypeInfo       types.IType
}

func (e *TypedSubtractionExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *TypedSubtractionExpr) GetTypeInfo() types.IType             { return e.TypeInfo }
func (e *TypedSubtractionExpr) isTypeExpression()                    {}

//=====================================================================================================================

// TypedTrailingDocumentationExpr represents lines of trailing documentation.
type TypedTrailingDocumentationExpr struct {
	SourcePosition parsing.SourcePos
	Text           string
	TypeInfo       types.IType
}

func (e *TypedTrailingDocumentationExpr) GetSourcePosition() parsing.SourcePos {
	return e.SourcePosition
}
func (e *TypedTrailingDocumentationExpr) GetTypeInfo() types.IType { return e.TypeInfo }
func (e *TypedTrailingDocumentationExpr) isTypeExpression()        {}

//=====================================================================================================================
