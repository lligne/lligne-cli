//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package parsing

//=====================================================================================================================

// IExpression is the interface to an expression AST node.
type IExpression interface {
	GetSourcePosition() SourcePos
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
	SourcePosition SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *AdditionExpr) GetSourcePosition() SourcePos { return e.SourcePosition }
func (e *AdditionExpr) isExpression()                {}

//=====================================================================================================================

// BooleanLiteralExpr represents a single boolean literal.
type BooleanLiteralExpr struct {
	SourcePosition SourcePos
	Value          bool
}

func (e *BooleanLiteralExpr) GetSourcePosition() SourcePos { return e.SourcePosition }
func (e *BooleanLiteralExpr) isExpression()                {}

//=====================================================================================================================

// DivisionExpr represents a division ("/") operation.
type DivisionExpr struct {
	SourcePosition SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *DivisionExpr) GetSourcePosition() SourcePos { return e.SourcePosition }
func (e *DivisionExpr) isExpression()                {}

//=====================================================================================================================

// DocumentExpr represents the pseudo operation o connecting an item to its documentation.
type DocumentExpr struct {
	SourcePosition SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *DocumentExpr) GetSourcePosition() SourcePos { return e.SourcePosition }
func (e *DocumentExpr) isExpression()                {}

//=====================================================================================================================

// EqualsExpr represents an equality comparison ("==") operation.
type EqualsExpr struct {
	SourcePosition SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *EqualsExpr) GetSourcePosition() SourcePos { return e.SourcePosition }
func (e *EqualsExpr) isExpression()                {}

//=====================================================================================================================

// FieldReferenceExpr represents a field reference (".") operation.
type FieldReferenceExpr struct {
	SourcePosition SourcePos
	Parent         IExpression
	Child          IExpression
}

func (e *FieldReferenceExpr) GetSourcePosition() SourcePos { return e.SourcePosition }
func (e *FieldReferenceExpr) isExpression()                {}

//=====================================================================================================================

// FloatingPointLiteralExpr represents a single integer literal.
type FloatingPointLiteralExpr struct {
	SourcePosition SourcePos
}

func (e *FloatingPointLiteralExpr) GetSourcePosition() SourcePos { return e.SourcePosition }
func (e *FloatingPointLiteralExpr) isExpression()                {}

//=====================================================================================================================

// FunctionArgumentsExpr represents a parenthesized, comma-separated sequence of expressions postfix to
// a function reference.
type FunctionArgumentsExpr struct {
	SourcePosition SourcePos
	Items          []IExpression
}

func (e *FunctionArgumentsExpr) GetSourcePosition() SourcePos { return e.SourcePosition }
func (e *FunctionArgumentsExpr) isExpression()                {}

//=====================================================================================================================

// FunctionArrowExpr represents a function call type with "->" operator.
type FunctionArrowExpr struct {
	SourcePosition SourcePos
	Argument       IExpression
	Result         IExpression
}

func (e *FunctionArrowExpr) GetSourcePosition() SourcePos { return e.SourcePosition }
func (e *FunctionArrowExpr) isExpression()                {}

//=====================================================================================================================

// FunctionCallExpr represents a function call (a function name followed by a parenthesized expression).
type FunctionCallExpr struct {
	SourcePosition    SourcePos
	FunctionReference IExpression
	Argument          IExpression
}

func (e *FunctionCallExpr) GetSourcePosition() SourcePos { return e.SourcePosition }
func (e *FunctionCallExpr) isExpression()                {}

//=====================================================================================================================

// GreaterThanExpr represents a greater than (">") comparison operation.
type GreaterThanExpr struct {
	SourcePosition SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *GreaterThanExpr) GetSourcePosition() SourcePos { return e.SourcePosition }
func (e *GreaterThanExpr) isExpression()                {}

//=====================================================================================================================

// GreaterThanOrEqualsExpr represents a greater than or equals (">=") comparison operation.
type GreaterThanOrEqualsExpr struct {
	SourcePosition SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *GreaterThanOrEqualsExpr) GetSourcePosition() SourcePos { return e.SourcePosition }
func (e *GreaterThanOrEqualsExpr) isExpression()                {}

//=====================================================================================================================

// IdentifierExpr represents a single identifier.
type IdentifierExpr struct {
	SourcePosition SourcePos
}

func (e *IdentifierExpr) GetSourcePosition() SourcePos { return e.SourcePosition }
func (e *IdentifierExpr) isExpression()                {}

//=====================================================================================================================

// InExpr represents a set membership "in" operation.
type InExpr struct {
	SourcePosition SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *InExpr) GetSourcePosition() SourcePos { return e.SourcePosition }
func (e *InExpr) isExpression()                {}

//=====================================================================================================================

