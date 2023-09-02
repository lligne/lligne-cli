//
// # Data types related to Lligne token scanning.
//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package formatting

import (
	"fmt"
	"lligne-cli/internal/lligne/code/parsing"
	"strings"
)

//=====================================================================================================================

func FormatCode(parseOutcome *parsing.Outcome) string {
	return formatCode(parseOutcome.SourceCode, parseOutcome.Model)
}

//=====================================================================================================================

func formatCode(origSourceCode string, expression parsing.IExpression) string {

	switch expr := expression.(type) {

	case *parsing.AdditionExpr:
		return formatAdditionExpr(origSourceCode, expr)
	case *parsing.ArrayLiteralExpr:
		return formatSequenceLiteralExpr(origSourceCode, expr)
	case *parsing.BooleanLiteralExpr:
		return formatBooleanLiteralExpr(expr)
	case *parsing.BuiltInTypeExpr:
		return formatBuiltInTypeExpr(origSourceCode, expr)
	case *parsing.DivisionExpr:
		return formatDivisionExpr(origSourceCode, expr)
	case *parsing.EqualsExpr:
		return formatEqualsExpr(origSourceCode, expr)
	case *parsing.FieldReferenceExpr:
		return formatFieldReferenceExpr(origSourceCode, expr)
	case *parsing.Float64LiteralExpr:
		return formatFloatingPointLiteralExpr(origSourceCode, expr)
	case *parsing.FunctionArgumentsExpr:
		return formatFunctionArgumentsExpr(origSourceCode, expr)
	case *parsing.FunctionArrowExpr:
		return formatFunctionArrowExpr(origSourceCode, expr)
	case *parsing.FunctionCallExpr:
		return formatFunctionCallExpr(origSourceCode, expr)
	case *parsing.GreaterThanExpr:
		return formatGreaterThanExpr(origSourceCode, expr)
	case *parsing.GreaterThanOrEqualsExpr:
		return formatGreaterThanOrEqualsExpr(origSourceCode, expr)
	case *parsing.IdentifierExpr:
		return formatIdentifierExpr(origSourceCode, expr)
	case *parsing.InExpr:
		return formatInExpr(origSourceCode, expr)
	case *parsing.IsExpr:
		return formatIsExpr(origSourceCode, expr)
	case *parsing.Int64LiteralExpr:
		return formatIntegerLiteralExpr(origSourceCode, expr)
	case *parsing.IntersectAssignValueExpr:
		return formatIntersectAssignValueExpr(origSourceCode, expr)
	case *parsing.IntersectExpr:
		return formatIntersectExpr(origSourceCode, expr)
	case *parsing.IntersectDefaultValueExpr:
		return formatIntersectDefaultValueExpr(origSourceCode, expr)
	case *parsing.IntersectLowPrecedenceExpr:
		return formatIntersectLowPrecedenceExpr(origSourceCode, expr)
	case *parsing.LessThanExpr:
		return formatLessThanExpr(origSourceCode, expr)
	case *parsing.LessThanOrEqualsExpr:
		return formatLessThanOrEqualsExpr(origSourceCode, expr)
	case *parsing.LogicalAndExpr:
		return formatLogicalAndExpr(origSourceCode, expr)
	case *parsing.LogicalNotOperationExpr:
		return formatLogicalNotOperationExpr(origSourceCode, expr)
	case *parsing.LogicalOrExpr:
		return formatLogicalOrExpr(origSourceCode, expr)
	case *parsing.MatchExpr:
		return formatMatchExpr(origSourceCode, expr)
	case *parsing.MultiplicationExpr:
		return formatMultiplicationExpr(origSourceCode, expr)
	case *parsing.NegationOperationExpr:
		return formatNegationOperationExpr(origSourceCode, expr)
	case *parsing.NotEqualsExpr:
		return formatNotEqualsExpr(origSourceCode, expr)
	case *parsing.NotMatchExpr:
		return formatNotMatchExpr(origSourceCode, expr)
	case *parsing.OptionalExpr:
		return formatOptionalExpr(origSourceCode, expr)
	case *parsing.ParenthesizedExpr:
		return formatParenthesizedExpr(origSourceCode, expr)
	case *parsing.QualifyExpr:
		return formatQualifyExpr(origSourceCode, expr)
	case *parsing.RangeExpr:
		return formatRangeExpr(origSourceCode, expr)
	case *parsing.RecordExpr:
		return formatRecordExpr(origSourceCode, expr)
	case *parsing.StringLiteralExpr:
		return formatStringLiteralExpr(origSourceCode, expr)
	case *parsing.SubtractionExpr:
		return formatSubtractionExpr(origSourceCode, expr)
	case *parsing.UnionExpr:
		return formatUnionExpr(origSourceCode, expr)
	case *parsing.UnitExpr:
		return formatUnitExpr(origSourceCode, expr)
	case *parsing.WhenExpr:
		return formatWhenExpr(origSourceCode, expr)
	case *parsing.WhereExpr:
		return formatWhereExpr(origSourceCode, expr)

	default:
		panic(fmt.Sprintf("Missing case in formatCode: %T\n", expression))

	}

}

