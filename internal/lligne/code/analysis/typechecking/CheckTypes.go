//
// # Data types related to Lligne token scanning.
//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package typechecking

import (
	"fmt"
	prior "lligne-cli/internal/lligne/code/analysis/structuring"
	"lligne-cli/internal/lligne/runtime/pools"
	"lligne-cli/internal/lligne/runtime/types"
)

//=====================================================================================================================

type Outcome struct {
	SourceCode      string
	NewLineOffsets  []uint32
	Model           IExpression
	StringConstants *pools.StringConstantPool
	IdentifierNames *pools.StringConstantPool
	TypeConstants   *types.TypeConstantPool
}

//=====================================================================================================================

func CheckTypes(priorOutcome *prior.Outcome) *Outcome {
	typePool := types.NewTypePool()
	model := checkTypes(priorOutcome.SourceCode, priorOutcome.Model, typePool)

	return &Outcome{
		SourceCode:      priorOutcome.SourceCode,
		NewLineOffsets:  priorOutcome.NewLineOffsets,
		Model:           model,
		StringConstants: priorOutcome.StringConstants,
		IdentifierNames: priorOutcome.IdentifierNames,
		TypeConstants:   typePool.Freeze(),
	}
}

//=====================================================================================================================

func checkTypes(sourceCode string, expression prior.IExpression, typePool *types.TypePool) IExpression {

	switch expr := expression.(type) {

	case *prior.AdditionExpr:
		return typeCheckAdditionExpr(sourceCode, expr, typePool)
	case *prior.BooleanLiteralExpr:
		return typeCheckBooleanLiteralExpr(expr)
	case *prior.BuiltInTypeExpr:
		return typeCheckBuiltInTypeExpr(sourceCode, expr, typePool)
	case *prior.DivisionExpr:
		return typeCheckDivisionExpr(sourceCode, expr, typePool)
	case *prior.EqualsExpr:
		return typeCheckEqualsExpr(sourceCode, expr, typePool)
	case *prior.Float64LiteralExpr:
		return typeCheckFloat64LiteralExpr(expr)
	case *prior.GreaterThanExpr:
		return typeCheckGreaterThanExpr(sourceCode, expr, typePool)
	case *prior.GreaterThanOrEqualsExpr:
		return typeCheckGreaterThanOrEqualsExpr(sourceCode, expr, typePool)
	case *prior.Int64LiteralExpr:
		return typeCheckInt64LiteralExpr(expr)
	case *prior.IsExpr:
		return typeCheckIsExpr(sourceCode, expr, typePool)
	case *prior.LessThanExpr:
		return typeCheckLessThanExpr(sourceCode, expr, typePool)
	case *prior.LessThanOrEqualsExpr:
		return typeCheckLessThanOrEqualsExpr(sourceCode, expr, typePool)
	case *prior.LogicalAndExpr:
		return typeCheckLogicalAndExpr(sourceCode, expr, typePool)
	case *prior.LogicalNotOperationExpr:
		return typeCheckLogicalNotOperationExpr(sourceCode, expr, typePool)
	case *prior.LogicalOrExpr:
		return typeCheckLogicalOrExpr(sourceCode, expr, typePool)
	case *prior.MultiplicationExpr:
		return typeCheckMultiplicationExpr(sourceCode, expr, typePool)
	case *prior.NegationOperationExpr:
		return typeCheckNegationOperationExpr(sourceCode, expr, typePool)
	case *prior.NotEqualsExpr:
		return typeCheckNotEqualsExpr(sourceCode, expr, typePool)
	case *prior.ParenthesizedExpr:
		return typeCheckParenthesizedExpr(sourceCode, expr, typePool)
	case *prior.RecordExpr:
		return typeCheckRecordExpr(sourceCode, expr, typePool)
	case *prior.StringLiteralExpr:
		return typeCheckStringLiteralExpr(expr)
	case *prior.SubtractionExpr:
		return typeCheckSubtractionExpr(sourceCode, expr, typePool)

	default:
		panic(fmt.Sprintf("Missing case in checkTypes: %T\n", expression))

	}

}

//=====================================================================================================================

