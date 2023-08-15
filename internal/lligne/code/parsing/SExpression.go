//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package parsing

import (
	"fmt"
)

//=====================================================================================================================

func SExpression(expression IExpression) string {
	switch expr := expression.(type) {

	case *AdditionExpr:
		return "(add " + SExpression(expr.Lhs) + " " + SExpression(expr.Rhs) + ")"

	case *BooleanLiteralExpr:
		if expr.Value {
			return "(bool true)"
		}
		return "(bool false)"

	case *DivisionExpr:
		return "(divide " + SExpression(expr.Lhs) + " " + SExpression(expr.Rhs) + ")"

	case *DocumentExpr:
		return "(doc " + SExpression(expr.Lhs) + " " + SExpression(expr.Rhs) + ")"

	case *EqualsExpr:
		return "(equals " + SExpression(expr.Lhs) + " " + SExpression(expr.Rhs) + ")"

	case *FieldReferenceExpr:
		return "(fieldref " + SExpression(expr.Parent) + " " + SExpression(expr.Child) + ")"

	case *FloatingPointLiteralExpr:
		return "(float " + expr.Text + ")"

	case *FunctionArrowExpr:
		result := "(arrow "
		result += SExpression(expr.Argument)
		result += " "
		result += SExpression(expr.Result)
		result += ")"
		return result

	case *FunctionCallExpr:
		result := "(call "
		result += SExpression(expr.FunctionReference)
		result += " "
		result += SExpression(expr.Argument)
		result += ")"
		return result

	case *GreaterThanExpr:
		return "(greater " + SExpression(expr.Lhs) + " " + SExpression(expr.Rhs) + ")"

	case *GreaterThanOrEqualsExpr:
		return "(greaterorequals " + SExpression(expr.Lhs) + " " + SExpression(expr.Rhs) + ")"

	case *IdentifierExpr:
		return "(id " + expr.Name + ")"

	case *InExpr:
		return "(in " + SExpression(expr.Lhs) + " " + SExpression(expr.Rhs) + ")"

	case *IntegerLiteralExpr:
		return "(int " + expr.Text + ")"

	case *IntersectExpr:
		return "(intersect " + SExpression(expr.Lhs) + " " + SExpression(expr.Rhs) + ")"

	case *IntersectAssignValueExpr:
		return "(intersectassign " + SExpression(expr.Lhs) + " " + SExpression(expr.Rhs) + ")"

	case *IntersectDefaultValueExpr:
		return "(intersectdefault " + SExpression(expr.Lhs) + " " + SExpression(expr.Rhs) + ")"

	case *IntersectLowPrecedenceExpr:
		return "(intersectlowprecedence " + SExpression(expr.Lhs) + " " + SExpression(expr.Rhs) + ")"

	case *IsExpr:
		return "(is " + SExpression(expr.Lhs) + " " + SExpression(expr.Rhs) + ")"

	case *LeadingDocumentationExpr:
		return "(leadingdoc\n" + expr.Text + ")"

	case *LessThanExpr:
		return "(less " + SExpression(expr.Lhs) + " " + SExpression(expr.Rhs) + ")"

	case *LessThanOrEqualsExpr:
		return "(lessorequals " + SExpression(expr.Lhs) + " " + SExpression(expr.Rhs) + ")"

	case *LogicalAndExpr:
		return "(and " + SExpression(expr.Lhs) + " " + SExpression(expr.Rhs) + ")"

	case *LogicalNotOperationExpr:
		return "(not " + SExpression(expr.Operand) + ")"

	case *LogicalOrExpr:
		return "(or " + SExpression(expr.Lhs) + " " + SExpression(expr.Rhs) + ")"

	case *MatchExpr:
		return "(match " + SExpression(expr.Lhs) + " " + SExpression(expr.Rhs) + ")"

	case *MultilineStringLiteralExpr:
		return "(multilinestr\n" + expr.Text + ")"

	case *MultiplicationExpr:
		return "(multiply " + SExpression(expr.Lhs) + " " + SExpression(expr.Rhs) + ")"

	case *NegationOperationExpr:
		return "(negate " + SExpression(expr.Operand) + ")"

	case *NotMatchExpr:
		return "(notmatch " + SExpression(expr.Lhs) + " " + SExpression(expr.Rhs) + ")"

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

	case *QualifyExpr:
		return "(qualify " + SExpression(expr.Lhs) + " " + SExpression(expr.Rhs) + ")"

	case *RangeExpr:
		return "(range " + SExpression(expr.First) + " " + SExpression(expr.Last) + ")"

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

	case *SubtractionExpr:
		return "(subtract " + SExpression(expr.Lhs) + " " + SExpression(expr.Rhs) + ")"

	case *TrailingDocumentationExpr:
		return "(trailingdoc\n" + expr.Text + ")"

	case *UnionExpr:
		return "(union " + SExpression(expr.Lhs) + " " + SExpression(expr.Rhs) + ")"

	case *WhenExpr:
		return "(when " + SExpression(expr.Lhs) + " " + SExpression(expr.Rhs) + ")"

	case *WhereExpr:
		return "(where " + SExpression(expr.Lhs) + " " + SExpression(expr.Rhs) + ")"

	default:
		panic(fmt.Sprintf("Missing case in SExpression: %T\n", expression))

	}

}

//=====================================================================================================================
