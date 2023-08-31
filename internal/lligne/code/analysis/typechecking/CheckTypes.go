//
// # Data types related to Lligne token scanning.
//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package typechecking

import (
	"fmt"
	prior "lligne-cli/internal/lligne/code/analysis/pooling"
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
	// TODO: TypesPool
}

//=====================================================================================================================

func CheckTypes(priorOutcome *prior.Outcome) *Outcome {
	model := checkTypes(priorOutcome.SourceCode, priorOutcome.Model)

	return &Outcome{
		SourceCode:      priorOutcome.SourceCode,
		NewLineOffsets:  priorOutcome.NewLineOffsets,
		Model:           model,
		StringConstants: priorOutcome.StringConstants,
		IdentifierNames: priorOutcome.IdentifierNames,
	}
}

//=====================================================================================================================

func checkTypes(sourceCode string, expression prior.IExpression) IExpression {

	switch expr := expression.(type) {

	case *prior.AdditionExpr:
		return typeCheckAdditionExpr(sourceCode, expr)
	case *prior.BooleanLiteralExpr:
		return typeCheckBooleanLiteralExpr(expr)
	case *prior.BuiltInTypeExpr:
		return typeCheckBuiltInTypeExpr(sourceCode, expr)
	case *prior.DivisionExpr:
		return typeCheckDivisionExpr(sourceCode, expr)
	case *prior.EqualsExpr:
		return typeCheckEqualsExpr(sourceCode, expr)
	case *prior.Float64LiteralExpr:
		return typeCheckFloat64LiteralExpr(expr)
	case *prior.GreaterThanExpr:
		return typeCheckGreaterThanExpr(sourceCode, expr)
	case *prior.GreaterThanOrEqualsExpr:
		return typeCheckGreaterThanOrEqualsExpr(sourceCode, expr)
	case *prior.Int64LiteralExpr:
		return typeCheckInt64LiteralExpr(expr)
	case *prior.IsExpr:
		return typeCheckIsExpr(sourceCode, expr)
	case *prior.LessThanExpr:
		return typeCheckLessThanExpr(sourceCode, expr)
	case *prior.LessThanOrEqualsExpr:
		return typeCheckLessThanOrEqualsExpr(sourceCode, expr)
	case *prior.LogicalAndExpr:
		return typeCheckLogicalAndExpr(sourceCode, expr)
	case *prior.LogicalNotOperationExpr:
		return typeCheckLogicalNotOperationExpr(sourceCode, expr)
	case *prior.LogicalOrExpr:
		return typeCheckLogicalOrExpr(sourceCode, expr)
	case *prior.MultiplicationExpr:
		return typeCheckMultiplicationExpr(sourceCode, expr)
	case *prior.NegationOperationExpr:
		return typeCheckNegationOperationExpr(sourceCode, expr)
	case *prior.NotEqualsExpr:
		return typeCheckNotEqualsExpr(sourceCode, expr)
	case *prior.ParenthesizedExpr:
		return typeCheckParenthesizedExpr(sourceCode, expr)
	case *prior.RecordExpr:
		return typeCheckRecordExpr(sourceCode, expr)
	case *prior.StringLiteralExpr:
		return typeCheckStringLiteralExpr(expr)
	case *prior.SubtractionExpr:
		return typeCheckSubtractionExpr(sourceCode, expr)

	default:
		panic(fmt.Sprintf("Missing case in checkTypes: %T\n", expression))

	}

}

//=====================================================================================================================

