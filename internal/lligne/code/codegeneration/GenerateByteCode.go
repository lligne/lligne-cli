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
		buildAdditionCodeBlock(codeBlock, expr)
	case *typechecking.TypedBooleanLiteralExpr:
		buildBooleanLiteralCodeBlock(codeBlock, expr)
	case *typechecking.TypedDivisionExpr:
		buildDivisionCodeBlock(codeBlock, expr)
	case *typechecking.TypedEqualsExpr:
		buildEqualsCodeBlock(codeBlock, expr)
	case *typechecking.TypedFloatingPointLiteralExpr:
		buildFloatingPointLiteralCodeBlock(codeBlock, expr)
	case *typechecking.TypedGreaterThanExpr:
		buildGreaterThanCodeBlock(codeBlock, expr)
	case *typechecking.TypedGreaterThanOrEqualsExpr:
		buildGreaterThanOrEqualsCodeBlock(codeBlock, expr)
	case *typechecking.TypedIntegerLiteralExpr:
		buildIntegerLiteralCodeBlock(codeBlock, expr)
	case *typechecking.TypedLessThanExpr:
		buildLessThanCodeBlock(codeBlock, expr)
	case *typechecking.TypedLessThanOrEqualsExpr:
		buildLessThanOrEqualsCodeBlock(codeBlock, expr)
	case *typechecking.TypedLogicalAndExpr:
		buildLogicalAndCodeBlock(codeBlock, expr)
	case *typechecking.TypedLogicalNotOperationExpr:
		buildLogicalNotCodeBlock(codeBlock, expr)
	case *typechecking.TypedLogicalOrExpr:
		buildLogicalOrCodeBlock(codeBlock, expr)
	case *typechecking.TypedMultiplicationExpr:
		buildMultiplicationCodeBlock(codeBlock, expr)
	case *typechecking.TypedNegationOperationExpr:
		buildNegationCodeBlock(codeBlock, expr)
	case *typechecking.TypedParenthesizedExpr:
		buildParenthesizedCodeBlock(codeBlock, expr)
	case *typechecking.TypedSubtractionExpr:
		buildSubtractionCodeBlock(codeBlock, expr)
	default:
		panic(fmt.Sprintf("unmatched node %s", expression))

	}

}

//=====================================================================================================================

func buildAdditionCodeBlock(codeBlock *bytecode.CodeBlock, expr *typechecking.TypedAdditionExpr) {

	if e, ok := expr.Lhs.(*typechecking.TypedIntegerLiteralExpr); ok && e.Text == "1" {
		buildCodeBlock(codeBlock, expr.Rhs)
		codeBlock.Int64Increment()
	} else if e, ok := expr.Rhs.(*typechecking.TypedIntegerLiteralExpr); ok && e.Text == "1" {
		buildCodeBlock(codeBlock, expr.Lhs)
		codeBlock.Int64Increment()
	} else {
		buildCodeBlock(codeBlock, expr.Lhs)
		buildCodeBlock(codeBlock, expr.Rhs)
		switch expr.TypeInfo.(type) {
		case *typechecking.Float64Type:
			codeBlock.Float64Add()
		case *typechecking.Int64Type:
			codeBlock.Int64Add()
		default:
			panic("Undefined addition type")
		}
	}
}

//=====================================================================================================================

func buildBooleanLiteralCodeBlock(codeBlock *bytecode.CodeBlock, expr *typechecking.TypedBooleanLiteralExpr) {
	if expr.Value {
		codeBlock.BoolLoadTrue()
	} else {
		codeBlock.BoolLoadFalse()
	}
}

//=====================================================================================================================

func buildDivisionCodeBlock(codeBlock *bytecode.CodeBlock, expr *typechecking.TypedDivisionExpr) {
	buildCodeBlock(codeBlock, expr.Lhs)
	buildCodeBlock(codeBlock, expr.Rhs)
	switch expr.TypeInfo.(type) {
	case *typechecking.Float64Type:
		codeBlock.Float64Divide()
	case *typechecking.Int64Type:
		codeBlock.Int64Divide()
	default:
		panic("Undefined division type")
	}
}

//=====================================================================================================================

func buildEqualsCodeBlock(codeBlock *bytecode.CodeBlock, expr *typechecking.TypedEqualsExpr) {
	buildCodeBlock(codeBlock, expr.Lhs)
	buildCodeBlock(codeBlock, expr.Rhs)
	switch expr.TypeInfo.(type) {
	case *typechecking.Float64Type:
		codeBlock.Float64Equals()
	case *typechecking.Int64Type:
		codeBlock.Int64Equals()
	default:
		panic("Undefined equality type")
	}
}

//=====================================================================================================================

func buildFloatingPointLiteralCodeBlock(codeBlock *bytecode.CodeBlock, expr *typechecking.TypedFloatingPointLiteralExpr) {
	value, _ := strconv.ParseFloat(expr.Text, 64)
	switch value {
	case 0:
		codeBlock.Float64LoadZero()
	case 1:
		codeBlock.Float64LoadOne()
	default:
		codeBlock.Float64LoadFloat64(value)
	}
}

//=====================================================================================================================

func buildGreaterThanCodeBlock(codeBlock *bytecode.CodeBlock, expr *typechecking.TypedGreaterThanExpr) {
	buildCodeBlock(codeBlock, expr.Lhs)
	buildCodeBlock(codeBlock, expr.Rhs)
	switch expr.TypeInfo.(type) {
	case *typechecking.Float64Type:
		codeBlock.Float64GreaterThan()
	case *typechecking.Int64Type:
		codeBlock.Int64GreaterThan()
	default:
		panic("Undefined greater than type")
	}
}

