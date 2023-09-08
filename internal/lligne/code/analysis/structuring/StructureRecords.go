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

	stringConstants := pools.NewStringPool()
	identifierNames := pools.NewStringPool()

	model := structureRecords(priorOutcome.SourceCode, priorOutcome.Model, stringConstants, identifierNames)

	return &Outcome{
		SourceCode:      priorOutcome.SourceCode,
		NewLineOffsets:  priorOutcome.NewLineOffsets,
		Model:           model,
		StringConstants: priorOutcome.StringConstants,
		IdentifierNames: priorOutcome.IdentifierNames,
	}
}

//=====================================================================================================================

func structureRecords(
	sourceCode string,
	expression prior.IExpression,
	stringConstants *pools.StringPool,
	identifierNames *pools.StringPool,
) IExpression {

	switch expr := expression.(type) {

	case *prior.AdditionExpr:
		return structureAdditionExpr(sourceCode, expr, stringConstants, identifierNames)
	case *prior.BooleanLiteralExpr:
		return structureBooleanLiteralExpr(expr)
	case *prior.BuiltInTypeExpr:
		return structureBuiltInTypeExpr(expr)
	case *prior.DivisionExpr:
		return structureDivisionExpr(sourceCode, expr, stringConstants, identifierNames)
	case *prior.EqualsExpr:
		return structureEqualsExpr(sourceCode, expr, stringConstants, identifierNames)
	case *prior.Float64LiteralExpr:
		return structureFloatingPointLiteralExpr(expr)
	case *prior.GreaterThanExpr:
		return structureGreaterThanExpr(sourceCode, expr, stringConstants, identifierNames)
	case *prior.GreaterThanOrEqualsExpr:
		return structureGreaterThanOrEqualsExpr(sourceCode, expr, stringConstants, identifierNames)
	case *prior.IdentifierExpr:
		return structureIdentifierExpr(expr)
	case *prior.Int64LiteralExpr:
		return structureIntegerLiteralExpr(expr)
	case *prior.IsExpr:
		return structureIsExpr(sourceCode, expr, stringConstants, identifierNames)
	case *prior.LessThanExpr:
		return structureLessThanExpr(sourceCode, expr, stringConstants, identifierNames)
	case *prior.LessThanOrEqualsExpr:
		return structureLessThanOrEqualsExpr(sourceCode, expr, stringConstants, identifierNames)
	case *prior.LogicalAndExpr:
		return structureLogicalAndExpr(sourceCode, expr, stringConstants, identifierNames)
	case *prior.LogicalNotOperationExpr:
		return structureLogicalNotOperationExpr(sourceCode, expr, stringConstants, identifierNames)
	case *prior.LogicalOrExpr:
		return structureLogicalOrExpr(sourceCode, expr, stringConstants, identifierNames)
	case *prior.MultiplicationExpr:
		return structureMultiplicationExpr(sourceCode, expr, stringConstants, identifierNames)
	case *prior.NegationOperationExpr:
		return structureNegationOperationExpr(sourceCode, expr, stringConstants, identifierNames)
	case *prior.NotEqualsExpr:
		return structureNotEqualsExpr(sourceCode, expr, stringConstants, identifierNames)
	case *prior.ParenthesizedExpr:
		return structureParenthesizedExpr(sourceCode, expr, stringConstants, identifierNames)
	case *prior.RecordExpr:
		return structureRecordExpr(sourceCode, expr, stringConstants, identifierNames)
	case *prior.StringLiteralExpr:
		return structureStringLiteralExpr(expr)
	case *prior.SubtractionExpr:
		return structureSubtractionExpr(sourceCode, expr, stringConstants, identifierNames)

	default:
		panic(fmt.Sprintf("Missing case in structureRecords: %T\n", expression))

	}

}

//=====================================================================================================================