func typeCheckAdditionExpr(sourceCode string, expr *prior.AdditionExpr) IExpression {
	lhs := checkTypes(sourceCode, expr.Lhs)
	rhs := checkTypes(sourceCode, expr.Rhs)
	switch lhs.GetTypeInfo().(type) {
	case *types.Float64Type, *types.Int64Type:
		// TODO: ensure they're the same
		// TODO: coerce integers
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

func typeCheckBuiltInTypeExpr(sourceCode string, expr *prior.BuiltInTypeExpr) IExpression {
	name := expr.SourcePosition.GetText(sourceCode)
	return &BuiltInTypeExpr{
		SourcePosition: expr.SourcePosition,
		Value:          types.BuiltInTypesByName[name],
	}
}

//=====================================================================================================================

func typeCheckDivisionExpr(sourceCode string, expr *prior.DivisionExpr) IExpression {
	lhs := checkTypes(sourceCode, expr.Lhs)
	rhs := checkTypes(sourceCode, expr.Rhs)
	// TODO: ensure they're the same
	// TODO: coerce integers
	return &DivisionExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
		TypeInfo:       lhs.GetTypeInfo(),
	}
}

//=====================================================================================================================

func typeCheckEqualsExpr(sourceCode string, expr *prior.EqualsExpr) IExpression {
	lhs := checkTypes(sourceCode, expr.Lhs)
	rhs := checkTypes(sourceCode, expr.Rhs)
	// TODO: ensure they're the same
	// TODO: coerce integers
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

func typeCheckGreaterThanExpr(sourceCode string, expr *prior.GreaterThanExpr) IExpression {
	lhs := checkTypes(sourceCode, expr.Lhs)
	rhs := checkTypes(sourceCode, expr.Rhs)
	// TODO: ensure they're the same
	// TODO: coerce integers
	return &GreaterThanExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func typeCheckGreaterThanOrEqualsExpr(sourceCode string, expr *prior.GreaterThanOrEqualsExpr) IExpression {
	lhs := checkTypes(sourceCode, expr.Lhs)
	rhs := checkTypes(sourceCode, expr.Rhs)
	// TODO: ensure they're the same
	// TODO: coerce integers
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

func typeCheckIsExpr(sourceCode string, expr *prior.IsExpr) IExpression {
	lhs := checkTypes(sourceCode, expr.Lhs)
	rhs := checkTypes(sourceCode, expr.Rhs)
	// TODO: ensure the lhs and rhs are compatible
	return &IsExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func typeCheckLessThanExpr(sourceCode string, expr *prior.LessThanExpr) IExpression {
	lhs := checkTypes(sourceCode, expr.Lhs)
	rhs := checkTypes(sourceCode, expr.Rhs)
	// TODO: ensure they're the same
	// TODO: coerce integers
	return &LessThanExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func typeCheckLessThanOrEqualsExpr(sourceCode string, expr *prior.LessThanOrEqualsExpr) IExpression {
	lhs := checkTypes(sourceCode, expr.Lhs)
	rhs := checkTypes(sourceCode, expr.Rhs)
	// TODO: ensure they're the same
	// TODO: coerce integers
	return &LessThanOrEqualsExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func typeCheckLogicalAndExpr(sourceCode string, expr *prior.LogicalAndExpr) IExpression {
	lhs := checkTypes(sourceCode, expr.Lhs)
	rhs := checkTypes(sourceCode, expr.Rhs)
	// TODO: ensure they're both boolean
	// TODO: coerce integers
	return &LogicalAndExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func typeCheckLogicalNotOperationExpr(sourceCode string, expr *prior.LogicalNotOperationExpr) IExpression {
	operand := checkTypes(sourceCode, expr.Operand)

	// TODO: validate that operands are boolean

	return &LogicalNotOperationExpr{
		SourcePosition: expr.SourcePosition,
		Operand:        operand,
	}
}

//=====================================================================================================================

func typeCheckLogicalOrExpr(sourceCode string, expr *prior.LogicalOrExpr) IExpression {
	lhs := checkTypes(sourceCode, expr.Lhs)
	rhs := checkTypes(sourceCode, expr.Rhs)
	// TODO: ensure they're both boolean
	// TODO: coerce integers
	return &LogicalOrExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func typeCheckMultiplicationExpr(sourceCode string, expr *prior.MultiplicationExpr) IExpression {
	lhs := checkTypes(sourceCode, expr.Lhs)
	rhs := checkTypes(sourceCode, expr.Rhs)
	// TODO: ensure they're the same
	// TODO: coerce integers
	return &MultiplicationExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
		TypeInfo:       lhs.GetTypeInfo(),
	}
}

//=====================================================================================================================

func typeCheckNegationOperationExpr(sourceCode string, expr *prior.NegationOperationExpr) IExpression {
	operand := checkTypes(sourceCode, expr.Operand)

	return &NegationOperationExpr{
		SourcePosition: expr.SourcePosition,
		Operand:        operand,
		TypeInfo:       operand.GetTypeInfo(),
	}
}

//=====================================================================================================================

func typeCheckNotEqualsExpr(sourceCode string, expr *prior.NotEqualsExpr) IExpression {
	lhs := checkTypes(sourceCode, expr.Lhs)
	rhs := checkTypes(sourceCode, expr.Rhs)
	// TODO: ensure they're the same
	// TODO: coerce integers
	return &NotEqualsExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func typeCheckParenthesizedExpr(sourceCode string, expr *prior.ParenthesizedExpr) IExpression {

	inner := checkTypes(sourceCode, expr.InnerExpr)

	return &ParenthesizedExpr{
		SourcePosition: expr.SourcePosition,
		InnerExpr:      inner,
		TypeInfo:       inner.GetTypeInfo(),
	}

}

//=====================================================================================================================

func typeCheckRecordExpr(sourceCode string, expr *prior.RecordExpr) IExpression {

	return &RecordExpr{
		SourcePosition: expr.SourcePosition,
		TypeInfo:       types.Int64TypeInstance, // TODO (obviously)
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

func typeCheckSubtractionExpr(sourceCode string, expr *prior.SubtractionExpr) IExpression {
	lhs := checkTypes(sourceCode, expr.Lhs)
	rhs := checkTypes(sourceCode, expr.Rhs)
	// TODO: ensure they're the same
	// TODO: coerce integers
	return &SubtractionExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
		TypeInfo:       lhs.GetTypeInfo(),
	}
}

//=====================================================================================================================
