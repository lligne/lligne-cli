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

	case *model.BooleanLiteralExpr:
		typeInfo := NewBoolType()
		return &TypedBooleanLiteralExpr{
			SourcePosition: expr.SourcePosition,
			Value:          expr.Value,
			TypeInfo:       typeInfo,
		}, typeInfo
	case *model.FloatingPointLiteralExpr:
		typeInfo := NewFloat64Type()
		return &TypedFloatingPointLiteralExpr{
			SourcePosition: expr.SourcePosition,
			Text:           expr.Text,
			TypeInfo:       typeInfo,
		}, typeInfo
	case *model.InfixOperationExpr:
		return typeCheckInfixOperationExpr(expr)
	case *model.IntegerLiteralExpr:
		typeInfo := NewInt64Type()
		return &TypedIntegerLiteralExpr{
			SourcePosition: expr.SourcePosition,
			Text:           expr.Text,
			TypeInfo:       typeInfo,
		}, typeInfo
	case *model.ParenthesizedExpr:
		return typeCheckParenthesizedExpr(expr)
	case *model.PrefixOperationExpr:
		return typeCheckPrefixOperationExpr(expr)

	}

	panic("Unhandled type check")
}

//=====================================================================================================================

func typeCheckInfixOperationExpr(expr *model.InfixOperationExpr) (ITypedExpression, IType) {
	lhs, lhsType := typeCheckExpr(expr.Lhs)
	rhs := TypeCheckExpr(expr.Rhs)
	return &TypedInfixOperationExpr{
		SourcePosition: expr.SourcePosition,
		Operator:       expr.Operator,
		Lhs:            lhs,
		Rhs:            rhs,
		TypeInfo:       lhsType,
	}, lhsType
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

func typeCheckPrefixOperationExpr(expr *model.PrefixOperationExpr) (ITypedExpression, IType) {
	operand, operandType := typeCheckExpr(expr.Operand)

	// TODO: lots more logic needed

	return &TypedPrefixOperationExpr{
		SourcePosition: expr.SourcePosition,
		Operator:       expr.Operator,
		Operand:        operand,
		TypeInfo:       operandType,
	}, operandType
}

//=====================================================================================================================