//=====================================================================================================================

func formatAdditionExpr(sourceCode string, expr *parsing.AdditionExpr) string {
	lhs := formatCode(sourceCode, expr.Lhs)
	rhs := formatCode(sourceCode, expr.Rhs)
	return lhs + " + " + rhs
}

//=====================================================================================================================

func formatBooleanLiteralExpr(expr *parsing.BooleanLiteralExpr) string {
	if expr.Value {
		return "true"
	}
	return "false"
}

//=====================================================================================================================

func formatBuiltInTypeExpr(sourceCode string, expr *parsing.BuiltInTypeExpr) string {
	return expr.SourcePosition.GetText(sourceCode)
}

//=====================================================================================================================

func formatDivisionExpr(sourceCode string, expr *parsing.DivisionExpr) string {
	lhs := formatCode(sourceCode, expr.Lhs)
	rhs := formatCode(sourceCode, expr.Rhs)
	return lhs + " / " + rhs
}

//=====================================================================================================================

func formatEqualsExpr(sourceCode string, expr *parsing.EqualsExpr) string {
	lhs := formatCode(sourceCode, expr.Lhs)
	rhs := formatCode(sourceCode, expr.Rhs)
	return lhs + " == " + rhs
}

//=====================================================================================================================

func formatFieldReferenceExpr(sourceCode string, expr *parsing.FieldReferenceExpr) string {
	lhs := formatCode(sourceCode, expr.Parent)
	rhs := formatCode(sourceCode, expr.Child)
	return lhs + "." + rhs
}

//=====================================================================================================================

func formatFloatingPointLiteralExpr(sourceCode string, expr *parsing.Float64LiteralExpr) string {
	return expr.SourcePosition.GetText(sourceCode)
}

//=====================================================================================================================

func formatFunctionArgumentsExpr(sourceCode string, expr *parsing.FunctionArgumentsExpr) string {

	sb := strings.Builder{}

	sb.WriteString("(")

	if len(expr.Items) > 0 {

		sb.WriteString(formatCode(sourceCode, expr.Items[0]))

		for _, item := range expr.Items[1:] {
			sb.WriteString(", ")
			sb.WriteString(formatCode(sourceCode, item))
		}

	}

	sb.WriteString(")")

	return sb.String()

}

//=====================================================================================================================

func formatFunctionArrowExpr(sourceCode string, expr *parsing.FunctionArrowExpr) string {
	arg := formatCode(sourceCode, expr.Argument)
	result := formatCode(sourceCode, expr.Result)
	return arg + " -> " + result
}

//=====================================================================================================================

func formatFunctionCallExpr(sourceCode string, expr *parsing.FunctionCallExpr) string {
	fun := formatCode(sourceCode, expr.FunctionReference)
	arg := formatCode(sourceCode, expr.Argument)
	return fun + arg
}

//=====================================================================================================================

func formatGreaterThanExpr(sourceCode string, expr *parsing.GreaterThanExpr) string {
	lhs := formatCode(sourceCode, expr.Lhs)
	rhs := formatCode(sourceCode, expr.Rhs)
	return lhs + " > " + rhs
}

//=====================================================================================================================

func formatGreaterThanOrEqualsExpr(sourceCode string, expr *parsing.GreaterThanOrEqualsExpr) string {
	lhs := formatCode(sourceCode, expr.Lhs)
	rhs := formatCode(sourceCode, expr.Rhs)
	return lhs + " >= " + rhs
}

