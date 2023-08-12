//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package model

import (
	"fmt"
	"strings"
)

//=====================================================================================================================

func SExpression(expression IExpression) string {
	switch expr := expression.(type) {

	case *BooleanLiteralExpr:
		if expr.Value {
			return "(bool true)"
		}
		return "(bool false)"

	case *FloatingPointLiteralExpr:
		return "(float " + expr.Text + ")"

	case *FunctionCallExpr:
		result := "(call "
		result += SExpression(expr.FunctionReference)
		result += " "
		result += SExpression(expr.Argument)
		result += ")"
		return result

	case *IdentifierExpr:
		return "(id " + expr.Name + ")"

	case *InfixOperationExpr:
		return "(" + strings.TrimSpace(expr.Operator.String()) + " " +
			SExpression(expr.Lhs) + " " + SExpression(expr.Rhs) + ")"

	case *IntegerLiteralExpr:
		return "(int " + expr.Text + ")"

	case *LeadingDocumentationExpr:
		return "(leadingdoc\n" + expr.Text + ")"

	case *MultilineStringLiteralExpr:
		return "(multilinestr\n" + expr.Text + ")"

	case *OptionalExpr:
		result := "(optional "
		result += SExpression(expr.Operand)
		result += ")"
		return result

	case *ParenthesizedExpr:
		var result string

		switch expr.Delimiters {
		case ParenExprDelimitersParentheses:
			result = "(parenthesized \"()\""
		case ParenExprDelimitersBraces:
			result = "(parenthesized \"{}\""
		case ParenExprDelimitersWholeFile:
			result = "(parenthesized \"\""
		}

		for _, item := range expr.Items {
			result += " "
			result += SExpression(item)
		}

		result += ")"

		return result

	case *PrefixOperationExpr:
		return "(prefix " + strings.TrimSpace(expr.Operator.String()) + " " + SExpression(expr.Operand) + ")"

	case *SequenceLiteralExpr:
		result := "(sequence"

		for _, element := range expr.Elements {
			result += " "
			result += SExpression(element)
		}

		result += ")"

		return result

	case *StringLiteralExpr:
		return "(string " + expr.Text + ")"

	case *TrailingDocumentationExpr:
		return "(trailingdoc\n" + expr.Text + ")"

	default:
		panic(fmt.Sprintf("Missing case in SExpression: %T\n", expression))

	}

}

//=====================================================================================================================
