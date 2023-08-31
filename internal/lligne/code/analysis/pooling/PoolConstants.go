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

	stringConstants := pools.NewStringPool()
	identifierNames := pools.NewStringPool()

	model := poolConstants(priorOutcome.SourceCode, priorOutcome.Model, stringConstants, identifierNames)

	return &Outcome{
		SourceCode:      priorOutcome.SourceCode,
		NewLineOffsets:  priorOutcome.NewLineOffsets,
		Model:           model,
		StringConstants: stringConstants.Freeze(),
		IdentifierNames: identifierNames.Freeze(),
	}
}

//=====================================================================================================================

func poolConstants(
	sourceCode string,
	expression prior.IExpression,
	stringConstants *pools.StringPool,
	identifierNames *pools.StringPool,
) IExpression {

	switch expr := expression.(type) {

	case *prior.AdditionExpr:
		return poolAdditionExpr(sourceCode, expr, stringConstants, identifierNames)
	case *prior.BooleanLiteralExpr:
		return poolBooleanLiteralExpr(expr)
	case *prior.BuiltInTypeExpr:
		return poolBuiltInTypeExpr(expr)
	case *prior.DivisionExpr:
		return poolDivisionExpr(sourceCode, expr, stringConstants, identifierNames)
	case *prior.EqualsExpr:
		return poolEqualsExpr(sourceCode, expr, stringConstants, identifierNames)
	case *prior.FloatingPointLiteralExpr:
		return poolFloatingPointLiteralExpr(expr)
	case *prior.GreaterThanExpr:
		return poolGreaterThanExpr(sourceCode, expr, stringConstants, identifierNames)
	case *prior.GreaterThanOrEqualsExpr:
		return poolGreaterThanOrEqualsExpr(sourceCode, expr, stringConstants, identifierNames)
	case *prior.IdentifierExpr:
		return poolIdentifierExpr(sourceCode, expr, identifierNames)
	case *prior.IntegerLiteralExpr:
		return poolIntegerLiteralExpr(expr)
	case *prior.IsExpr:
		return poolIsExpr(sourceCode, expr, stringConstants, identifierNames)
	case *prior.LessThanExpr:
		return poolLessThanExpr(sourceCode, expr, stringConstants, identifierNames)
	case *prior.LessThanOrEqualsExpr:
		return poolLessThanOrEqualsExpr(sourceCode, expr, stringConstants, identifierNames)
	case *prior.LogicalAndExpr:
		return poolLogicalAndExpr(sourceCode, expr, stringConstants, identifierNames)
	case *prior.LogicalNotOperationExpr:
		return poolLogicalNotOperationExpr(sourceCode, expr, stringConstants, identifierNames)
	case *prior.LogicalOrExpr:
		return poolLogicalOrExpr(sourceCode, expr, stringConstants, identifierNames)
	case *prior.MultiplicationExpr:
		return poolMultiplicationExpr(sourceCode, expr, stringConstants, identifierNames)
	case *prior.NegationOperationExpr:
		return poolNegationOperationExpr(sourceCode, expr, stringConstants, identifierNames)
	case *prior.NotEqualsExpr:
		return poolNotEqualsExpr(sourceCode, expr, stringConstants, identifierNames)
	case *prior.ParenthesizedExpr:
		return poolParenthesizedExpr(sourceCode, expr, stringConstants, identifierNames)
	case *prior.RecordExpr:
		return poolRecordExpr(sourceCode, expr, stringConstants, identifierNames)
	case *prior.StringLiteralExpr:
		return poolStringLiteralExpr(sourceCode, expr, stringConstants)
	case *prior.SubtractionExpr:
		return poolSubtractionExpr(sourceCode, expr, stringConstants, identifierNames)

	default:
		panic(fmt.Sprintf("Missing case in poolConstants: %T\n", expression))

	}

}

//=====================================================================================================================

