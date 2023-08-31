//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package pooling

import (
	"lligne-cli/internal/lligne/code/parsing"
)

//=====================================================================================================================

// IExpression is the interface to an expression AST node with literal strings and identifier names pooled.
type IExpression interface {
	GetSourcePosition() parsing.SourcePos
	isPooledExpression()
}

//=====================================================================================================================

// AdditionExpr represents an addition operation.
type AdditionExpr struct {
	SourcePosition parsing.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *AdditionExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *AdditionExpr) isPooledExpression()                  {}

//=====================================================================================================================

// ArrayLiteralExpr represents a parenthesized expression or comma-separated sequence of expressions.
type ArrayLiteralExpr struct {
	SourcePosition parsing.SourcePos
	Elements       []IExpression
}

func (e *ArrayLiteralExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *ArrayLiteralExpr) isPooledExpression()                  {}

//=====================================================================================================================

// BooleanLiteralExpr represents a single boolean literal.
type BooleanLiteralExpr struct {
	SourcePosition parsing.SourcePos
	Value          bool
}

func (e *BooleanLiteralExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *BooleanLiteralExpr) isPooledExpression()                  {}

//=====================================================================================================================

// BuiltInTypeExpr represents a pre-defined base type.
type BuiltInTypeExpr struct {
	SourcePosition parsing.SourcePos
}

func (e *BuiltInTypeExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *BuiltInTypeExpr) isPooledExpression()                  {}

//=====================================================================================================================

// DivisionExpr represents a division operation.
type DivisionExpr struct {
	SourcePosition parsing.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *DivisionExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *DivisionExpr) isPooledExpression()                  {}

//=====================================================================================================================

// EqualsExpr represents a equals operation.
type EqualsExpr struct {
	SourcePosition parsing.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *EqualsExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *EqualsExpr) isPooledExpression()                  {}

//=====================================================================================================================

// Float64LiteralExpr represents a single 64-bit floating point literal.
type Float64LiteralExpr struct {
	SourcePosition parsing.SourcePos
	Value          float64
}

func (e *Float64LiteralExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *Float64LiteralExpr) isPooledExpression()                  {}

//=====================================================================================================================

// FunctionCallExpr represents a function call (a function name followed by a parenthesized expression).
type FunctionCallExpr struct {
	SourcePosition    parsing.SourcePos
	FunctionReference IExpression
	Argument          IExpression
}

func (e *FunctionCallExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *FunctionCallExpr) isPooledExpression()                  {}

//=====================================================================================================================

// GreaterThanExpr represents a greater than operation.
type GreaterThanExpr struct {
	SourcePosition parsing.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *GreaterThanExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *GreaterThanExpr) isPooledExpression()                  {}

//=====================================================================================================================

// GreaterThanOrEqualsExpr represents a greater than operation.
type GreaterThanOrEqualsExpr struct {
	SourcePosition parsing.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *GreaterThanOrEqualsExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *GreaterThanOrEqualsExpr) isPooledExpression()                  {}

//=====================================================================================================================

// IdentifierExpr represents a single identifier.
type IdentifierExpr struct {
	SourcePosition parsing.SourcePos
	NameIndex      uint64
}

func (e *IdentifierExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *IdentifierExpr) isPooledExpression()                  {}

//=====================================================================================================================

// Int64LiteralExpr represents a single 64-bit integer literal.
type Int64LiteralExpr struct {
	SourcePosition parsing.SourcePos
	Value          int64
}

func (e *Int64LiteralExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *Int64LiteralExpr) isPooledExpression()                  {}

//=====================================================================================================================

// IsExpr represents an "is" test.
type IsExpr struct {
	SourcePosition parsing.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *IsExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *IsExpr) isPooledExpression()                  {}

//=====================================================================================================================

// LeadingDocumentationExpr represents lines of leading documentation.
type LeadingDocumentationExpr struct {
	SourcePosition parsing.SourcePos
	Text           string
}

func (e *LeadingDocumentationExpr) GetSourcePosition() parsing.SourcePos {
	return e.SourcePosition
}
func (e *LeadingDocumentationExpr) isPooledExpression() {}

//=====================================================================================================================

// LessThanExpr represents a less than operation.
type LessThanExpr struct {
	SourcePosition parsing.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *LessThanExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *LessThanExpr) isPooledExpression()                  {}

