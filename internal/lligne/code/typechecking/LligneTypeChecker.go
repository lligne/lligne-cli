//
// # Data types related to Lligne token scanning.
//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package typechecking

import "lligne-cli/internal/lligne/code/model"

//=====================================================================================================================

func DetermineTypes(expression *model.ILligneExpression) {

	switch (*expression).ExprType() {

	case model.ExprTypeBooleanLiteral:
		(*expression).SetTypeInfo(model.NewBoolType())
	case model.ExprTypeFloatingPointLiteral:
		(*expression).SetTypeInfo(model.NewFloat64Type())
	case model.ExprTypeInfixOperation:
		determineInfixOperationTypes((*expression).(*model.LligneInfixOperationExpr))
	case model.ExprTypeIntegerLiteral:
		(*expression).SetTypeInfo(model.NewInt64Type())
	case model.ExprTypeParenthesized:
		determineParenthesizedTypes((*expression).(*model.LligneParenthesizedExpr))
	case model.ExprTypePrefixOperation:
		determinePrefixOperationTypes((*expression).(*model.LlignePrefixOperationExpr))

	}
}

//=====================================================================================================================

func determineInfixOperationTypes(expression *model.LligneInfixOperationExpr) {
	DetermineTypes(&expression.Lhs)
	DetermineTypes(&expression.Rhs)

	// TODO: lots more logic needed

	expression.SetTypeInfo(expression.Lhs.TypeInfo())
}

//=====================================================================================================================

func determineParenthesizedTypes(expression *model.LligneParenthesizedExpr) {
	for _, item := range expression.Items {
		DetermineTypes(&item)
	}

	// TODO: lots more logic needed

	expression.SetTypeInfo(expression.Items[0].TypeInfo())
}

//=====================================================================================================================

func determinePrefixOperationTypes(expression *model.LlignePrefixOperationExpr) {
	DetermineTypes(&expression.Operand)

	// TODO: lots more logic needed

	expression.SetTypeInfo(expression.Operand.TypeInfo())
}

//=====================================================================================================================
