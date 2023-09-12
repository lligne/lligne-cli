//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package typechecking

import (
	"lligne-cli/internal/lligne/code/util"
	"lligne-cli/internal/lligne/runtime/types"
)

//=====================================================================================================================

// IExpression is the interface to an expression AST node with types added.
type IExpression interface {
	GetSourcePosition() util.SourcePos
	GetTypeIndex() uint64
	isTypeExpression()
}

//=====================================================================================================================

// AdditionExpr represents an addition operation.
type AdditionExpr struct {
	SourcePosition util.SourcePos
	Lhs            IExpression
	Rhs            IExpression
	TypeIndex      uint64
}

func (e *AdditionExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *AdditionExpr) GetTypeIndex() uint64              { return e.TypeIndex }
func (e *AdditionExpr) isTypeExpression()                 {}

//=====================================================================================================================

// ArrayLiteralExpr represents a parenthesized expression or comma-separated sequence of expressions.
type ArrayLiteralExpr struct {
	SourcePosition util.SourcePos
	Elements       []IExpression
	TypeIndex      uint64 // TODO: Should be element type
}

func (e *ArrayLiteralExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *ArrayLiteralExpr) GetTypeIndex() uint64              { return e.TypeIndex }
func (e *ArrayLiteralExpr) isTypeExpression()                 {}

//=====================================================================================================================

// BooleanLiteralExpr represents a single boolean literal.
type BooleanLiteralExpr struct {
	SourcePosition util.SourcePos
	Value          bool
}

func (e *BooleanLiteralExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *BooleanLiteralExpr) GetTypeIndex() uint64              { return types.BuiltInTypeIndexBool }
func (e *BooleanLiteralExpr) isTypeExpression()                 {}

//=====================================================================================================================

// BuiltInTypeExpr represents a pre-defined base type.
type BuiltInTypeExpr struct {
	SourcePosition util.SourcePos
	ValueIndex     uint64
}

func (e *BuiltInTypeExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *BuiltInTypeExpr) GetTypeIndex() uint64              { return types.BuiltInTypeIndexType }
func (e *BuiltInTypeExpr) isTypeExpression()                 {}

//=====================================================================================================================

// DivisionExpr represents a division operation.
type DivisionExpr struct {
	SourcePosition util.SourcePos
	Lhs            IExpression
	Rhs            IExpression
	TypeIndex      uint64
}

func (e *DivisionExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *DivisionExpr) GetTypeIndex() uint64              { return e.TypeIndex }
func (e *DivisionExpr) isTypeExpression()                 {}

//=====================================================================================================================

// EqualsExpr represents a equals operation.
type EqualsExpr struct {
	SourcePosition util.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *EqualsExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *EqualsExpr) GetTypeIndex() uint64              { return types.BuiltInTypeIndexBool }
func (e *EqualsExpr) isTypeExpression()                 {}

//=====================================================================================================================

// FieldReferenceExpr represents a field reference (".") operation.
type FieldReferenceExpr struct {
	SourcePosition util.SourcePos
	Parent         IExpression
	Child          IExpression
}

func (e *FieldReferenceExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *FieldReferenceExpr) GetTypeIndex() uint64              { return e.Child.GetTypeIndex() }
func (e *FieldReferenceExpr) isTypeExpression()                 {}

//=====================================================================================================================

// Float64LiteralExpr represents a single 64-bit floating point literal.
type Float64LiteralExpr struct {
	SourcePosition util.SourcePos
	Value          float64
}

func (e *Float64LiteralExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *Float64LiteralExpr) GetTypeIndex() uint64              { return types.BuiltInTypeIndexFloat64 }
func (e *Float64LiteralExpr) isTypeExpression()                 {}

//=====================================================================================================================

// FunctionCallExpr represents a function call (a function name followed by a parenthesized expression).
type FunctionCallExpr struct {
	SourcePosition    util.SourcePos
	FunctionReference IExpression
	Argument          IExpression
	TypeIndex         uint64
}

func (e *FunctionCallExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *FunctionCallExpr) GetTypeIndex() uint64              { return e.TypeIndex }
func (e *FunctionCallExpr) isTypeExpression()                 {}

//=====================================================================================================================

