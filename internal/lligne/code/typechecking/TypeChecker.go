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
	"strconv"
)

//=====================================================================================================================

func TypeCheckExpr(expression parsing.IExpression) ITypedExpression {

	switch expr := expression.(type) {

	case *parsing.AdditionExpr:
		return typeCheckAdditionExpr(expr)
	case *parsing.BooleanLiteralExpr:
		return typeCheckBooleanLiteralExpr(expr)
	case *parsing.DivisionExpr:
		return typeCheckDivisionExpr(expr)
	case *parsing.EqualsExpr:
		return typeCheckEqualsExpr(expr)
	case *parsing.FloatingPointLiteralExpr:
		return typeCheckFloatingPointLiteralExpr(expr)
	case *parsing.GreaterThanExpr:
		return typeCheckGreaterThanExpr(expr)
	case *parsing.GreaterThanOrEqualsExpr:
		return typeCheckGreaterThanOrEqualsExpr(expr)
	case *parsing.IntegerLiteralExpr:
		return typeCheckIntegerLiteralExpr(expr)
	case *parsing.LessThanExpr:
		return typeCheckLessThanExpr(expr)
	case *parsing.LessThanOrEqualsExpr:
		return typeCheckLessThanOrEqualsExpr(expr)
	case *parsing.LogicalAndExpr:
		return typeCheckLogicalAndExpr(expr)
	case *parsing.LogicalNotOperationExpr:
		return typeCheckLogicalNotOperationExpr(expr)
	case *parsing.LogicalOrExpr:
		return typeCheckLogicalOrExpr(expr)
	case *parsing.MultiplicationExpr:
		return typeCheckMultiplicationExpr(expr)
	case *parsing.NegationOperationExpr:
		return typeCheckNegationOperationExpr(expr)
	case *parsing.ParenthesizedExpr:
		return typeCheckParenthesizedExpr(expr)
	case *parsing.StringLiteralExpr:
		return typeCheckStringLiteralExpr(expr)
	case *parsing.SubtractionExpr:
		return typeCheckSubtractionExpr(expr)

	default:
		panic(fmt.Sprintf("Missing case in TypeCheckExpr: %T\n", expression))

	}

}

//=====================================================================================================================

