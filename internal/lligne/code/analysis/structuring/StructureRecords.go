//
// # Data types related to Lligne token scanning.
//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package structuring

import (
	"fmt"
	prior "lligne-cli/internal/lligne/code/analysis/pooling"
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

func StructureRecords(priorOutcome *prior.Outcome) *Outcome {

	s := newStructurer(priorOutcome)

	model := s.structureRecords(priorOutcome, priorOutcome.Model)

	return &Outcome{
		SourceCode:      priorOutcome.SourceCode,
		NewLineOffsets:  priorOutcome.NewLineOffsets,
		Model:           model,
		StringConstants: priorOutcome.StringConstants,
		IdentifierNames: priorOutcome.IdentifierNames,
	}
}

//=====================================================================================================================

type structurer struct {
	SourceCode      string
	NewLineOffsets  []uint32
	StringConstants *pools.StringConstantPool
	IdentifierNames *pools.StringConstantPool
}

//---------------------------------------------------------------------------------------------------------------------

func newStructurer(priorOutcome *prior.Outcome) *structurer {
	return &structurer{
		SourceCode:      priorOutcome.SourceCode,
		NewLineOffsets:  priorOutcome.NewLineOffsets,
		StringConstants: priorOutcome.StringConstants,
		IdentifierNames: priorOutcome.IdentifierNames,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *structurer) structureRecords(
	priorOutcome *prior.Outcome,
	expression prior.IExpression,
) IExpression {

	switch expr := expression.(type) {

	case *prior.AdditionExpr:
		return s.structureAdditionExpr(priorOutcome, expr)
	case *prior.BooleanLiteralExpr:
		return s.structureBooleanLiteralExpr(expr)
	case *prior.BuiltInTypeExpr:
		return s.structureBuiltInTypeExpr(expr)
	case *prior.DivisionExpr:
		return s.structureDivisionExpr(priorOutcome, expr)
	case *prior.EqualsExpr:
		return s.structureEqualsExpr(priorOutcome, expr)
	case *prior.FieldReferenceExpr:
		return s.structureFieldReferenceExpr(priorOutcome, expr)
	case *prior.Float64LiteralExpr:
		return s.structureFloatingPointLiteralExpr(expr)
	case *prior.GreaterThanExpr:
		return s.structureGreaterThanExpr(priorOutcome, expr)
	case *prior.GreaterThanOrEqualsExpr:
		return s.structureGreaterThanOrEqualsExpr(priorOutcome, expr)
	case *prior.IdentifierExpr:
		return s.structureIdentifierExpr(expr)
	case *prior.Int64LiteralExpr:
		return s.structureIntegerLiteralExpr(expr)
	case *prior.IsExpr:
		return s.structureIsExpr(priorOutcome, expr)
	case *prior.LessThanExpr:
		return s.structureLessThanExpr(priorOutcome, expr)
	case *prior.LessThanOrEqualsExpr:
		return s.structureLessThanOrEqualsExpr(priorOutcome, expr)
	case *prior.LogicalAndExpr:
		return s.structureLogicalAndExpr(priorOutcome, expr)
	case *prior.LogicalNotOperationExpr:
		return s.structureLogicalNotOperationExpr(priorOutcome, expr)
	case *prior.LogicalOrExpr:
		return s.structureLogicalOrExpr(priorOutcome, expr)
	case *prior.MultiplicationExpr:
		return s.structureMultiplicationExpr(priorOutcome, expr)
	case *prior.NegationOperationExpr:
		return s.structureNegationOperationExpr(priorOutcome, expr)
	case *prior.NotEqualsExpr:
		return s.structureNotEqualsExpr(priorOutcome, expr)
	case *prior.ParenthesizedExpr:
		return s.structureParenthesizedExpr(priorOutcome, expr)
	case *prior.RecordExpr:
		return s.structureRecordExpr(priorOutcome, expr)
	case *prior.StringLiteralExpr:
		return s.structureStringLiteralExpr(expr)
	case *prior.SubtractionExpr:
		return s.structureSubtractionExpr(priorOutcome, expr)

	default:
		panic(fmt.Sprintf("Missing case in structureRecords: %T\n", expression))

	}

}

//---------------------------------------------------------------------------------------------------------------------

func (s *structurer) structureAdditionExpr(
	priorOutcome *prior.Outcome,
	expr *prior.AdditionExpr,
) IExpression {
	lhs := s.structureRecords(priorOutcome, expr.Lhs)
	rhs := s.structureRecords(priorOutcome, expr.Rhs)
	return &AdditionExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *structurer) structureBooleanLiteralExpr(expr *prior.BooleanLiteralExpr) IExpression {
	return &BooleanLiteralExpr{
		SourcePosition: expr.SourcePosition,
		Value:          expr.Value,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *structurer) structureBuiltInTypeExpr(expr *prior.BuiltInTypeExpr) IExpression {
	return &BuiltInTypeExpr{
		SourcePosition: expr.SourcePosition,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *structurer) structureDivisionExpr(
	priorOutcome *prior.Outcome,
	expr *prior.DivisionExpr,
) IExpression {
	lhs := s.structureRecords(priorOutcome, expr.Lhs)
	rhs := s.structureRecords(priorOutcome, expr.Rhs)
	return &DivisionExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *structurer) structureEqualsExpr(
	priorOutcome *prior.Outcome,
	expr *prior.EqualsExpr,
) IExpression {
	lhs := s.structureRecords(priorOutcome, expr.Lhs)
	rhs := s.structureRecords(priorOutcome, expr.Rhs)
	return &EqualsExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *structurer) structureFieldReferenceExpr(
	priorOutcome *prior.Outcome,
	expr *prior.FieldReferenceExpr,
) IExpression {
	parent := s.structureRecords(priorOutcome, expr.Parent)
	child := s.structureRecords(priorOutcome, expr.Child)
	return &FieldReferenceExpr{
		SourcePosition: expr.SourcePosition,
		Parent:         parent,
		Child:          child,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *structurer) structureFloatingPointLiteralExpr(expr *prior.Float64LiteralExpr) IExpression {
	return &Float64LiteralExpr{
		SourcePosition: expr.SourcePosition,
		Value:          expr.Value,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *structurer) structureGreaterThanExpr(
	priorOutcome *prior.Outcome,
	expr *prior.GreaterThanExpr,
) IExpression {
	lhs := s.structureRecords(priorOutcome, expr.Lhs)
	rhs := s.structureRecords(priorOutcome, expr.Rhs)
	return &GreaterThanExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *structurer) structureGreaterThanOrEqualsExpr(
	priorOutcome *prior.Outcome,
	expr *prior.GreaterThanOrEqualsExpr,
) IExpression {
	lhs := s.structureRecords(priorOutcome, expr.Lhs)
	rhs := s.structureRecords(priorOutcome, expr.Rhs)
	return &GreaterThanOrEqualsExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *structurer) structureIdentifierExpr(
	expr *prior.IdentifierExpr,
) IExpression {
	return &IdentifierExpr{
		SourcePosition: expr.SourcePosition,
		NameIndex:      expr.NameIndex,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *structurer) structureIntegerLiteralExpr(expr *prior.Int64LiteralExpr) IExpression {
	return &Int64LiteralExpr{
		SourcePosition: expr.SourcePosition,
		Value:          expr.Value,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *structurer) structureIsExpr(
	priorOutcome *prior.Outcome,
	expr *prior.IsExpr,
) IExpression {
	lhs := s.structureRecords(priorOutcome, expr.Lhs)
	rhs := s.structureRecords(priorOutcome, expr.Rhs)
	return &IsExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *structurer) structureLessThanExpr(
	priorOutcome *prior.Outcome,
	expr *prior.LessThanExpr,
) IExpression {
	lhs := s.structureRecords(priorOutcome, expr.Lhs)
	rhs := s.structureRecords(priorOutcome, expr.Rhs)
	return &LessThanExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *structurer) structureLessThanOrEqualsExpr(
	priorOutcome *prior.Outcome,
	expr *prior.LessThanOrEqualsExpr,
) IExpression {
	lhs := s.structureRecords(priorOutcome, expr.Lhs)
	rhs := s.structureRecords(priorOutcome, expr.Rhs)
	return &LessThanOrEqualsExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *structurer) structureLogicalAndExpr(
	priorOutcome *prior.Outcome,
	expr *prior.LogicalAndExpr,
) IExpression {
	lhs := s.structureRecords(priorOutcome, expr.Lhs)
	rhs := s.structureRecords(priorOutcome, expr.Rhs)
	return &LogicalAndExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *structurer) structureLogicalNotOperationExpr(
	priorOutcome *prior.Outcome,
	expr *prior.LogicalNotOperationExpr,
) IExpression {
	operand := s.structureRecords(priorOutcome, expr.Operand)
	return &LogicalNotOperationExpr{
		SourcePosition: expr.SourcePosition,
		Operand:        operand,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *structurer) structureLogicalOrExpr(
	priorOutcome *prior.Outcome,
	expr *prior.LogicalOrExpr,
) IExpression {
	lhs := s.structureRecords(priorOutcome, expr.Lhs)
	rhs := s.structureRecords(priorOutcome, expr.Rhs)
	return &LogicalOrExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *structurer) structureMultiplicationExpr(
	priorOutcome *prior.Outcome,
	expr *prior.MultiplicationExpr,
) IExpression {
	lhs := s.structureRecords(priorOutcome, expr.Lhs)
	rhs := s.structureRecords(priorOutcome, expr.Rhs)
	return &MultiplicationExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *structurer) structureNegationOperationExpr(
	priorOutcome *prior.Outcome,
	expr *prior.NegationOperationExpr,
) IExpression {
	operand := s.structureRecords(priorOutcome, expr.Operand)
	return &NegationOperationExpr{
		SourcePosition: expr.SourcePosition,
		Operand:        operand,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *structurer) structureNotEqualsExpr(
	priorOutcome *prior.Outcome,
	expr *prior.NotEqualsExpr,
) IExpression {
	lhs := s.structureRecords(priorOutcome, expr.Lhs)
	rhs := s.structureRecords(priorOutcome, expr.Rhs)
	return &NotEqualsExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *structurer) structureParenthesizedExpr(
	priorOutcome *prior.Outcome,
	expr *prior.ParenthesizedExpr,
) IExpression {
	inner := s.structureRecords(priorOutcome, expr.InnerExpr)
	return &ParenthesizedExpr{
		SourcePosition: expr.SourcePosition,
		InnerExpr:      inner,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *structurer) structureRecordExpr(
	priorOutcome *prior.Outcome,
	expr *prior.RecordExpr,
) IExpression {
	items := make([]*RecordFieldExpr, 0)
	for _, item := range expr.Items {
		fieldExpr := s.structureRecordFieldExpr(priorOutcome, item)
		items = append(items, fieldExpr)
	}

	return &RecordExpr{
		SourcePosition: expr.SourcePosition,
		Fields:         items,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *structurer) structureRecordFieldExpr(
	priorOutcome *prior.Outcome,
	expr prior.IExpression,
) *RecordFieldExpr {

	// TODO: Qualify and other intersection expressions

	switch fieldExpr := expr.(type) {
	case *prior.IntersectAssignValueExpr:
		fieldNameIndex := fieldExpr.Lhs.(*prior.IdentifierExpr).NameIndex
		value := s.structureRecords(priorOutcome, fieldExpr.Rhs)
		return &RecordFieldExpr{
			SourcePosition: expr.GetSourcePosition(),
			FieldNameIndex: fieldNameIndex,
			FieldValue:     value,
		}
	default:
		panic(fmt.Sprintf("Missing case in structureRecordFieldExpr: %T\n", expr))
	}

}

//---------------------------------------------------------------------------------------------------------------------

func (s *structurer) structureStringLiteralExpr(
	expr *prior.StringLiteralExpr,
) IExpression {
	return &StringLiteralExpr{
		SourcePosition: expr.SourcePosition,
		ValueIndex:     expr.ValueIndex,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *structurer) structureSubtractionExpr(
	priorOutcome *prior.Outcome,
	expr *prior.SubtractionExpr,
) IExpression {
	lhs := s.structureRecords(priorOutcome, expr.Lhs)
	rhs := s.structureRecords(priorOutcome, expr.Rhs)
	return &SubtractionExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------
