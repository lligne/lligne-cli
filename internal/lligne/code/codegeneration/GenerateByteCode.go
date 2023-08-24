//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package codegeneration

import (
	"fmt"
	"lligne-cli/internal/lligne/code/typechecking"
	"lligne-cli/internal/lligne/runtime/bytecode"
	"lligne-cli/internal/lligne/runtime/types"
)

//=====================================================================================================================

func GenerateByteCode(expression typechecking.ITypedExpression) *bytecode.CodeBlock {
	result := bytecode.NewCodeBlock()

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
	case *typechecking.TypedFloat64LiteralExpr:
		buildFloat64LiteralCodeBlock(codeBlock, expr)
	case *typechecking.TypedGreaterThanExpr:
		buildGreaterThanCodeBlock(codeBlock, expr)
	case *typechecking.TypedGreaterThanOrEqualsExpr:
		buildGreaterThanOrEqualsCodeBlock(codeBlock, expr)
	case *typechecking.TypedInt64LiteralExpr:
		buildInt64LiteralCodeBlock(codeBlock, expr)
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
	case *typechecking.TypedNotEqualsExpr:
		buildNotEqualsCodeBlock(codeBlock, expr)
	case *typechecking.TypedParenthesizedExpr:
		buildParenthesizedCodeBlock(codeBlock, expr)
	case *typechecking.TypedStringConcatenationExpr:
		buildStringConcatenationCodeBlock(codeBlock, expr)
	case *typechecking.TypedStringLiteralExpr:
		buildStringLiteralCodeBlock(codeBlock, expr)
	case *typechecking.TypedSubtractionExpr:
		buildSubtractionCodeBlock(codeBlock, expr)
	default:
		panic(fmt.Sprintf("Missing case in buildCodeBlock: %T\n", expression))

	}

}

//=====================================================================================================================

