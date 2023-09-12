//
// # Data types related to Lligne token scanning.
//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package typechecking

import (
	"fmt"
	prior "lligne-cli/internal/lligne/code/analysis/structuring"
	"lligne-cli/internal/lligne/runtime/pools"
	"lligne-cli/internal/lligne/runtime/types"
)

//=====================================================================================================================

type Outcome struct {
	SourceCode      string
	NewLineOffsets  []uint32
	Model           IExpression
	StringConstants *pools.StringConstantPool
	IdentifierNames *pools.StringConstantPool
	TypeConstants   *types.TypeConstantPool
}

//=====================================================================================================================

func CheckTypes(priorOutcome *prior.Outcome) *Outcome {
	typePool := types.NewTypePool()
	model := checkTypes(priorOutcome.SourceCode, priorOutcome.Model, typePool, make([]uint64, 0))

	return &Outcome{
		SourceCode:      priorOutcome.SourceCode,
		NewLineOffsets:  priorOutcome.NewLineOffsets,
		Model:           model,
		StringConstants: priorOutcome.StringConstants,
		IdentifierNames: priorOutcome.IdentifierNames,
		TypeConstants:   typePool.Freeze(),
	}
}

//=====================================================================================================================

func checkTypes(
	sourceCode string,
	expression prior.IExpression,
	typePool *types.TypePool,
	idContexts []uint64,
) IExpression {

	switch expr := expression.(type) {

	case *prior.AdditionExpr:
		return typeCheckAdditionExpr(sourceCode, expr, typePool, idContexts)
	case *prior.BooleanLiteralExpr:
		return typeCheckBooleanLiteralExpr(expr)
	case *prior.BuiltInTypeExpr:
		return typeCheckBuiltInTypeExpr(sourceCode, expr, typePool)
	case *prior.DivisionExpr:
		return typeCheckDivisionExpr(sourceCode, expr, typePool, idContexts)
	case *prior.EqualsExpr:
		return typeCheckEqualsExpr(sourceCode, expr, typePool, idContexts)
	case *prior.FieldReferenceExpr:
		return typeCheckFieldReferenceExpr(sourceCode, expr, typePool, idContexts)
	case *prior.Float64LiteralExpr:
		return typeCheckFloat64LiteralExpr(expr)
	case *prior.GreaterThanExpr:
		return typeCheckGreaterThanExpr(sourceCode, expr, typePool, idContexts)
	case *prior.GreaterThanOrEqualsExpr:
		return typeCheckGreaterThanOrEqualsExpr(sourceCode, expr, typePool, idContexts)
	case *prior.IdentifierExpr:
		return typeCheckIdentifierExpr(expr, typePool, idContexts)
	case *prior.Int64LiteralExpr:
		return typeCheckInt64LiteralExpr(expr)
	case *prior.IsExpr:
		return typeCheckIsExpr(sourceCode, expr, typePool, idContexts)
	case *prior.LessThanExpr:
		return typeCheckLessThanExpr(sourceCode, expr, typePool, idContexts)
	case *prior.LessThanOrEqualsExpr:
		return typeCheckLessThanOrEqualsExpr(sourceCode, expr, typePool, idContexts)
	case *prior.LogicalAndExpr:
		return typeCheckLogicalAndExpr(sourceCode, expr, typePool, idContexts)
	case *prior.LogicalNotOperationExpr:
		return typeCheckLogicalNotOperationExpr(sourceCode, expr, typePool, idContexts)
	case *prior.LogicalOrExpr:
		return typeCheckLogicalOrExpr(sourceCode, expr, typePool, idContexts)
	case *prior.MultiplicationExpr:
		return typeCheckMultiplicationExpr(sourceCode, expr, typePool, idContexts)
	case *prior.NegationOperationExpr:
		return typeCheckNegationOperationExpr(sourceCode, expr, typePool, idContexts)
	case *prior.NotEqualsExpr:
		return typeCheckNotEqualsExpr(sourceCode, expr, typePool, idContexts)
	case *prior.ParenthesizedExpr:
		return typeCheckParenthesizedExpr(sourceCode, expr, typePool, idContexts)
	case *prior.RecordExpr:
		return typeCheckRecordExpr(sourceCode, expr, typePool, idContexts)
	case *prior.StringLiteralExpr:
		return typeCheckStringLiteralExpr(expr)
	case *prior.SubtractionExpr:
		return typeCheckSubtractionExpr(sourceCode, expr, typePool, idContexts)

	default:
		panic(fmt.Sprintf("Missing case in checkTypes: %T\n", expression))

	}

}

