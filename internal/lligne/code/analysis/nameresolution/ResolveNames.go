//
// # Data types related to Lligne token scanning.
//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package nameresolution

import (
	"fmt"
	prior "lligne-cli/internal/lligne/code/analysis/structuring"
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

func ResolveNames(priorOutcome *prior.Outcome) *Outcome {

	s := newNameResolver(priorOutcome)
	context := newNameResolutionContext()

	model := s.resolveNames(priorOutcome.Model, context)

	return &Outcome{
		SourceCode:      priorOutcome.SourceCode,
		NewLineOffsets:  priorOutcome.NewLineOffsets,
		Model:           model,
		StringConstants: priorOutcome.StringConstants,
		IdentifierNames: priorOutcome.IdentifierNames,
	}
}

//=====================================================================================================================

type nameResolver struct {
	SourceCode      string
	NewLineOffsets  []uint32
	StringConstants *pools.StringConstantPool
	IdentifierNames *pools.StringConstantPool
}

//---------------------------------------------------------------------------------------------------------------------

func newNameResolver(priorOutcome *prior.Outcome) *nameResolver {
	return &nameResolver{
		SourceCode:      priorOutcome.SourceCode,
		NewLineOffsets:  priorOutcome.NewLineOffsets,
		StringConstants: priorOutcome.StringConstants,
		IdentifierNames: priorOutcome.IdentifierNames,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *nameResolver) resolveNames(
	expression prior.IExpression,
	context *nameResolutionContext,
) IExpression {

	switch expr := expression.(type) {

	case *prior.AdditionExpr:
		return s.resolveAdditionExpr(expr, context)
	case *prior.BooleanLiteralExpr:
		return s.resolveBooleanLiteralExpr(expr)
	case *prior.BuiltInTypeExpr:
		return s.resolveBuiltInTypeExpr(expr)
	case *prior.DivisionExpr:
		return s.resolveDivisionExpr(expr, context)
	case *prior.EqualsExpr:
		return s.resolveEqualsExpr(expr, context)
	case *prior.FieldReferenceExpr:
		return s.resolveFieldReferenceExpr(expr, context)
	case *prior.Float64LiteralExpr:
		return s.resolveFloatingPointLiteralExpr(expr)
	case *prior.GreaterThanExpr:
		return s.resolveGreaterThanExpr(expr, context)
	case *prior.GreaterThanOrEqualsExpr:
		return s.resolveGreaterThanOrEqualsExpr(expr, context)
	case *prior.IdentifierExpr:
		return s.resolveIdentifierExpr(expr, context)
	case *prior.Int64LiteralExpr:
		return s.resolveIntegerLiteralExpr(expr)
	case *prior.IsExpr:
		return s.resolveIsExpr(expr, context)
	case *prior.LessThanExpr:
		return s.resolveLessThanExpr(expr, context)
	case *prior.LessThanOrEqualsExpr:
		return s.resolveLessThanOrEqualsExpr(expr, context)
	case *prior.LogicalAndExpr:
		return s.resolveLogicalAndExpr(expr, context)
	case *prior.LogicalNotOperationExpr:
		return s.resolveLogicalNotOperationExpr(expr, context)
	case *prior.LogicalOrExpr:
		return s.resolveLogicalOrExpr(expr, context)
	case *prior.MultiplicationExpr:
		return s.resolveMultiplicationExpr(expr, context)
	case *prior.NegationOperationExpr:
		return s.resolveNegationOperationExpr(expr, context)
	case *prior.NotEqualsExpr:
		return s.resolveNotEqualsExpr(expr, context)
	case *prior.ParenthesizedExpr:
		return s.resolveParenthesizedExpr(expr, context)
	case *prior.RecordExpr:
		return s.resolveRecordExpr(expr, context)
	case *prior.StringLiteralExpr:
		return s.resolveStringLiteralExpr(expr)
	case *prior.SubtractionExpr:
		return s.resolveSubtractionExpr(expr, context)
	case *prior.WhereExpr:
		return s.resolveWhereExpr(expr, context)

	default:
		panic(fmt.Sprintf("Missing case in resolveNames: %T\n", expression))

	}

}

//---------------------------------------------------------------------------------------------------------------------

func (s *nameResolver) resolveAdditionExpr(
	expr *prior.AdditionExpr,
	context *nameResolutionContext,
) IExpression {
	lhs := s.resolveNames(expr.Lhs, context)
	rhs := s.resolveNames(expr.Rhs, context)
	return &AdditionExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *nameResolver) resolveBooleanLiteralExpr(expr *prior.BooleanLiteralExpr) IExpression {
	return &BooleanLiteralExpr{
		SourcePosition: expr.SourcePosition,
		Value:          expr.Value,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *nameResolver) resolveBuiltInTypeExpr(expr *prior.BuiltInTypeExpr) IExpression {
	return &BuiltInTypeExpr{
		SourcePosition: expr.SourcePosition,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *nameResolver) resolveDivisionExpr(
	expr *prior.DivisionExpr,
	context *nameResolutionContext,
) IExpression {
	lhs := s.resolveNames(expr.Lhs, context)
	rhs := s.resolveNames(expr.Rhs, context)
	return &DivisionExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *nameResolver) resolveEqualsExpr(
	expr *prior.EqualsExpr,
	context *nameResolutionContext,
) IExpression {
	lhs := s.resolveNames(expr.Lhs, context)
	rhs := s.resolveNames(expr.Rhs, context)
	return &EqualsExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *nameResolver) resolveFieldReferenceExpr(
	expr *prior.FieldReferenceExpr,
	context *nameResolutionContext,
) IExpression {
	parent := s.resolveNames(expr.Parent, context)
	child := s.resolveNames(expr.Child, context.withFieldReferenceLhs(parent))
	return &FieldReferenceExpr{
		SourcePosition: expr.SourcePosition,
		Parent:         parent,
		Child:          child,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *nameResolver) resolveFloatingPointLiteralExpr(expr *prior.Float64LiteralExpr) IExpression {
	return &Float64LiteralExpr{
		SourcePosition: expr.SourcePosition,
		Value:          expr.Value,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *nameResolver) resolveGreaterThanExpr(
	expr *prior.GreaterThanExpr,
	context *nameResolutionContext,
) IExpression {
	lhs := s.resolveNames(expr.Lhs, context)
	rhs := s.resolveNames(expr.Rhs, context)
	return &GreaterThanExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *nameResolver) resolveGreaterThanOrEqualsExpr(
	expr *prior.GreaterThanOrEqualsExpr,
	context *nameResolutionContext,
) IExpression {
	lhs := s.resolveNames(expr.Lhs, context)
	rhs := s.resolveNames(expr.Rhs, context)
	return &GreaterThanOrEqualsExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *nameResolver) resolveIdentifierExpr(
	expr *prior.IdentifierExpr,
	context *nameResolutionContext,
) IExpression {
	return &IdentifierExpr{
		SourcePosition: expr.SourcePosition,
		NameIndex:      expr.NameIndex,
		// TODO: fields from resolving the name
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *nameResolver) resolveIntegerLiteralExpr(expr *prior.Int64LiteralExpr) IExpression {
	return &Int64LiteralExpr{
		SourcePosition: expr.SourcePosition,
		Value:          expr.Value,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *nameResolver) resolveIsExpr(
	expr *prior.IsExpr,
	context *nameResolutionContext,
) IExpression {
	lhs := s.resolveNames(expr.Lhs, context)
	rhs := s.resolveNames(expr.Rhs, context)
	return &IsExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *nameResolver) resolveLessThanExpr(
	expr *prior.LessThanExpr,
	context *nameResolutionContext,
) IExpression {
	lhs := s.resolveNames(expr.Lhs, context)
	rhs := s.resolveNames(expr.Rhs, context)
	return &LessThanExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *nameResolver) resolveLessThanOrEqualsExpr(
	expr *prior.LessThanOrEqualsExpr,
	context *nameResolutionContext,
) IExpression {
	lhs := s.resolveNames(expr.Lhs, context)
	rhs := s.resolveNames(expr.Rhs, context)
	return &LessThanOrEqualsExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *nameResolver) resolveLogicalAndExpr(
	expr *prior.LogicalAndExpr,
	context *nameResolutionContext,
) IExpression {
	lhs := s.resolveNames(expr.Lhs, context)
	rhs := s.resolveNames(expr.Rhs, context)
	return &LogicalAndExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *nameResolver) resolveLogicalNotOperationExpr(
	expr *prior.LogicalNotOperationExpr,
	context *nameResolutionContext,
) IExpression {
	operand := s.resolveNames(expr.Operand, context)
	return &LogicalNotOperationExpr{
		SourcePosition: expr.SourcePosition,
		Operand:        operand,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *nameResolver) resolveLogicalOrExpr(
	expr *prior.LogicalOrExpr,
	context *nameResolutionContext,
) IExpression {
	lhs := s.resolveNames(expr.Lhs, context)
	rhs := s.resolveNames(expr.Rhs, context)
	return &LogicalOrExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *nameResolver) resolveMultiplicationExpr(
	expr *prior.MultiplicationExpr,
	context *nameResolutionContext,
) IExpression {
	lhs := s.resolveNames(expr.Lhs, context)
	rhs := s.resolveNames(expr.Rhs, context)
	return &MultiplicationExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *nameResolver) resolveNegationOperationExpr(
	expr *prior.NegationOperationExpr,
	context *nameResolutionContext,
) IExpression {
	operand := s.resolveNames(expr.Operand, context)
	return &NegationOperationExpr{
		SourcePosition: expr.SourcePosition,
		Operand:        operand,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *nameResolver) resolveNotEqualsExpr(
	expr *prior.NotEqualsExpr,
	context *nameResolutionContext,
) IExpression {
	lhs := s.resolveNames(expr.Lhs, context)
	rhs := s.resolveNames(expr.Rhs, context)
	return &NotEqualsExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *nameResolver) resolveParenthesizedExpr(
	expr *prior.ParenthesizedExpr,
	context *nameResolutionContext,
) IExpression {
	inner := s.resolveNames(expr.InnerExpr, context)
	return &ParenthesizedExpr{
		SourcePosition: expr.SourcePosition,
		InnerExpr:      inner,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *nameResolver) resolveRecordExpr(
	expr *prior.RecordExpr,
	context *nameResolutionContext,
) IExpression {
	fields := make([]*RecordFieldExpr, 0)
	fieldNameIndexes := make([]uint64, 0)

	for _, field := range expr.Fields {
		fieldNameIndexes = append(fieldNameIndexes, field.FieldNameIndex)
		fieldExpr := s.resolveRecordFieldExpr(field, context)
		fields = append(fields, fieldExpr)
	}

	return &RecordExpr{
		SourcePosition:   expr.SourcePosition,
		FieldNameIndexes: fieldNameIndexes,
		Fields:           fields,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *nameResolver) resolveRecordFieldExpr(
	expr *prior.RecordFieldExpr,
	context *nameResolutionContext,
) *RecordFieldExpr {
	return &RecordFieldExpr{
		SourcePosition: expr.GetSourcePosition(),
		FieldNameIndex: expr.FieldNameIndex,
		FieldValue:     s.resolveNames(expr.FieldValue, context),
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *nameResolver) resolveStringLiteralExpr(
	expr *prior.StringLiteralExpr,
) IExpression {
	return &StringLiteralExpr{
		SourcePosition: expr.SourcePosition,
		ValueIndex:     expr.ValueIndex,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *nameResolver) resolveSubtractionExpr(
	expr *prior.SubtractionExpr,
	context *nameResolutionContext,
) IExpression {
	lhs := s.resolveNames(expr.Lhs, context)
	rhs := s.resolveNames(expr.Rhs, context)
	return &SubtractionExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *nameResolver) resolveWhereExpr(
	expr *prior.WhereExpr,
	context *nameResolutionContext,
) IExpression {
	rhs := s.resolveNames(expr.Rhs, context)
	lhs := s.resolveNames(expr.Lhs, context.withWhereRhs(rhs))
	return &WhereExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

// TODO: make these maps instead of arrays

type nameResolutionContext struct {
	fieldReferenceNameIndexes           []uint64
	whereNameIndexes                    []uint64
	recordsUnderConstructionNameIndexes [][]uint64
	topLevelNameIndexes                 []uint64
}

//---------------------------------------------------------------------------------------------------------------------

func newNameResolutionContext() *nameResolutionContext {
	return &nameResolutionContext{
		fieldReferenceNameIndexes:           make([]uint64, 0),
		whereNameIndexes:                    make([]uint64, 0),
		recordsUnderConstructionNameIndexes: make([][]uint64, 0),
		topLevelNameIndexes:                 make([]uint64, 0),
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (c *nameResolutionContext) withFieldReferenceLhs(fieldReferenceLhs IExpression) *nameResolutionContext {
	return &nameResolutionContext{
		fieldReferenceNameIndexes:           fieldReferenceLhs.GetFieldNameIndexes(),
		whereNameIndexes:                    c.whereNameIndexes,
		recordsUnderConstructionNameIndexes: c.recordsUnderConstructionNameIndexes,
		topLevelNameIndexes:                 c.topLevelNameIndexes,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (c *nameResolutionContext) withWhereRhs(whereRhs IExpression) *nameResolutionContext {
	return &nameResolutionContext{
		fieldReferenceNameIndexes:           c.fieldReferenceNameIndexes,
		whereNameIndexes:                    whereRhs.GetFieldNameIndexes(),
		recordsUnderConstructionNameIndexes: c.recordsUnderConstructionNameIndexes,
		topLevelNameIndexes:                 c.topLevelNameIndexes,
	}
}

//=====================================================================================================================