func typeCheckAdditionExpr(sourceCode string, expr *prior.AdditionExpr, typePool *types.TypePool) IExpression {
	lhs := checkTypes(sourceCode, expr.Lhs, typePool)
	rhs := checkTypes(sourceCode, expr.Rhs, typePool)
	switch lhs.GetTypeInfo().(type) {
	case *types.Float64Type, *types.Int64Type:
		// TODO: ensure they're the same
		return &AdditionExpr{
			SourcePosition: expr.SourcePosition,
			Lhs:            lhs,
			Rhs:            rhs,
			TypeInfo:       lhs.GetTypeInfo(),
		}
	case *types.StringType:
		// TODO: ensure both strings
		return &StringConcatenationExpr{
			SourcePosition: expr.SourcePosition,
			Lhs:            lhs,
			Rhs:            rhs,
		}
	default:
		panic(fmt.Sprintf("Missing case in typeCheckAdditionExpr: %T\n", lhs.GetTypeInfo()))
	}
}

//=====================================================================================================================

func typeCheckBooleanLiteralExpr(expr *prior.BooleanLiteralExpr) IExpression {
	return &BooleanLiteralExpr{
		SourcePosition: expr.SourcePosition,
		Value:          expr.Value,
	}
}

//=====================================================================================================================

func typeCheckBuiltInTypeExpr(sourceCode string, expr *prior.BuiltInTypeExpr, typePool *types.TypePool) IExpression {
	name := expr.SourcePosition.GetText(sourceCode)
	return &BuiltInTypeExpr{
		SourcePosition: expr.SourcePosition,
		ValueIndex:     typePool.GetIndexByName(name),
	}
}

//=====================================================================================================================

func typeCheckDivisionExpr(sourceCode string, expr *prior.DivisionExpr, typePool *types.TypePool) IExpression {
	lhs := checkTypes(sourceCode, expr.Lhs, typePool)
	rhs := checkTypes(sourceCode, expr.Rhs, typePool)
	// TODO: ensure they're the same
	return &DivisionExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
		TypeInfo:       lhs.GetTypeInfo(),
	}
}

//=====================================================================================================================

