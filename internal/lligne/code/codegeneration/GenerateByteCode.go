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
	TypeConstants   *types.TypeConstantPool
	CodeBlock       *bytecode.CodeBlock
}

//=====================================================================================================================

func GenerateByteCode(priorOutcome *prior.Outcome) *Outcome {
	codeBlock := bytecode.NewCodeBlock()

	buildCodeBlock(codeBlock, priorOutcome.Model, priorOutcome.TypeConstants)

	codeBlock.Stop()

	return &Outcome{
		SourceCode:      priorOutcome.SourceCode,
		NewLineOffsets:  priorOutcome.NewLineOffsets,
		Model:           priorOutcome.Model,
		StringConstants: priorOutcome.StringConstants,
		IdentifierNames: priorOutcome.IdentifierNames,
		TypeConstants:   priorOutcome.TypeConstants,
		CodeBlock:       codeBlock,
	}
}

//=====================================================================================================================

func buildCodeBlock(codeBlock *bytecode.CodeBlock, expression prior.IExpression, typePool *types.TypeConstantPool) {

	switch expr := expression.(type) {

	case *prior.AdditionExpr:
		buildAdditionCodeBlock(codeBlock, expr, typePool)
	case *prior.BooleanLiteralExpr:
		buildBooleanLiteralCodeBlock(codeBlock, expr)
	case *prior.BuiltInTypeExpr:
		buildBuiltInTypeCodeBlock(codeBlock, expr)
	case *prior.DivisionExpr:
		buildDivisionCodeBlock(codeBlock, expr, typePool)
	case *prior.EqualsExpr:
		buildEqualsCodeBlock(codeBlock, expr, typePool)
	case *prior.Float64LiteralExpr:
		buildFloat64LiteralCodeBlock(codeBlock, expr)
	case *prior.GreaterThanExpr:
		buildGreaterThanCodeBlock(codeBlock, expr, typePool)
	case *prior.GreaterThanOrEqualsExpr:
		buildGreaterThanOrEqualsCodeBlock(codeBlock, expr, typePool)
	case *prior.Int64LiteralExpr:
		buildInt64LiteralCodeBlock(codeBlock, expr)
	case *prior.IsExpr:
		buildIsCodeBlock(codeBlock, expr, typePool)
	case *prior.LessThanExpr:
		buildLessThanCodeBlock(codeBlock, expr, typePool)
	case *prior.LessThanOrEqualsExpr:
		buildLessThanOrEqualsCodeBlock(codeBlock, expr, typePool)
	case *prior.LogicalAndExpr:
		buildLogicalAndCodeBlock(codeBlock, expr, typePool)
	case *prior.LogicalNotOperationExpr:
		buildLogicalNotCodeBlock(codeBlock, expr, typePool)
	case *prior.LogicalOrExpr:
		buildLogicalOrCodeBlock(codeBlock, expr, typePool)
	case *prior.MultiplicationExpr:
		buildMultiplicationCodeBlock(codeBlock, expr, typePool)
	case *prior.NegationOperationExpr:
		buildNegationCodeBlock(codeBlock, expr, typePool)
	case *prior.NotEqualsExpr:
		buildNotEqualsCodeBlock(codeBlock, expr, typePool)
	case *prior.ParenthesizedExpr:
		buildParenthesizedCodeBlock(codeBlock, expr, typePool)
	case *prior.StringConcatenationExpr:
		buildStringConcatenationCodeBlock(codeBlock, expr, typePool)
	case *prior.StringLiteralExpr:
		buildStringLiteralCodeBlock(codeBlock, expr)
	case *prior.SubtractionExpr:
		buildSubtractionCodeBlock(codeBlock, expr, typePool)
	default:
		panic(fmt.Sprintf("Missing case in buildCodeBlock: %T\n", expression))

	}

}

//=====================================================================================================================

