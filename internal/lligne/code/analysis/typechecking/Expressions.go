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

// IExpression is the interface to an expression AST node with types added.
type IExpression interface {
	GetSourcePosition() parsing.SourcePos
	GetTypeInfo() types.IType
	isTypeExpression()
}

//=====================================================================================================================

// AdditionExpr represents an addition operation.
type AdditionExpr struct {
	SourcePosition parsing.SourcePos
	Lhs            IExpression
	Rhs            IExpression
	TypeInfo       types.IType
}

func (e *AdditionExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *AdditionExpr) GetTypeInfo() types.IType             { return e.TypeInfo }
func (e *AdditionExpr) isTypeExpression()                    {}

//=====================================================================================================================

// ArrayLiteralExpr represents a parenthesized expression or comma-separated sequence of expressions.
type ArrayLiteralExpr struct {
	SourcePosition parsing.SourcePos
	Elements       []IExpression
	TypeInfo       types.IType // TODO: Should be element type
}

func (e *ArrayLiteralExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *ArrayLiteralExpr) GetTypeInfo() types.IType             { return e.TypeInfo }
func (e *ArrayLiteralExpr) isTypeExpression()                    {}

//=====================================================================================================================

// BooleanLiteralExpr represents a single boolean literal.
type BooleanLiteralExpr struct {
	SourcePosition parsing.SourcePos
	Value          bool
}

func (e *BooleanLiteralExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *BooleanLiteralExpr) GetTypeInfo() types.IType             { return types.BoolTypeInstance }
func (e *BooleanLiteralExpr) isTypeExpression()                    {}

//=====================================================================================================================

// BuiltInTypeExpr represents a pre-defined base type.
type BuiltInTypeExpr struct {
	SourcePosition parsing.SourcePos
	Value          types.BuiltInType
}

func (e *BuiltInTypeExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *BuiltInTypeExpr) GetTypeInfo() types.IType             { return types.TypeTypeInstance }
func (e *BuiltInTypeExpr) isTypeExpression()                    {}

//=====================================================================================================================

// DivisionExpr represents a division operation.
type DivisionExpr struct {
	SourcePosition parsing.SourcePos
	Lhs            IExpression
	Rhs            IExpression
	TypeInfo       types.IType
}

func (e *DivisionExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *DivisionExpr) GetTypeInfo() types.IType             { return e.TypeInfo }
func (e *DivisionExpr) isTypeExpression()                    {}

//=====================================================================================================================

// EqualsExpr represents a equals operation.
type EqualsExpr struct {
	SourcePosition parsing.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *EqualsExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *EqualsExpr) GetTypeInfo() types.IType             { return types.BoolTypeInstance }
func (e *EqualsExpr) isTypeExpression()                    {}

//=====================================================================================================================

// Float64LiteralExpr represents a single 64-bit floating point literal.
type Float64LiteralExpr struct {
	SourcePosition parsing.SourcePos
	Value          float64
}

func (e *Float64LiteralExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *Float64LiteralExpr) GetTypeInfo() types.IType             { return types.Float64TypeInstance }
func (e *Float64LiteralExpr) isTypeExpression()                    {}

//=====================================================================================================================

// FunctionCallExpr represents a function call (a function name followed by a parenthesized expression).
type FunctionCallExpr struct {
	SourcePosition    parsing.SourcePos
	FunctionReference IExpression
	Argument          IExpression
	TypeInfo          types.IType
}

func (e *FunctionCallExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *FunctionCallExpr) GetTypeInfo() types.IType             { return e.TypeInfo }
func (e *FunctionCallExpr) isTypeExpression()                    {}

//=====================================================================================================================

// GreaterThanExpr represents a greater than operation.
type GreaterThanExpr struct {
	SourcePosition parsing.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *GreaterThanExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *GreaterThanExpr) GetTypeInfo() types.IType             { return types.BoolTypeInstance }
func (e *GreaterThanExpr) isTypeExpression()                    {}

//=====================================================================================================================

// GreaterThanOrEqualsExpr represents a greater than operation.
type GreaterThanOrEqualsExpr struct {
	SourcePosition parsing.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *GreaterThanOrEqualsExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *GreaterThanOrEqualsExpr) GetTypeInfo() types.IType             { return types.BoolTypeInstance }
