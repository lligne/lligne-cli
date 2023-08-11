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

func DetermineTypes(expression *model.IExpression) {

	switch (*expression).(type) {

	case *model.BooleanLiteralExpr:
		(*expression).SetTypeInfo(model.NewBoolType())
	case *model.FloatingPointLiteralExpr:
		(*expression).SetTypeInfo(model.NewFloat64Type())
	case *model.InfixOperationExpr:
		determineInfixOperationTypes((*expression).(*model.InfixOperationExpr))
	case *model.IntegerLiteralExpr:
		(*expression).SetTypeInfo(model.NewInt64Type())
	case *model.ParenthesizedExpr:
		determineParenthesizedTypes((*expression).(*model.ParenthesizedExpr))
	case *model.PrefixOperationExpr:
		determinePrefixOperationTypes((*expression).(*model.PrefixOperationExpr))

	}
}

//=====================================================================================================================

func determineInfixOperationTypes(expression *model.InfixOperationExpr) {
	DetermineTypes(&expression.Lhs)
	DetermineTypes(&expression.Rhs)

	// TODO: lots more logic needed

	expression.SetTypeInfo(expression.Lhs.TypeInfo())
}

//=====================================================================================================================

func determineParenthesizedTypes(expression *model.ParenthesizedExpr) {
	for _, item := range expression.Items {
		DetermineTypes(&item)
	}

	// TODO: lots more logic needed

	expression.SetTypeInfo(expression.Items[0].TypeInfo())
}

//=====================================================================================================================

func determinePrefixOperationTypes(expression *model.PrefixOperationExpr) {
	DetermineTypes(&expression.Operand)

	// TODO: lots more logic needed

	expression.SetTypeInfo(expression.Operand.TypeInfo())
}

//=====================================================================================================================