//=====================================================================================================================

func typeCheckAdditionExpr(sourceCode string, expr *prior.AdditionExpr, typePool *types.TypePool, idContexts []uint64) IExpression {
	lhs := checkTypes(sourceCode, expr.Lhs, typePool, idContexts)
	rhs := checkTypes(sourceCode, expr.Rhs, typePool, idContexts)
	switch lhs.GetTypeIndex() {
	case types.BuiltInTypeIndexFloat64, types.BuiltInTypeIndexInt64:
		// TODO: ensure they're the same
		return &AdditionExpr{
			SourcePosition: expr.SourcePosition,
			Lhs:            lhs,
			Rhs:            rhs,
			TypeIndex:      lhs.GetTypeIndex(),
		}
	case types.BuiltInTypeIndexString:
		// TODO: ensure both strings
		return &StringConcatenationExpr{
			SourcePosition: expr.SourcePosition,
			Lhs:            lhs,
			Rhs:            rhs,
		}
	default:
		panic(fmt.Sprintf("Missing case in typeCheckAdditionExpr: %d\n", lhs.GetTypeIndex()))
	}
}

//=====================================================================================================================

func typeCheckBooleanLiteralExpr(expr *prior.BooleanLiteralExpr) IExpression {
	return &BooleanLiteralExpr{
		SourcePosition: expr.SourcePosition,
		Value:          expr.Value,
	}
}

//=====================================================================================================================

func typeCheckBuiltInTypeExpr(sourceCode string, expr *prior.BuiltInTypeExpr, typePool *types.TypePool) IExpression {
	name := expr.SourcePosition.GetText(sourceCode)
	return &BuiltInTypeExpr{
		SourcePosition: expr.SourcePosition,
		ValueIndex:     typePool.GetIndexByName(name),
	}
}

//=====================================================================================================================

func typeCheckDivisionExpr(sourceCode string, expr *prior.DivisionExpr, typePool *types.TypePool, idContexts []uint64) IExpression {
	lhs := checkTypes(sourceCode, expr.Lhs, typePool, idContexts)
	rhs := checkTypes(sourceCode, expr.Rhs, typePool, idContexts)
	// TODO: ensure they're the same
	return &DivisionExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
		TypeIndex:      lhs.GetTypeIndex(),
	}
}

//=====================================================================================================================

