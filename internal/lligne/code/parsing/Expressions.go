//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package parsing

import "lligne-cli/internal/lligne/code/util"

//=====================================================================================================================

// IExpression is the interface to an expression AST node.
type IExpression interface {
	GetSourcePosition() util.SourcePos
	isExpression()
}

//=====================================================================================================================

// StringDelimiters is an enumeration of start/stop delimiters for string literal expressions.
type StringDelimiters int

const (
	StringDelimitersSingleQuotes StringDelimiters = 1 + iota
	StringDelimitersDoubleQuotes
	StringDelimitersBackTicks
	StringDelimitersSingleQuotesMultiline
	StringDelimitersDoubleQuotesMultiline
	StringDelimitersBackTicksMultiline
)

//=====================================================================================================================

// AdditionExpr represents an addition ("+") operation.
type AdditionExpr struct {
	SourcePosition util.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *AdditionExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *AdditionExpr) isExpression()                     {}

//=====================================================================================================================

// ArrayLiteralExpr represents an array literal.
type ArrayLiteralExpr struct {
	SourcePosition util.SourcePos
	Elements       []IExpression
}

func (e *ArrayLiteralExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *ArrayLiteralExpr) isExpression()                     {}

//=====================================================================================================================

// BooleanLiteralExpr represents a single boolean literal.
type BooleanLiteralExpr struct {
	SourcePosition util.SourcePos
	Value          bool
}

func (e *BooleanLiteralExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *BooleanLiteralExpr) isExpression()                     {}

//=====================================================================================================================

// BuiltInTypeExpr represents a single fundamental type name.
type BuiltInTypeExpr struct {
	SourcePosition util.SourcePos
}

func (e *BuiltInTypeExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *BuiltInTypeExpr) isExpression()                     {}

//=====================================================================================================================

// DivisionExpr represents a division ("/") operation.
type DivisionExpr struct {
	SourcePosition util.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *DivisionExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *DivisionExpr) isExpression()                     {}

//=====================================================================================================================

// DocumentExpr represents the pseudo operation o connecting an item to its documentation.
type DocumentExpr struct {
	SourcePosition util.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *DocumentExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *DocumentExpr) isExpression()                     {}

//=====================================================================================================================

// EqualsExpr represents an equality comparison ("==") operation.
type EqualsExpr struct {
	SourcePosition util.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *EqualsExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *EqualsExpr) isExpression()                     {}

//=====================================================================================================================

// FieldReferenceExpr represents a field reference (".") operation.
type FieldReferenceExpr struct {
	SourcePosition util.SourcePos
	Parent         IExpression
	Child          IExpression
}

func (e *FieldReferenceExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *FieldReferenceExpr) isExpression()                     {}

//=====================================================================================================================

// Float64LiteralExpr represents a single integer literal.
type Float64LiteralExpr struct {
	SourcePosition util.SourcePos
	Value          float64
}

func (e *Float64LiteralExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *Float64LiteralExpr) isExpression()                     {}

//=====================================================================================================================

// FunctionArgumentsExpr represents a parenthesized, comma-separated sequence of expressions postfix to
// a function reference.
type FunctionArgumentsExpr struct {
	SourcePosition util.SourcePos
	Items          []IExpression
}

func (e *FunctionArgumentsExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *FunctionArgumentsExpr) isExpression()                     {}

//=====================================================================================================================

// FunctionArrowExpr represents a function call type with "->" operator.
type FunctionArrowExpr struct {
	SourcePosition util.SourcePos
	Argument       IExpression
	Result         IExpression
}

func (e *FunctionArrowExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *FunctionArrowExpr) isExpression()                     {}

//=====================================================================================================================

// FunctionCallExpr represents a function call (a function name followed by a parenthesized expression).
type FunctionCallExpr struct {
	SourcePosition    util.SourcePos
	FunctionReference IExpression
	Argument          IExpression
}

func (e *FunctionCallExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *FunctionCallExpr) isExpression()                     {}

//=====================================================================================================================