// IntegerLiteralExpr represents a single integer literal.
type IntegerLiteralExpr struct {
	SourcePosition SourcePos
}

func (e *IntegerLiteralExpr) GetSourcePosition() SourcePos { return e.SourcePosition }
func (e *IntegerLiteralExpr) isExpression()                {}

//=====================================================================================================================

// IntersectExpr represents a type/value intersection "&" operation.
type IntersectExpr struct {
	SourcePosition SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *IntersectExpr) GetSourcePosition() SourcePos { return e.SourcePosition }
func (e *IntersectExpr) isExpression()                {}

//=====================================================================================================================

// IntersectAssignValueExpr represents a type/value intersection value assignment "=" operation.
type IntersectAssignValueExpr struct {
	SourcePosition SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *IntersectAssignValueExpr) GetSourcePosition() SourcePos { return e.SourcePosition }
func (e *IntersectAssignValueExpr) isExpression()                {}

//=====================================================================================================================

// IntersectDefaultValueExpr represents a type/value intersection default value "?:" operation.
type IntersectDefaultValueExpr struct {
	SourcePosition SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *IntersectDefaultValueExpr) GetSourcePosition() SourcePos { return e.SourcePosition }
func (e *IntersectDefaultValueExpr) isExpression()                {}

//=====================================================================================================================

// IntersectLowPrecedenceExpr represents a low precedence type/value intersection "&&" operation.
type IntersectLowPrecedenceExpr struct {
	SourcePosition SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *IntersectLowPrecedenceExpr) GetSourcePosition() SourcePos { return e.SourcePosition }
func (e *IntersectLowPrecedenceExpr) isExpression()                {}

//=====================================================================================================================

// IsExpr represents a type membership "is" operation.
type IsExpr struct {
	SourcePosition SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *IsExpr) GetSourcePosition() SourcePos { return e.SourcePosition }
func (e *IsExpr) isExpression()                {}

//=====================================================================================================================

// LeadingDocumentationExpr represents lines of leading documentation.
type LeadingDocumentationExpr struct {
	SourcePosition SourcePos
}

func (e *LeadingDocumentationExpr) GetSourcePosition() SourcePos { return e.SourcePosition }
func (e *LeadingDocumentationExpr) isExpression()                {}

//=====================================================================================================================

// LessThanExpr represents a less than ("<") comparison operation.
type LessThanExpr struct {
	SourcePosition SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *LessThanExpr) GetSourcePosition() SourcePos { return e.SourcePosition }
func (e *LessThanExpr) isExpression()                {}

//=====================================================================================================================

// LessThanOrEqualsExpr represents a less than or equals ("<=") comparison operation.
type LessThanOrEqualsExpr struct {
	SourcePosition SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *LessThanOrEqualsExpr) GetSourcePosition() SourcePos { return e.SourcePosition }
func (e *LessThanOrEqualsExpr) isExpression()                {}

//=====================================================================================================================

// LogicalAndExpr represents a conjunction "and" operation.
type LogicalAndExpr struct {
	SourcePosition SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *LogicalAndExpr) GetSourcePosition() SourcePos { return e.SourcePosition }
func (e *LogicalAndExpr) isExpression()                {}

//=====================================================================================================================

// LogicalNotOperationExpr represents the logical not prefix operation.
type LogicalNotOperationExpr struct {
	SourcePosition SourcePos
	Operand        IExpression
}

func (e *LogicalNotOperationExpr) GetSourcePosition() SourcePos { return e.SourcePosition }
func (e *LogicalNotOperationExpr) isExpression()                {}

//=====================================================================================================================

// LogicalOrExpr represents a disjunction "or" operation.
type LogicalOrExpr struct {
	SourcePosition SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *LogicalOrExpr) GetSourcePosition() SourcePos { return e.SourcePosition }
func (e *LogicalOrExpr) isExpression()                {}

//=====================================================================================================================

// MatchExpr represents a pattern match "=~" operation.
type MatchExpr struct {
	SourcePosition SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *MatchExpr) GetSourcePosition() SourcePos { return e.SourcePosition }
func (e *MatchExpr) isExpression()                {}

//=====================================================================================================================

// MultiplicationExpr represents a multiplication ("*") operation.
type MultiplicationExpr struct {
	SourcePosition SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *MultiplicationExpr) GetSourcePosition() SourcePos { return e.SourcePosition }
func (e *MultiplicationExpr) isExpression()                {}

//=====================================================================================================================

// NegationOperationExpr represents the arithmetic negation prefix ("-") operation.
type NegationOperationExpr struct {
	SourcePosition SourcePos
	Operand        IExpression
}

func (e *NegationOperationExpr) GetSourcePosition() SourcePos { return e.SourcePosition }
func (e *NegationOperationExpr) isExpression()                {}

//=====================================================================================================================