// GreaterThanExpr represents a greater than operation.
type GreaterThanExpr struct {
	SourcePosition util.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *GreaterThanExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *GreaterThanExpr) GetTypeIndex() uint64              { return types.BuiltInTypeIndexBool }
func (e *GreaterThanExpr) isTypeExpression()                 {}

//=====================================================================================================================

// GreaterThanOrEqualsExpr represents a greater than operation.
type GreaterThanOrEqualsExpr struct {
	SourcePosition util.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *GreaterThanOrEqualsExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *GreaterThanOrEqualsExpr) GetTypeIndex() uint64              { return types.BuiltInTypeIndexBool }
func (e *GreaterThanOrEqualsExpr) isTypeExpression()                 {}

//=====================================================================================================================

// IdentifierExpr represents a single identifier.
type IdentifierExpr struct {
	SourcePosition util.SourcePos
	NameIndex      uint64
	FieldIndex     uint64
	TypeIndex      uint64
}

func (e *IdentifierExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *IdentifierExpr) GetTypeIndex() uint64              { return e.TypeIndex }
func (e *IdentifierExpr) isTypeExpression()                 {}

//=====================================================================================================================

// Int64LiteralExpr represents a single 64-bit integer literal.
type Int64LiteralExpr struct {
	SourcePosition util.SourcePos
	Value          int64
}

func (e *Int64LiteralExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *Int64LiteralExpr) GetTypeIndex() uint64              { return types.BuiltInTypeIndexInt64 }
func (e *Int64LiteralExpr) isTypeExpression()                 {}

//=====================================================================================================================

// IsExpr represents an "is" test.
type IsExpr struct {
	SourcePosition util.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *IsExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *IsExpr) GetTypeIndex() uint64              { return types.BuiltInTypeIndexBool }
func (e *IsExpr) isTypeExpression()                 {}

//=====================================================================================================================

// LeadingDocumentationExpr represents lines of leading documentation.
type LeadingDocumentationExpr struct {
	SourcePosition util.SourcePos
	Text           string
}

func (e *LeadingDocumentationExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *LeadingDocumentationExpr) GetTypeIndex() uint64              { return types.BuiltInTypeIndexUnit }
func (e *LeadingDocumentationExpr) isTypeExpression()                 {}

//=====================================================================================================================

// LessThanExpr represents a less than operation.
type LessThanExpr struct {
	SourcePosition util.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *LessThanExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *LessThanExpr) GetTypeIndex() uint64              { return types.BuiltInTypeIndexBool }
func (e *LessThanExpr) isTypeExpression()                 {}

//=====================================================================================================================

// LessThanOrEqualsExpr represents a less than operation.
type LessThanOrEqualsExpr struct {
	SourcePosition util.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *LessThanOrEqualsExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *LessThanOrEqualsExpr) GetTypeIndex() uint64              { return types.BuiltInTypeIndexBool }
func (e *LessThanOrEqualsExpr) isTypeExpression()                 {}

//=====================================================================================================================

// LogicalAndExpr represents a logical "and" operation.
type LogicalAndExpr struct {
	SourcePosition util.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *LogicalAndExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *LogicalAndExpr) GetTypeIndex() uint64              { return types.BuiltInTypeIndexBool }
func (e *LogicalAndExpr) isTypeExpression()                 {}

//=====================================================================================================================

// LogicalNotOperationExpr represents a logical "not" operation.
type LogicalNotOperationExpr struct {
	SourcePosition util.SourcePos
	Operand        IExpression
}

func (e *LogicalNotOperationExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *LogicalNotOperationExpr) GetTypeIndex() uint64              { return types.BuiltInTypeIndexBool }
func (e *LogicalNotOperationExpr) isTypeExpression()                 {}

//=====================================================================================================================

// LogicalOrExpr represents a logical "or" operation.
type LogicalOrExpr struct {
	SourcePosition util.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *LogicalOrExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *LogicalOrExpr) GetTypeIndex() uint64              { return types.BuiltInTypeIndexBool }
func (e *LogicalOrExpr) isTypeExpression()                 {}

//=====================================================================================================================

// MultiplicationExpr represents a multiplication operation.
type MultiplicationExpr struct {
	SourcePosition util.SourcePos
	Lhs            IExpression
	Rhs            IExpression
	TypeIndex      uint64
}