func structureAdditionExpr(
	sourceCode string,
	expr *prior.AdditionExpr,
	stringConstants *pools.StringPool,
	identifierNames *pools.StringPool,
) IExpression {
	lhs := structureRecords(sourceCode, expr.Lhs, stringConstants, identifierNames)
	rhs := structureRecords(sourceCode, expr.Rhs, stringConstants, identifierNames)
	return &AdditionExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func structureBooleanLiteralExpr(expr *prior.BooleanLiteralExpr) IExpression {
	return &BooleanLiteralExpr{
		SourcePosition: expr.SourcePosition,
		Value:          expr.Value,
	}
}

//=====================================================================================================================

func structureBuiltInTypeExpr(expr *prior.BuiltInTypeExpr) IExpression {
	return &BuiltInTypeExpr{
		SourcePosition: expr.SourcePosition,
	}
}

//=====================================================================================================================

func structureDivisionExpr(
	sourceCode string,
	expr *prior.DivisionExpr,
	stringConstants *pools.StringPool,
	identifierNames *pools.StringPool,
) IExpression {
	lhs := structureRecords(sourceCode, expr.Lhs, stringConstants, identifierNames)
	rhs := structureRecords(sourceCode, expr.Rhs, stringConstants, identifierNames)
	return &DivisionExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func structureEqualsExpr(
	sourceCode string,
	expr *prior.EqualsExpr,
	stringConstants *pools.StringPool,
	identifierNames *pools.StringPool,
) IExpression {
	lhs := structureRecords(sourceCode, expr.Lhs, stringConstants, identifierNames)
	rhs := structureRecords(sourceCode, expr.Rhs, stringConstants, identifierNames)
	return &EqualsExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func structureFloatingPointLiteralExpr(expr *prior.Float64LiteralExpr) IExpression {
	return &Float64LiteralExpr{
		SourcePosition: expr.SourcePosition,
		Value:          expr.Value,
	}
}

//=====================================================================================================================

func structureGreaterThanExpr(
	sourceCode string,
	expr *prior.GreaterThanExpr,
	stringConstants *pools.StringPool,
	identifierNames *pools.StringPool,
) IExpression {
	lhs := structureRecords(sourceCode, expr.Lhs, stringConstants, identifierNames)
	rhs := structureRecords(sourceCode, expr.Rhs, stringConstants, identifierNames)
	return &GreaterThanExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func structureGreaterThanOrEqualsExpr(
	sourceCode string,
	expr *prior.GreaterThanOrEqualsExpr,
	stringConstants *pools.StringPool,
	identifierNames *pools.StringPool,
) IExpression {
	lhs := structureRecords(sourceCode, expr.Lhs, stringConstants, identifierNames)
	rhs := structureRecords(sourceCode, expr.Rhs, stringConstants, identifierNames)
	return &GreaterThanOrEqualsExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func structureIdentifierExpr(
	expr *prior.IdentifierExpr,
) IExpression {
	return &IdentifierExpr{
		SourcePosition: expr.SourcePosition,
		NameIndex:      expr.NameIndex,
	}
}

//=====================================================================================================================

func structureIntegerLiteralExpr(expr *prior.Int64LiteralExpr) IExpression {
	return &Int64LiteralExpr{
		SourcePosition: expr.SourcePosition,
		Value:          expr.Value,
	}
}

//=====================================================================================================================

func structureIsExpr(
	sourceCode string,
	expr *prior.IsExpr,
	stringConstants *pools.StringPool,
	identifierNames *pools.StringPool,
) IExpression {
	lhs := structureRecords(sourceCode, expr.Lhs, stringConstants, identifierNames)
	rhs := structureRecords(sourceCode, expr.Rhs, stringConstants, identifierNames)
	return &IsExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func structureLessThanExpr(
	sourceCode string,
	expr *prior.LessThanExpr,
	stringConstants *pools.StringPool,
	identifierNames *pools.StringPool,
) IExpression {
	lhs := structureRecords(sourceCode, expr.Lhs, stringConstants, identifierNames)
	rhs := structureRecords(sourceCode, expr.Rhs, stringConstants, identifierNames)
	return &LessThanExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func structureLessThanOrEqualsExpr(
	sourceCode string,
	expr *prior.LessThanOrEqualsExpr,
	stringConstants *pools.StringPool,
	identifierNames *pools.StringPool,
) IExpression {
	lhs := structureRecords(sourceCode, expr.Lhs, stringConstants, identifierNames)
	rhs := structureRecords(sourceCode, expr.Rhs, stringConstants, identifierNames)
	return &LessThanOrEqualsExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func structureLogicalAndExpr(
	sourceCode string,
	expr *prior.LogicalAndExpr,
	stringConstants *pools.StringPool,
	identifierNames *pools.StringPool,
) IExpression {
	lhs := structureRecords(sourceCode, expr.Lhs, stringConstants, identifierNames)
	rhs := structureRecords(sourceCode, expr.Rhs, stringConstants, identifierNames)
	return &LogicalAndExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func structureLogicalNotOperationExpr(
	sourceCode string,
	expr *prior.LogicalNotOperationExpr,
	stringConstants *pools.StringPool,
	identifierNames *pools.StringPool,
) IExpression {
	operand := structureRecords(sourceCode, expr.Operand, stringConstants, identifierNames)
	return &LogicalNotOperationExpr{
		SourcePosition: expr.SourcePosition,
		Operand:        operand,
	}
}

//=====================================================================================================================

func structureLogicalOrExpr(
	sourceCode string,
	expr *prior.LogicalOrExpr,
	stringConstants *pools.StringPool,
	identifierNames *pools.StringPool,
) IExpression {
	lhs := structureRecords(sourceCode, expr.Lhs, stringConstants, identifierNames)
	rhs := structureRecords(sourceCode, expr.Rhs, stringConstants, identifierNames)
	return &LogicalOrExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func structureMultiplicationExpr(
	sourceCode string,
	expr *prior.MultiplicationExpr,
	stringConstants *pools.StringPool,
	identifierNames *pools.StringPool,
) IExpression {
	lhs := structureRecords(sourceCode, expr.Lhs, stringConstants, identifierNames)
	rhs := structureRecords(sourceCode, expr.Rhs, stringConstants, identifierNames)
	return &MultiplicationExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func structureNegationOperationExpr(
	sourceCode string,
	expr *prior.NegationOperationExpr,
	stringConstants *pools.StringPool,
	identifierNames *pools.StringPool,
) IExpression {
	operand := structureRecords(sourceCode, expr.Operand, stringConstants, identifierNames)
	return &NegationOperationExpr{
		SourcePosition: expr.SourcePosition,
		Operand:        operand,
	}
}

//=====================================================================================================================

func structureNotEqualsExpr(
	sourceCode string,
	expr *prior.NotEqualsExpr,
	stringConstants *pools.StringPool,
	identifierNames *pools.StringPool,
) IExpression {
	lhs := structureRecords(sourceCode, expr.Lhs, stringConstants, identifierNames)
	rhs := structureRecords(sourceCode, expr.Rhs, stringConstants, identifierNames)
	return &NotEqualsExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func structureParenthesizedExpr(
	sourceCode string,
	expr *prior.ParenthesizedExpr,
	stringConstants *pools.StringPool,
	identifierNames *pools.StringPool,
) IExpression {
	inner := structureRecords(sourceCode, expr.InnerExpr, stringConstants, identifierNames)
	return &ParenthesizedExpr{
		SourcePosition: expr.SourcePosition,
		InnerExpr:      inner,
	}
}

//=====================================================================================================================

func structureRecordExpr(
	sourceCode string,
	expr *prior.RecordExpr,
	stringConstants *pools.StringPool,
	identifierNames *pools.StringPool,
) IExpression {
	items := make([]*RecordFieldExpr, 0)
	for _, item := range expr.Items {
		fieldExpr := structureRecordFieldExpr(sourceCode, item, stringConstants, identifierNames)
		items = append(items, fieldExpr)
	}

	return &RecordExpr{
		SourcePosition: expr.SourcePosition,
		Fields:         items,
	}
}

//=====================================================================================================================

func structureRecordFieldExpr(
	sourceCode string,
	expr prior.IExpression,
	stringConstants *pools.StringPool,
	identifierNames *pools.StringPool,
) *RecordFieldExpr {

	// TODO: Qualify and other intersection expressions

	switch fieldExpr := expr.(type) {
	case *prior.IntersectAssignValueExpr:
		fieldNameIndex := fieldExpr.Lhs.(*prior.IdentifierExpr).NameIndex
		value := structureRecords(sourceCode, fieldExpr.Rhs, stringConstants, identifierNames)
		return &RecordFieldExpr{
			SourcePosition: expr.GetSourcePosition(),
			FieldNameIndex: fieldNameIndex,
			FieldValue:     value,
		}
	default:
		panic(fmt.Sprintf("Missing case in structureRecordFieldExpr: %T\n", expr))
	}

}

//=====================================================================================================================

func structureStringLiteralExpr(
	expr *prior.StringLiteralExpr,
) IExpression {
	return &StringLiteralExpr{
		SourcePosition: expr.SourcePosition,
		ValueIndex:     expr.ValueIndex,
	}
}

//=====================================================================================================================

func structureSubtractionExpr(
	sourceCode string,
	expr *prior.SubtractionExpr,
	stringConstants *pools.StringPool,
	identifierNames *pools.StringPool,
) IExpression {
	lhs := structureRecords(sourceCode, expr.Lhs, stringConstants, identifierNames)
	rhs := structureRecords(sourceCode, expr.Rhs, stringConstants, identifierNames)
	return &SubtractionExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================
