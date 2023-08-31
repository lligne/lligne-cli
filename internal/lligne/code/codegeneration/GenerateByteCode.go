//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package codegeneration

import (
	"fmt"
	prior "lligne-cli/internal/lligne/code/analysis/typechecking"
	"lligne-cli/internal/lligne/runtime/bytecode"
	"lligne-cli/internal/lligne/runtime/pools"
	"lligne-cli/internal/lligne/runtime/types"
)

//=====================================================================================================================

type Outcome struct {
	SourceCode      string
	NewLineOffsets  []uint32
	Model           prior.IExpression
	StringConstants *pools.StringConstantPool
	IdentifierNames *pools.StringConstantPool
	// TODO: TypesPool
	CodeBlock *bytecode.CodeBlock
}

//=====================================================================================================================

func GenerateByteCode(priorOutcome *prior.Outcome) *Outcome {
	codeBlock := bytecode.NewCodeBlock()

	buildCodeBlock(codeBlock, priorOutcome.Model)

	codeBlock.Stop()

	return &Outcome{
		SourceCode:      priorOutcome.SourceCode,
		NewLineOffsets:  priorOutcome.NewLineOffsets,
		Model:           priorOutcome.Model,
		StringConstants: priorOutcome.StringConstants,
		IdentifierNames: priorOutcome.IdentifierNames,
		CodeBlock:       codeBlock,
	}
}

//=====================================================================================================================

func buildCodeBlock(codeBlock *bytecode.CodeBlock, expression prior.IExpression) {

	switch expr := expression.(type) {

	case *prior.AdditionExpr:
		buildAdditionCodeBlock(codeBlock, expr)
	case *prior.BooleanLiteralExpr:
		buildBooleanLiteralCodeBlock(codeBlock, expr)
	case *prior.BuiltInTypeExpr:
		buildBuiltInTypeCodeBlock(codeBlock, expr)
	case *prior.DivisionExpr:
		buildDivisionCodeBlock(codeBlock, expr)
	case *prior.EqualsExpr:
		buildEqualsCodeBlock(codeBlock, expr)
	case *prior.Float64LiteralExpr:
		buildFloat64LiteralCodeBlock(codeBlock, expr)
	case *prior.GreaterThanExpr:
		buildGreaterThanCodeBlock(codeBlock, expr)
	case *prior.GreaterThanOrEqualsExpr:
		buildGreaterThanOrEqualsCodeBlock(codeBlock, expr)
	case *prior.Int64LiteralExpr:
		buildInt64LiteralCodeBlock(codeBlock, expr)
	case *prior.IsExpr:
		buildIsCodeBlock(codeBlock, expr)
	case *prior.LessThanExpr:
		buildLessThanCodeBlock(codeBlock, expr)
	case *prior.LessThanOrEqualsExpr:
		buildLessThanOrEqualsCodeBlock(codeBlock, expr)
	case *prior.LogicalAndExpr:
		buildLogicalAndCodeBlock(codeBlock, expr)
	case *prior.LogicalNotOperationExpr:
		buildLogicalNotCodeBlock(codeBlock, expr)
	case *prior.LogicalOrExpr:
		buildLogicalOrCodeBlock(codeBlock, expr)
	case *prior.MultiplicationExpr:
		buildMultiplicationCodeBlock(codeBlock, expr)
	case *prior.NegationOperationExpr:
		buildNegationCodeBlock(codeBlock, expr)
	case *prior.NotEqualsExpr:
		buildNotEqualsCodeBlock(codeBlock, expr)
	case *prior.ParenthesizedExpr:
		buildParenthesizedCodeBlock(codeBlock, expr)
	case *prior.StringConcatenationExpr:
		buildStringConcatenationCodeBlock(codeBlock, expr)
	case *prior.StringLiteralExpr:
		buildStringLiteralCodeBlock(codeBlock, expr)
	case *prior.SubtractionExpr:
		buildSubtractionCodeBlock(codeBlock, expr)
	default:
		panic(fmt.Sprintf("Missing case in buildCodeBlock: %T\n", expression))

	}

}

