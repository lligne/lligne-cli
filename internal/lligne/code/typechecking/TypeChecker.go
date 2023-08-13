//
// # Data types related to Lligne token scanning.
//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package typechecking

import (
	"lligne-cli/internal/lligne/code/model"
)

//=====================================================================================================================

func TypeCheckExpr(expression model.IExpression) ITypedExpression {
	result, _ := typeCheckExpr(expression)
	return result
}

//=====================================================================================================================

func typeCheckExpr(expression model.IExpression) (ITypedExpression, IType) {

	switch expr := expression.(type) {

	case *model.AdditionExpr:
		return typeCheckAdditionExpr(expr)

	case *model.BooleanLiteralExpr:
		typeInfo := NewBoolType()
		return &TypedBooleanLiteralExpr{
			SourcePosition: expr.SourcePosition,
			Value:          expr.Value,
			TypeInfo:       typeInfo,
		}, typeInfo

	case *model.DivisionExpr:
		return typeCheckDivisionExpr(expr)

	case *model.EqualsExpr:
		return typeCheckEqualsExpr(expr)

	case *model.FloatingPointLiteralExpr:
		typeInfo := NewFloat64Type()
		return &TypedFloatingPointLiteralExpr{
			SourcePosition: expr.SourcePosition,
			Text:           expr.Text,
			TypeInfo:       typeInfo,
		}, typeInfo

	case *model.GreaterThanExpr:
		return typeCheckGreaterThanExpr(expr)

	case *model.GreaterThanOrEqualsExpr:
		return typeCheckGreaterThanOrEqualsExpr(expr)

	case *model.IntegerLiteralExpr:
		typeInfo := NewInt64Type()
		return &TypedIntegerLiteralExpr{
			SourcePosition: expr.SourcePosition,
			Text:           expr.Text,
			TypeInfo:       typeInfo,
		}, typeInfo

	case *model.LessThanExpr:
		return typeCheckLessThanExpr(expr)

	case *model.LessThanOrEqualsExpr:
		return typeCheckLessThanOrEqualsExpr(expr)

	case *model.LogicalAndExpr:
		return typeCheckLogicalAndExpr(expr)

	case *model.LogicalNotOperationExpr:
		return typeCheckLogicalNotOperationExpr(expr)

	case *model.LogicalOrExpr:
		return typeCheckLogicalOrExpr(expr)

	case *model.MultiplicationExpr:
		return typeCheckMultiplicationExpr(expr)

	case *model.NegationOperationExpr:
		return typeCheckNegationOperationExpr(expr)

	case *model.ParenthesizedExpr:
		return typeCheckParenthesizedExpr(expr)

	case *model.SubtractionExpr:
		return typeCheckSubtractionExpr(expr)

	}

	panic("Unhandled type check")
}

//=====================================================================================================================

func typeCheckAdditionExpr(expr *model.AdditionExpr) (ITypedExpression, IType) {
	lhs, lhsType := typeCheckExpr(expr.Lhs)
	rhs := TypeCheckExpr(expr.Rhs)
	// TODO: ensure they're the same
	// TODO: coerce integers
	return &TypedAdditionExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
		TypeInfo:       lhsType,
	}, lhsType
}

//=====================================================================================================================

func typeCheckDivisionExpr(expr *model.DivisionExpr) (ITypedExpression, IType) {
	lhs, lhsType := typeCheckExpr(expr.Lhs)
	rhs := TypeCheckExpr(expr.Rhs)
	// TODO: ensure they're the same
	// TODO: coerce integers
	return &TypedDivisionExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
		TypeInfo:       lhsType,
	}, lhsType
}

//=====================================================================================================================

func typeCheckEqualsExpr(expr *model.EqualsExpr) (ITypedExpression, IType) {
	lhs, lhsType := typeCheckExpr(expr.Lhs)
	rhs := TypeCheckExpr(expr.Rhs)
	// TODO: ensure they're the same
	// TODO: coerce integers
	return &TypedEqualsExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
		TypeInfo:       lhsType,
	}, lhsType
}

//=====================================================================================================================

func typeCheckGreaterThanExpr(expr *model.GreaterThanExpr) (ITypedExpression, IType) {
	lhs, lhsType := typeCheckExpr(expr.Lhs)
	rhs := TypeCheckExpr(expr.Rhs)
	// TODO: ensure they're the same
	// TODO: coerce integers
	return &TypedGreaterThanExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
		TypeInfo:       lhsType,
	}, lhsType
}

//=====================================================================================================================