func (e *MultiplicationExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *MultiplicationExpr) GetTypeIndex() uint64              { return e.TypeIndex }
func (e *MultiplicationExpr) isTypeExpression()                 {}

//=====================================================================================================================

// NegationOperationExpr represents an arithmetic negation operation.
type NegationOperationExpr struct {
	SourcePosition util.SourcePos
	Operand        IExpression
	TypeIndex      uint64
}

func (e *NegationOperationExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *NegationOperationExpr) GetTypeIndex() uint64              { return e.TypeIndex }
func (e *NegationOperationExpr) isTypeExpression()                 {}

//=====================================================================================================================

// NotEqualsExpr represents a equals operation.
type NotEqualsExpr struct {
	SourcePosition util.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *NotEqualsExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *NotEqualsExpr) GetTypeIndex() uint64              { return types.BuiltInTypeIndexBool }
func (e *NotEqualsExpr) isTypeExpression()                 {}

//=====================================================================================================================

// OptionalExpr represents a parenthesized expression or comma-separated sequence of expressions.
type OptionalExpr struct {
	SourcePosition util.SourcePos
	Operand        IExpression
	TypeIndex      uint64
}

func (e *OptionalExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *OptionalExpr) GetTypeIndex() uint64              { return e.TypeIndex }
func (e *OptionalExpr) isTypeExpression()                 {}

//=====================================================================================================================

// ParenthesizedExpr represents a parenthesized expression or comma-separated sequence of expressions.
type ParenthesizedExpr struct {
	SourcePosition util.SourcePos
	InnerExpr      IExpression
	TypeIndex      uint64
}

func (e *ParenthesizedExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *ParenthesizedExpr) GetTypeIndex() uint64              { return e.TypeIndex }
func (e *ParenthesizedExpr) isTypeExpression()                 {}

//=====================================================================================================================

// RecordExpr represents a record.
type RecordExpr struct {
	SourcePosition util.SourcePos
	Fields         []*RecordFieldExpr
	TypeIndex      uint64
}

func (e *RecordExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *RecordExpr) GetTypeIndex() uint64              { return e.TypeIndex }
func (e *RecordExpr) isTypeExpression()                 {}

//=====================================================================================================================

// RecordFieldExpr represents a record field.
type RecordFieldExpr struct {
	SourcePosition util.SourcePos
	FieldNameIndex uint64 // TODO: this is redundant with record type information
	FieldValue     IExpression
}

func (e *RecordFieldExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *RecordFieldExpr) GetTypeIndex() uint64              { return e.FieldValue.GetTypeIndex() }
func (e *RecordFieldExpr) isTypeExpression()                 {}

//=====================================================================================================================

// StringConcatenationExpr represents concatenation of two strings.
type StringConcatenationExpr struct {
	SourcePosition util.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *StringConcatenationExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *StringConcatenationExpr) GetTypeIndex() uint64              { return types.BuiltInTypeIndexString }
func (e *StringConcatenationExpr) isTypeExpression()                 {}

//=====================================================================================================================

// StringLiteralExpr represents a single string literal.
type StringLiteralExpr struct {
	SourcePosition util.SourcePos
	ValueIndex     uint64
}

func (e *StringLiteralExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *StringLiteralExpr) GetTypeIndex() uint64              { return types.BuiltInTypeIndexString }
func (e *StringLiteralExpr) isTypeExpression()                 {}

//=====================================================================================================================

// SubtractionExpr represents a subtraction operation.
type SubtractionExpr struct {
	SourcePosition util.SourcePos
	Lhs            IExpression
	Rhs            IExpression
	TypeIndex      uint64
}

func (e *SubtractionExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *SubtractionExpr) GetTypeIndex() uint64              { return e.TypeIndex }
func (e *SubtractionExpr) isTypeExpression()                 {}

//=====================================================================================================================

// TrailingDocumentationExpr represents lines of trailing documentation.
type TrailingDocumentationExpr struct {
	SourcePosition util.SourcePos
	Text           string
}

func (e *TrailingDocumentationExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *TrailingDocumentationExpr) GetTypeIndex() uint64              { return types.BuiltInTypeIndexUnit }
func (e *TrailingDocumentationExpr) isTypeExpression()                 {}

//=====================================================================================================================
