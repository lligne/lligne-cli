//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package nameresolution

import (
	"lligne-cli/internal/lligne/code/util"
)

//=====================================================================================================================

// ResolutionMechanism is an enumeration of how identifiers link to their origin.
//
// Name resolution rules:
// 1. When inside the right hand side of a field reference expression, find the name inside the left hand side of the expression or fail.
//
// 2. When inside the left hand side of a where expression, find the name inside the right hand side of the expression or continue.
// 3. When inside a record, find the name as a sibling field in the record or continue.
// 4. When inside a nested record, recursively find the name as a field of the parent record or continue.
// 5. Find the name inside the top level.
type ResolutionMechanism uint16

const (
	ResolutionMechanismUndefined ResolutionMechanism = iota
	ResolutionMechanismFieldReference
	ResolutionMechanismWhereField
	ResolutionMechanismRecordField
	ResolutionMechanismTopLevel
)

//=====================================================================================================================

// IExpression is the interface to an expression AST node with identifier names linked to their source.
type IExpression interface {
	GetFieldNameIndexes() []uint64
	GetSourcePosition() util.SourcePos
	isStructuredExpression()
}

//=====================================================================================================================

// AdditionExpr represents an addition operation.
type AdditionExpr struct {
	SourcePosition util.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *AdditionExpr) GetFieldNameIndexes() []uint64     { return nil }
func (e *AdditionExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *AdditionExpr) isStructuredExpression()           {}

//=====================================================================================================================

// ArrayLiteralExpr represents a parenthesized expression or comma-separated sequence of expressions.
type ArrayLiteralExpr struct {
	SourcePosition util.SourcePos
	Elements       []IExpression
}

func (e *ArrayLiteralExpr) GetFieldNameIndexes() []uint64     { return nil }
func (e *ArrayLiteralExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *ArrayLiteralExpr) isStructuredExpression()           {}

//=====================================================================================================================

// BooleanLiteralExpr represents a single boolean literal.
type BooleanLiteralExpr struct {
	SourcePosition util.SourcePos
	Value          bool
}

func (e *BooleanLiteralExpr) GetFieldNameIndexes() []uint64     { return nil }
func (e *BooleanLiteralExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *BooleanLiteralExpr) isStructuredExpression()           {}

//=====================================================================================================================

// BuiltInTypeExpr represents a pre-defined base type.
type BuiltInTypeExpr struct {
	SourcePosition util.SourcePos
}

func (e *BuiltInTypeExpr) GetFieldNameIndexes() []uint64     { return nil }
func (e *BuiltInTypeExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *BuiltInTypeExpr) isStructuredExpression()           {}

//=====================================================================================================================

// DivisionExpr represents a division operation.
type DivisionExpr struct {
	SourcePosition util.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *DivisionExpr) GetFieldNameIndexes() []uint64     { return nil }
func (e *DivisionExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *DivisionExpr) isStructuredExpression()           {}

//=====================================================================================================================

// EqualsExpr represents a equals operation.
type EqualsExpr struct {
	SourcePosition util.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *EqualsExpr) GetFieldNameIndexes() []uint64     { return nil }
func (e *EqualsExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *EqualsExpr) isStructuredExpression()           {}

//=====================================================================================================================

// FieldReferenceExpr represents a field reference (".") operation.
type FieldReferenceExpr struct {
	SourcePosition util.SourcePos
	Parent         IExpression
	Child          IExpression
}

func (e *FieldReferenceExpr) GetFieldNameIndexes() []uint64     { return nil }
func (e *FieldReferenceExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *FieldReferenceExpr) isStructuredExpression()           {}

//=====================================================================================================================

// Float64LiteralExpr represents a single 64-bit floating point literal.
type Float64LiteralExpr struct {
	SourcePosition util.SourcePos
	Value          float64
}

func (e *Float64LiteralExpr) GetFieldNameIndexes() []uint64     { return nil }
func (e *Float64LiteralExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *Float64LiteralExpr) isStructuredExpression()           {}

//=====================================================================================================================

// FunctionCallExpr represents a function call (a function name followed by a parenthesized expression).
type FunctionCallExpr struct {
	SourcePosition    util.SourcePos
	FunctionReference IExpression
	Argument          IExpression
}

