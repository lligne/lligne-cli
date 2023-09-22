//
// # Data types related to Lligne token scanning.
//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package formatting

import (
	"fmt"
	prior "lligne-cli/internal/lligne/code/parsing"
	"lligne-cli/internal/lligne/runtime/pools"
	"strings"
)

//=====================================================================================================================

func FormatCode(parseOutcome *prior.Outcome) string {
	formatter := newFormatter(parseOutcome)
	return formatter.formatCode(parseOutcome.Model)
}

//=====================================================================================================================

type formatter struct {
	SourceCode      string
	NewLineOffsets  []uint32
	StringConstants *pools.StringPool
	IdentifierNames *pools.NamePool
}

//---------------------------------------------------------------------------------------------------------------------

func newFormatter(priorOutcome *prior.Outcome) *formatter {
	return &formatter{
		SourceCode:      priorOutcome.SourceCode,
		NewLineOffsets:  priorOutcome.NewLineOffsets,
		StringConstants: pools.NewStringPool(),
		IdentifierNames: pools.NewNamePool(),
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (f *formatter) formatCode(expression prior.IExpression) string {

	switch expr := expression.(type) {

	case *prior.AdditionExpr:
		return f.formatAdditionExpr(expr)
	case *prior.ArrayLiteralExpr:
		return f.formatSequenceLiteralExpr(expr)
	case *prior.BooleanLiteralExpr:
		return f.formatBooleanLiteralExpr(expr)
	case *prior.BuiltInTypeExpr:
		return f.formatBuiltInTypeExpr(expr)
	case *prior.DivisionExpr:
		return f.formatDivisionExpr(expr)
	case *prior.EqualsExpr:
		return f.formatEqualsExpr(expr)
	case *prior.FieldReferenceExpr:
		return f.formatFieldReferenceExpr(expr)
	case *prior.Float64LiteralExpr:
		return f.formatFloatingPointLiteralExpr(expr)
	case *prior.FunctionArgumentsExpr:
		return f.formatFunctionArgumentsExpr(expr)
	case *prior.FunctionArrowExpr:
		return f.formatFunctionArrowExpr(expr)
	case *prior.FunctionCallExpr:
		return f.formatFunctionCallExpr(expr)
	case *prior.GreaterThanExpr:
		return f.formatGreaterThanExpr(expr)
	case *prior.GreaterThanOrEqualsExpr:
		return f.formatGreaterThanOrEqualsExpr(expr)
	case *prior.IdentifierExpr:
		return f.formatIdentifierExpr(expr)
	case *prior.InExpr:
		return f.formatInExpr(expr)
	case *prior.IsExpr:
		return f.formatIsExpr(expr)
	case *prior.Int64LiteralExpr:
		return f.formatIntegerLiteralExpr(expr)
	case *prior.IntersectAssignValueExpr:
		return f.formatIntersectAssignValueExpr(expr)
	case *prior.IntersectExpr:
		return f.formatIntersectExpr(expr)
	case *prior.IntersectDefaultValueExpr:
		return f.formatIntersectDefaultValueExpr(expr)
	case *prior.IntersectLowPrecedenceExpr:
		return f.formatIntersectLowPrecedenceExpr(expr)
	case *prior.LessThanExpr:
		return f.formatLessThanExpr(expr)
	case *prior.LessThanOrEqualsExpr:
		return f.formatLessThanOrEqualsExpr(expr)
	case *prior.LogicalAndExpr:
		return f.formatLogicalAndExpr(expr)
	case *prior.LogicalNotOperationExpr:
		return f.formatLogicalNotOperationExpr(expr)
	case *prior.LogicalOrExpr:
		return f.formatLogicalOrExpr(expr)
	case *prior.MatchExpr:
		return f.formatMatchExpr(expr)
	case *prior.MultiplicationExpr:
		return f.formatMultiplicationExpr(expr)
	case *prior.NegationOperationExpr:
		return f.formatNegationOperationExpr(expr)
	case *prior.NotEqualsExpr:
		return f.formatNotEqualsExpr(expr)
	case *prior.NotMatchExpr:
		return f.formatNotMatchExpr(expr)
	case *prior.OptionalExpr:
		return f.formatOptionalExpr(expr)
	case *prior.ParenthesizedExpr:
		return f.formatParenthesizedExpr(expr)
	case *prior.QualifyExpr:
		return f.formatQualifyExpr(expr)
	case *prior.RangeExpr:
		return f.formatRangeExpr(expr)
	case *prior.RecordExpr:
		return f.formatRecordExpr(expr)
	case *prior.StringLiteralExpr:
		return f.formatStringLiteralExpr(expr)
	case *prior.SubtractionExpr:
		return f.formatSubtractionExpr(expr)
	case *prior.UnionExpr:
		return f.formatUnionExpr(expr)
	case *prior.UnitExpr:
		return f.formatUnitExpr()
	case *prior.WhenExpr:
		return f.formatWhenExpr(expr)
	case *prior.WhereExpr:
		return f.formatWhereExpr(expr)

	default:
		panic(fmt.Sprintf("Missing case in formatCode: %T\n", expression))

	}

}

//---------------------------------------------------------------------------------------------------------------------

func (f *formatter) formatAdditionExpr(expr *prior.AdditionExpr) string {
	lhs := f.formatCode(expr.Lhs)
	rhs := f.formatCode(expr.Rhs)
	return lhs + " + " + rhs
}

//---------------------------------------------------------------------------------------------------------------------

func (f *formatter) formatBooleanLiteralExpr(expr *prior.BooleanLiteralExpr) string {
	if expr.Value {
		return "true"
	}
	return "false"
}

//---------------------------------------------------------------------------------------------------------------------

func (f *formatter) formatBuiltInTypeExpr(expr *prior.BuiltInTypeExpr) string {
	return expr.SourcePosition.GetText(f.SourceCode)
}

//---------------------------------------------------------------------------------------------------------------------

func (f *formatter) formatDivisionExpr(expr *prior.DivisionExpr) string {
	lhs := f.formatCode(expr.Lhs)
	rhs := f.formatCode(expr.Rhs)
	return lhs + " / " + rhs
}

//---------------------------------------------------------------------------------------------------------------------

func (f *formatter) formatEqualsExpr(expr *prior.EqualsExpr) string {
	lhs := f.formatCode(expr.Lhs)
	rhs := f.formatCode(expr.Rhs)
	return lhs + " == " + rhs
}

//---------------------------------------------------------------------------------------------------------------------

func (f *formatter) formatFieldReferenceExpr(expr *prior.FieldReferenceExpr) string {
	lhs := f.formatCode(expr.Parent)
	rhs := f.formatCode(expr.Child)
	return lhs + "." + rhs
}

//---------------------------------------------------------------------------------------------------------------------

func (f *formatter) formatFloatingPointLiteralExpr(expr *prior.Float64LiteralExpr) string {
	return expr.SourcePosition.GetText(f.SourceCode)
}

//---------------------------------------------------------------------------------------------------------------------

func (f *formatter) formatFunctionArgumentsExpr(expr *prior.FunctionArgumentsExpr) string {

	sb := strings.Builder{}

	sb.WriteString("(")

	if len(expr.Items) > 0 {

		sb.WriteString(f.formatCode(expr.Items[0]))

		for _, item := range expr.Items[1:] {
			sb.WriteString(", ")
			sb.WriteString(f.formatCode(item))
		}

	}

	sb.WriteString(")")

	return sb.String()

}

//---------------------------------------------------------------------------------------------------------------------

func (f *formatter) formatFunctionArrowExpr(expr *prior.FunctionArrowExpr) string {
	arg := f.formatCode(expr.Argument)
	result := f.formatCode(expr.Result)
	return arg + " -> " + result
}

//---------------------------------------------------------------------------------------------------------------------

func (f *formatter) formatFunctionCallExpr(expr *prior.FunctionCallExpr) string {
	fun := f.formatCode(expr.FunctionReference)
	arg := f.formatCode(expr.Argument)
	return fun + arg
}

//---------------------------------------------------------------------------------------------------------------------

func (f *formatter) formatGreaterThanExpr(expr *prior.GreaterThanExpr) string {
	lhs := f.formatCode(expr.Lhs)
	rhs := f.formatCode(expr.Rhs)
	return lhs + " > " + rhs
}

//---------------------------------------------------------------------------------------------------------------------

func (f *formatter) formatGreaterThanOrEqualsExpr(expr *prior.GreaterThanOrEqualsExpr) string {
	lhs := f.formatCode(expr.Lhs)
	rhs := f.formatCode(expr.Rhs)
	return lhs + " >= " + rhs
}

//---------------------------------------------------------------------------------------------------------------------

func (f *formatter) formatIdentifierExpr(expr *prior.IdentifierExpr) string {
	return expr.SourcePosition.GetText(f.SourceCode)
}

//---------------------------------------------------------------------------------------------------------------------

func (f *formatter) formatInExpr(expr *prior.InExpr) string {
	lhs := f.formatCode(expr.Lhs)
	rhs := f.formatCode(expr.Rhs)
	return lhs + " in " + rhs
}

//---------------------------------------------------------------------------------------------------------------------

func (f *formatter) formatIsExpr(expr *prior.IsExpr) string {
	lhs := f.formatCode(expr.Lhs)
	rhs := f.formatCode(expr.Rhs)
	return lhs + " is " + rhs
}

//---------------------------------------------------------------------------------------------------------------------

func (f *formatter) formatIntegerLiteralExpr(expr *prior.Int64LiteralExpr) string {
	return expr.SourcePosition.GetText(f.SourceCode)
}

//---------------------------------------------------------------------------------------------------------------------

func (f *formatter) formatIntersectAssignValueExpr(expr *prior.IntersectAssignValueExpr) string {
	lhs := f.formatCode(expr.Lhs)
	rhs := f.formatCode(expr.Rhs)
	return lhs + " = " + rhs
}

//---------------------------------------------------------------------------------------------------------------------

func (f *formatter) formatIntersectExpr(expr *prior.IntersectExpr) string {
	lhs := f.formatCode(expr.Lhs)
	rhs := f.formatCode(expr.Rhs)
	return lhs + " & " + rhs
}

//---------------------------------------------------------------------------------------------------------------------

func (f *formatter) formatIntersectDefaultValueExpr(expr *prior.IntersectDefaultValueExpr) string {
	lhs := f.formatCode(expr.Lhs)
	rhs := f.formatCode(expr.Rhs)
	return lhs + " ?: " + rhs
}

//---------------------------------------------------------------------------------------------------------------------

func (f *formatter) formatIntersectLowPrecedenceExpr(expr *prior.IntersectLowPrecedenceExpr) string {
	lhs := f.formatCode(expr.Lhs)
	rhs := f.formatCode(expr.Rhs)
	return lhs + " && " + rhs
}

//---------------------------------------------------------------------------------------------------------------------

func (f *formatter) formatLessThanExpr(expr *prior.LessThanExpr) string {
	lhs := f.formatCode(expr.Lhs)
	rhs := f.formatCode(expr.Rhs)
	return lhs + " < " + rhs
}

//---------------------------------------------------------------------------------------------------------------------

func (f *formatter) formatLessThanOrEqualsExpr(expr *prior.LessThanOrEqualsExpr) string {
	lhs := f.formatCode(expr.Lhs)
	rhs := f.formatCode(expr.Rhs)
	return lhs + " <= " + rhs
}

//---------------------------------------------------------------------------------------------------------------------

func (f *formatter) formatLogicalAndExpr(expr *prior.LogicalAndExpr) string {
	lhs := f.formatCode(expr.Lhs)
	rhs := f.formatCode(expr.Rhs)
	return lhs + " and " + rhs
}

//---------------------------------------------------------------------------------------------------------------------

func (f *formatter) formatLogicalNotOperationExpr(expr *prior.LogicalNotOperationExpr) string {
	return "not " + f.formatCode(expr.Operand)
}

//---------------------------------------------------------------------------------------------------------------------

func (f *formatter) formatLogicalOrExpr(expr *prior.LogicalOrExpr) string {
	lhs := f.formatCode(expr.Lhs)
	rhs := f.formatCode(expr.Rhs)
	return lhs + " or " + rhs
}

//---------------------------------------------------------------------------------------------------------------------

func (f *formatter) formatMatchExpr(expr *prior.MatchExpr) string {
	lhs := f.formatCode(expr.Lhs)
	rhs := f.formatCode(expr.Rhs)
	return lhs + " =~ " + rhs
}

//---------------------------------------------------------------------------------------------------------------------

func (f *formatter) formatMultiplicationExpr(expr *prior.MultiplicationExpr) string {
	lhs := f.formatCode(expr.Lhs)
	rhs := f.formatCode(expr.Rhs)
	return lhs + " * " + rhs
}

//---------------------------------------------------------------------------------------------------------------------

func (f *formatter) formatNegationOperationExpr(expr *prior.NegationOperationExpr) string {
	return "-" + f.formatCode(expr.Operand)
}

//---------------------------------------------------------------------------------------------------------------------

func (f *formatter) formatNotEqualsExpr(expr *prior.NotEqualsExpr) string {
	lhs := f.formatCode(expr.Lhs)
	rhs := f.formatCode(expr.Rhs)
	return lhs + " != " + rhs
}

//---------------------------------------------------------------------------------------------------------------------

func (f *formatter) formatNotMatchExpr(expr *prior.NotMatchExpr) string {
	lhs := f.formatCode(expr.Lhs)
	rhs := f.formatCode(expr.Rhs)
	return lhs + " !~ " + rhs
}

//---------------------------------------------------------------------------------------------------------------------

func (f *formatter) formatOptionalExpr(expr *prior.OptionalExpr) string {
	return f.formatCode(expr.Operand) + "?"
}

//---------------------------------------------------------------------------------------------------------------------

func (f *formatter) formatParenthesizedExpr(expr *prior.ParenthesizedExpr) string {

	sb := strings.Builder{}

	sb.WriteString("(")

	sb.WriteString(f.formatCode(expr.InnerExpr))

	sb.WriteString(")")

	return sb.String()

}

//---------------------------------------------------------------------------------------------------------------------

func (f *formatter) formatQualifyExpr(expr *prior.QualifyExpr) string {
	lhs := f.formatCode(expr.Lhs)
	rhs := f.formatCode(expr.Rhs)
	return lhs + ": " + rhs
}

//---------------------------------------------------------------------------------------------------------------------

func (f *formatter) formatRangeExpr(expr *prior.RangeExpr) string {
	first := f.formatCode(expr.First)
	last := f.formatCode(expr.Last)
	return first + ".." + last
}

//---------------------------------------------------------------------------------------------------------------------

func (f *formatter) formatRecordExpr(expr *prior.RecordExpr) string {

	sb := strings.Builder{}

	sb.WriteString("{")

	if len(expr.Items) > 0 {

		sb.WriteString(f.formatCode(expr.Items[0]))

		for _, item := range expr.Items[1:] {
			sb.WriteString(", ")
			sb.WriteString(f.formatCode(item))
		}

	}

	sb.WriteString("}")

	return sb.String()

}

//---------------------------------------------------------------------------------------------------------------------

func (f *formatter) formatSequenceLiteralExpr(expr *prior.ArrayLiteralExpr) string {

	sb := strings.Builder{}

	sb.WriteString("[")

	if len(expr.Elements) > 0 {

		sb.WriteString(f.formatCode(expr.Elements[0]))

		for _, item := range expr.Elements[1:] {
			sb.WriteString(", ")
			sb.WriteString(f.formatCode(item))
		}

	}

	sb.WriteString("]")

	return sb.String()

}

//---------------------------------------------------------------------------------------------------------------------

func (f *formatter) formatStringLiteralExpr(expr *prior.StringLiteralExpr) string {
	return expr.SourcePosition.GetText(f.SourceCode)
}

//---------------------------------------------------------------------------------------------------------------------

func (f *formatter) formatSubtractionExpr(expr *prior.SubtractionExpr) string {
	lhs := f.formatCode(expr.Lhs)
	rhs := f.formatCode(expr.Rhs)
	return lhs + " - " + rhs
}

//---------------------------------------------------------------------------------------------------------------------

func (f *formatter) formatUnionExpr(expr *prior.UnionExpr) string {
	lhs := f.formatCode(expr.Lhs)
	rhs := f.formatCode(expr.Rhs)
	return lhs + " | " + rhs
}

//---------------------------------------------------------------------------------------------------------------------

func (f *formatter) formatUnitExpr() string {
	return "()"
}

//---------------------------------------------------------------------------------------------------------------------

func (f *formatter) formatWhenExpr(expr *prior.WhenExpr) string {
	lhs := f.formatCode(expr.Lhs)
	rhs := f.formatCode(expr.Rhs)
	return lhs + " when " + rhs
}

//---------------------------------------------------------------------------------------------------------------------

func (f *formatter) formatWhereExpr(expr *prior.WhereExpr) string {
	lhs := f.formatCode(expr.Lhs)
	rhs := f.formatCode(expr.Rhs)
	return lhs + " where " + rhs
}

//---------------------------------------------------------------------------------------------------------------------
