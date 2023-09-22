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
	IdentifierNames *pools.NameConstantPool
}

//=====================================================================================================================

func StructureRecords(priorOutcome *prior.Outcome) *Outcome {

	s := newStructurer(priorOutcome)

	model := s.structureRecords(priorOutcome.Model)

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
	IdentifierNames *pools.NameConstantPool
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
	expression prior.IExpression,
) IExpression {

	switch expr := expression.(type) {

	case *prior.AdditionExpr:
		return s.structureAdditionExpr(expr)
	case *prior.BooleanLiteralExpr:
		return s.structureBooleanLiteralExpr(expr)
	case *prior.BuiltInTypeExpr:
		return s.structureBuiltInTypeExpr(expr)
	case *prior.DivisionExpr:
		return s.structureDivisionExpr(expr)
	case *prior.EqualsExpr:
		return s.structureEqualsExpr(expr)
	case *prior.FieldReferenceExpr:
		return s.structureFieldReferenceExpr(expr)
	case *prior.Float64LiteralExpr:
		return s.structureFloatingPointLiteralExpr(expr)
	case *prior.GreaterThanExpr:
		return s.structureGreaterThanExpr(expr)
	case *prior.GreaterThanOrEqualsExpr:
		return s.structureGreaterThanOrEqualsExpr(expr)
	case *prior.IdentifierExpr:
		return s.structureIdentifierExpr(expr)
	case *prior.Int64LiteralExpr:
		return s.structureIntegerLiteralExpr(expr)
	case *prior.IsExpr:
		return s.structureIsExpr(expr)
	case *prior.LessThanExpr:
		return s.structureLessThanExpr(expr)
	case *prior.LessThanOrEqualsExpr:
		return s.structureLessThanOrEqualsExpr(expr)
	case *prior.LogicalAndExpr:
		return s.structureLogicalAndExpr(expr)
	case *prior.LogicalNotOperationExpr:
		return s.structureLogicalNotOperationExpr(expr)
	case *prior.LogicalOrExpr:
		return s.structureLogicalOrExpr(expr)
	case *prior.MultiplicationExpr:
		return s.structureMultiplicationExpr(expr)
	case *prior.NegationOperationExpr:
		return s.structureNegationOperationExpr(expr)
	case *prior.NotEqualsExpr:
		return s.structureNotEqualsExpr(expr)
	case *prior.ParenthesizedExpr:
		return s.structureParenthesizedExpr(expr)
	case *prior.RecordExpr:
		return s.structureRecordExpr(expr)
	case *prior.StringLiteralExpr:
		return s.structureStringLiteralExpr(expr)
	case *prior.SubtractionExpr:
		return s.structureSubtractionExpr(expr)
	case *prior.WhereExpr:
		return s.structureWhereExpr(expr)

	default:
		panic(fmt.Sprintf("Missing case in structureRecords: %T\n", expression))

	}

}

//---------------------------------------------------------------------------------------------------------------------