// GreaterThanExpr represents a greater than (">") comparison operation.
type GreaterThanExpr struct {
	SourcePosition util.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *GreaterThanExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *GreaterThanExpr) isExpression()                     {}

//=====================================================================================================================

// GreaterThanOrEqualsExpr represents a greater than or equals (">=") comparison operation.
type GreaterThanOrEqualsExpr struct {
	SourcePosition util.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *GreaterThanOrEqualsExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *GreaterThanOrEqualsExpr) isExpression()                     {}

//=====================================================================================================================

// IdentifierExpr represents a single identifier.
type IdentifierExpr struct {
	SourcePosition util.SourcePos
}

func (e *IdentifierExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *IdentifierExpr) isExpression()                     {}

//=====================================================================================================================

// InExpr represents a set membership "in" operation.
type InExpr struct {
	SourcePosition util.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *InExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *InExpr) isExpression()                     {}

//=====================================================================================================================

// Int64LiteralExpr represents a single integer literal.
type Int64LiteralExpr struct {
	SourcePosition util.SourcePos
	Value          int64
}

func (e *Int64LiteralExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *Int64LiteralExpr) isExpression()                     {}

//=====================================================================================================================

// IntersectExpr represents a type/value intersection "&" operation.
type IntersectExpr struct {
	SourcePosition util.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *IntersectExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *IntersectExpr) isExpression()                     {}

//=====================================================================================================================

// IntersectAssignValueExpr represents a type/value intersection value assignment "=" operation.
type IntersectAssignValueExpr struct {
	SourcePosition util.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *IntersectAssignValueExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *IntersectAssignValueExpr) isExpression()                     {}

//=====================================================================================================================

// IntersectDefaultValueExpr represents a type/value intersection default value "?:" operation.
type IntersectDefaultValueExpr struct {
	SourcePosition util.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *IntersectDefaultValueExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *IntersectDefaultValueExpr) isExpression()                     {}

//=====================================================================================================================

// IntersectLowPrecedenceExpr represents a low precedence type/value intersection "&&" operation.
type IntersectLowPrecedenceExpr struct {
	SourcePosition util.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *IntersectLowPrecedenceExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *IntersectLowPrecedenceExpr) isExpression()                     {}

//=====================================================================================================================

// IsExpr represents a type membership "is" operation.
type IsExpr struct {
	SourcePosition util.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *IsExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *IsExpr) isExpression()                     {}

//=====================================================================================================================

// LeadingDocumentationExpr represents lines of leading documentation.
type LeadingDocumentationExpr struct {
	SourcePosition util.SourcePos
}

func (e *LeadingDocumentationExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *LeadingDocumentationExpr) isExpression()                     {}

//=====================================================================================================================

// LessThanExpr represents a less than ("<") comparison operation.
type LessThanExpr struct {
	SourcePosition util.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *LessThanExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *LessThanExpr) isExpression()                     {}

//=====================================================================================================================

// LessThanOrEqualsExpr represents a less than or equals ("<=") comparison operation.
type LessThanOrEqualsExpr struct {
	SourcePosition util.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *LessThanOrEqualsExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *LessThanOrEqualsExpr) isExpression()                     {}

//=====================================================================================================================

// LogicalAndExpr represents a conjunction "and" operation.
type LogicalAndExpr struct {
	SourcePosition util.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *LogicalAndExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *LogicalAndExpr) isExpression()                     {}

//=====================================================================================================================

// LogicalNotOperationExpr represents the logical not prefix operation.
type LogicalNotOperationExpr struct {
	SourcePosition util.SourcePos
	Operand        IExpression
}

func (e *LogicalNotOperationExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *LogicalNotOperationExpr) isExpression()                     {}

//=====================================================================================================================

// LogicalOrExpr represents a disjunction "or" operation.
type LogicalOrExpr struct {
	SourcePosition util.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *LogicalOrExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *LogicalOrExpr) isExpression()                     {}

//=====================================================================================================================

// MatchExpr represents a pattern match "=~" operation.
type MatchExpr struct {
	SourcePosition util.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *MatchExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *MatchExpr) isExpression()                     {}