//=====================================================================================================================

// LessThanOrEqualsExpr represents a less than operation.
type LessThanOrEqualsExpr struct {
	SourcePosition parsing.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *LessThanOrEqualsExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *LessThanOrEqualsExpr) isPooledExpression()                  {}

//=====================================================================================================================

// LogicalAndExpr represents a logical "and" operation.
type LogicalAndExpr struct {
	SourcePosition parsing.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *LogicalAndExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *LogicalAndExpr) isPooledExpression()                  {}

//=====================================================================================================================

// LogicalNotOperationExpr represents a logical "not" operation.
type LogicalNotOperationExpr struct {
	SourcePosition parsing.SourcePos
	Operand        IExpression
}

func (e *LogicalNotOperationExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *LogicalNotOperationExpr) isPooledExpression()                  {}

//=====================================================================================================================

// LogicalOrExpr represents a logical "or" operation.
type LogicalOrExpr struct {
	SourcePosition parsing.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *LogicalOrExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *LogicalOrExpr) isPooledExpression()                  {}

//=====================================================================================================================

// MultilineStringLiteralExpr represents a multiline (back-ticked) string literal.
type MultilineStringLiteralExpr struct {
	SourcePosition parsing.SourcePos
	Text           string
}

func (e *MultilineStringLiteralExpr) GetSourcePosition() parsing.SourcePos {
	return e.SourcePosition
}
func (e *MultilineStringLiteralExpr) isPooledExpression() {}

//=====================================================================================================================

// MultiplicationExpr represents a multiplication operation.
type MultiplicationExpr struct {
	SourcePosition parsing.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *MultiplicationExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *MultiplicationExpr) isPooledExpression()                  {}

//=====================================================================================================================

// NegationOperationExpr represents an arithmetic negation operation.
type NegationOperationExpr struct {
	SourcePosition parsing.SourcePos
	Operand        IExpression
}

func (e *NegationOperationExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *NegationOperationExpr) isPooledExpression()                  {}

//=====================================================================================================================

// NotEqualsExpr represents a equals operation.
type NotEqualsExpr struct {
	SourcePosition parsing.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *NotEqualsExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *NotEqualsExpr) isPooledExpression()                  {}

//=====================================================================================================================

// OptionalExpr represents a parenthesized expression or comma-separated sequence of expressions.
type OptionalExpr struct {
	SourcePosition parsing.SourcePos
	Operand        IExpression
}

func (e *OptionalExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *OptionalExpr) isPooledExpression()                  {}

//=====================================================================================================================

// ParenthesizedExpr represents a parenthesized expression or comma-separated sequence of expressions.
type ParenthesizedExpr struct {
	SourcePosition parsing.SourcePos
	InnerExpr      IExpression
}

func (e *ParenthesizedExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *ParenthesizedExpr) isPooledExpression()                  {}

//=====================================================================================================================

// RecordExpr represents a record.
type RecordExpr struct {
	SourcePosition parsing.SourcePos
}

func (e *RecordExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *RecordExpr) isPooledExpression()                  {}

//=====================================================================================================================

// StringConcatenationExpr represents concatenation of two strings.
type StringConcatenationExpr struct {
	SourcePosition parsing.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *StringConcatenationExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *StringConcatenationExpr) isPooledExpression()                  {}

//=====================================================================================================================

// StringLiteralExpr represents a single string literal.
type StringLiteralExpr struct {
	SourcePosition parsing.SourcePos
	ValueIndex     uint64
}

func (e *StringLiteralExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *StringLiteralExpr) isPooledExpression()                  {}

//=====================================================================================================================

// SubtractionExpr represents a subtraction operation.
type SubtractionExpr struct {
	SourcePosition parsing.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *SubtractionExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *SubtractionExpr) isPooledExpression()                  {}

//=====================================================================================================================

// TrailingDocumentationExpr represents lines of trailing documentation.
type TrailingDocumentationExpr struct {
	SourcePosition parsing.SourcePos
	Text           string
}

func (e *TrailingDocumentationExpr) GetSourcePosition() parsing.SourcePos {
	return e.SourcePosition
}
func (e *TrailingDocumentationExpr) isPooledExpression() {}

//=====================================================================================================================