func (e *GreaterThanOrEqualsExpr) isTypeExpression()                    {}

//=====================================================================================================================

// IdentifierExpr represents a single identifier.
type IdentifierExpr struct {
	SourcePosition parsing.SourcePos
	Name           string
	TypeInfo       types.IType
}

func (e *IdentifierExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *IdentifierExpr) GetTypeInfo() types.IType             { return e.TypeInfo }
func (e *IdentifierExpr) isTypeExpression()                    {}

//=====================================================================================================================

// Int64LiteralExpr represents a single 64-bit integer literal.
type Int64LiteralExpr struct {
	SourcePosition parsing.SourcePos
	Value          int64
}

func (e *Int64LiteralExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *Int64LiteralExpr) GetTypeInfo() types.IType             { return types.Int64TypeInstance }
func (e *Int64LiteralExpr) isTypeExpression()                    {}

//=====================================================================================================================

// IsExpr represents an "is" test.
type IsExpr struct {
	SourcePosition parsing.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *IsExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *IsExpr) GetTypeInfo() types.IType             { return types.BoolTypeInstance }
func (e *IsExpr) isTypeExpression()                    {}

//=====================================================================================================================

// LeadingDocumentationExpr represents lines of leading documentation.
type LeadingDocumentationExpr struct {
	SourcePosition parsing.SourcePos
	Text           string
	TypeInfo       types.IType
}

func (e *LeadingDocumentationExpr) GetSourcePosition() parsing.SourcePos {
	return e.SourcePosition
}
func (e *LeadingDocumentationExpr) GetTypeInfo() types.IType { return e.TypeInfo }
func (e *LeadingDocumentationExpr) isTypeExpression()        {}

//=====================================================================================================================

// LessThanExpr represents a less than operation.
type LessThanExpr struct {
	SourcePosition parsing.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *LessThanExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *LessThanExpr) GetTypeInfo() types.IType             { return types.BoolTypeInstance }
func (e *LessThanExpr) isTypeExpression()                    {}

//=====================================================================================================================

// LessThanOrEqualsExpr represents a less than operation.
type LessThanOrEqualsExpr struct {
	SourcePosition parsing.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *LessThanOrEqualsExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *LessThanOrEqualsExpr) GetTypeInfo() types.IType             { return types.BoolTypeInstance }
func (e *LessThanOrEqualsExpr) isTypeExpression()                    {}

//=====================================================================================================================

// LogicalAndExpr represents a logical "and" operation.
type LogicalAndExpr struct {
	SourcePosition parsing.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *LogicalAndExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *LogicalAndExpr) GetTypeInfo() types.IType             { return types.BoolTypeInstance }
func (e *LogicalAndExpr) isTypeExpression()                    {}

//=====================================================================================================================

// LogicalNotOperationExpr represents a logical "not" operation.
type LogicalNotOperationExpr struct {
	SourcePosition parsing.SourcePos
	Operand        IExpression
}

func (e *LogicalNotOperationExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *LogicalNotOperationExpr) GetTypeInfo() types.IType             { return types.BoolTypeInstance }
func (e *LogicalNotOperationExpr) isTypeExpression()                    {}

//=====================================================================================================================

// LogicalOrExpr represents a logical "or" operation.
type LogicalOrExpr struct {
	SourcePosition parsing.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *LogicalOrExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *LogicalOrExpr) GetTypeInfo() types.IType             { return types.BoolTypeInstance }
func (e *LogicalOrExpr) isTypeExpression()                    {}

//=====================================================================================================================

// MultilineStringLiteralExpr represents a multiline (back-ticked) string literal.
type MultilineStringLiteralExpr struct {
	SourcePosition parsing.SourcePos
	Text           string
	TypeInfo       types.IType
}

func (e *MultilineStringLiteralExpr) GetSourcePosition() parsing.SourcePos {
	return e.SourcePosition
}
func (e *MultilineStringLiteralExpr) GetTypeInfo() types.IType { return e.TypeInfo }
func (e *MultilineStringLiteralExpr) isTypeExpression()        {}

//=====================================================================================================================