func typeCheckGreaterThanOrEqualsExpr(expr *model.GreaterThanOrEqualsExpr) (ITypedExpression, IType) {
	lhs, lhsType := typeCheckExpr(expr.Lhs)
	rhs := TypeCheckExpr(expr.Rhs)
	// TODO: ensure they're the same
	// TODO: coerce integers
	return &TypedGreaterThanOrEqualsExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
		TypeInfo:       lhsType,
	}, lhsType
}

//=====================================================================================================================

func typeCheckLessThanExpr(expr *model.LessThanExpr) (ITypedExpression, IType) {
	lhs, lhsType := typeCheckExpr(expr.Lhs)
	rhs := TypeCheckExpr(expr.Rhs)
	// TODO: ensure they're the same
	// TODO: coerce integers
	return &TypedLessThanExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
		TypeInfo:       lhsType,
	}, lhsType
}

//=====================================================================================================================

func typeCheckLessThanOrEqualsExpr(expr *model.LessThanOrEqualsExpr) (ITypedExpression, IType) {
	lhs, lhsType := typeCheckExpr(expr.Lhs)
	rhs := TypeCheckExpr(expr.Rhs)
	// TODO: ensure they're the same
	// TODO: coerce integers
	return &TypedLessThanOrEqualsExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
		TypeInfo:       lhsType,
	}, lhsType
}

//=====================================================================================================================

func typeCheckLogicalAndExpr(expr *model.LogicalAndExpr) (ITypedExpression, IType) {
	lhs, _ := typeCheckExpr(expr.Lhs)
	rhs := TypeCheckExpr(expr.Rhs)
	exprType := NewBoolType()
	// TODO: ensure they're both boolean
	// TODO: coerce integers
	return &TypedLogicalAndExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
		TypeInfo:       exprType,
	}, exprType
}

//=====================================================================================================================

func typeCheckLogicalNotOperationExpr(expr *model.LogicalNotOperationExpr) (ITypedExpression, IType) {
	operand, operandType := typeCheckExpr(expr.Operand)

	// TODO: validate that operands are boolean

	return &TypedLogicalNotOperationExpr{
		SourcePosition: expr.SourcePosition,
		Operand:        operand,
		TypeInfo:       operandType,
	}, operandType
}

//=====================================================================================================================

func typeCheckLogicalOrExpr(expr *model.LogicalOrExpr) (ITypedExpression, IType) {
	lhs, _ := typeCheckExpr(expr.Lhs)
	rhs := TypeCheckExpr(expr.Rhs)
	exprType := NewBoolType()
	// TODO: ensure they're both boolean
	// TODO: coerce integers
	return &TypedLogicalOrExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
		TypeInfo:       exprType,
	}, exprType
}

//=====================================================================================================================

func typeCheckMultiplicationExpr(expr *model.MultiplicationExpr) (ITypedExpression, IType) {
	lhs, lhsType := typeCheckExpr(expr.Lhs)
	rhs := TypeCheckExpr(expr.Rhs)
	// TODO: ensure they're the same
	// TODO: coerce integers
	return &TypedMultiplicationExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
		TypeInfo:       lhsType,
	}, lhsType
}

//=====================================================================================================================

func typeCheckNegationOperationExpr(expr *model.NegationOperationExpr) (ITypedExpression, IType) {
	operand, operandType := typeCheckExpr(expr.Operand)

	return &TypedNegationOperationExpr{
		SourcePosition: expr.SourcePosition,
		Operand:        operand,
		TypeInfo:       operandType,
	}, operandType
}

//=====================================================================================================================

func typeCheckParenthesizedExpr(expr *model.ParenthesizedExpr) (ITypedExpression, IType) {

	var items []ITypedExpression
	var itemTypes []IType

	for _, item0 := range expr.Items {
		item, itemType := typeCheckExpr(item0)
		items = append(items, item)
		itemTypes = append(itemTypes, itemType)
	}

	// TODO: lots more logic needed

	return &TypedParenthesizedExpr{
		SourcePosition: expr.SourcePosition,
		Delimiters:     expr.Delimiters,
		Items:          items,
		TypeInfo:       itemTypes[0],
	}, itemTypes[0]

}

//=====================================================================================================================

func typeCheckSubtractionExpr(expr *model.SubtractionExpr) (ITypedExpression, IType) {
	lhs, lhsType := typeCheckExpr(expr.Lhs)
	rhs := TypeCheckExpr(expr.Rhs)
	// TODO: ensure they're the same
	// TODO: coerce integers
	return &TypedSubtractionExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
		TypeInfo:       lhsType,
	}, lhsType
}

//=====================================================================================================================