func typeCheckEqualsExpr(sourceCode string, expr *prior.EqualsExpr, typePool *types.TypePool) IExpression {
	lhs := checkTypes(sourceCode, expr.Lhs, typePool)
	rhs := checkTypes(sourceCode, expr.Rhs, typePool)
	// TODO: ensure they're the same
	return &EqualsExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func typeCheckFloat64LiteralExpr(expr *prior.Float64LiteralExpr) IExpression {
	return &Float64LiteralExpr{
		SourcePosition: expr.SourcePosition,
		Value:          expr.Value,
	}
}

//=====================================================================================================================

func typeCheckGreaterThanExpr(sourceCode string, expr *prior.GreaterThanExpr, typePool *types.TypePool) IExpression {
	lhs := checkTypes(sourceCode, expr.Lhs, typePool)
	rhs := checkTypes(sourceCode, expr.Rhs, typePool)
	// TODO: ensure they're the same
	return &GreaterThanExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func typeCheckGreaterThanOrEqualsExpr(sourceCode string, expr *prior.GreaterThanOrEqualsExpr, typePool *types.TypePool) IExpression {
	lhs := checkTypes(sourceCode, expr.Lhs, typePool)
	rhs := checkTypes(sourceCode, expr.Rhs, typePool)
	// TODO: ensure they're the same
	return &GreaterThanOrEqualsExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func typeCheckInt64LiteralExpr(expr *prior.Int64LiteralExpr) IExpression {
	return &Int64LiteralExpr{
		SourcePosition: expr.SourcePosition,
		Value:          expr.Value,
	}
}

//=====================================================================================================================

func typeCheckIsExpr(sourceCode string, expr *prior.IsExpr, typePool *types.TypePool) IExpression {
	lhs := checkTypes(sourceCode, expr.Lhs, typePool)
	rhs := checkTypes(sourceCode, expr.Rhs, typePool)
	// TODO: ensure the lhs and rhs are compatible
	return &IsExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func typeCheckLessThanExpr(sourceCode string, expr *prior.LessThanExpr, typePool *types.TypePool) IExpression {
	lhs := checkTypes(sourceCode, expr.Lhs, typePool)
	rhs := checkTypes(sourceCode, expr.Rhs, typePool)
	// TODO: ensure they're the same
	return &LessThanExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func typeCheckLessThanOrEqualsExpr(sourceCode string, expr *prior.LessThanOrEqualsExpr, typePool *types.TypePool) IExpression {
	lhs := checkTypes(sourceCode, expr.Lhs, typePool)
	rhs := checkTypes(sourceCode, expr.Rhs, typePool)
	// TODO: ensure they're the same
	return &LessThanOrEqualsExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func typeCheckLogicalAndExpr(sourceCode string, expr *prior.LogicalAndExpr, typePool *types.TypePool) IExpression {
	lhs := checkTypes(sourceCode, expr.Lhs, typePool)
	rhs := checkTypes(sourceCode, expr.Rhs, typePool)
	// TODO: ensure they're both boolean
	return &LogicalAndExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func typeCheckLogicalNotOperationExpr(sourceCode string, expr *prior.LogicalNotOperationExpr, typePool *types.TypePool) IExpression {
	operand := checkTypes(sourceCode, expr.Operand, typePool)

	// TODO: validate that operands are boolean

	return &LogicalNotOperationExpr{
		SourcePosition: expr.SourcePosition,
		Operand:        operand,
	}
}

//=====================================================================================================================

func typeCheckLogicalOrExpr(sourceCode string, expr *prior.LogicalOrExpr, typePool *types.TypePool) IExpression {
	lhs := checkTypes(sourceCode, expr.Lhs, typePool)
	rhs := checkTypes(sourceCode, expr.Rhs, typePool)
	// TODO: ensure they're both boolean
	return &LogicalOrExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func typeCheckMultiplicationExpr(sourceCode string, expr *prior.MultiplicationExpr, typePool *types.TypePool) IExpression {
	lhs := checkTypes(sourceCode, expr.Lhs, typePool)
	rhs := checkTypes(sourceCode, expr.Rhs, typePool)
	// TODO: ensure they're the same
	return &MultiplicationExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
		TypeInfo:       lhs.GetTypeInfo(),
	}
}

//=====================================================================================================================

func typeCheckNegationOperationExpr(sourceCode string, expr *prior.NegationOperationExpr, typePool *types.TypePool) IExpression {
	operand := checkTypes(sourceCode, expr.Operand, typePool)

	return &NegationOperationExpr{
		SourcePosition: expr.SourcePosition,
		Operand:        operand,
		TypeInfo:       operand.GetTypeInfo(),
	}
}

//=====================================================================================================================

func typeCheckNotEqualsExpr(sourceCode string, expr *prior.NotEqualsExpr, typePool *types.TypePool) IExpression {
	lhs := checkTypes(sourceCode, expr.Lhs, typePool)
	rhs := checkTypes(sourceCode, expr.Rhs, typePool)
	// TODO: ensure they're the same
	return &NotEqualsExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func typeCheckParenthesizedExpr(sourceCode string, expr *prior.ParenthesizedExpr, typePool *types.TypePool) IExpression {

	inner := checkTypes(sourceCode, expr.InnerExpr, typePool)

	return &ParenthesizedExpr{
		SourcePosition: expr.SourcePosition,
		InnerExpr:      inner,
		TypeInfo:       inner.GetTypeInfo(),
	}

}

//=====================================================================================================================

func typeCheckRecordExpr(sourceCode string, expr *prior.RecordExpr, typePool *types.TypePool) IExpression {

	items := make([]*RecordFieldExpr, 0)
	for _, item := range expr.Items {
		items = append(items, typeCheckRecordFieldExpr(sourceCode, item, typePool))
	}

	return &RecordExpr{
		SourcePosition: expr.SourcePosition,
		TypeInfo:       types.Int64TypeInstance, // TODO (obviously)
	}

}

//=====================================================================================================================

func typeCheckRecordFieldExpr(sourceCode string, expr *prior.RecordFieldExpr, typePool *types.TypePool) *RecordFieldExpr {
	return &RecordFieldExpr{
		SourcePosition: expr.SourcePosition,
		FieldNameIndex: expr.FieldNameIndex,
		FieldValue:     checkTypes(sourceCode, expr.FieldValue, typePool),
	}
}

//=====================================================================================================================

func typeCheckStringLiteralExpr(expr *prior.StringLiteralExpr) IExpression {
	return &StringLiteralExpr{
		SourcePosition: expr.SourcePosition,
		ValueIndex:     expr.ValueIndex,
	}
}

//=====================================================================================================================

func typeCheckSubtractionExpr(sourceCode string, expr *prior.SubtractionExpr, typePool *types.TypePool) IExpression {
	lhs := checkTypes(sourceCode, expr.Lhs, typePool)
	rhs := checkTypes(sourceCode, expr.Rhs, typePool)
	// TODO: ensure they're the same
	return &SubtractionExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
		TypeInfo:       lhs.GetTypeInfo(),
	}
}

//=====================================================================================================================
