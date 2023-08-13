//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package codegeneration

import (
	"fmt"
	"lligne-cli/internal/lligne/code/typechecking"
	"lligne-cli/internal/lligne/runtime/bytecode"
	"strconv"
)

//=====================================================================================================================

func GenerateByteCode(expression typechecking.ITypedExpression) *bytecode.CodeBlock {
	result := &bytecode.CodeBlock{}

	buildCodeBlock(result, expression)

	result.Stop()

	return result
}

//=====================================================================================================================

func buildCodeBlock(codeBlock *bytecode.CodeBlock, expression typechecking.ITypedExpression) {

	switch expr := expression.(type) {

	case *typechecking.TypedAdditionExpr:
		buildCodeBlock(codeBlock, expr.Lhs)
		buildCodeBlock(codeBlock, expr.Rhs)
		if expr.TypeInfo.BaseType() == typechecking.BaseTypeInt64 {
			codeBlock.Int64Add()
		} else {
			codeBlock.Float64Add()
		}

	case *typechecking.TypedBooleanLiteralExpr:
		if expr.Value {
			codeBlock.BoolLoadTrue()
		} else {
			codeBlock.BoolLoadFalse()
		}

	case *typechecking.TypedDivisionExpr:
		buildCodeBlock(codeBlock, expr.Lhs)
		buildCodeBlock(codeBlock, expr.Rhs)
		if expr.TypeInfo.BaseType() == typechecking.BaseTypeInt64 {
			codeBlock.Int64Divide()
		} else {
			codeBlock.Float64Divide()
		}

	case *typechecking.TypedEqualsExpr:
		buildCodeBlock(codeBlock, expr.Lhs)
		buildCodeBlock(codeBlock, expr.Rhs)
		if expr.TypeInfo.BaseType() == typechecking.BaseTypeInt64 {
			codeBlock.Int64Equals()
		} else {
			codeBlock.Float64Equals()
		}

	case *typechecking.TypedFloatingPointLiteralExpr:
		value, _ := strconv.ParseFloat(expr.Text, 64)
		switch value {
		case 0:
			codeBlock.Float64LoadZero()
		case 1:
			codeBlock.Float64LoadOne()
		default:
			codeBlock.Float64LoadFloat64(value)
		}

	case *typechecking.TypedGreaterThanExpr:
		buildCodeBlock(codeBlock, expr.Lhs)
		buildCodeBlock(codeBlock, expr.Rhs)
		if expr.TypeInfo.BaseType() == typechecking.BaseTypeInt64 {
			codeBlock.Int64GreaterThan()
		} else {
			codeBlock.Float64GreaterThan()
		}

	case *typechecking.TypedGreaterThanOrEqualsExpr:
		buildCodeBlock(codeBlock, expr.Lhs)
		buildCodeBlock(codeBlock, expr.Rhs)
		if expr.TypeInfo.BaseType() == typechecking.BaseTypeInt64 {
			codeBlock.Int64GreaterThanOrEquals()
		} else {
			codeBlock.Float64GreaterThanOrEquals()
		}

	case *typechecking.TypedIntegerLiteralExpr:
		value, _ := strconv.Atoi(expr.Text)
		switch value {
		case 0:
			codeBlock.Int64LoadZero()
		case 1:
			codeBlock.Int64LoadOne()
		default:
			codeBlock.Int64LoadInt16(int16(value))
		}

	case *typechecking.TypedLessThanExpr:
		buildCodeBlock(codeBlock, expr.Lhs)
		buildCodeBlock(codeBlock, expr.Rhs)
		if expr.TypeInfo.BaseType() == typechecking.BaseTypeInt64 {
			codeBlock.Int64LessThan()
		} else {
			codeBlock.Float64LessThan()
		}

	case *typechecking.TypedLessThanOrEqualsExpr:
		buildCodeBlock(codeBlock, expr.Lhs)
		buildCodeBlock(codeBlock, expr.Rhs)
		if expr.TypeInfo.BaseType() == typechecking.BaseTypeInt64 {
			codeBlock.Int64LessThanOrEquals()
		} else {
			codeBlock.Float64LessThanOrEquals()
		}

	case *typechecking.TypedLogicalAndExpr:
		buildCodeBlock(codeBlock, expr.Lhs)
		buildCodeBlock(codeBlock, expr.Rhs)
		codeBlock.BoolAnd()

	case *typechecking.TypedLogicalNotOperationExpr:
		buildCodeBlock(codeBlock, expr.Operand)
		codeBlock.BoolNot()

	case *typechecking.TypedLogicalOrExpr:
		buildCodeBlock(codeBlock, expr.Lhs)
		buildCodeBlock(codeBlock, expr.Rhs)
		codeBlock.BoolOr()

	case *typechecking.TypedMultiplicationExpr:
		buildCodeBlock(codeBlock, expr.Lhs)
		buildCodeBlock(codeBlock, expr.Rhs)
		if expr.TypeInfo.BaseType() == typechecking.BaseTypeInt64 {
			codeBlock.Int64Multiply()
		} else {
			codeBlock.Float64Multiply()
		}

	case *typechecking.TypedNegationOperationExpr:
		buildCodeBlock(codeBlock, expr.Operand)
		if expr.TypeInfo.BaseType() == typechecking.BaseTypeInt64 {
			codeBlock.Int64Negate()
		} else {
			codeBlock.Float64Negate()
		}

	case *typechecking.TypedParenthesizedExpr:
		if len(expr.Items) == 1 {
			buildCodeBlock(codeBlock, expr.Items[0])
		} else {
			panic("Records not yet handled")
		}

	case *typechecking.TypedSubtractionExpr:
		buildCodeBlock(codeBlock, expr.Lhs)
		buildCodeBlock(codeBlock, expr.Rhs)
		if expr.TypeInfo.BaseType() == typechecking.BaseTypeInt64 {
			codeBlock.Int64Subtract()
		} else {
			codeBlock.Float64Subtract()
		}

	default:
		panic(fmt.Sprintf("unmatched node %s", expression))

	}

}

//=====================================================================================================================