func (e *FunctionCallExpr) GetFieldNameIndexes() []uint64     { return nil }
func (e *FunctionCallExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *FunctionCallExpr) isStructuredExpression()           {}

//=====================================================================================================================

// GreaterThanExpr represents a greater than operation.
type GreaterThanExpr struct {
	SourcePosition util.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *GreaterThanExpr) GetFieldNameIndexes() []uint64     { return nil }
func (e *GreaterThanExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *GreaterThanExpr) isStructuredExpression()           {}

//=====================================================================================================================

// GreaterThanOrEqualsExpr represents a greater than operation.
type GreaterThanOrEqualsExpr struct {
	SourcePosition util.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *GreaterThanOrEqualsExpr) GetFieldNameIndexes() []uint64     { return nil }
func (e *GreaterThanOrEqualsExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *GreaterThanOrEqualsExpr) isStructuredExpression()           {}

//=====================================================================================================================

// IdentifierExpr represents a single identifier.
type IdentifierExpr struct {
	SourcePosition util.SourcePos
	NameIndex      uint64
	RefMechanism   ResolutionMechanism
	FieldIndex     uint64
}

func (e *IdentifierExpr) GetFieldNameIndexes() []uint64     { return nil }
func (e *IdentifierExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *IdentifierExpr) isStructuredExpression()           {}

//=====================================================================================================================

// Int64LiteralExpr represents a single 64-bit integer literal.
type Int64LiteralExpr struct {
	SourcePosition util.SourcePos
	Value          int64
}

func (e *Int64LiteralExpr) GetFieldNameIndexes() []uint64     { return nil }
func (e *Int64LiteralExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *Int64LiteralExpr) isStructuredExpression()           {}

//=====================================================================================================================

// IsExpr represents an "is" test.
type IsExpr struct {
	SourcePosition util.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *IsExpr) GetFieldNameIndexes() []uint64     { return nil }
func (e *IsExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *IsExpr) isStructuredExpression()           {}

//=====================================================================================================================

// LeadingDocumentationExpr represents lines of leading documentation.
type LeadingDocumentationExpr struct {
	SourcePosition util.SourcePos
	Text           string
}

func (e *LeadingDocumentationExpr) GetFieldNameIndexes() []uint64 { return nil }
func (e *LeadingDocumentationExpr) GetSourcePosition() util.SourcePos {
	return e.SourcePosition
}
func (e *LeadingDocumentationExpr) isStructuredExpression() {}

//=====================================================================================================================

// LessThanExpr represents a less than operation.
type LessThanExpr struct {
	SourcePosition util.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *LessThanExpr) GetFieldNameIndexes() []uint64     { return nil }
func (e *LessThanExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *LessThanExpr) isStructuredExpression()           {}

//=====================================================================================================================

// LessThanOrEqualsExpr represents a less than operation.
type LessThanOrEqualsExpr struct {
	SourcePosition util.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *LessThanOrEqualsExpr) GetFieldNameIndexes() []uint64     { return nil }
func (e *LessThanOrEqualsExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *LessThanOrEqualsExpr) isStructuredExpression()           {}

//=====================================================================================================================

// LogicalAndExpr represents a logical "and" operation.
type LogicalAndExpr struct {
	SourcePosition util.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *LogicalAndExpr) GetFieldNameIndexes() []uint64     { return nil }
func (e *LogicalAndExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *LogicalAndExpr) isStructuredExpression()           {}

//=====================================================================================================================

// LogicalNotOperationExpr represents a logical "not" operation.
type LogicalNotOperationExpr struct {
	SourcePosition util.SourcePos
	Operand        IExpression
}

func (e *LogicalNotOperationExpr) GetFieldNameIndexes() []uint64     { return nil }
func (e *LogicalNotOperationExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *LogicalNotOperationExpr) isStructuredExpression()           {}

//=====================================================================================================================

// LogicalOrExpr represents a logical "or" operation.
type LogicalOrExpr struct {
	SourcePosition util.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *LogicalOrExpr) GetFieldNameIndexes() []uint64     { return nil }
func (e *LogicalOrExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *LogicalOrExpr) isStructuredExpression()           {}

//=====================================================================================================================

// MultiplicationExpr represents a multiplication operation.
type MultiplicationExpr struct {
	SourcePosition util.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *MultiplicationExpr) GetFieldNameIndexes() []uint64     { return nil }
