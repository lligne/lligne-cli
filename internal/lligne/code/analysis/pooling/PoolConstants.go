//
// # Data types related to Lligne token scanning.
//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package pooling

import (
	"fmt"
	prior "lligne-cli/internal/lligne/code/parsing"
	"lligne-cli/internal/lligne/runtime/pools"
)

//=====================================================================================================================

type Outcome struct {
	SourceCode      string
	NewLineOffsets  []uint32
	Model           IExpression
	StringConstants *pools.StringConstantPool
	IdentifierNames *pools.StringConstantPool
}

//=====================================================================================================================

func PoolConstants(priorOutcome *prior.Outcome) *Outcome {

	pooler := newPooler(priorOutcome)
	model := pooler.poolConstants(priorOutcome.Model)

	return &Outcome{
		SourceCode:      priorOutcome.SourceCode,
		NewLineOffsets:  priorOutcome.NewLineOffsets,
		Model:           model,
		StringConstants: pooler.StringConstants.Freeze(),
		IdentifierNames: pooler.IdentifierNames.Freeze(),
	}
}

//=====================================================================================================================

type pooler struct {
	SourceCode      string
	NewLineOffsets  []uint32
	StringConstants *pools.StringPool
	IdentifierNames *pools.StringPool
}

//---------------------------------------------------------------------------------------------------------------------