//=====================================================================================================================

func formatIdentifierExpr(sourceCode string, expr *parsing.IdentifierExpr) string {
	return expr.SourcePosition.GetText(sourceCode)
}

//=====================================================================================================================

func formatInExpr(sourceCode string, expr *parsing.InExpr) string {
	lhs := formatCode(sourceCode, expr.Lhs)
	rhs := formatCode(sourceCode, expr.Rhs)
	return lhs + " in " + rhs
}

//=====================================================================================================================

func formatIsExpr(sourceCode string, expr *parsing.IsExpr) string {
	lhs := formatCode(sourceCode, expr.Lhs)
	rhs := formatCode(sourceCode, expr.Rhs)
	return lhs + " is " + rhs
}

//=====================================================================================================================

func formatIntegerLiteralExpr(sourceCode string, expr *parsing.Int64LiteralExpr) string {
	return expr.SourcePosition.GetText(sourceCode)
}

//=====================================================================================================================

func formatIntersectAssignValueExpr(sourceCode string, expr *parsing.IntersectAssignValueExpr) string {
	lhs := formatCode(sourceCode, expr.Lhs)
	rhs := formatCode(sourceCode, expr.Rhs)
	return lhs + " = " + rhs
}

//=====================================================================================================================

func formatIntersectExpr(sourceCode string, expr *parsing.IntersectExpr) string {
	lhs := formatCode(sourceCode, expr.Lhs)
	rhs := formatCode(sourceCode, expr.Rhs)
	return lhs + " & " + rhs
}

//=====================================================================================================================

func formatIntersectDefaultValueExpr(sourceCode string, expr *parsing.IntersectDefaultValueExpr) string {
	lhs := formatCode(sourceCode, expr.Lhs)
	rhs := formatCode(sourceCode, expr.Rhs)
	return lhs + " ?: " + rhs
}

//=====================================================================================================================

func formatIntersectLowPrecedenceExpr(sourceCode string, expr *parsing.IntersectLowPrecedenceExpr) string {
	lhs := formatCode(sourceCode, expr.Lhs)
	rhs := formatCode(sourceCode, expr.Rhs)
	return lhs + " && " + rhs
}

//=====================================================================================================================

func formatLessThanExpr(sourceCode string, expr *parsing.LessThanExpr) string {
	lhs := formatCode(sourceCode, expr.Lhs)
	rhs := formatCode(sourceCode, expr.Rhs)
	return lhs + " < " + rhs
}

//=====================================================================================================================

func formatLessThanOrEqualsExpr(sourceCode string, expr *parsing.LessThanOrEqualsExpr) string {
	lhs := formatCode(sourceCode, expr.Lhs)
	rhs := formatCode(sourceCode, expr.Rhs)
	return lhs + " <= " + rhs
}

//=====================================================================================================================

func formatLogicalAndExpr(sourceCode string, expr *parsing.LogicalAndExpr) string {
	lhs := formatCode(sourceCode, expr.Lhs)
	rhs := formatCode(sourceCode, expr.Rhs)
	return lhs + " and " + rhs
}

//=====================================================================================================================

func formatLogicalNotOperationExpr(sourceCode string, expr *parsing.LogicalNotOperationExpr) string {
	return "not " + formatCode(sourceCode, expr.Operand)
}

//=====================================================================================================================

func formatLogicalOrExpr(sourceCode string, expr *parsing.LogicalOrExpr) string {
	lhs := formatCode(sourceCode, expr.Lhs)
	rhs := formatCode(sourceCode, expr.Rhs)
	return lhs + " or " + rhs
}

//=====================================================================================================================

func formatMatchExpr(sourceCode string, expr *parsing.MatchExpr) string {
	lhs := formatCode(sourceCode, expr.Lhs)
	rhs := formatCode(sourceCode, expr.Rhs)
	return lhs + " =~ " + rhs
}

//=====================================================================================================================

func formatMultiplicationExpr(sourceCode string, expr *parsing.MultiplicationExpr) string {
	lhs := formatCode(sourceCode, expr.Lhs)
	rhs := formatCode(sourceCode, expr.Rhs)
	return lhs + " * " + rhs
}

//=====================================================================================================================