func buildAdditionCodeBlock(codeBlock *bytecode.CodeBlock, expr *typechecking.TypedAdditionExpr) {

	if e, ok := expr.Lhs.(*typechecking.TypedInt64LiteralExpr); ok && e.Value == 1 {
		buildCodeBlock(codeBlock, expr.Rhs)
		codeBlock.Int64Increment()
	} else if e, ok := expr.Rhs.(*typechecking.TypedInt64LiteralExpr); ok && e.Value == 1 {
		buildCodeBlock(codeBlock, expr.Lhs)
		codeBlock.Int64Increment()
	} else {
		buildCodeBlock(codeBlock, expr.Lhs)
		buildCodeBlock(codeBlock, expr.Rhs)
		switch expr.TypeInfo.(type) {
		case *types.Float64Type:
			codeBlock.Float64Add()
		case *types.Int64Type:
			codeBlock.Int64Add()
		default:
			panic(fmt.Sprintf("Missing case in buildAdditionCodeBlock: %T\n", expr.TypeInfo))
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
	case *types.Float64Type:
		codeBlock.Float64Divide()
	case *types.Int64Type:
		codeBlock.Int64Divide()
	default:
		panic("Undefined division type")
	}
}

//=====================================================================================================================

func buildEqualsCodeBlock(codeBlock *bytecode.CodeBlock, expr *typechecking.TypedEqualsExpr) {
	buildCodeBlock(codeBlock, expr.Lhs)
	buildCodeBlock(codeBlock, expr.Rhs)
	switch expr.Lhs.GetTypeInfo().(type) {
	case *types.Float64Type:
		codeBlock.Float64Equals()
	case *types.Int64Type:
		codeBlock.Int64Equals()
	case *types.StringType:
		codeBlock.StringEquals()
	default:
		panic("Undefined equality type")
	}
}

//=====================================================================================================================

func buildFloat64LiteralCodeBlock(codeBlock *bytecode.CodeBlock, expr *typechecking.TypedFloat64LiteralExpr) {
	switch expr.Value {
	case 0:
		codeBlock.Float64LoadZero()
	case 1:
		codeBlock.Float64LoadOne()
	default:
		codeBlock.Float64Load(expr.Value)
	}
}

//=====================================================================================================================

func buildGreaterThanCodeBlock(codeBlock *bytecode.CodeBlock, expr *typechecking.TypedGreaterThanExpr) {
	buildCodeBlock(codeBlock, expr.Lhs)
	buildCodeBlock(codeBlock, expr.Rhs)
	switch expr.Lhs.GetTypeInfo().(type) {
	case *types.Float64Type:
		codeBlock.Float64GreaterThan()
	case *types.Int64Type:
		codeBlock.Int64GreaterThan()
	default:
		panic("Undefined greater than type")
	}
}

//=====================================================================================================================

func buildGreaterThanOrEqualsCodeBlock(codeBlock *bytecode.CodeBlock, expr *typechecking.TypedGreaterThanOrEqualsExpr) {
	buildCodeBlock(codeBlock, expr.Lhs)
	buildCodeBlock(codeBlock, expr.Rhs)
	switch expr.Lhs.GetTypeInfo().(type) {
	case *types.Float64Type:
		codeBlock.Float64GreaterThanOrEquals()
	case *types.Int64Type:
		codeBlock.Int64GreaterThanOrEquals()
	default:
		panic("Undefined greater than or equals type")
	}
}

//=====================================================================================================================

func buildInt64LiteralCodeBlock(codeBlock *bytecode.CodeBlock, expr *typechecking.TypedInt64LiteralExpr) {
	switch expr.Value {
	case 0:
		codeBlock.Int64LoadZero()
	case 1:
		codeBlock.Int64LoadOne()
	default:
		codeBlock.Int64Load(expr.Value)
	}
}

//=====================================================================================================================

func buildLessThanCodeBlock(codeBlock *bytecode.CodeBlock, expr *typechecking.TypedLessThanExpr) {
	buildCodeBlock(codeBlock, expr.Lhs)
	buildCodeBlock(codeBlock, expr.Rhs)
	switch expr.Lhs.GetTypeInfo().(type) {
	case *types.Float64Type:
		codeBlock.Float64LessThan()
	case *types.Int64Type:
		codeBlock.Int64LessThan()
	default:
		panic("Undefined less than type")
	}
}

//=====================================================================================================================

func buildLessThanOrEqualsCodeBlock(codeBlock *bytecode.CodeBlock, expr *typechecking.TypedLessThanOrEqualsExpr) {
	buildCodeBlock(codeBlock, expr.Lhs)
	buildCodeBlock(codeBlock, expr.Rhs)
	switch expr.Lhs.GetTypeInfo().(type) {
	case *types.Float64Type:
		codeBlock.Float64LessThanOrEquals()
	case *types.Int64Type:
		codeBlock.Int64LessThanOrEquals()
	default:
		panic(fmt.Sprintf("Missing case in buildLessThanOrEqualsCodeBlock: %T\n", expr.Lhs.GetTypeInfo()))
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
	case *types.Float64Type:
		codeBlock.Float64Multiply()
	case *types.Int64Type:
		codeBlock.Int64Multiply()
	default:
		panic(fmt.Sprintf("Missing case in buildMultiplicationCodeBlock: %T\n", expr.TypeInfo))
	}
}

//=====================================================================================================================

func buildNegationCodeBlock(codeBlock *bytecode.CodeBlock, expr *typechecking.TypedNegationOperationExpr) {
	buildCodeBlock(codeBlock, expr.Operand)
	switch expr.TypeInfo.(type) {
	case *types.Float64Type:
		codeBlock.Float64Negate()
	case *types.Int64Type:
		codeBlock.Int64Negate()
	default:
		panic(fmt.Sprintf("Missing case in buildNegationCodeBlock: %T\n", expr.TypeInfo))
	}
}

//=====================================================================================================================

func buildNotEqualsCodeBlock(codeBlock *bytecode.CodeBlock, expr *typechecking.TypedNotEqualsExpr) {
	buildCodeBlock(codeBlock, expr.Lhs)
	buildCodeBlock(codeBlock, expr.Rhs)
	switch expr.Lhs.GetTypeInfo().(type) {
	case *types.Float64Type:
		codeBlock.Float64NotEquals()
	case *types.Int64Type:
		codeBlock.Int64NotEquals()
	case *types.StringType:
		codeBlock.StringNotEquals()
	default:
		panic("Undefined inequality type")
	}
}

//=====================================================================================================================

func buildParenthesizedCodeBlock(codeBlock *bytecode.CodeBlock, expr *typechecking.TypedParenthesizedExpr) {
	buildCodeBlock(codeBlock, expr.InnerExpr)
}

//=====================================================================================================================

func buildStringConcatenationCodeBlock(codeBlock *bytecode.CodeBlock, expr *typechecking.TypedStringConcatenationExpr) {
	buildCodeBlock(codeBlock, expr.Lhs)
	buildCodeBlock(codeBlock, expr.Rhs)
	codeBlock.StringConcatenate()
}

//=====================================================================================================================

func buildStringLiteralCodeBlock(codeBlock *bytecode.CodeBlock, expr *typechecking.TypedStringLiteralExpr) {
	codeBlock.StringLoad(expr.Value)
}

//=====================================================================================================================

func buildSubtractionCodeBlock(codeBlock *bytecode.CodeBlock, expr *typechecking.TypedSubtractionExpr) {
	buildCodeBlock(codeBlock, expr.Lhs)

	if e, ok := expr.Rhs.(*typechecking.TypedInt64LiteralExpr); ok && e.Value == 1 {
		codeBlock.Int64Decrement()
	} else {
		buildCodeBlock(codeBlock, expr.Rhs)
		switch expr.TypeInfo.(type) {
		case *types.Float64Type:
			codeBlock.Float64Subtract()
		case *types.Int64Type:
			codeBlock.Int64Subtract()
		default:
			panic(fmt.Sprintf("Missing case in buildSubtractionCodeBlock: %T\n", expr.TypeInfo))
		}
	}
}

//=====================================================================================================================