func newPooler(priorOutcome *prior.Outcome) *pooler {
	return &pooler{
		SourceCode:      priorOutcome.SourceCode,
		NewLineOffsets:  priorOutcome.NewLineOffsets,
		StringConstants: pools.NewStringPool(),
		IdentifierNames: pools.NewStringPool(),
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (p *pooler) poolConstants(expression prior.IExpression) IExpression {

	switch expr := expression.(type) {

	case *prior.AdditionExpr:
		return p.poolAdditionExpr(expr)
	case *prior.BooleanLiteralExpr:
		return p.poolBooleanLiteralExpr(expr)
	case *prior.BuiltInTypeExpr:
		return p.poolBuiltInTypeExpr(expr)
	case *prior.DivisionExpr:
		return p.poolDivisionExpr(expr)
	case *prior.EqualsExpr:
		return p.poolEqualsExpr(expr)
	case *prior.FieldReferenceExpr:
		return p.poolFieldReferenceExpr(expr)
	case *prior.Float64LiteralExpr:
		return p.poolFloatingPointLiteralExpr(expr)
	case *prior.GreaterThanExpr:
		return p.poolGreaterThanExpr(expr)
	case *prior.GreaterThanOrEqualsExpr:
		return p.poolGreaterThanOrEqualsExpr(expr)
	case *prior.IdentifierExpr:
		return p.poolIdentifierExpr(expr)
	case *prior.Int64LiteralExpr:
		return p.poolIntegerLiteralExpr(expr)
	case *prior.IntersectAssignValueExpr:
		return p.poolIntersectAssignValueExpr(expr)
	case *prior.IsExpr:
		return p.poolIsExpr(expr)
	case *prior.LessThanExpr:
		return p.poolLessThanExpr(expr)
	case *prior.LessThanOrEqualsExpr:
		return p.poolLessThanOrEqualsExpr(expr)
	case *prior.LogicalAndExpr:
		return p.poolLogicalAndExpr(expr)
	case *prior.LogicalNotOperationExpr:
		return p.poolLogicalNotOperationExpr(expr)
	case *prior.LogicalOrExpr:
		return p.poolLogicalOrExpr(expr)
	case *prior.MultiplicationExpr:
		return p.poolMultiplicationExpr(expr)
	case *prior.NegationOperationExpr:
		return p.poolNegationOperationExpr(expr)
	case *prior.NotEqualsExpr:
		return p.poolNotEqualsExpr(expr)
	case *prior.ParenthesizedExpr:
		return p.poolParenthesizedExpr(expr)
	case *prior.RecordExpr:
		return p.poolRecordExpr(expr)
	case *prior.StringLiteralExpr:
		return p.poolStringLiteralExpr(expr)
	case *prior.SubtractionExpr:
		return p.poolSubtractionExpr(expr)
	case *prior.WhereExpr:
		return p.poolWhereExpr(expr)

	default:
		panic(fmt.Sprintf("Missing case in poolConstants: %T\n", expression))

	}

}

//---------------------------------------------------------------------------------------------------------------------

func (p *pooler) poolAdditionExpr(expr *prior.AdditionExpr) IExpression {
	lhs := p.poolConstants(expr.Lhs)
	rhs := p.poolConstants(expr.Rhs)
	return &AdditionExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (p *pooler) poolBooleanLiteralExpr(expr *prior.BooleanLiteralExpr) IExpression {
	return &BooleanLiteralExpr{
		SourcePosition: expr.SourcePosition,
		Value:          expr.Value,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (p *pooler) poolBuiltInTypeExpr(expr *prior.BuiltInTypeExpr) IExpression {
	return &BuiltInTypeExpr{
		SourcePosition: expr.SourcePosition,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (p *pooler) poolDivisionExpr(expr *prior.DivisionExpr) IExpression {
	lhs := p.poolConstants(expr.Lhs)
	rhs := p.poolConstants(expr.Rhs)
	return &DivisionExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (p *pooler) poolEqualsExpr(expr *prior.EqualsExpr) IExpression {
	lhs := p.poolConstants(expr.Lhs)
	rhs := p.poolConstants(expr.Rhs)
	return &EqualsExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (p *pooler) poolFieldReferenceExpr(expr *prior.FieldReferenceExpr) IExpression {
	parent := p.poolConstants(expr.Parent)
	child := p.poolConstants(expr.Child)
	return &FieldReferenceExpr{
		SourcePosition: expr.SourcePosition,
		Parent:         parent,
		Child:          child,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (p *pooler) poolFloatingPointLiteralExpr(expr *prior.Float64LiteralExpr) IExpression {
	return &Float64LiteralExpr{
		SourcePosition: expr.SourcePosition,
		Value:          expr.Value,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (p *pooler) poolGreaterThanExpr(expr *prior.GreaterThanExpr) IExpression {
	lhs := p.poolConstants(expr.Lhs)
	rhs := p.poolConstants(expr.Rhs)
	return &GreaterThanExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (p *pooler) poolGreaterThanOrEqualsExpr(expr *prior.GreaterThanOrEqualsExpr) IExpression {
	lhs := p.poolConstants(expr.Lhs)
	rhs := p.poolConstants(expr.Rhs)
	return &GreaterThanOrEqualsExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (p *pooler) poolIdentifierExpr(expr *prior.IdentifierExpr) IExpression {
	name := expr.SourcePosition.GetText(p.SourceCode)
	nameIndex := p.IdentifierNames.Put(name)

	return &IdentifierExpr{
		SourcePosition: expr.SourcePosition,
		NameIndex:      nameIndex,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (p *pooler) poolIntegerLiteralExpr(expr *prior.Int64LiteralExpr) IExpression {
	return &Int64LiteralExpr{
		SourcePosition: expr.SourcePosition,
		Value:          expr.Value,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (p *pooler) poolIntersectAssignValueExpr(expr *prior.IntersectAssignValueExpr) IExpression {
	lhs := p.poolConstants(expr.Lhs)
	rhs := p.poolConstants(expr.Rhs)
	return &IntersectAssignValueExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (p *pooler) poolIsExpr(expr *prior.IsExpr) IExpression {
	lhs := p.poolConstants(expr.Lhs)
	rhs := p.poolConstants(expr.Rhs)
	return &IsExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (p *pooler) poolLessThanExpr(expr *prior.LessThanExpr) IExpression {
	lhs := p.poolConstants(expr.Lhs)
	rhs := p.poolConstants(expr.Rhs)
	return &LessThanExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (p *pooler) poolLessThanOrEqualsExpr(expr *prior.LessThanOrEqualsExpr) IExpression {
	lhs := p.poolConstants(expr.Lhs)
	rhs := p.poolConstants(expr.Rhs)
	return &LessThanOrEqualsExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (p *pooler) poolLogicalAndExpr(expr *prior.LogicalAndExpr) IExpression {
	lhs := p.poolConstants(expr.Lhs)
	rhs := p.poolConstants(expr.Rhs)
	return &LogicalAndExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (p *pooler) poolLogicalNotOperationExpr(expr *prior.LogicalNotOperationExpr) IExpression {
	operand := p.poolConstants(expr.Operand)
	return &LogicalNotOperationExpr{
		SourcePosition: expr.SourcePosition,
		Operand:        operand,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (p *pooler) poolLogicalOrExpr(expr *prior.LogicalOrExpr) IExpression {
	lhs := p.poolConstants(expr.Lhs)
	rhs := p.poolConstants(expr.Rhs)
	return &LogicalOrExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (p *pooler) poolMultiplicationExpr(expr *prior.MultiplicationExpr) IExpression {
	lhs := p.poolConstants(expr.Lhs)
	rhs := p.poolConstants(expr.Rhs)
	return &MultiplicationExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (p *pooler) poolNegationOperationExpr(expr *prior.NegationOperationExpr) IExpression {
	operand := p.poolConstants(expr.Operand)
	return &NegationOperationExpr{
		SourcePosition: expr.SourcePosition,
		Operand:        operand,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (p *pooler) poolNotEqualsExpr(expr *prior.NotEqualsExpr) IExpression {
	lhs := p.poolConstants(expr.Lhs)
	rhs := p.poolConstants(expr.Rhs)
	return &NotEqualsExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (p *pooler) poolParenthesizedExpr(expr *prior.ParenthesizedExpr) IExpression {
	inner := p.poolConstants(expr.InnerExpr)
	return &ParenthesizedExpr{
		SourcePosition: expr.SourcePosition,
		InnerExpr:      inner,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (p *pooler) poolRecordExpr(expr *prior.RecordExpr) IExpression {
	items := make([]IExpression, 0)
	for _, item := range expr.Items {
		items = append(items, p.poolConstants(item))
	}

	return &RecordExpr{
		SourcePosition: expr.SourcePosition,
		Items:          items,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (p *pooler) poolStringLiteralExpr(expr *prior.StringLiteralExpr) IExpression {
	text := expr.SourcePosition.GetText(p.SourceCode)
	var value string

	switch expr.Delimiters {
	case prior.StringDelimitersDoubleQuotes:
		value = text[1 : len(text)-1]
	case prior.StringDelimitersSingleQuotes:
		value = text[1 : len(text)-1]
	default:
		panic("TODO: Unhandled string delimiters")
	}

	valueIndex := p.StringConstants.Put(value)

	// TODO: escape chars
	return &StringLiteralExpr{
		SourcePosition: expr.SourcePosition,
		ValueIndex:     valueIndex,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (p *pooler) poolSubtractionExpr(expr *prior.SubtractionExpr) IExpression {
	lhs := p.poolConstants(expr.Lhs)
	rhs := p.poolConstants(expr.Rhs)
	return &SubtractionExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (p *pooler) poolWhereExpr(expr *prior.WhereExpr) IExpression {
	lhs := p.poolConstants(expr.Lhs)
	rhs := p.poolConstants(expr.Rhs)
	return &WhereExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------