//=====================================================================================================================

func buildGreaterThanOrEqualsCodeBlock(codeBlock *bytecode.CodeBlock, expr *typechecking.TypedGreaterThanOrEqualsExpr) {
	buildCodeBlock(codeBlock, expr.Lhs)
	buildCodeBlock(codeBlock, expr.Rhs)
	switch expr.TypeInfo.(type) {
	case *typechecking.Float64Type:
		codeBlock.Float64GreaterThanOrEquals()
	case *typechecking.Int64Type:
		codeBlock.Int64GreaterThanOrEquals()
	default:
		panic("Undefined greater than or equals type")
	}
}

//=====================================================================================================================

func buildIntegerLiteralCodeBlock(codeBlock *bytecode.CodeBlock, expr *typechecking.TypedIntegerLiteralExpr) {
	value, _ := strconv.Atoi(expr.Text)
	switch value {
	case 0:
		codeBlock.Int64LoadZero()
	case 1:
		codeBlock.Int64LoadOne()
	default:
		codeBlock.Int64LoadInt16(int16(value))
	}
}

//=====================================================================================================================

func buildLessThanCodeBlock(codeBlock *bytecode.CodeBlock, expr *typechecking.TypedLessThanExpr) {
	buildCodeBlock(codeBlock, expr.Lhs)
	buildCodeBlock(codeBlock, expr.Rhs)
	switch expr.TypeInfo.(type) {
	case *typechecking.Float64Type:
		codeBlock.Float64LessThan()
	case *typechecking.Int64Type:
		codeBlock.Int64LessThan()
	default:
		panic("Undefined less than type")
	}
}

//=====================================================================================================================

func buildLessThanOrEqualsCodeBlock(codeBlock *bytecode.CodeBlock, expr *typechecking.TypedLessThanOrEqualsExpr) {
	buildCodeBlock(codeBlock, expr.Lhs)
	buildCodeBlock(codeBlock, expr.Rhs)
	switch expr.TypeInfo.(type) {
	case *typechecking.Float64Type:
		codeBlock.Float64LessThanOrEquals()
	case *typechecking.Int64Type:
		codeBlock.Int64LessThanOrEquals()
	default:
		panic("Undefined less than or equals type")
	}
}

//=====================================================================================================================

func buildLogicalAndCodeBlock(codeBlock *bytecode.CodeBlock, expr *typechecking.TypedLogicalAndExpr) {
	buildCodeBlock(codeBlock, expr.Lhs)
	buildCodeBlock(codeBlock, expr.Rhs)
	codeBlock.BoolAnd()
}

//=====================================================================================================================

func buildLogicalNotCodeBlock(codeBlock *bytecode.CodeBlock, expr *typechecking.TypedLogicalNotOperationExpr) {
	buildCodeBlock(codeBlock, expr.Operand)
	codeBlock.BoolNot()
}

//=====================================================================================================================

func buildLogicalOrCodeBlock(codeBlock *bytecode.CodeBlock, expr *typechecking.TypedLogicalOrExpr) {
	buildCodeBlock(codeBlock, expr.Lhs)
	buildCodeBlock(codeBlock, expr.Rhs)
	codeBlock.BoolOr()
}

//=====================================================================================================================

func buildMultiplicationCodeBlock(codeBlock *bytecode.CodeBlock, expr *typechecking.TypedMultiplicationExpr) {
	buildCodeBlock(codeBlock, expr.Lhs)
	buildCodeBlock(codeBlock, expr.Rhs)
	switch expr.TypeInfo.(type) {
	case *typechecking.Float64Type:
		codeBlock.Float64Multiply()
	case *typechecking.Int64Type:
		codeBlock.Int64Multiply()
	default:
		panic("Undefined multiplication type")
	}
}

//=====================================================================================================================

func buildNegationCodeBlock(codeBlock *bytecode.CodeBlock, expr *typechecking.TypedNegationOperationExpr) {
	buildCodeBlock(codeBlock, expr.Operand)
	switch expr.TypeInfo.(type) {
	case *typechecking.Float64Type:
		codeBlock.Float64Negate()
	case *typechecking.Int64Type:
		codeBlock.Int64Negate()
	default:
		panic("Undefined negation type")
	}
}

//=====================================================================================================================

func buildParenthesizedCodeBlock(codeBlock *bytecode.CodeBlock, expr *typechecking.TypedParenthesizedExpr) {
	if len(expr.Items) == 1 {
		buildCodeBlock(codeBlock, expr.Items[0])
	} else {
		panic("Records not yet handled")
	}
}

//=====================================================================================================================

func buildSubtractionCodeBlock(codeBlock *bytecode.CodeBlock, expr *typechecking.TypedSubtractionExpr) {
	buildCodeBlock(codeBlock, expr.Lhs)

	if e, ok := expr.Rhs.(*typechecking.TypedIntegerLiteralExpr); ok && e.Text == "1" {
		codeBlock.Int64Decrement()
	} else {
		buildCodeBlock(codeBlock, expr.Rhs)
		switch expr.TypeInfo.(type) {
		case *typechecking.Float64Type:
			codeBlock.Float64Subtract()
		case *typechecking.Int64Type:
			codeBlock.Int64Subtract()
		default:
			panic("Undefined subtraction type")
		}
	}
}

//=====================================================================================================================