func typeCheckAdditionExpr(expr *parsing.AdditionExpr) ITypedExpression {
	lhs := TypeCheckExpr(expr.Lhs)
	rhs := TypeCheckExpr(expr.Rhs)
	// TODO: ensure they're the same
	// TODO: coerce integers
	return &TypedAdditionExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
		TypeInfo:       lhs.GetTypeInfo(),
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

func typeCheckDivisionExpr(expr *parsing.DivisionExpr) ITypedExpression {
	lhs := TypeCheckExpr(expr.Lhs)
	rhs := TypeCheckExpr(expr.Rhs)
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

func typeCheckEqualsExpr(expr *parsing.EqualsExpr) ITypedExpression {
	lhs := TypeCheckExpr(expr.Lhs)
	rhs := TypeCheckExpr(expr.Rhs)
	// TODO: ensure they're the same
	// TODO: coerce integers
	return &TypedEqualsExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func typeCheckFloatingPointLiteralExpr(expr *parsing.FloatingPointLiteralExpr) ITypedExpression {
	value, _ := strconv.ParseFloat(expr.Text, 64)
	return &TypedFloat64LiteralExpr{
		SourcePosition: expr.SourcePosition,
		Value:          value,
	}
}

//=====================================================================================================================

func typeCheckGreaterThanExpr(expr *parsing.GreaterThanExpr) ITypedExpression {
	lhs := TypeCheckExpr(expr.Lhs)
	rhs := TypeCheckExpr(expr.Rhs)
	// TODO: ensure they're the same
	// TODO: coerce integers
	return &TypedGreaterThanExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func typeCheckGreaterThanOrEqualsExpr(expr *parsing.GreaterThanOrEqualsExpr) ITypedExpression {
	lhs := TypeCheckExpr(expr.Lhs)
	rhs := TypeCheckExpr(expr.Rhs)
	// TODO: ensure they're the same
	// TODO: coerce integers
	return &TypedGreaterThanOrEqualsExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func typeCheckIntegerLiteralExpr(expr *parsing.IntegerLiteralExpr) ITypedExpression {
	value, _ := strconv.ParseInt(expr.Text, 10, 64)
	return &TypedInt64LiteralExpr{
		SourcePosition: expr.SourcePosition,
		Value:          value,
	}
}

//=====================================================================================================================

func typeCheckLessThanExpr(expr *parsing.LessThanExpr) ITypedExpression {
	lhs := TypeCheckExpr(expr.Lhs)
	rhs := TypeCheckExpr(expr.Rhs)
	// TODO: ensure they're the same
	// TODO: coerce integers
	return &TypedLessThanExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func typeCheckLessThanOrEqualsExpr(expr *parsing.LessThanOrEqualsExpr) ITypedExpression {
	lhs := TypeCheckExpr(expr.Lhs)
	rhs := TypeCheckExpr(expr.Rhs)
	// TODO: ensure they're the same
	// TODO: coerce integers
	return &TypedLessThanOrEqualsExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func typeCheckLogicalAndExpr(expr *parsing.LogicalAndExpr) ITypedExpression {
	lhs := TypeCheckExpr(expr.Lhs)
	rhs := TypeCheckExpr(expr.Rhs)
	// TODO: ensure they're both boolean
	// TODO: coerce integers
	return &TypedLogicalAndExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func typeCheckLogicalNotOperationExpr(expr *parsing.LogicalNotOperationExpr) ITypedExpression {
	operand := TypeCheckExpr(expr.Operand)

	// TODO: validate that operands are boolean

	return &TypedLogicalNotOperationExpr{
		SourcePosition: expr.SourcePosition,
		Operand:        operand,
	}
}

//=====================================================================================================================

func typeCheckLogicalOrExpr(expr *parsing.LogicalOrExpr) ITypedExpression {
	lhs := TypeCheckExpr(expr.Lhs)
	rhs := TypeCheckExpr(expr.Rhs)
	// TODO: ensure they're both boolean
	// TODO: coerce integers
	return &TypedLogicalOrExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func typeCheckMultiplicationExpr(expr *parsing.MultiplicationExpr) ITypedExpression {
	lhs := TypeCheckExpr(expr.Lhs)
	rhs := TypeCheckExpr(expr.Rhs)
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

func typeCheckNegationOperationExpr(expr *parsing.NegationOperationExpr) ITypedExpression {
	operand := TypeCheckExpr(expr.Operand)

	return &TypedNegationOperationExpr{
		SourcePosition: expr.SourcePosition,
		Operand:        operand,
		TypeInfo:       operand.GetTypeInfo(),
	}
}

//=====================================================================================================================

func typeCheckParenthesizedExpr(expr *parsing.ParenthesizedExpr) ITypedExpression {

	var items []ITypedExpression

	for _, item0 := range expr.Items {
		item := TypeCheckExpr(item0)
		items = append(items, item)
	}

	// TODO: lots more logic needed

	return &TypedParenthesizedExpr{
		SourcePosition: expr.SourcePosition,
		Delimiters:     expr.Delimiters,
		Items:          items,
		TypeInfo:       items[0].GetTypeInfo(),
	}

}

//=====================================================================================================================

func typeCheckStringLiteralExpr(expr *parsing.StringLiteralExpr) ITypedExpression {
	// TODO: escape chars
	value := expr.Text[1 : len(expr.Text)-1]
	return &TypedStringLiteralExpr{
		SourcePosition: expr.SourcePosition,
		Value:          value,
	}
}

//=====================================================================================================================

func typeCheckSubtractionExpr(expr *parsing.SubtractionExpr) ITypedExpression {
	lhs := TypeCheckExpr(expr.Lhs)
	rhs := TypeCheckExpr(expr.Rhs)
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