// MultiplicationExpr represents a multiplication operation.
type MultiplicationExpr struct {
	SourcePosition parsing.SourcePos
	Lhs            IExpression
	Rhs            IExpression
	TypeInfo       types.IType
}

func (e *MultiplicationExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *MultiplicationExpr) GetTypeInfo() types.IType             { return e.TypeInfo }
func (e *MultiplicationExpr) isTypeExpression()                    {}

//=====================================================================================================================

// NegationOperationExpr represents an arithmetic negation operation.
type NegationOperationExpr struct {
	SourcePosition parsing.SourcePos
	Operand        IExpression
	TypeInfo       types.IType
}

func (e *NegationOperationExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *NegationOperationExpr) GetTypeInfo() types.IType             { return e.TypeInfo }
func (e *NegationOperationExpr) isTypeExpression()                    {}

//=====================================================================================================================

// NotEqualsExpr represents a equals operation.
type NotEqualsExpr struct {
	SourcePosition parsing.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *NotEqualsExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *NotEqualsExpr) GetTypeInfo() types.IType             { return types.BoolTypeInstance }
func (e *NotEqualsExpr) isTypeExpression()                    {}

//=====================================================================================================================

// OptionalExpr represents a parenthesized expression or comma-separated sequence of expressions.
type OptionalExpr struct {
	SourcePosition parsing.SourcePos
	Operand        IExpression
	TypeInfo       types.IType
}

func (e *OptionalExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *OptionalExpr) GetTypeInfo() types.IType             { return e.TypeInfo }
func (e *OptionalExpr) isTypeExpression()                    {}

//=====================================================================================================================

// ParenthesizedExpr represents a parenthesized expression or comma-separated sequence of expressions.
type ParenthesizedExpr struct {
	SourcePosition parsing.SourcePos
	InnerExpr      IExpression
	TypeInfo       types.IType
}

func (e *ParenthesizedExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *ParenthesizedExpr) GetTypeInfo() types.IType             { return e.TypeInfo }
func (e *ParenthesizedExpr) isTypeExpression()                    {}

//=====================================================================================================================

// RecordExpr represents a record.
type RecordExpr struct {
	SourcePosition parsing.SourcePos
	TypeInfo       types.IType
}

func (e *RecordExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *RecordExpr) GetTypeInfo() types.IType             { return e.TypeInfo }
func (e *RecordExpr) isTypeExpression()                    {}

//=====================================================================================================================

// StringConcatenationExpr represents concatenation of two strings.
type StringConcatenationExpr struct {
	SourcePosition parsing.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *StringConcatenationExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *StringConcatenationExpr) GetTypeInfo() types.IType             { return types.StringTypeInstance }
func (e *StringConcatenationExpr) isTypeExpression()                    {}

//=====================================================================================================================

// StringLiteralExpr represents a single string literal.
type StringLiteralExpr struct {
	SourcePosition parsing.SourcePos
	ValueIndex     uint64
}

func (e *StringLiteralExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *StringLiteralExpr) GetTypeInfo() types.IType             { return types.StringTypeInstance }
func (e *StringLiteralExpr) isTypeExpression()                    {}

//=====================================================================================================================

// SubtractionExpr represents a subtraction operation.
type SubtractionExpr struct {
	SourcePosition parsing.SourcePos
	Lhs            IExpression
	Rhs            IExpression
	TypeInfo       types.IType
}

func (e *SubtractionExpr) GetSourcePosition() parsing.SourcePos { return e.SourcePosition }
func (e *SubtractionExpr) GetTypeInfo() types.IType             { return e.TypeInfo }
func (e *SubtractionExpr) isTypeExpression()                    {}

//=====================================================================================================================

// TrailingDocumentationExpr represents lines of trailing documentation.
type TrailingDocumentationExpr struct {
	SourcePosition parsing.SourcePos
	Text           string
	TypeInfo       types.IType
}

func (e *TrailingDocumentationExpr) GetSourcePosition() parsing.SourcePos {
	return e.SourcePosition
}
func (e *TrailingDocumentationExpr) GetTypeInfo() types.IType { return e.TypeInfo }
func (e *TrailingDocumentationExpr) isTypeExpression()        {}

//=====================================================================================================================