func typeCheckEqualsExpr(sourceCode string, expr *prior.EqualsExpr, typePool *types.TypePool, idContexts []uint64) IExpression {
	lhs := checkTypes(sourceCode, expr.Lhs, typePool, idContexts)
	rhs := checkTypes(sourceCode, expr.Rhs, typePool, idContexts)
	// TODO: ensure they're the same
	return &EqualsExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func typeCheckFieldReferenceExpr(sourceCode string, expr *prior.FieldReferenceExpr, typePool *types.TypePool, idContexts []uint64) IExpression {
	parent := checkTypes(sourceCode, expr.Parent, typePool, idContexts)
	parentTypeIndex := parent.GetTypeIndex()
	child := checkTypes(sourceCode, expr.Child, typePool, append(idContexts, parentTypeIndex))
	// TODO: ensure the parent is a record
	return &FieldReferenceExpr{
		SourcePosition: expr.SourcePosition,
		Parent:         parent,
		Child:          child,
	}
}

//=====================================================================================================================

func typeCheckFloat64LiteralExpr(expr *prior.Float64LiteralExpr) IExpression {
	return &Float64LiteralExpr{
		SourcePosition: expr.SourcePosition,
		Value:          expr.Value,
	}
}

//=====================================================================================================================

func typeCheckGreaterThanExpr(sourceCode string, expr *prior.GreaterThanExpr, typePool *types.TypePool, idContexts []uint64) IExpression {
	lhs := checkTypes(sourceCode, expr.Lhs, typePool, idContexts)
	rhs := checkTypes(sourceCode, expr.Rhs, typePool, idContexts)
	// TODO: ensure they're the same
	return &GreaterThanExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func typeCheckGreaterThanOrEqualsExpr(sourceCode string, expr *prior.GreaterThanOrEqualsExpr, typePool *types.TypePool, idContexts []uint64) IExpression {
	lhs := checkTypes(sourceCode, expr.Lhs, typePool, idContexts)
	rhs := checkTypes(sourceCode, expr.Rhs, typePool, idContexts)
	// TODO: ensure they're the same
	return &GreaterThanOrEqualsExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func typeCheckIdentifierExpr(expr *prior.IdentifierExpr, typePool *types.TypePool, idContexts []uint64) IExpression {

	fieldIndex := uint64(0xFFFFFFFF)
	typeIndex := uint64(0xFFFFFFFF)
outer:
	for i := len(idContexts) - 1; i >= 0; i-- {
		recordType := typePool.Get(idContexts[i]).(*types.RecordType)

		for j, fieldNameIndex := range recordType.FieldNameIndexes {
			if expr.NameIndex == fieldNameIndex {
				fieldIndex = uint64(j)
				typeIndex = recordType.FieldTypeIndexes[j]
				break outer
			}
		}
	}

	if typeIndex == 0xFFFFFFFF {
		panic("Identifier type not found")
	}

	return &IdentifierExpr{
		SourcePosition: expr.SourcePosition,
		NameIndex:      expr.NameIndex,
		FieldIndex:     fieldIndex,
		TypeIndex:      typeIndex,
	}
}

//=====================================================================================================================

func typeCheckInt64LiteralExpr(expr *prior.Int64LiteralExpr) IExpression {
	return &Int64LiteralExpr{
		SourcePosition: expr.SourcePosition,
		Value:          expr.Value,
	}
}

//=====================================================================================================================

func typeCheckIsExpr(sourceCode string, expr *prior.IsExpr, typePool *types.TypePool, idContexts []uint64) IExpression {
	lhs := checkTypes(sourceCode, expr.Lhs, typePool, idContexts)
	rhs := checkTypes(sourceCode, expr.Rhs, typePool, idContexts)
	// TODO: ensure the lhs and rhs are compatible
	return &IsExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func typeCheckLessThanExpr(sourceCode string, expr *prior.LessThanExpr, typePool *types.TypePool, idContexts []uint64) IExpression {
	lhs := checkTypes(sourceCode, expr.Lhs, typePool, idContexts)
	rhs := checkTypes(sourceCode, expr.Rhs, typePool, idContexts)
	// TODO: ensure they're the same
	return &LessThanExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func typeCheckLessThanOrEqualsExpr(sourceCode string, expr *prior.LessThanOrEqualsExpr, typePool *types.TypePool, idContexts []uint64) IExpression {
	lhs := checkTypes(sourceCode, expr.Lhs, typePool, idContexts)
	rhs := checkTypes(sourceCode, expr.Rhs, typePool, idContexts)
	// TODO: ensure they're the same
	return &LessThanOrEqualsExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func typeCheckLogicalAndExpr(sourceCode string, expr *prior.LogicalAndExpr, typePool *types.TypePool, idContexts []uint64) IExpression {
	lhs := checkTypes(sourceCode, expr.Lhs, typePool, idContexts)
	rhs := checkTypes(sourceCode, expr.Rhs, typePool, idContexts)
	// TODO: ensure they're both boolean
	return &LogicalAndExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func typeCheckLogicalNotOperationExpr(sourceCode string, expr *prior.LogicalNotOperationExpr, typePool *types.TypePool, idContexts []uint64) IExpression {
	operand := checkTypes(sourceCode, expr.Operand, typePool, idContexts)

	// TODO: validate that operands are boolean

	return &LogicalNotOperationExpr{
		SourcePosition: expr.SourcePosition,
		Operand:        operand,
	}
}

//=====================================================================================================================

func typeCheckLogicalOrExpr(sourceCode string, expr *prior.LogicalOrExpr, typePool *types.TypePool, idContexts []uint64) IExpression {
	lhs := checkTypes(sourceCode, expr.Lhs, typePool, idContexts)
	rhs := checkTypes(sourceCode, expr.Rhs, typePool, idContexts)
	// TODO: ensure they're both boolean
	return &LogicalOrExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func typeCheckMultiplicationExpr(sourceCode string, expr *prior.MultiplicationExpr, typePool *types.TypePool, idContexts []uint64) IExpression {
	lhs := checkTypes(sourceCode, expr.Lhs, typePool, idContexts)
	rhs := checkTypes(sourceCode, expr.Rhs, typePool, idContexts)
	// TODO: ensure they're the same
	return &MultiplicationExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
		TypeIndex:      lhs.GetTypeIndex(),
	}
}

//=====================================================================================================================

func typeCheckNegationOperationExpr(sourceCode string, expr *prior.NegationOperationExpr, typePool *types.TypePool, idContexts []uint64) IExpression {
	operand := checkTypes(sourceCode, expr.Operand, typePool, idContexts)

	return &NegationOperationExpr{
		SourcePosition: expr.SourcePosition,
		Operand:        operand,
		TypeIndex:      operand.GetTypeIndex(),
	}
}

//=====================================================================================================================

func typeCheckNotEqualsExpr(sourceCode string, expr *prior.NotEqualsExpr, typePool *types.TypePool, idContexts []uint64) IExpression {
	lhs := checkTypes(sourceCode, expr.Lhs, typePool, idContexts)
	rhs := checkTypes(sourceCode, expr.Rhs, typePool, idContexts)
	// TODO: ensure they're the same
	return &NotEqualsExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//=====================================================================================================================

func typeCheckParenthesizedExpr(sourceCode string, expr *prior.ParenthesizedExpr, typePool *types.TypePool, idContexts []uint64) IExpression {

	inner := checkTypes(sourceCode, expr.InnerExpr, typePool, idContexts)

	return &ParenthesizedExpr{
		SourcePosition: expr.SourcePosition,
		InnerExpr:      inner,
		TypeIndex:      inner.GetTypeIndex(),
	}

}

//=====================================================================================================================

func typeCheckRecordExpr(sourceCode string, expr *prior.RecordExpr, typePool *types.TypePool, idContexts []uint64) IExpression {

	fields := make([]*RecordFieldExpr, 0)
	for _, field := range expr.Fields {
		fields = append(fields, typeCheckRecordFieldExpr(sourceCode, field, typePool, idContexts))
	}

	// TODO: make sure fields are in the same order as the record type

	fieldNameIndexes := make([]uint64, 0)
	fieldTypeIndexes := make([]uint64, 0)

	for _, field := range fields {
		fieldTypeIndex := field.FieldValue.GetTypeIndex()
		fieldNameIndexes = append(fieldNameIndexes, field.FieldNameIndex)
		fieldTypeIndexes = append(fieldTypeIndexes, fieldTypeIndex)
	}

	recordType := &types.RecordType{
		FieldNameIndexes: fieldNameIndexes,
		FieldTypeIndexes: fieldTypeIndexes,
	}
	typeIndex := typePool.Put(recordType)

	return &RecordExpr{
		SourcePosition: expr.SourcePosition,
		Fields:         fields,
		TypeIndex:      typeIndex,
	}

}

//=====================================================================================================================

func typeCheckRecordFieldExpr(sourceCode string, expr *prior.RecordFieldExpr, typePool *types.TypePool, idContexts []uint64) *RecordFieldExpr {
	return &RecordFieldExpr{
		SourcePosition: expr.SourcePosition,
		FieldNameIndex: expr.FieldNameIndex,
		FieldValue:     checkTypes(sourceCode, expr.FieldValue, typePool, idContexts),
	}
}

//=====================================================================================================================

func typeCheckStringLiteralExpr(expr *prior.StringLiteralExpr) IExpression {
	return &StringLiteralExpr{
		SourcePosition: expr.SourcePosition,
		ValueIndex:     expr.ValueIndex,
	}
}

//=====================================================================================================================

func typeCheckSubtractionExpr(sourceCode string, expr *prior.SubtractionExpr, typePool *types.TypePool, idContexts []uint64) IExpression {
	lhs := checkTypes(sourceCode, expr.Lhs, typePool, idContexts)
	rhs := checkTypes(sourceCode, expr.Rhs, typePool, idContexts)
	// TODO: ensure they're the same
	return &SubtractionExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
		TypeIndex:      lhs.GetTypeIndex(),
	}
}

//=====================================================================================================================
