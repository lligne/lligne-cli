//
// # Data types related to Lligne token scanning.
//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package typechecking

import (
	"fmt"
	"lligne-cli/internal/lligne/code/parsing"
	"lligne-cli/internal/lligne/runtime/types"
	"strconv"
)

//=====================================================================================================================

func TypeCheckExpr(sourceCode string, expression parsing.IExpression) ITypedExpression {

	switch expr := expression.(type) {

	case *parsing.AdditionExpr:
		return typeCheckAdditionExpr(sourceCode, expr)
	case *parsing.BooleanLiteralExpr:
		return typeCheckBooleanLiteralExpr(expr)
	case *parsing.BuiltInTypeExpr:
		return typeCheckBuiltInTypeExpr(sourceCode, expr)
	case *parsing.DivisionExpr:
		return typeCheckDivisionExpr(sourceCode, expr)
	case *parsing.EqualsExpr:
		return typeCheckEqualsExpr(sourceCode, expr)
	case *parsing.FloatingPointLiteralExpr:
		return typeCheckFloatingPointLiteralExpr(sourceCode, expr)
	case *parsing.GreaterThanExpr:
		return typeCheckGreaterThanExpr(sourceCode, expr)
	case *parsing.GreaterThanOrEqualsExpr:
		return typeCheckGreaterThanOrEqualsExpr(sourceCode, expr)
	case *parsing.IntegerLiteralExpr:
		return typeCheckIntegerLiteralExpr(sourceCode, expr)
	case *parsing.IsExpr:
		return typeCheckIsExpr(sourceCode, expr)
	case *parsing.LessThanExpr:
		return typeCheckLessThanExpr(sourceCode, expr)
	case *parsing.LessThanOrEqualsExpr:
		return typeCheckLessThanOrEqualsExpr(sourceCode, expr)
	case *parsing.LogicalAndExpr:
		return typeCheckLogicalAndExpr(sourceCode, expr)
	case *parsing.LogicalNotOperationExpr:
		return typeCheckLogicalNotOperationExpr(sourceCode, expr)
	case *parsing.LogicalOrExpr:
		return typeCheckLogicalOrExpr(sourceCode, expr)
	case *parsing.MultiplicationExpr:
		return typeCheckMultiplicationExpr(sourceCode, expr)
	case *parsing.NegationOperationExpr:
		return typeCheckNegationOperationExpr(sourceCode, expr)
	case *parsing.NotEqualsExpr:
		return typeCheckNotEqualsExpr(sourceCode, expr)
	case *parsing.ParenthesizedExpr:
		return typeCheckParenthesizedExpr(sourceCode, expr)
	case *parsing.RecordExpr:
		return typeCheckRecordExpr(sourceCode, expr)
	case *parsing.StringLiteralExpr:
		return typeCheckStringLiteralExpr(sourceCode, expr)
	case *parsing.SubtractionExpr:
		return typeCheckSubtractionExpr(sourceCode, expr)

	default:
		panic(fmt.Sprintf("Missing case in TypeCheckExpr: %T\n", expression))

	}

}

//=====================================================================================================================

func typeCheckAdditionExpr(sourceCode string, expr *parsing.AdditionExpr) ITypedExpression {
	lhs := TypeCheckExpr(sourceCode, expr.Lhs)
	rhs := TypeCheckExpr(sourceCode, expr.Rhs)
	switch lhs.GetTypeInfo().(type) {
	case *types.Float64Type, *types.Int64Type:
		// TODO: ensure they're the same
		// TODO: coerce integers
		return &TypedAdditionExpr{
			SourcePosition: expr.SourcePosition,
			Lhs:            lhs,
			Rhs:            rhs,
			TypeInfo:       lhs.GetTypeInfo(),
		}
	case *types.StringType:
		// TODO: ensure both strings
		return &TypedStringConcatenationExpr{
			SourcePosition: expr.SourcePosition,
			Lhs:            lhs,
			Rhs:            rhs,
		}
	default:
		panic(fmt.Sprintf("Missing case in typeCheckAdditionExpr: %T\n", lhs.GetTypeInfo()))
	}
}

//=====================================================================================================================