func (s *structurer) structureAdditionExpr(
	expr *prior.AdditionExpr,
) IExpression {
	lhs := s.structureRecords(expr.Lhs)
	rhs := s.structureRecords(expr.Rhs)
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
	expr *prior.DivisionExpr,
) IExpression {
	lhs := s.structureRecords(expr.Lhs)
	rhs := s.structureRecords(expr.Rhs)
	return &DivisionExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *structurer) structureEqualsExpr(
	expr *prior.EqualsExpr,
) IExpression {
	lhs := s.structureRecords(expr.Lhs)
	rhs := s.structureRecords(expr.Rhs)
	return &EqualsExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *structurer) structureFieldReferenceExpr(
	expr *prior.FieldReferenceExpr,
) IExpression {
	parent := s.structureRecords(expr.Parent)
	child := s.structureRecords(expr.Child)
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
	expr *prior.GreaterThanExpr,
) IExpression {
	lhs := s.structureRecords(expr.Lhs)
	rhs := s.structureRecords(expr.Rhs)
	return &GreaterThanExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *structurer) structureGreaterThanOrEqualsExpr(
	expr *prior.GreaterThanOrEqualsExpr,
) IExpression {
	lhs := s.structureRecords(expr.Lhs)
	rhs := s.structureRecords(expr.Rhs)
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
	expr *prior.IsExpr,
) IExpression {
	lhs := s.structureRecords(expr.Lhs)
	rhs := s.structureRecords(expr.Rhs)
	return &IsExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *structurer) structureLessThanExpr(
	expr *prior.LessThanExpr,
) IExpression {
	lhs := s.structureRecords(expr.Lhs)
	rhs := s.structureRecords(expr.Rhs)
	return &LessThanExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *structurer) structureLessThanOrEqualsExpr(
	expr *prior.LessThanOrEqualsExpr,
) IExpression {
	lhs := s.structureRecords(expr.Lhs)
	rhs := s.structureRecords(expr.Rhs)
	return &LessThanOrEqualsExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *structurer) structureLogicalAndExpr(
	expr *prior.LogicalAndExpr,
) IExpression {
	lhs := s.structureRecords(expr.Lhs)
	rhs := s.structureRecords(expr.Rhs)
	return &LogicalAndExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *structurer) structureLogicalNotOperationExpr(
	expr *prior.LogicalNotOperationExpr,
) IExpression {
	operand := s.structureRecords(expr.Operand)
	return &LogicalNotOperationExpr{
		SourcePosition: expr.SourcePosition,
		Operand:        operand,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *structurer) structureLogicalOrExpr(
	expr *prior.LogicalOrExpr,
) IExpression {
	lhs := s.structureRecords(expr.Lhs)
	rhs := s.structureRecords(expr.Rhs)
	return &LogicalOrExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *structurer) structureMultiplicationExpr(
	expr *prior.MultiplicationExpr,
) IExpression {
	lhs := s.structureRecords(expr.Lhs)
	rhs := s.structureRecords(expr.Rhs)
	return &MultiplicationExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *structurer) structureNegationOperationExpr(
	expr *prior.NegationOperationExpr,
) IExpression {
	operand := s.structureRecords(expr.Operand)
	return &NegationOperationExpr{
		SourcePosition: expr.SourcePosition,
		Operand:        operand,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *structurer) structureNotEqualsExpr(
	expr *prior.NotEqualsExpr,
) IExpression {
	lhs := s.structureRecords(expr.Lhs)
	rhs := s.structureRecords(expr.Rhs)
	return &NotEqualsExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *structurer) structureParenthesizedExpr(
	expr *prior.ParenthesizedExpr,
) IExpression {
	inner := s.structureRecords(expr.InnerExpr)
	return &ParenthesizedExpr{
		SourcePosition: expr.SourcePosition,
		InnerExpr:      inner,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *structurer) structureRecordExpr(
	expr *prior.RecordExpr,
) IExpression {
	items := make([]*RecordFieldExpr, 0)
	for _, item := range expr.Items {
		fieldExpr := s.structureRecordFieldExpr(item)
		items = append(items, fieldExpr)
	}

	return &RecordExpr{
		SourcePosition: expr.SourcePosition,
		Fields:         items,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *structurer) structureRecordFieldExpr(
	expr prior.IExpression,
) *RecordFieldExpr {

	// TODO: Qualify and other intersection expressions

	switch fieldExpr := expr.(type) {
	case *prior.IntersectAssignValueExpr:
		fieldNameIndex := fieldExpr.Lhs.(*prior.IdentifierExpr).NameIndex
		value := s.structureRecords(fieldExpr.Rhs)
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
	expr *prior.SubtractionExpr,
) IExpression {
	lhs := s.structureRecords(expr.Lhs)
	rhs := s.structureRecords(expr.Rhs)
	return &SubtractionExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *structurer) structureWhereExpr(
	expr *prior.WhereExpr,
) IExpression {
	lhs := s.structureRecords(expr.Lhs)
	rhs := s.structureRecords(expr.Rhs)
	return &WhereExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------