// NotEqualsExpr represents an equality comparison ("==") operation.
type NotEqualsExpr struct {
	SourcePosition SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *NotEqualsExpr) GetSourcePosition() SourcePos { return e.SourcePosition }
func (e *NotEqualsExpr) isExpression()                {}

//=====================================================================================================================

// NotMatchExpr represents a pattern nonmatch ("!~") operation.
type NotMatchExpr struct {
	SourcePosition SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *NotMatchExpr) GetSourcePosition() SourcePos { return e.SourcePosition }
func (e *NotMatchExpr) isExpression()                {}

//=====================================================================================================================

// OptionalExpr represents an Optional(of:X) expression using "?" suffix.
type OptionalExpr struct {
	SourcePosition SourcePos
	Operand        IExpression
}

func (e *OptionalExpr) GetSourcePosition() SourcePos { return e.SourcePosition }
func (e *OptionalExpr) isExpression()                {}

//=====================================================================================================================

// ParenthesizedExpr represents a parenthesized expression or comma-separated sequence of expressions.
type ParenthesizedExpr struct {
	SourcePosition SourcePos
	InnerExpr      IExpression
}

func (e *ParenthesizedExpr) GetSourcePosition() SourcePos { return e.SourcePosition }
func (e *ParenthesizedExpr) isExpression()                {}

//=====================================================================================================================

// QualifyExpr represents a type qualification (":") operation.
type QualifyExpr struct {
	SourcePosition SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *QualifyExpr) GetSourcePosition() SourcePos { return e.SourcePosition }
func (e *QualifyExpr) isExpression()                {}

//=====================================================================================================================

// RangeExpr represents a range ("..") operation.
type RangeExpr struct {
	SourcePosition SourcePos
	First          IExpression
	Last           IExpression
}

func (e *RangeExpr) GetSourcePosition() SourcePos { return e.SourcePosition }
func (e *RangeExpr) isExpression()                {}

//=====================================================================================================================

// RecordExpr represents a record literal expression.
type RecordExpr struct {
	SourcePosition SourcePos
	Items          []IExpression
}

func (e *RecordExpr) GetSourcePosition() SourcePos { return e.SourcePosition }
func (e *RecordExpr) isExpression()                {}

//=====================================================================================================================

// SequenceLiteralExpr represents an array literal.
type SequenceLiteralExpr struct {
	SourcePosition SourcePos
	Elements       []IExpression
}

func (e *SequenceLiteralExpr) GetSourcePosition() SourcePos { return e.SourcePosition }
func (e *SequenceLiteralExpr) isExpression()                {}

//=====================================================================================================================

// StringLiteralExpr represents a single string literal.
type StringLiteralExpr struct {
	SourcePosition SourcePos
	Delimiters     StringDelimiters
}

func (e *StringLiteralExpr) GetSourcePosition() SourcePos { return e.SourcePosition }
func (e *StringLiteralExpr) isExpression()                {}

//=====================================================================================================================

// SubtractionExpr represents a subtraction ("-") operation.
type SubtractionExpr struct {
	SourcePosition SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *SubtractionExpr) GetSourcePosition() SourcePos { return e.SourcePosition }
func (e *SubtractionExpr) isExpression()                {}

//=====================================================================================================================

// TrailingDocumentationExpr represents lines of trailing documentation.
type TrailingDocumentationExpr struct {
	SourcePosition SourcePos
}

func (e *TrailingDocumentationExpr) GetSourcePosition() SourcePos { return e.SourcePosition }
func (e *TrailingDocumentationExpr) isExpression()                {}

//=====================================================================================================================

// UnionExpr represents a type union ("|") operation.
type UnionExpr struct {
	SourcePosition SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *UnionExpr) GetSourcePosition() SourcePos { return e.SourcePosition }
func (e *UnionExpr) isExpression()                {}

//=====================================================================================================================

// UnitExpr represents a parenthesized expression with nothing in it.
type UnitExpr struct {
	SourcePosition SourcePos
}

func (e *UnitExpr) GetSourcePosition() SourcePos { return e.SourcePosition }
func (e *UnitExpr) isExpression()                {}

//=====================================================================================================================

// WhenExpr represents a when ("when") operation.
type WhenExpr struct {
	SourcePosition SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *WhenExpr) GetSourcePosition() SourcePos { return e.SourcePosition }
func (e *WhenExpr) isExpression()                {}

//=====================================================================================================================

// WhereExpr represents a type where ("where") operation.
type WhereExpr struct {
	SourcePosition SourcePos
	Lhs            IExpression
	Rhs            IExpression
}

func (e *WhereExpr) GetSourcePosition() SourcePos { return e.SourcePosition }
func (e *WhereExpr) isExpression()                {}

//=====================================================================================================================