func buildAdditionCodeBlock(codeBlock *bytecode.CodeBlock, expr *prior.AdditionExpr, typePool *types.TypeConstantPool) {

	if e, ok := expr.Lhs.(*prior.Int64LiteralExpr); ok && e.Value == 1 {
		buildCodeBlock(codeBlock, expr.Rhs, typePool)
		codeBlock.Int64Increment()
	} else if e, ok := expr.Rhs.(*prior.Int64LiteralExpr); ok && e.Value == 1 {
		buildCodeBlock(codeBlock, expr.Lhs, typePool)
		codeBlock.Int64Increment()
	} else {
		buildCodeBlock(codeBlock, expr.Lhs, typePool)
		buildCodeBlock(codeBlock, expr.Rhs, typePool)
		switch expr.TypeIndex {
		case types.BuiltInTypeIndexFloat64:
			codeBlock.Float64Add()
		case types.BuiltInTypeIndexInt64:
			codeBlock.Int64Add()
		default:
			panic(fmt.Sprintf("Missing case in buildAdditionCodeBlock: %d\n", expr.TypeIndex))
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
	codeBlock.TypeLoad(expr.ValueIndex)
}

//=====================================================================================================================

func buildDivisionCodeBlock(codeBlock *bytecode.CodeBlock, expr *prior.DivisionExpr, typePool *types.TypeConstantPool) {
	buildCodeBlock(codeBlock, expr.Lhs, typePool)
	buildCodeBlock(codeBlock, expr.Rhs, typePool)
	switch expr.TypeIndex {
	case types.BuiltInTypeIndexFloat64:
		codeBlock.Float64Divide()
	case types.BuiltInTypeIndexInt64:
		codeBlock.Int64Divide()
	default:
		panic("Undefined division type")
	}
}

//=====================================================================================================================

func buildEqualsCodeBlock(codeBlock *bytecode.CodeBlock, expr *prior.EqualsExpr, typePool *types.TypeConstantPool) {
	buildCodeBlock(codeBlock, expr.Lhs, typePool)
	buildCodeBlock(codeBlock, expr.Rhs, typePool)
	switch expr.Lhs.GetTypeIndex() {
	case types.BuiltInTypeIndexFloat64:
		codeBlock.Float64Equals()
	case types.BuiltInTypeIndexInt64:
		codeBlock.Int64Equals()
	case types.BuiltInTypeIndexString:
		codeBlock.StringEquals()
	case types.BuiltInTypeIndexType:
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

func buildGreaterThanCodeBlock(codeBlock *bytecode.CodeBlock, expr *prior.GreaterThanExpr, typePool *types.TypeConstantPool) {
	buildCodeBlock(codeBlock, expr.Lhs, typePool)
	buildCodeBlock(codeBlock, expr.Rhs, typePool)
	switch expr.Lhs.GetTypeIndex() {
	case types.BuiltInTypeIndexFloat64:
		codeBlock.Float64GreaterThan()
	case types.BuiltInTypeIndexInt64:
		codeBlock.Int64GreaterThan()
	default:
		panic("Undefined greater than type")
	}
}

//=====================================================================================================================

func buildGreaterThanOrEqualsCodeBlock(codeBlock *bytecode.CodeBlock, expr *prior.GreaterThanOrEqualsExpr, typePool *types.TypeConstantPool) {
	buildCodeBlock(codeBlock, expr.Lhs, typePool)
	buildCodeBlock(codeBlock, expr.Rhs, typePool)
	switch expr.Lhs.GetTypeIndex() {
	case types.BuiltInTypeIndexFloat64:
		codeBlock.Float64GreaterThanOrEquals()
	case types.BuiltInTypeIndexInt64:
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
func buildIsCodeBlock(codeBlock *bytecode.CodeBlock, expr *prior.IsExpr, typePool *types.TypeConstantPool) {
	buildCodeBlock(codeBlock, expr.Lhs, typePool)
	buildCodeBlock(codeBlock, expr.Rhs, typePool)
	switch expr.Lhs.GetTypeIndex() {
	case types.BuiltInTypeIndexBool:
		switch rhs := expr.Rhs.(type) {
		case *prior.BuiltInTypeExpr:
			if rhs.ValueIndex == types.BuiltInTypeIndexBool {
				codeBlock.BoolLoadTrue()
			} else {
				codeBlock.BoolLoadFalse()
			}
		default:
			panic(fmt.Sprintf("Missing case in buildIsCodeBlock for BoolType: %T\n", expr.Rhs))
		}
	case types.BuiltInTypeIndexFloat64:
		switch rhs := expr.Rhs.(type) {
		case *prior.BuiltInTypeExpr:
			if rhs.ValueIndex == types.BuiltInTypeIndexFloat64 {
				codeBlock.BoolLoadTrue()
			} else {
				codeBlock.BoolLoadFalse()
			}
		default:
			panic(fmt.Sprintf("Missing case in buildIsCodeBlock for Float64Type: %T\n", expr.Rhs))
		}
	case types.BuiltInTypeIndexInt64:
		switch rhs := expr.Rhs.(type) {
		case *prior.BuiltInTypeExpr:
			if rhs.ValueIndex == types.BuiltInTypeIndexInt64 {
				codeBlock.BoolLoadTrue()
			} else {
				codeBlock.BoolLoadFalse()
			}
		default:
			panic(fmt.Sprintf("Missing case in buildIsCodeBlock for Int64Type: %T\n", expr.Rhs))
		}
	case types.BuiltInTypeIndexString:
		switch rhs := expr.Rhs.(type) {
		case *prior.BuiltInTypeExpr:
			if rhs.ValueIndex == types.BuiltInTypeIndexString {
				codeBlock.BoolLoadTrue()
			} else {
				codeBlock.BoolLoadFalse()
			}
		default:
			panic(fmt.Sprintf("Missing case in buildIsCodeBlock for StringType: %T\n", expr.Rhs))
		}
	default:
		panic(fmt.Sprintf("Missing case in buildIsCodeBlock: %d\n", expr.Lhs.GetTypeIndex()))
	}
}

//=====================================================================================================================

func buildLessThanCodeBlock(codeBlock *bytecode.CodeBlock, expr *prior.LessThanExpr, typePool *types.TypeConstantPool) {
	buildCodeBlock(codeBlock, expr.Lhs, typePool)
	buildCodeBlock(codeBlock, expr.Rhs, typePool)
	switch expr.Lhs.GetTypeIndex() {
	case types.BuiltInTypeIndexFloat64:
		codeBlock.Float64LessThan()
	case types.BuiltInTypeIndexInt64:
		codeBlock.Int64LessThan()
	default:
		panic("Undefined less than type")
	}
}

//=====================================================================================================================

func buildLessThanOrEqualsCodeBlock(codeBlock *bytecode.CodeBlock, expr *prior.LessThanOrEqualsExpr, typePool *types.TypeConstantPool) {
	buildCodeBlock(codeBlock, expr.Lhs, typePool)
	buildCodeBlock(codeBlock, expr.Rhs, typePool)
	switch expr.Lhs.GetTypeIndex() {
	case types.BuiltInTypeIndexFloat64:
		codeBlock.Float64LessThanOrEquals()
	case types.BuiltInTypeIndexInt64:
		codeBlock.Int64LessThanOrEquals()
	default:
		panic(fmt.Sprintf("Missing case in buildLessThanOrEqualsCodeBlock: %d\n", expr.Lhs.GetTypeIndex()))
	}
}

//=====================================================================================================================

func buildLogicalAndCodeBlock(codeBlock *bytecode.CodeBlock, expr *prior.LogicalAndExpr, typePool *types.TypeConstantPool) {
	buildCodeBlock(codeBlock, expr.Lhs, typePool)
	buildCodeBlock(codeBlock, expr.Rhs, typePool)
	codeBlock.BoolAnd()
}

//=====================================================================================================================

func buildLogicalNotCodeBlock(codeBlock *bytecode.CodeBlock, expr *prior.LogicalNotOperationExpr, typePool *types.TypeConstantPool) {
	buildCodeBlock(codeBlock, expr.Operand, typePool)
	codeBlock.BoolNot()
}

//=====================================================================================================================

func buildLogicalOrCodeBlock(codeBlock *bytecode.CodeBlock, expr *prior.LogicalOrExpr, typePool *types.TypeConstantPool) {
	buildCodeBlock(codeBlock, expr.Lhs, typePool)
	buildCodeBlock(codeBlock, expr.Rhs, typePool)
	codeBlock.BoolOr()
}

//=====================================================================================================================

func buildMultiplicationCodeBlock(codeBlock *bytecode.CodeBlock, expr *prior.MultiplicationExpr, typePool *types.TypeConstantPool) {
	buildCodeBlock(codeBlock, expr.Lhs, typePool)
	buildCodeBlock(codeBlock, expr.Rhs, typePool)
	switch expr.TypeIndex {
	case types.BuiltInTypeIndexFloat64:
		codeBlock.Float64Multiply()
	case types.BuiltInTypeIndexInt64:
		codeBlock.Int64Multiply()
	default:
		panic(fmt.Sprintf("Missing case in buildMultiplicationCodeBlock: %d\n", expr.TypeIndex))
	}
}

//=====================================================================================================================

func buildNegationCodeBlock(codeBlock *bytecode.CodeBlock, expr *prior.NegationOperationExpr, typePool *types.TypeConstantPool) {
	buildCodeBlock(codeBlock, expr.Operand, typePool)
	switch expr.TypeIndex {
	case types.BuiltInTypeIndexFloat64:
		codeBlock.Float64Negate()
	case types.BuiltInTypeIndexInt64:
		codeBlock.Int64Negate()
	default:
		panic(fmt.Sprintf("Missing case in buildNegationCodeBlock: %d\n", expr.TypeIndex))
	}
}

//=====================================================================================================================

func buildNotEqualsCodeBlock(codeBlock *bytecode.CodeBlock, expr *prior.NotEqualsExpr, typePool *types.TypeConstantPool) {
	buildCodeBlock(codeBlock, expr.Lhs, typePool)
	buildCodeBlock(codeBlock, expr.Rhs, typePool)
	switch expr.Lhs.GetTypeIndex() {
	case types.BuiltInTypeIndexFloat64:
		codeBlock.Float64NotEquals()
	case types.BuiltInTypeIndexInt64:
		codeBlock.Int64NotEquals()
	case types.BuiltInTypeIndexString:
		codeBlock.StringNotEquals()
	case types.BuiltInTypeIndexType:
		codeBlock.TypeNotEquals()
	default:
		panic("Undefined inequality type")
	}
}

//=====================================================================================================================

func buildParenthesizedCodeBlock(codeBlock *bytecode.CodeBlock, expr *prior.ParenthesizedExpr, typePool *types.TypeConstantPool) {
	buildCodeBlock(codeBlock, expr.InnerExpr, typePool)
}

//=====================================================================================================================

func buildStringConcatenationCodeBlock(codeBlock *bytecode.CodeBlock, expr *prior.StringConcatenationExpr, typePool *types.TypeConstantPool) {
	buildCodeBlock(codeBlock, expr.Lhs, typePool)
	buildCodeBlock(codeBlock, expr.Rhs, typePool)
	codeBlock.StringConcatenate()
}

//=====================================================================================================================

func buildStringLiteralCodeBlock(codeBlock *bytecode.CodeBlock, expr *prior.StringLiteralExpr) {
	codeBlock.StringLoad(expr.ValueIndex)
}

//=====================================================================================================================

func buildSubtractionCodeBlock(codeBlock *bytecode.CodeBlock, expr *prior.SubtractionExpr, typePool *types.TypeConstantPool) {
	buildCodeBlock(codeBlock, expr.Lhs, typePool)

	if e, ok := expr.Rhs.(*prior.Int64LiteralExpr); ok && e.Value == 1 {
		codeBlock.Int64Decrement()
	} else {
		buildCodeBlock(codeBlock, expr.Rhs, typePool)
		switch expr.TypeIndex {
		case types.BuiltInTypeIndexFloat64:
			codeBlock.Float64Subtract()
		case types.BuiltInTypeIndexInt64:
			codeBlock.Int64Subtract()
		default:
			panic(fmt.Sprintf("Missing case in buildSubtractionCodeBlock: %d\n", expr.TypeIndex))
		}
	}
}

//=====================================================================================================================
