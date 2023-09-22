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
	IdentifierNames *pools.NameConstantPool
	TypeConstants   *types.TypeConstantPool
	CodeBlock       *bytecode.CodeBlock
}

//=====================================================================================================================

func GenerateByteCode(priorOutcome *prior.Outcome) *Outcome {
	generator := newGenerator(priorOutcome)
	generator.buildCodeBlock(priorOutcome.Model)

	generator.CodeBlock.Stop()

	return &Outcome{
		SourceCode:      priorOutcome.SourceCode,
		NewLineOffsets:  priorOutcome.NewLineOffsets,
		Model:           priorOutcome.Model,
		StringConstants: priorOutcome.StringConstants,
		IdentifierNames: priorOutcome.IdentifierNames,
		TypeConstants:   priorOutcome.TypeConstants,
		CodeBlock:       generator.CodeBlock,
	}
}

//=====================================================================================================================

type generator struct {
	SourceCode      string
	NewLineOffsets  []uint32
	StringConstants *pools.StringPool
	IdentifierNames *pools.NamePool
	TypeConstants   *types.TypeConstantPool
	CodeBlock       *bytecode.CodeBlock
}

//---------------------------------------------------------------------------------------------------------------------

func newGenerator(priorOutcome *prior.Outcome) *generator {
	return &generator{
		SourceCode:      priorOutcome.SourceCode,
		NewLineOffsets:  priorOutcome.NewLineOffsets,
		StringConstants: pools.NewStringPool(),
		IdentifierNames: pools.NewNamePool(),
		TypeConstants:   priorOutcome.TypeConstants,
		CodeBlock:       bytecode.NewCodeBlock(),
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (g *generator) buildCodeBlock(expression prior.IExpression) {

	switch expr := expression.(type) {

	case *prior.AdditionExpr:
		g.buildAdditionCodeBlock(expr)
	case *prior.BooleanLiteralExpr:
		g.buildBooleanLiteralCodeBlock(expr)
	case *prior.BuiltInTypeExpr:
		g.buildBuiltInTypeCodeBlock(expr)
	case *prior.DivisionExpr:
		g.buildDivisionCodeBlock(expr)
	case *prior.EqualsExpr:
		g.buildEqualsCodeBlock(expr)
	case *prior.FieldReferenceExpr:
		g.buildFieldReferenceCodeBlock(expr)
	case *prior.Float64LiteralExpr:
		g.buildFloat64LiteralCodeBlock(expr)
	case *prior.GreaterThanExpr:
		g.buildGreaterThanCodeBlock(expr)
	case *prior.GreaterThanOrEqualsExpr:
		g.buildGreaterThanOrEqualsCodeBlock(expr)
	case *prior.IdentifierExpr:
		g.buildIdentifierCodeBlock(expr)
	case *prior.Int64LiteralExpr:
		g.buildInt64LiteralCodeBlock(expr)
	case *prior.IsExpr:
		g.buildIsCodeBlock(expr)
	case *prior.LessThanExpr:
		g.buildLessThanCodeBlock(expr)
	case *prior.LessThanOrEqualsExpr:
		g.buildLessThanOrEqualsCodeBlock(expr)
	case *prior.LogicalAndExpr:
		g.buildLogicalAndCodeBlock(expr)
	case *prior.LogicalNotOperationExpr:
		g.buildLogicalNotCodeBlock(expr)
	case *prior.LogicalOrExpr:
		g.buildLogicalOrCodeBlock(expr)
	case *prior.MultiplicationExpr:
		g.buildMultiplicationCodeBlock(expr)
	case *prior.NegationOperationExpr:
		g.buildNegationCodeBlock(expr)
	case *prior.NotEqualsExpr:
		g.buildNotEqualsCodeBlock(expr)
	case *prior.ParenthesizedExpr:
		g.buildParenthesizedCodeBlock(expr)
	case *prior.RecordExpr:
		g.buildRecordCodeBlock(expr)
	case *prior.StringConcatenationExpr:
		g.buildStringConcatenationCodeBlock(expr)
	case *prior.StringLiteralExpr:
		g.buildStringLiteralCodeBlock(expr)
	case *prior.SubtractionExpr:
		g.buildSubtractionCodeBlock(expr)
	case *prior.WhereExpr:
		g.buildWhereCodeBlock(expr)
	default:
		panic(fmt.Sprintf("Missing case in buildCodeBlock: %T\n", expression))

	}

}

//---------------------------------------------------------------------------------------------------------------------

func (g *generator) buildAdditionCodeBlock(expr *prior.AdditionExpr) {

	if e, ok := expr.Lhs.(*prior.Int64LiteralExpr); ok && e.Value == 1 {
		g.buildCodeBlock(expr.Rhs)
		g.CodeBlock.Int64Increment()
	} else if e, ok := expr.Rhs.(*prior.Int64LiteralExpr); ok && e.Value == 1 {
		g.buildCodeBlock(expr.Lhs)
		g.CodeBlock.Int64Increment()
	} else {
		g.buildCodeBlock(expr.Lhs)
		g.buildCodeBlock(expr.Rhs)
		switch expr.TypeIndex {
		case types.BuiltInTypeIndexFloat64:
			g.CodeBlock.Float64Add()
		case types.BuiltInTypeIndexInt64:
			g.CodeBlock.Int64Add()
		default:
			panic(fmt.Sprintf("Missing case in buildAdditionCodeBlock: %d\n", expr.TypeIndex))
		}
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (g *generator) buildBooleanLiteralCodeBlock(expr *prior.BooleanLiteralExpr) {
	if expr.Value {
		g.CodeBlock.BoolLoadTrue()
	} else {
		g.CodeBlock.BoolLoadFalse()
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (g *generator) buildBuiltInTypeCodeBlock(expr *prior.BuiltInTypeExpr) {
	g.CodeBlock.TypeLoad(expr.ValueIndex)
}

//---------------------------------------------------------------------------------------------------------------------

func (g *generator) buildDivisionCodeBlock(expr *prior.DivisionExpr) {
	g.buildCodeBlock(expr.Lhs)
	g.buildCodeBlock(expr.Rhs)
	switch expr.TypeIndex {
	case types.BuiltInTypeIndexFloat64:
		g.CodeBlock.Float64Divide()
	case types.BuiltInTypeIndexInt64:
		g.CodeBlock.Int64Divide()
	default:
		panic("Undefined division type")
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (g *generator) buildEqualsCodeBlock(expr *prior.EqualsExpr) {
	g.buildCodeBlock(expr.Lhs)
	g.buildCodeBlock(expr.Rhs)
	switch expr.Lhs.GetTypeIndex() {
	case types.BuiltInTypeIndexFloat64:
		g.CodeBlock.Float64Equals()
	case types.BuiltInTypeIndexInt64:
		g.CodeBlock.Int64Equals()
	case types.BuiltInTypeIndexString:
		g.CodeBlock.StringEquals()
	case types.BuiltInTypeIndexType:
		g.CodeBlock.TypeEquals()
	default:
		typ := g.TypeConstants.Get(expr.Lhs.GetTypeIndex())

		switch typ.Category() {
		case types.TypeCategoryRecord:
			g.CodeBlock.RecordEquals()
		default:
			panic("Undefined equality type")
		}
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (g *generator) buildFieldReferenceCodeBlock(expr *prior.FieldReferenceExpr) {
	g.buildCodeBlock(expr.Parent)
	g.buildCodeBlock(expr.Child)
	g.CodeBlock.RecordFieldReference()
}

//---------------------------------------------------------------------------------------------------------------------

func (g *generator) buildFloat64LiteralCodeBlock(expr *prior.Float64LiteralExpr) {
	switch expr.Value {
	case 0:
		g.CodeBlock.Float64LoadZero()
	case 1:
		g.CodeBlock.Float64LoadOne()
	default:
		g.CodeBlock.Float64Load(expr.Value)
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (g *generator) buildGreaterThanCodeBlock(expr *prior.GreaterThanExpr) {
	g.buildCodeBlock(expr.Lhs)
	g.buildCodeBlock(expr.Rhs)
	switch expr.Lhs.GetTypeIndex() {
	case types.BuiltInTypeIndexFloat64:
		g.CodeBlock.Float64GreaterThan()
	case types.BuiltInTypeIndexInt64:
		g.CodeBlock.Int64GreaterThan()
	default:
		panic("Undefined greater than type")
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (g *generator) buildGreaterThanOrEqualsCodeBlock(expr *prior.GreaterThanOrEqualsExpr) {
	g.buildCodeBlock(expr.Lhs)
	g.buildCodeBlock(expr.Rhs)
	switch expr.Lhs.GetTypeIndex() {
	case types.BuiltInTypeIndexFloat64:
		g.CodeBlock.Float64GreaterThanOrEquals()
	case types.BuiltInTypeIndexInt64:
		g.CodeBlock.Int64GreaterThanOrEquals()
	default:
		panic("Undefined greater than or equals type")
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (g *generator) buildIdentifierCodeBlock(expr *prior.IdentifierExpr) {
	g.CodeBlock.RecordFieldIndexLoad(expr.FieldIndex)
}

//---------------------------------------------------------------------------------------------------------------------

func (g *generator) buildInt64LiteralCodeBlock(expr *prior.Int64LiteralExpr) {
	switch expr.Value {
	case 0:
		g.CodeBlock.Int64LoadZero()
	case 1:
		g.CodeBlock.Int64LoadOne()
	default:
		g.CodeBlock.Int64Load(expr.Value)
	}
}

//---------------------------------------------------------------------------------------------------------------------

// TODO: This will evolve into a module unto itself
func (g *generator) buildIsCodeBlock(expr *prior.IsExpr) {
	g.buildCodeBlock(expr.Lhs)
	g.buildCodeBlock(expr.Rhs)
	switch expr.Lhs.GetTypeIndex() {
	case types.BuiltInTypeIndexBool:
		switch rhs := expr.Rhs.(type) {
		case *prior.BuiltInTypeExpr:
			if rhs.ValueIndex == types.BuiltInTypeIndexBool {
				g.CodeBlock.BoolLoadTrue()
			} else {
				g.CodeBlock.BoolLoadFalse()
			}
		default:
			panic(fmt.Sprintf("Missing case in buildIsCodeBlock for BoolType: %T\n", expr.Rhs))
		}
	case types.BuiltInTypeIndexFloat64:
		switch rhs := expr.Rhs.(type) {
		case *prior.BuiltInTypeExpr:
			if rhs.ValueIndex == types.BuiltInTypeIndexFloat64 {
				g.CodeBlock.BoolLoadTrue()
			} else {
				g.CodeBlock.BoolLoadFalse()
			}
		default:
			panic(fmt.Sprintf("Missing case in buildIsCodeBlock for Float64Type: %T\n", expr.Rhs))
		}
	case types.BuiltInTypeIndexInt64:
		switch rhs := expr.Rhs.(type) {
		case *prior.BuiltInTypeExpr:
			if rhs.ValueIndex == types.BuiltInTypeIndexInt64 {
				g.CodeBlock.BoolLoadTrue()
			} else {
				g.CodeBlock.BoolLoadFalse()
			}
		default:
			panic(fmt.Sprintf("Missing case in buildIsCodeBlock for Int64Type: %T\n", expr.Rhs))
		}
	case types.BuiltInTypeIndexString:
		switch rhs := expr.Rhs.(type) {
		case *prior.BuiltInTypeExpr:
			if rhs.ValueIndex == types.BuiltInTypeIndexString {
				g.CodeBlock.BoolLoadTrue()
			} else {
				g.CodeBlock.BoolLoadFalse()
			}
		default:
			panic(fmt.Sprintf("Missing case in buildIsCodeBlock for StringType: %T\n", expr.Rhs))
		}
	default:
		panic(fmt.Sprintf("Missing case in buildIsCodeBlock: %d\n", expr.Lhs.GetTypeIndex()))
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (g *generator) buildLessThanCodeBlock(expr *prior.LessThanExpr) {
	g.buildCodeBlock(expr.Lhs)
	g.buildCodeBlock(expr.Rhs)
	switch expr.Lhs.GetTypeIndex() {
	case types.BuiltInTypeIndexFloat64:
		g.CodeBlock.Float64LessThan()
	case types.BuiltInTypeIndexInt64:
		g.CodeBlock.Int64LessThan()
	default:
		panic("Undefined less than type")
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (g *generator) buildLessThanOrEqualsCodeBlock(expr *prior.LessThanOrEqualsExpr) {
	g.buildCodeBlock(expr.Lhs)
	g.buildCodeBlock(expr.Rhs)
	switch expr.Lhs.GetTypeIndex() {
	case types.BuiltInTypeIndexFloat64:
		g.CodeBlock.Float64LessThanOrEquals()
	case types.BuiltInTypeIndexInt64:
		g.CodeBlock.Int64LessThanOrEquals()
	default:
		panic(fmt.Sprintf("Missing case in buildLessThanOrEqualsCodeBlock: %d\n", expr.Lhs.GetTypeIndex()))
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (g *generator) buildLogicalAndCodeBlock(expr *prior.LogicalAndExpr) {
	g.buildCodeBlock(expr.Lhs)
	g.buildCodeBlock(expr.Rhs)
	g.CodeBlock.BoolAnd()
}

//---------------------------------------------------------------------------------------------------------------------

func (g *generator) buildLogicalNotCodeBlock(expr *prior.LogicalNotOperationExpr) {
	g.buildCodeBlock(expr.Operand)
	g.CodeBlock.BoolNot()
}

//---------------------------------------------------------------------------------------------------------------------

func (g *generator) buildLogicalOrCodeBlock(expr *prior.LogicalOrExpr) {
	g.buildCodeBlock(expr.Lhs)
	g.buildCodeBlock(expr.Rhs)
	g.CodeBlock.BoolOr()
}

//---------------------------------------------------------------------------------------------------------------------

func (g *generator) buildMultiplicationCodeBlock(expr *prior.MultiplicationExpr) {
	g.buildCodeBlock(expr.Lhs)
	g.buildCodeBlock(expr.Rhs)
	switch expr.TypeIndex {
	case types.BuiltInTypeIndexFloat64:
		g.CodeBlock.Float64Multiply()
	case types.BuiltInTypeIndexInt64:
		g.CodeBlock.Int64Multiply()
	default:
		panic(fmt.Sprintf("Missing case in buildMultiplicationCodeBlock: %d\n", expr.TypeIndex))
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (g *generator) buildNegationCodeBlock(expr *prior.NegationOperationExpr) {
	g.buildCodeBlock(expr.Operand)
	switch expr.TypeIndex {
	case types.BuiltInTypeIndexFloat64:
		g.CodeBlock.Float64Negate()
	case types.BuiltInTypeIndexInt64:
		g.CodeBlock.Int64Negate()
	default:
		panic(fmt.Sprintf("Missing case in buildNegationCodeBlock: %d\n", expr.TypeIndex))
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (g *generator) buildNotEqualsCodeBlock(expr *prior.NotEqualsExpr) {
	g.buildCodeBlock(expr.Lhs)
	g.buildCodeBlock(expr.Rhs)
	switch expr.Lhs.GetTypeIndex() {
	case types.BuiltInTypeIndexFloat64:
		g.CodeBlock.Float64NotEquals()
	case types.BuiltInTypeIndexInt64:
		g.CodeBlock.Int64NotEquals()
	case types.BuiltInTypeIndexString:
		g.CodeBlock.StringNotEquals()
	case types.BuiltInTypeIndexType:
		g.CodeBlock.TypeNotEquals()
	default:
		typ := g.TypeConstants.Get(expr.Lhs.GetTypeIndex())

		switch typ.Category() {
		case types.TypeCategoryRecord:
			g.CodeBlock.RecordNotEquals()
		default:
			panic("Undefined equality type")
		}
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (g *generator) buildParenthesizedCodeBlock(expr *prior.ParenthesizedExpr) {
	g.buildCodeBlock(expr.InnerExpr)
}

//---------------------------------------------------------------------------------------------------------------------

func (g *generator) buildRecordCodeBlock(expr *prior.RecordExpr) {
	// Load the type index on the stack
	g.CodeBlock.TypeLoad(expr.TypeIndex)

	// Evaluate each field until fields are in order on the stack
	for _, field := range expr.Fields {
		g.buildCodeBlock(field.FieldValue)
	}

	// Copy from the stack into the record pool together with record type index (leave the record pool index on the stack).
	g.CodeBlock.RecordStore(len(expr.Fields))
}

//---------------------------------------------------------------------------------------------------------------------

func (g *generator) buildStringConcatenationCodeBlock(expr *prior.StringConcatenationExpr) {
	g.buildCodeBlock(expr.Lhs)
	g.buildCodeBlock(expr.Rhs)
	g.CodeBlock.StringConcatenate()
}

//---------------------------------------------------------------------------------------------------------------------

func (g *generator) buildStringLiteralCodeBlock(expr *prior.StringLiteralExpr) {
	g.CodeBlock.StringLoad(expr.ValueIndex)
}

//---------------------------------------------------------------------------------------------------------------------

func (g *generator) buildSubtractionCodeBlock(expr *prior.SubtractionExpr) {
	g.buildCodeBlock(expr.Lhs)

	if e, ok := expr.Rhs.(*prior.Int64LiteralExpr); ok && e.Value == 1 {
		g.CodeBlock.Int64Decrement()
	} else {
		g.buildCodeBlock(expr.Rhs)
		switch expr.TypeIndex {
		case types.BuiltInTypeIndexFloat64:
			g.CodeBlock.Float64Subtract()
		case types.BuiltInTypeIndexInt64:
			g.CodeBlock.Int64Subtract()
		default:
			panic(fmt.Sprintf("Missing case in buildSubtractionCodeBlock: %d\n", expr.TypeIndex))
		}
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (g *generator) buildWhereCodeBlock(expr *prior.WhereExpr) {
	g.buildCodeBlock(expr.Rhs)
	g.buildCodeBlock(expr.Lhs)
	g.CodeBlock.StackPopSecond()
}

//---------------------------------------------------------------------------------------------------------------------