//=====================================================================================================================

// MultiplicationExpr represents a multiplication ("*") operation.
type MultiplicationExpr struct {
	SourcePosition util.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *MultiplicationExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *MultiplicationExpr) isExpression()                     {}

//=====================================================================================================================

// NegationOperationExpr represents the arithmetic negation prefix ("-") operation.
type NegationOperationExpr struct {
	SourcePosition util.SourcePos
	Operand        IExpression
}

func (e *NegationOperationExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *NegationOperationExpr) isExpression()                     {}

//=====================================================================================================================

// NotEqualsExpr represents an equality comparison ("==") operation.
type NotEqualsExpr struct {
	SourcePosition util.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *NotEqualsExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *NotEqualsExpr) isExpression()                     {}

//=====================================================================================================================

// NotMatchExpr represents a pattern nonmatch ("!~") operation.
type NotMatchExpr struct {
	SourcePosition util.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *NotMatchExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *NotMatchExpr) isExpression()                     {}

//=====================================================================================================================

// OptionalExpr represents an Optional(of:X) expression using "?" suffix.
type OptionalExpr struct {
	SourcePosition util.SourcePos
	Operand        IExpression
}

func (e *OptionalExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *OptionalExpr) isExpression()                     {}

//=====================================================================================================================

// ParenthesizedExpr represents a parenthesized expression or comma-separated sequence of expressions.
type ParenthesizedExpr struct {
	SourcePosition util.SourcePos
	InnerExpr      IExpression
}

func (e *ParenthesizedExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *ParenthesizedExpr) isExpression()                     {}

//=====================================================================================================================

// QualifyExpr represents a type qualification (":") operation.
type QualifyExpr struct {
	SourcePosition util.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *QualifyExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *QualifyExpr) isExpression()                     {}

//=====================================================================================================================

// RangeExpr represents a range ("..") operation.
type RangeExpr struct {
	SourcePosition util.SourcePos
	First          IExpression
	Last           IExpression
}

func (e *RangeExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *RangeExpr) isExpression()                     {}

//=====================================================================================================================

// RecordExpr represents a record literal expression.
type RecordExpr struct {
	SourcePosition util.SourcePos
	Items          []IExpression
}

func (e *RecordExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *RecordExpr) isExpression()                     {}

//=====================================================================================================================

// StringLiteralExpr represents a single string literal.
type StringLiteralExpr struct {
	SourcePosition util.SourcePos
	Delimiters     StringDelimiters
}

func (e *StringLiteralExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *StringLiteralExpr) isExpression()                     {}

//=====================================================================================================================

// SubtractionExpr represents a subtraction ("-") operation.
type SubtractionExpr struct {
	SourcePosition util.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *SubtractionExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *SubtractionExpr) isExpression()                     {}

//=====================================================================================================================

// TrailingDocumentationExpr represents lines of trailing documentation.
type TrailingDocumentationExpr struct {
	SourcePosition util.SourcePos
}

func (e *TrailingDocumentationExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *TrailingDocumentationExpr) isExpression()                     {}

//=====================================================================================================================

// UnionExpr represents a type union ("|") operation.
type UnionExpr struct {
	SourcePosition util.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *UnionExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *UnionExpr) isExpression()                     {}

//=====================================================================================================================

// UnitExpr represents a parenthesized expression with nothing in it.
type UnitExpr struct {
	SourcePosition util.SourcePos
}

func (e *UnitExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *UnitExpr) isExpression()                     {}

//=====================================================================================================================

// WhenExpr represents a when ("when") operation.
type WhenExpr struct {
	SourcePosition util.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *WhenExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *WhenExpr) isExpression()                     {}

//=====================================================================================================================

// WhereExpr represents a type where ("where") operation.
type WhereExpr struct {
	SourcePosition util.SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *WhereExpr) GetSourcePosition() util.SourcePos { return e.SourcePosition }
func (e *WhereExpr) isExpression()                     {}

//=====================================================================================================================