func typeCheckBooleanLiteralExpr(expr *parsing.BooleanLiteralExpr) ITypedExpression {
	return &TypedBooleanLiteralExpr{
		SourcePosition: expr.SourcePosition,
		Value:          expr.Value,
	}
}

//=====================================================================================================================

func typeCheckBuiltInTypeExpr(sourceCode string, expr *parsing.BuiltInTypeExpr) ITypedExpression {
	name := expr.SourcePosition.GetText(sourceCode)
	return &TypedBuiltInTypeExpr{
		SourcePosition: expr.SourcePosition,
		Value:          types.BuiltInTypesByName[name],
	}
}

//=====================================================================================================================

func typeCheckDivisionExpr(sourceCode string, expr *parsing.DivisionExpr) ITypedExpression {
	lhs := TypeCheckExpr(sourceCode, expr.Lhs)
	rhs := TypeCheckExpr(sourceCode, expr.Rhs)
	// TODO: ensure they're the same
	// TODO: coerce integers
	return &TypedDivisionExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
		TypeInfo:       lhs.GetTypeInfo(),
	}
}

//=====================================================================================================================

func typeCheckEqualsExpr(sourceCode string, expr *parsing.EqualsExpr) ITypedExpression {
	lhs := TypeCheckExpr(sourceCode, expr.Lhs)
	rhs := TypeCheckExpr(sourceCode, expr.Rhs)
	// TODO: ensure they're the same
	// TODO: coerce integers
	return &TypedEqualsExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func typeCheckFloatingPointLiteralExpr(sourceCode string, expr *parsing.FloatingPointLiteralExpr) ITypedExpression {
	valueStr := expr.SourcePosition.GetText(sourceCode)
	value, _ := strconv.ParseFloat(valueStr, 64)
	return &TypedFloat64LiteralExpr{
		SourcePosition: expr.SourcePosition,
		Value:          value,
	}
}

//=====================================================================================================================

func typeCheckGreaterThanExpr(sourceCode string, expr *parsing.GreaterThanExpr) ITypedExpression {
	lhs := TypeCheckExpr(sourceCode, expr.Lhs)
	rhs := TypeCheckExpr(sourceCode, expr.Rhs)
	// TODO: ensure they're the same
	// TODO: coerce integers
	return &TypedGreaterThanExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func typeCheckGreaterThanOrEqualsExpr(sourceCode string, expr *parsing.GreaterThanOrEqualsExpr) ITypedExpression {
	lhs := TypeCheckExpr(sourceCode, expr.Lhs)
	rhs := TypeCheckExpr(sourceCode, expr.Rhs)
	// TODO: ensure they're the same
	// TODO: coerce integers
	return &TypedGreaterThanOrEqualsExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func typeCheckIntegerLiteralExpr(sourceCode string, expr *parsing.IntegerLiteralExpr) ITypedExpression {
	valueStr := expr.SourcePosition.GetText(sourceCode)
	value, _ := strconv.ParseInt(valueStr, 10, 64)
	return &TypedInt64LiteralExpr{
		SourcePosition: expr.SourcePosition,
		Value:          value,
	}
}

//=====================================================================================================================

func typeCheckIsExpr(sourceCode string, expr *parsing.IsExpr) ITypedExpression {
	lhs := TypeCheckExpr(sourceCode, expr.Lhs)
	rhs := TypeCheckExpr(sourceCode, expr.Rhs)
	// TODO: ensure the lhs and rhs are compatible
	return &TypedIsExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func typeCheckLessThanExpr(sourceCode string, expr *parsing.LessThanExpr) ITypedExpression {
	lhs := TypeCheckExpr(sourceCode, expr.Lhs)
	rhs := TypeCheckExpr(sourceCode, expr.Rhs)
	// TODO: ensure they're the same
	// TODO: coerce integers
	return &TypedLessThanExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func typeCheckLessThanOrEqualsExpr(sourceCode string, expr *parsing.LessThanOrEqualsExpr) ITypedExpression {
	lhs := TypeCheckExpr(sourceCode, expr.Lhs)
	rhs := TypeCheckExpr(sourceCode, expr.Rhs)
	// TODO: ensure they're the same
	// TODO: coerce integers
	return &TypedLessThanOrEqualsExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func typeCheckLogicalAndExpr(sourceCode string, expr *parsing.LogicalAndExpr) ITypedExpression {
	lhs := TypeCheckExpr(sourceCode, expr.Lhs)
	rhs := TypeCheckExpr(sourceCode, expr.Rhs)
	// TODO: ensure they're both boolean
	// TODO: coerce integers
	return &TypedLogicalAndExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func typeCheckLogicalNotOperationExpr(sourceCode string, expr *parsing.LogicalNotOperationExpr) ITypedExpression {
	operand := TypeCheckExpr(sourceCode, expr.Operand)

	// TODO: validate that operands are boolean

	return &TypedLogicalNotOperationExpr{
		SourcePosition: expr.SourcePosition,
		Operand:        operand,
	}
}

//=====================================================================================================================

func typeCheckLogicalOrExpr(sourceCode string, expr *parsing.LogicalOrExpr) ITypedExpression {
	lhs := TypeCheckExpr(sourceCode, expr.Lhs)
	rhs := TypeCheckExpr(sourceCode, expr.Rhs)
	// TODO: ensure they're both boolean
	// TODO: coerce integers
	return &TypedLogicalOrExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func typeCheckMultiplicationExpr(sourceCode string, expr *parsing.MultiplicationExpr) ITypedExpression {
	lhs := TypeCheckExpr(sourceCode, expr.Lhs)
	rhs := TypeCheckExpr(sourceCode, expr.Rhs)
	// TODO: ensure they're the same
	// TODO: coerce integers
	return &TypedMultiplicationExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
		TypeInfo:       lhs.GetTypeInfo(),
	}
}

//=====================================================================================================================

func typeCheckNegationOperationExpr(sourceCode string, expr *parsing.NegationOperationExpr) ITypedExpression {
	operand := TypeCheckExpr(sourceCode, expr.Operand)

	return &TypedNegationOperationExpr{
		SourcePosition: expr.SourcePosition,
		Operand:        operand,
		TypeInfo:       operand.GetTypeInfo(),
	}
}

//=====================================================================================================================

func typeCheckNotEqualsExpr(sourceCode string, expr *parsing.NotEqualsExpr) ITypedExpression {
	lhs := TypeCheckExpr(sourceCode, expr.Lhs)
	rhs := TypeCheckExpr(sourceCode, expr.Rhs)
	// TODO: ensure they're the same
	// TODO: coerce integers
	return &TypedNotEqualsExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func typeCheckParenthesizedExpr(sourceCode string, expr *parsing.ParenthesizedExpr) ITypedExpression {

	inner := TypeCheckExpr(sourceCode, expr.InnerExpr)

	return &TypedParenthesizedExpr{
		SourcePosition: expr.SourcePosition,
		InnerExpr:      inner,
		TypeInfo:       inner.GetTypeInfo(),
	}

}

//=====================================================================================================================

func typeCheckRecordExpr(sourceCode string, expr *parsing.RecordExpr) ITypedExpression {

	return &TypedRecordExpr{
		SourcePosition: expr.SourcePosition,
		TypeInfo:       types.Int64TypeInstance, // TODO (obviously)
	}

}

//=====================================================================================================================

func typeCheckStringLiteralExpr(sourceCode string, expr *parsing.StringLiteralExpr) ITypedExpression {
	text := expr.SourcePosition.GetText(sourceCode)
	var value string

	switch expr.Delimiters {
	case parsing.StringDelimitersDoubleQuotes:
		value = text[1 : len(text)-1]
	case parsing.StringDelimitersSingleQuotes:
		value = text[1 : len(text)-1]
	default:
		panic("TODO: Unhandled string delimiters")
	}

	// TODO: escape chars
	return &TypedStringLiteralExpr{
		SourcePosition: expr.SourcePosition,
		Value:          value,
	}
}

//=====================================================================================================================

func typeCheckSubtractionExpr(sourceCode string, expr *parsing.SubtractionExpr) ITypedExpression {
	lhs := TypeCheckExpr(sourceCode, expr.Lhs)
	rhs := TypeCheckExpr(sourceCode, expr.Rhs)
	// TODO: ensure they're the same
	// TODO: coerce integers
	return &TypedSubtractionExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
		TypeInfo:       lhs.GetTypeInfo(),
	}
}

//=====================================================================================================================