//=====================================================================================================================

func buildAdditionCodeBlock(codeBlock *bytecode.CodeBlock, expr *prior.AdditionExpr) {

	if e, ok := expr.Lhs.(*prior.Int64LiteralExpr); ok && e.Value == 1 {
		buildCodeBlock(codeBlock, expr.Rhs)
		codeBlock.Int64Increment()
	} else if e, ok := expr.Rhs.(*prior.Int64LiteralExpr); ok && e.Value == 1 {
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

func buildBooleanLiteralCodeBlock(codeBlock *bytecode.CodeBlock, expr *prior.BooleanLiteralExpr) {
	if expr.Value {
		codeBlock.BoolLoadTrue()
	} else {
		codeBlock.BoolLoadFalse()
	}
}

//=====================================================================================================================

func buildBuiltInTypeCodeBlock(codeBlock *bytecode.CodeBlock, expr *prior.BuiltInTypeExpr) {
	codeBlock.TypeLoad(expr.Value)
}

//=====================================================================================================================

func buildDivisionCodeBlock(codeBlock *bytecode.CodeBlock, expr *prior.DivisionExpr) {
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

func buildEqualsCodeBlock(codeBlock *bytecode.CodeBlock, expr *prior.EqualsExpr) {
	buildCodeBlock(codeBlock, expr.Lhs)
	buildCodeBlock(codeBlock, expr.Rhs)
	switch expr.Lhs.GetTypeInfo().(type) {
	case *types.Float64Type:
		codeBlock.Float64Equals()
	case *types.Int64Type:
		codeBlock.Int64Equals()
	case *types.StringType:
		codeBlock.StringEquals()
	case *types.TypeType:
		codeBlock.TypeEquals()
	default:
		panic("Undefined equality type")
	}
}

//=====================================================================================================================

func buildFloat64LiteralCodeBlock(codeBlock *bytecode.CodeBlock, expr *prior.Float64LiteralExpr) {
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

func buildGreaterThanCodeBlock(codeBlock *bytecode.CodeBlock, expr *prior.GreaterThanExpr) {
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

func buildGreaterThanOrEqualsCodeBlock(codeBlock *bytecode.CodeBlock, expr *prior.GreaterThanOrEqualsExpr) {
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

func buildInt64LiteralCodeBlock(codeBlock *bytecode.CodeBlock, expr *prior.Int64LiteralExpr) {
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

// TODO: This will evolve into a module unto itself
func buildIsCodeBlock(codeBlock *bytecode.CodeBlock, expr *prior.IsExpr) {
	buildCodeBlock(codeBlock, expr.Lhs)
	buildCodeBlock(codeBlock, expr.Rhs)
	switch expr.Lhs.GetTypeInfo().(type) {
	case *types.BoolType:
		switch rhs := expr.Rhs.(type) {
		case *prior.BuiltInTypeExpr:
			switch rhs.Value {
			case types.BuiltInTypeBool:
				codeBlock.BoolLoadTrue()
			default:
				codeBlock.BoolLoadFalse()
			}
		default:
			panic(fmt.Sprintf("Missing case in buildIsCodeBlock for BoolType: %T\n", expr.Rhs))
		}
	case *types.Float64Type:
		switch rhs := expr.Rhs.(type) {
		case *prior.BuiltInTypeExpr:
			switch rhs.Value {
			case types.BuiltInTypeFloat64:
				codeBlock.BoolLoadTrue()
			default:
				codeBlock.BoolLoadFalse()
			}
		default:
			panic(fmt.Sprintf("Missing case in buildIsCodeBlock for Float64Type: %T\n", expr.Rhs))
		}
	case *types.Int64Type:
		switch rhs := expr.Rhs.(type) {
		case *prior.BuiltInTypeExpr:
			switch rhs.Value {
			case types.BuiltInTypeInt64:
				codeBlock.BoolLoadTrue()
			default:
				codeBlock.BoolLoadFalse()
			}
		default:
			panic(fmt.Sprintf("Missing case in buildIsCodeBlock for Int64Type: %T\n", expr.Rhs))
		}
	case *types.StringType:
		switch rhs := expr.Rhs.(type) {
		case *prior.BuiltInTypeExpr:
			switch rhs.Value {
			case types.BuiltInTypeString:
				codeBlock.BoolLoadTrue()
			default:
				codeBlock.BoolLoadFalse()
			}
		default:
			panic(fmt.Sprintf("Missing case in buildIsCodeBlock for StringType: %T\n", expr.Rhs))
		}
	default:
		panic(fmt.Sprintf("Missing case in buildIsCodeBlock: %T\n", expr.Lhs.GetTypeInfo()))
	}
}

//=====================================================================================================================

func buildLessThanCodeBlock(codeBlock *bytecode.CodeBlock, expr *prior.LessThanExpr) {
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

func buildLessThanOrEqualsCodeBlock(codeBlock *bytecode.CodeBlock, expr *prior.LessThanOrEqualsExpr) {
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

func buildLogicalAndCodeBlock(codeBlock *bytecode.CodeBlock, expr *prior.LogicalAndExpr) {
	buildCodeBlock(codeBlock, expr.Lhs)
	buildCodeBlock(codeBlock, expr.Rhs)
	codeBlock.BoolAnd()
}

//=====================================================================================================================

func buildLogicalNotCodeBlock(codeBlock *bytecode.CodeBlock, expr *prior.LogicalNotOperationExpr) {
	buildCodeBlock(codeBlock, expr.Operand)
	codeBlock.BoolNot()
}

//=====================================================================================================================

func buildLogicalOrCodeBlock(codeBlock *bytecode.CodeBlock, expr *prior.LogicalOrExpr) {
	buildCodeBlock(codeBlock, expr.Lhs)
	buildCodeBlock(codeBlock, expr.Rhs)
	codeBlock.BoolOr()
}

//=====================================================================================================================

func buildMultiplicationCodeBlock(codeBlock *bytecode.CodeBlock, expr *prior.MultiplicationExpr) {
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

func buildNegationCodeBlock(codeBlock *bytecode.CodeBlock, expr *prior.NegationOperationExpr) {
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

func buildNotEqualsCodeBlock(codeBlock *bytecode.CodeBlock, expr *prior.NotEqualsExpr) {
	buildCodeBlock(codeBlock, expr.Lhs)
	buildCodeBlock(codeBlock, expr.Rhs)
	switch expr.Lhs.GetTypeInfo().(type) {
	case *types.Float64Type:
		codeBlock.Float64NotEquals()
	case *types.Int64Type:
		codeBlock.Int64NotEquals()
	case *types.StringType:
		codeBlock.StringNotEquals()
	case *types.TypeType:
		codeBlock.TypeNotEquals()
	default:
		panic("Undefined inequality type")
	}
}

//=====================================================================================================================

func buildParenthesizedCodeBlock(codeBlock *bytecode.CodeBlock, expr *prior.ParenthesizedExpr) {
	buildCodeBlock(codeBlock, expr.InnerExpr)
}

//=====================================================================================================================

func buildStringConcatenationCodeBlock(codeBlock *bytecode.CodeBlock, expr *prior.StringConcatenationExpr) {
	buildCodeBlock(codeBlock, expr.Lhs)
	buildCodeBlock(codeBlock, expr.Rhs)
	codeBlock.StringConcatenate()
}

//=====================================================================================================================

func buildStringLiteralCodeBlock(codeBlock *bytecode.CodeBlock, expr *prior.StringLiteralExpr) {
	codeBlock.StringLoad(expr.ValueIndex)
}

//=====================================================================================================================

func buildSubtractionCodeBlock(codeBlock *bytecode.CodeBlock, expr *prior.SubtractionExpr) {
	buildCodeBlock(codeBlock, expr.Lhs)

	if e, ok := expr.Rhs.(*prior.Int64LiteralExpr); ok && e.Value == 1 {
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