func poolAdditionExpr(
	sourceCode string,
	expr *prior.AdditionExpr,
	stringConstants *pools.StringPool,
	identifierNames *pools.StringPool,
) IExpression {
	lhs := poolConstants(sourceCode, expr.Lhs, stringConstants, identifierNames)
	rhs := poolConstants(sourceCode, expr.Rhs, stringConstants, identifierNames)
	return &AdditionExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func poolBooleanLiteralExpr(expr *prior.BooleanLiteralExpr) IExpression {
	return &BooleanLiteralExpr{
		SourcePosition: expr.SourcePosition,
		Value:          expr.Value,
	}
}

//=====================================================================================================================

func poolBuiltInTypeExpr(expr *prior.BuiltInTypeExpr) IExpression {
	return &BuiltInTypeExpr{
		SourcePosition: expr.SourcePosition,
	}
}

//=====================================================================================================================

func poolDivisionExpr(
	sourceCode string,
	expr *prior.DivisionExpr,
	stringConstants *pools.StringPool,
	identifierNames *pools.StringPool,
) IExpression {
	lhs := poolConstants(sourceCode, expr.Lhs, stringConstants, identifierNames)
	rhs := poolConstants(sourceCode, expr.Rhs, stringConstants, identifierNames)
	return &DivisionExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func poolEqualsExpr(
	sourceCode string,
	expr *prior.EqualsExpr,
	stringConstants *pools.StringPool,
	identifierNames *pools.StringPool,
) IExpression {
	lhs := poolConstants(sourceCode, expr.Lhs, stringConstants, identifierNames)
	rhs := poolConstants(sourceCode, expr.Rhs, stringConstants, identifierNames)
	return &EqualsExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func poolFloatingPointLiteralExpr(expr *prior.FloatingPointLiteralExpr) IExpression {
	return &Float64LiteralExpr{
		SourcePosition: expr.SourcePosition,
		Value:          expr.Value,
	}
}

//=====================================================================================================================

func poolGreaterThanExpr(
	sourceCode string,
	expr *prior.GreaterThanExpr,
	stringConstants *pools.StringPool,
	identifierNames *pools.StringPool,
) IExpression {
	lhs := poolConstants(sourceCode, expr.Lhs, stringConstants, identifierNames)
	rhs := poolConstants(sourceCode, expr.Rhs, stringConstants, identifierNames)
	return &GreaterThanExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func poolGreaterThanOrEqualsExpr(
	sourceCode string,
	expr *prior.GreaterThanOrEqualsExpr,
	stringConstants *pools.StringPool,
	identifierNames *pools.StringPool,
) IExpression {
	lhs := poolConstants(sourceCode, expr.Lhs, stringConstants, identifierNames)
	rhs := poolConstants(sourceCode, expr.Rhs, stringConstants, identifierNames)
	return &GreaterThanOrEqualsExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func poolIdentifierExpr(
	sourceCode string,
	expr *prior.IdentifierExpr,
	identifierNames *pools.StringPool,
) IExpression {
	name := expr.SourcePosition.GetText(sourceCode)
	nameIndex := identifierNames.Put(name)

	return &IdentifierExpr{
		SourcePosition: expr.SourcePosition,
		NameIndex:      nameIndex,
	}
}

//=====================================================================================================================

func poolIntegerLiteralExpr(expr *prior.IntegerLiteralExpr) IExpression {
	return &Int64LiteralExpr{
		SourcePosition: expr.SourcePosition,
		Value:          expr.Value,
	}
}

//=====================================================================================================================

func poolIsExpr(
	sourceCode string,
	expr *prior.IsExpr,
	stringConstants *pools.StringPool,
	identifierNames *pools.StringPool,
) IExpression {
	lhs := poolConstants(sourceCode, expr.Lhs, stringConstants, identifierNames)
	rhs := poolConstants(sourceCode, expr.Rhs, stringConstants, identifierNames)
	return &IsExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func poolLessThanExpr(
	sourceCode string,
	expr *prior.LessThanExpr,
	stringConstants *pools.StringPool,
	identifierNames *pools.StringPool,
) IExpression {
	lhs := poolConstants(sourceCode, expr.Lhs, stringConstants, identifierNames)
	rhs := poolConstants(sourceCode, expr.Rhs, stringConstants, identifierNames)
	return &LessThanExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func poolLessThanOrEqualsExpr(
	sourceCode string,
	expr *prior.LessThanOrEqualsExpr,
	stringConstants *pools.StringPool,
	identifierNames *pools.StringPool,
) IExpression {
	lhs := poolConstants(sourceCode, expr.Lhs, stringConstants, identifierNames)
	rhs := poolConstants(sourceCode, expr.Rhs, stringConstants, identifierNames)
	return &LessThanOrEqualsExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func poolLogicalAndExpr(
	sourceCode string,
	expr *prior.LogicalAndExpr,
	stringConstants *pools.StringPool,
	identifierNames *pools.StringPool,
) IExpression {
	lhs := poolConstants(sourceCode, expr.Lhs, stringConstants, identifierNames)
	rhs := poolConstants(sourceCode, expr.Rhs, stringConstants, identifierNames)
	return &LogicalAndExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func poolLogicalNotOperationExpr(
	sourceCode string,
	expr *prior.LogicalNotOperationExpr,
	stringConstants *pools.StringPool,
	identifierNames *pools.StringPool,
) IExpression {
	operand := poolConstants(sourceCode, expr.Operand, stringConstants, identifierNames)
	return &LogicalNotOperationExpr{
		SourcePosition: expr.SourcePosition,
		Operand:        operand,
	}
}

//=====================================================================================================================

func poolLogicalOrExpr(
	sourceCode string,
	expr *prior.LogicalOrExpr,
	stringConstants *pools.StringPool,
	identifierNames *pools.StringPool,
) IExpression {
	lhs := poolConstants(sourceCode, expr.Lhs, stringConstants, identifierNames)
	rhs := poolConstants(sourceCode, expr.Rhs, stringConstants, identifierNames)
	return &LogicalOrExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func poolMultiplicationExpr(
	sourceCode string,
	expr *prior.MultiplicationExpr,
	stringConstants *pools.StringPool,
	identifierNames *pools.StringPool,
) IExpression {
	lhs := poolConstants(sourceCode, expr.Lhs, stringConstants, identifierNames)
	rhs := poolConstants(sourceCode, expr.Rhs, stringConstants, identifierNames)
	return &MultiplicationExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func poolNegationOperationExpr(
	sourceCode string,
	expr *prior.NegationOperationExpr,
	stringConstants *pools.StringPool,
	identifierNames *pools.StringPool,
) IExpression {
	operand := poolConstants(sourceCode, expr.Operand, stringConstants, identifierNames)
	return &NegationOperationExpr{
		SourcePosition: expr.SourcePosition,
		Operand:        operand,
	}
}

//=====================================================================================================================

func poolNotEqualsExpr(
	sourceCode string,
	expr *prior.NotEqualsExpr,
	stringConstants *pools.StringPool,
	identifierNames *pools.StringPool,
) IExpression {
	lhs := poolConstants(sourceCode, expr.Lhs, stringConstants, identifierNames)
	rhs := poolConstants(sourceCode, expr.Rhs, stringConstants, identifierNames)
	return &NotEqualsExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func poolParenthesizedExpr(
	sourceCode string,
	expr *prior.ParenthesizedExpr,
	stringConstants *pools.StringPool,
	identifierNames *pools.StringPool,
) IExpression {
	inner := poolConstants(sourceCode, expr.InnerExpr, stringConstants, identifierNames)
	return &ParenthesizedExpr{
		SourcePosition: expr.SourcePosition,
		InnerExpr:      inner,
	}
}

//=====================================================================================================================

func poolRecordExpr(
	sourceCode string,
	expr *prior.RecordExpr,
	stringConstants *pools.StringPool,
	identifierNames *pools.StringPool,
) IExpression {
	return &RecordExpr{
		SourcePosition: expr.SourcePosition,
	}
}

//=====================================================================================================================

func poolStringLiteralExpr(
	sourceCode string,
	expr *prior.StringLiteralExpr,
	stringConstants *pools.StringPool,
) IExpression {
	text := expr.SourcePosition.GetText(sourceCode)
	var value string

	switch expr.Delimiters {
	case prior.StringDelimitersDoubleQuotes:
		value = text[1 : len(text)-1]
	case prior.StringDelimitersSingleQuotes:
		value = text[1 : len(text)-1]
	default:
		panic("TODO: Unhandled string delimiters")
	}

	valueIndex := stringConstants.Put(value)

	// TODO: escape chars
	return &StringLiteralExpr{
		SourcePosition: expr.SourcePosition,
		ValueIndex:     valueIndex,
	}
}

//=====================================================================================================================

func poolSubtractionExpr(
	sourceCode string,
	expr *prior.SubtractionExpr,
	stringConstants *pools.StringPool,
	identifierNames *pools.StringPool,
) IExpression {
	lhs := poolConstants(sourceCode, expr.Lhs, stringConstants, identifierNames)
	rhs := poolConstants(sourceCode, expr.Rhs, stringConstants, identifierNames)
	return &SubtractionExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================