func formatNegationOperationExpr(sourceCode string, expr *parsing.NegationOperationExpr) string {
	return "-" + formatCode(sourceCode, expr.Operand)
}

//=====================================================================================================================

func formatNotEqualsExpr(sourceCode string, expr *parsing.NotEqualsExpr) string {
	lhs := formatCode(sourceCode, expr.Lhs)
	rhs := formatCode(sourceCode, expr.Rhs)
	return lhs + " != " + rhs
}

//=====================================================================================================================

func formatNotMatchExpr(sourceCode string, expr *parsing.NotMatchExpr) string {
	lhs := formatCode(sourceCode, expr.Lhs)
	rhs := formatCode(sourceCode, expr.Rhs)
	return lhs + " !~ " + rhs
}

//=====================================================================================================================

func formatOptionalExpr(sourceCode string, expr *parsing.OptionalExpr) string {
	return formatCode(sourceCode, expr.Operand) + "?"
}

//=====================================================================================================================

func formatParenthesizedExpr(sourceCode string, expr *parsing.ParenthesizedExpr) string {

	sb := strings.Builder{}

	sb.WriteString("(")

	sb.WriteString(formatCode(sourceCode, expr.InnerExpr))

	sb.WriteString(")")

	return sb.String()

}

//=====================================================================================================================

func formatQualifyExpr(sourceCode string, expr *parsing.QualifyExpr) string {
	lhs := formatCode(sourceCode, expr.Lhs)
	rhs := formatCode(sourceCode, expr.Rhs)
	return lhs + ": " + rhs
}

//=====================================================================================================================

func formatRangeExpr(sourceCode string, expr *parsing.RangeExpr) string {
	first := formatCode(sourceCode, expr.First)
	last := formatCode(sourceCode, expr.Last)
	return first + ".." + last
}

//=====================================================================================================================

func formatRecordExpr(sourceCode string, expr *parsing.RecordExpr) string {

	sb := strings.Builder{}

	sb.WriteString("{")

	if len(expr.Items) > 0 {

		sb.WriteString(formatCode(sourceCode, expr.Items[0]))

		for _, item := range expr.Items[1:] {
			sb.WriteString(", ")
			sb.WriteString(formatCode(sourceCode, item))
		}

	}

	sb.WriteString("}")

	return sb.String()

}

//=====================================================================================================================

func formatSequenceLiteralExpr(sourceCode string, expr *parsing.ArrayLiteralExpr) string {

	sb := strings.Builder{}

	sb.WriteString("[")

	if len(expr.Elements) > 0 {

		sb.WriteString(formatCode(sourceCode, expr.Elements[0]))

		for _, item := range expr.Elements[1:] {
			sb.WriteString(", ")
			sb.WriteString(formatCode(sourceCode, item))
		}

	}

	sb.WriteString("]")

	return sb.String()

}

//=====================================================================================================================

func formatStringLiteralExpr(sourceCode string, expr *parsing.StringLiteralExpr) string {
	return expr.SourcePosition.GetText(sourceCode)
}

//=====================================================================================================================

func formatSubtractionExpr(sourceCode string, expr *parsing.SubtractionExpr) string {
	lhs := formatCode(sourceCode, expr.Lhs)
	rhs := formatCode(sourceCode, expr.Rhs)
	return lhs + " - " + rhs
}

//=====================================================================================================================

func formatUnionExpr(sourceCode string, expr *parsing.UnionExpr) string {
	lhs := formatCode(sourceCode, expr.Lhs)
	rhs := formatCode(sourceCode, expr.Rhs)
	return lhs + " | " + rhs
}

//=====================================================================================================================

func formatUnitExpr(sourceCode string, expr *parsing.UnitExpr) string {
	return "()"
}

//=====================================================================================================================

func formatWhenExpr(sourceCode string, expr *parsing.WhenExpr) string {
	lhs := formatCode(sourceCode, expr.Lhs)
	rhs := formatCode(sourceCode, expr.Rhs)
	return lhs + " when " + rhs
}

//=====================================================================================================================

func formatWhereExpr(sourceCode string, expr *parsing.WhereExpr) string {
	lhs := formatCode(sourceCode, expr.Lhs)
	rhs := formatCode(sourceCode, expr.Rhs)
	return lhs + " where " + rhs
}

//=====================================================================================================================
