//
// # Data types related to Lligne token scanning.
//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package typechecking

import (
	"lligne-cli/internal/lligne/code/parsing"
)

//=====================================================================================================================

func TypeCheckExpr(expression parsing.IExpression) ITypedExpression {
	result, _ := typeCheckExpr(expression)
	return result
}

//=====================================================================================================================

func typeCheckExpr(expression parsing.IExpression) (ITypedExpression, IType) {

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
	case *parsing.SubtractionExpr:
		return typeCheckSubtractionExpr(expr)
	default:
		panic("Unhandled type check")

	}

}

//=====================================================================================================================

func typeCheckBooleanLiteralExpr(expr *parsing.BooleanLiteralExpr) (ITypedExpression, IType) {
	typeInfo := BoolTypeInstance
	return &TypedBooleanLiteralExpr{
		SourcePosition: expr.SourcePosition,
		Value:          expr.Value,
		TypeInfo:       typeInfo,
	}, typeInfo
}

//=====================================================================================================================

func typeCheckAdditionExpr(expr *parsing.AdditionExpr) (ITypedExpression, IType) {
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

func typeCheckDivisionExpr(expr *parsing.DivisionExpr) (ITypedExpression, IType) {
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

func typeCheckEqualsExpr(expr *parsing.EqualsExpr) (ITypedExpression, IType) {
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

func typeCheckFloatingPointLiteralExpr(expr *parsing.FloatingPointLiteralExpr) (ITypedExpression, IType) {
	typeInfo := Float64TypeInstance
	return &TypedFloatingPointLiteralExpr{
		SourcePosition: expr.SourcePosition,
		Text:           expr.Text,
		TypeInfo:       typeInfo,
	}, typeInfo
}

//=====================================================================================================================

func typeCheckGreaterThanExpr(expr *parsing.GreaterThanExpr) (ITypedExpression, IType) {
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

func typeCheckGreaterThanOrEqualsExpr(expr *parsing.GreaterThanOrEqualsExpr) (ITypedExpression, IType) {
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

func typeCheckIntegerLiteralExpr(expr *parsing.IntegerLiteralExpr) (ITypedExpression, IType) {
	typeInfo := Int64TypeInstance
	return &TypedIntegerLiteralExpr{
		SourcePosition: expr.SourcePosition,
		Text:           expr.Text,
		TypeInfo:       typeInfo,
	}, typeInfo
}

//=====================================================================================================================

func typeCheckLessThanExpr(expr *parsing.LessThanExpr) (ITypedExpression, IType) {
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

func typeCheckLessThanOrEqualsExpr(expr *parsing.LessThanOrEqualsExpr) (ITypedExpression, IType) {
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

func typeCheckLogicalAndExpr(expr *parsing.LogicalAndExpr) (ITypedExpression, IType) {
	lhs, _ := typeCheckExpr(expr.Lhs)
	rhs := TypeCheckExpr(expr.Rhs)
	exprType := BoolTypeInstance
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

func typeCheckLogicalNotOperationExpr(expr *parsing.LogicalNotOperationExpr) (ITypedExpression, IType) {
	operand, operandType := typeCheckExpr(expr.Operand)

	// TODO: validate that operands are boolean

	return &TypedLogicalNotOperationExpr{
		SourcePosition: expr.SourcePosition,
		Operand:        operand,
		TypeInfo:       operandType,
	}, operandType
}

//=====================================================================================================================

func typeCheckLogicalOrExpr(expr *parsing.LogicalOrExpr) (ITypedExpression, IType) {
	lhs, _ := typeCheckExpr(expr.Lhs)
	rhs := TypeCheckExpr(expr.Rhs)
	exprType := BoolTypeInstance
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

func typeCheckMultiplicationExpr(expr *parsing.MultiplicationExpr) (ITypedExpression, IType) {
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

func typeCheckNegationOperationExpr(expr *parsing.NegationOperationExpr) (ITypedExpression, IType) {
	operand, operandType := typeCheckExpr(expr.Operand)

	return &TypedNegationOperationExpr{
		SourcePosition: expr.SourcePosition,
		Operand:        operand,
		TypeInfo:       operandType,
	}, operandType
}

//=====================================================================================================================

func typeCheckParenthesizedExpr(expr *parsing.ParenthesizedExpr) (ITypedExpression, IType) {

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

func typeCheckSubtractionExpr(expr *parsing.SubtractionExpr) (ITypedExpression, IType) {
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