func (e *MultiplicationExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *MultiplicationExpr) isStructuredExpression()           {}

//=====================================================================================================================

// NegationOperationExpr represents an arithmetic negation operation.
type NegationOperationExpr struct {
	SourcePosition util.SourcePos
	Operand        IExpression
}

func (e *NegationOperationExpr) GetFieldNameIndexes() []uint64     { return nil }
func (e *NegationOperationExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *NegationOperationExpr) isStructuredExpression()           {}

//=====================================================================================================================

// NotEqualsExpr represents a equals operation.
type NotEqualsExpr struct {
	SourcePosition util.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *NotEqualsExpr) GetFieldNameIndexes() []uint64     { return nil }
func (e *NotEqualsExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *NotEqualsExpr) isStructuredExpression()           {}

//=====================================================================================================================

// OptionalExpr represents a parenthesized expression or comma-separated sequence of expressions.
type OptionalExpr struct {
	SourcePosition util.SourcePos
	Operand        IExpression
}

func (e *OptionalExpr) GetFieldNameIndexes() []uint64     { return nil }
func (e *OptionalExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *OptionalExpr) isStructuredExpression()           {}

//=====================================================================================================================

// ParenthesizedExpr represents a parenthesized expression or comma-separated sequence of expressions.
type ParenthesizedExpr struct {
	SourcePosition util.SourcePos
	InnerExpr      IExpression
}

func (e *ParenthesizedExpr) GetFieldNameIndexes() []uint64     { return nil }
func (e *ParenthesizedExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *ParenthesizedExpr) isStructuredExpression()           {}

//=====================================================================================================================

// RecordExpr represents a record.
type RecordExpr struct {
	SourcePosition   util.SourcePos
	FieldNameIndexes []uint64
	Fields           []*RecordFieldExpr
}

func (e *RecordExpr) GetFieldNameIndexes() []uint64     { return e.FieldNameIndexes }
func (e *RecordExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *RecordExpr) isStructuredExpression()           {}

//=====================================================================================================================

// RecordFieldExpr represents a record field.
type RecordFieldExpr struct {
	SourcePosition util.SourcePos
	FieldNameIndex uint64
	FieldValue     IExpression
}

func (e *RecordFieldExpr) GetFieldNameIndexes() []uint64     { return nil }
func (e *RecordFieldExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *RecordFieldExpr) isStructuredExpression()           {}

//=====================================================================================================================

// StringConcatenationExpr represents concatenation of two strings.
type StringConcatenationExpr struct {
	SourcePosition util.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *StringConcatenationExpr) GetFieldNameIndexes() []uint64     { return nil }
func (e *StringConcatenationExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *StringConcatenationExpr) isStructuredExpression()           {}

//=====================================================================================================================

// StringLiteralExpr represents a single string literal.
type StringLiteralExpr struct {
	SourcePosition util.SourcePos
	ValueIndex     uint64
}

func (e *StringLiteralExpr) GetFieldNameIndexes() []uint64     { return nil }
func (e *StringLiteralExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *StringLiteralExpr) isStructuredExpression()           {}

//=====================================================================================================================

// SubtractionExpr represents a subtraction operation.
type SubtractionExpr struct {
	SourcePosition util.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *SubtractionExpr) GetFieldNameIndexes() []uint64     { return nil }
func (e *SubtractionExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *SubtractionExpr) isStructuredExpression()           {}

//=====================================================================================================================

// TrailingDocumentationExpr represents lines of trailing documentation.
type TrailingDocumentationExpr struct {
	SourcePosition util.SourcePos
	Text           string
}

func (e *TrailingDocumentationExpr) GetFieldNameIndexes() []uint64 { return nil }
func (e *TrailingDocumentationExpr) GetSourcePosition() util.SourcePos {
	return e.SourcePosition
}
func (e *TrailingDocumentationExpr) isStructuredExpression() {}

//=====================================================================================================================

// WhereExpr represents a subtraction operation.
type WhereExpr struct {
	SourcePosition util.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *WhereExpr) GetFieldNameIndexes() []uint64     { return nil }
func (e *WhereExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *WhereExpr) isStructuredExpression()           {}

//=====================================================================================================================
