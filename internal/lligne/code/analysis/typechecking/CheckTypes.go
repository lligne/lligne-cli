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
	checker := newTypeChecker(priorOutcome)
	model := checker.checkTypes(priorOutcome.Model, make([]uint64, 0))

	return &Outcome{
		SourceCode:      priorOutcome.SourceCode,
		NewLineOffsets:  priorOutcome.NewLineOffsets,
		Model:           model,
		StringConstants: priorOutcome.StringConstants,
		IdentifierNames: priorOutcome.IdentifierNames,
		TypeConstants:   checker.TypePool.Freeze(),
	}
}

//=====================================================================================================================

type typeChecker struct {
	SourceCode      string
	NewLineOffsets  []uint32
	StringConstants *pools.StringConstantPool
	IdentifierNames *pools.StringConstantPool
	TypePool        *types.TypePool
}

//---------------------------------------------------------------------------------------------------------------------

func newTypeChecker(priorOutcome *prior.Outcome) *typeChecker {
	return &typeChecker{
		SourceCode:      priorOutcome.SourceCode,
		NewLineOffsets:  priorOutcome.NewLineOffsets,
		StringConstants: priorOutcome.StringConstants,
		IdentifierNames: priorOutcome.IdentifierNames,
		TypePool:        types.NewTypePool(),
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (t *typeChecker) checkTypes(
	expression prior.IExpression,
	idContexts []uint64,
) IExpression {

	switch expr := expression.(type) {

	case *prior.AdditionExpr:
		return t.typeCheckAdditionExpr(expr, idContexts)
	case *prior.BooleanLiteralExpr:
		return t.typeCheckBooleanLiteralExpr(expr)
	case *prior.BuiltInTypeExpr:
		return t.typeCheckBuiltInTypeExpr(expr)
	case *prior.DivisionExpr:
		return t.typeCheckDivisionExpr(expr, idContexts)
	case *prior.EqualsExpr:
		return t.typeCheckEqualsExpr(expr, idContexts)
	case *prior.FieldReferenceExpr:
		return t.typeCheckFieldReferenceExpr(expr, idContexts)
	case *prior.Float64LiteralExpr:
		return t.typeCheckFloat64LiteralExpr(expr)
	case *prior.GreaterThanExpr:
		return t.typeCheckGreaterThanExpr(expr, idContexts)
	case *prior.GreaterThanOrEqualsExpr:
		return t.typeCheckGreaterThanOrEqualsExpr(expr, idContexts)
	case *prior.IdentifierExpr:
		return t.typeCheckIdentifierExpr(expr, idContexts)
	case *prior.Int64LiteralExpr:
		return t.typeCheckInt64LiteralExpr(expr)
	case *prior.IsExpr:
		return t.typeCheckIsExpr(expr, idContexts)
	case *prior.LessThanExpr:
		return t.typeCheckLessThanExpr(expr, idContexts)
	case *prior.LessThanOrEqualsExpr:
		return t.typeCheckLessThanOrEqualsExpr(expr, idContexts)
	case *prior.LogicalAndExpr:
		return t.typeCheckLogicalAndExpr(expr, idContexts)
	case *prior.LogicalNotOperationExpr:
		return t.typeCheckLogicalNotOperationExpr(expr, idContexts)
	case *prior.LogicalOrExpr:
		return t.typeCheckLogicalOrExpr(expr, idContexts)
	case *prior.MultiplicationExpr:
		return t.typeCheckMultiplicationExpr(expr, idContexts)
	case *prior.NegationOperationExpr:
		return t.typeCheckNegationOperationExpr(expr, idContexts)
	case *prior.NotEqualsExpr:
		return t.typeCheckNotEqualsExpr(expr, idContexts)
	case *prior.ParenthesizedExpr:
		return t.typeCheckParenthesizedExpr(expr, idContexts)
	case *prior.RecordExpr:
		return t.typeCheckRecordExpr(expr, idContexts)
	case *prior.StringLiteralExpr:
		return t.typeCheckStringLiteralExpr(expr)
	case *prior.SubtractionExpr:
		return t.typeCheckSubtractionExpr(expr, idContexts)

	default:
		panic(fmt.Sprintf("Missing case in checkTypes: %T\n", expression))

	}

}

//---------------------------------------------------------------------------------------------------------------------

func (t *typeChecker) typeCheckAdditionExpr(expr *prior.AdditionExpr, idContexts []uint64) IExpression {
	lhs := t.checkTypes(expr.Lhs, idContexts)
	rhs := t.checkTypes(expr.Rhs, idContexts)
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

//---------------------------------------------------------------------------------------------------------------------

func (t *typeChecker) typeCheckBooleanLiteralExpr(expr *prior.BooleanLiteralExpr) IExpression {
	return &BooleanLiteralExpr{
		SourcePosition: expr.SourcePosition,
		Value:          expr.Value,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (t *typeChecker) typeCheckBuiltInTypeExpr(expr *prior.BuiltInTypeExpr) IExpression {
	name := expr.SourcePosition.GetText(t.SourceCode)
	return &BuiltInTypeExpr{
		SourcePosition: expr.SourcePosition,
		ValueIndex:     t.TypePool.GetIndexByName(name),
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (t *typeChecker) typeCheckDivisionExpr(expr *prior.DivisionExpr, idContexts []uint64) IExpression {
	lhs := t.checkTypes(expr.Lhs, idContexts)
	rhs := t.checkTypes(expr.Rhs, idContexts)
	// TODO: ensure they're the same
	return &DivisionExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
		TypeIndex:      lhs.GetTypeIndex(),
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (t *typeChecker) typeCheckEqualsExpr(expr *prior.EqualsExpr, idContexts []uint64) IExpression {
	lhs := t.checkTypes(expr.Lhs, idContexts)
	rhs := t.checkTypes(expr.Rhs, idContexts)
	// TODO: ensure they're the same
	return &EqualsExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (t *typeChecker) typeCheckFieldReferenceExpr(expr *prior.FieldReferenceExpr, idContexts []uint64) IExpression {
	parent := t.checkTypes(expr.Parent, idContexts)
	parentTypeIndex := parent.GetTypeIndex()
	child := t.checkTypes(expr.Child, append(idContexts, parentTypeIndex))
	// TODO: ensure the parent is a record
	return &FieldReferenceExpr{
		SourcePosition: expr.SourcePosition,
		Parent:         parent,
		Child:          child,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (t *typeChecker) typeCheckFloat64LiteralExpr(expr *prior.Float64LiteralExpr) IExpression {
	return &Float64LiteralExpr{
		SourcePosition: expr.SourcePosition,
		Value:          expr.Value,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (t *typeChecker) typeCheckGreaterThanExpr(expr *prior.GreaterThanExpr, idContexts []uint64) IExpression {
	lhs := t.checkTypes(expr.Lhs, idContexts)
	rhs := t.checkTypes(expr.Rhs, idContexts)
	// TODO: ensure they're the same
	return &GreaterThanExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (t *typeChecker) typeCheckGreaterThanOrEqualsExpr(expr *prior.GreaterThanOrEqualsExpr, idContexts []uint64) IExpression {
	lhs := t.checkTypes(expr.Lhs, idContexts)
	rhs := t.checkTypes(expr.Rhs, idContexts)
	// TODO: ensure they're the same
	return &GreaterThanOrEqualsExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (t *typeChecker) typeCheckIdentifierExpr(expr *prior.IdentifierExpr, idContexts []uint64) IExpression {

	fieldIndex := uint64(0xFFFFFFFF)
	typeIndex := uint64(0xFFFFFFFF)
outer:
	for i := len(idContexts) - 1; i >= 0; i-- {
		recordType := t.TypePool.Get(idContexts[i]).(*types.RecordType)

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

//---------------------------------------------------------------------------------------------------------------------

func (t *typeChecker) typeCheckInt64LiteralExpr(expr *prior.Int64LiteralExpr) IExpression {
	return &Int64LiteralExpr{
		SourcePosition: expr.SourcePosition,
		Value:          expr.Value,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (t *typeChecker) typeCheckIsExpr(expr *prior.IsExpr, idContexts []uint64) IExpression {
	lhs := t.checkTypes(expr.Lhs, idContexts)
	rhs := t.checkTypes(expr.Rhs, idContexts)
	// TODO: ensure the lhs and rhs are compatible
	return &IsExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (t *typeChecker) typeCheckLessThanExpr(expr *prior.LessThanExpr, idContexts []uint64) IExpression {
	lhs := t.checkTypes(expr.Lhs, idContexts)
	rhs := t.checkTypes(expr.Rhs, idContexts)
	// TODO: ensure they're the same
	return &LessThanExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (t *typeChecker) typeCheckLessThanOrEqualsExpr(expr *prior.LessThanOrEqualsExpr, idContexts []uint64) IExpression {
	lhs := t.checkTypes(expr.Lhs, idContexts)
	rhs := t.checkTypes(expr.Rhs, idContexts)
	// TODO: ensure they're the same
	return &LessThanOrEqualsExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (t *typeChecker) typeCheckLogicalAndExpr(expr *prior.LogicalAndExpr, idContexts []uint64) IExpression {
	lhs := t.checkTypes(expr.Lhs, idContexts)
	rhs := t.checkTypes(expr.Rhs, idContexts)
	// TODO: ensure they're both boolean
	return &LogicalAndExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (t *typeChecker) typeCheckLogicalNotOperationExpr(expr *prior.LogicalNotOperationExpr, idContexts []uint64) IExpression {
	operand := t.checkTypes(expr.Operand, idContexts)

	// TODO: validate that operands are boolean

	return &LogicalNotOperationExpr{
		SourcePosition: expr.SourcePosition,
		Operand:        operand,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (t *typeChecker) typeCheckLogicalOrExpr(expr *prior.LogicalOrExpr, idContexts []uint64) IExpression {
	lhs := t.checkTypes(expr.Lhs, idContexts)
	rhs := t.checkTypes(expr.Rhs, idContexts)
	// TODO: ensure they're both boolean
	return &LogicalOrExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (t *typeChecker) typeCheckMultiplicationExpr(expr *prior.MultiplicationExpr, idContexts []uint64) IExpression {
	lhs := t.checkTypes(expr.Lhs, idContexts)
	rhs := t.checkTypes(expr.Rhs, idContexts)
	// TODO: ensure they're the same
	return &MultiplicationExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
		TypeIndex:      lhs.GetTypeIndex(),
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (t *typeChecker) typeCheckNegationOperationExpr(expr *prior.NegationOperationExpr, idContexts []uint64) IExpression {
	operand := t.checkTypes(expr.Operand, idContexts)

	return &NegationOperationExpr{
		SourcePosition: expr.SourcePosition,
		Operand:        operand,
		TypeIndex:      operand.GetTypeIndex(),
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (t *typeChecker) typeCheckNotEqualsExpr(expr *prior.NotEqualsExpr, idContexts []uint64) IExpression {
	lhs := t.checkTypes(expr.Lhs, idContexts)
	rhs := t.checkTypes(expr.Rhs, idContexts)
	// TODO: ensure they're the same
	return &NotEqualsExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (t *typeChecker) typeCheckParenthesizedExpr(expr *prior.ParenthesizedExpr, idContexts []uint64) IExpression {

	inner := t.checkTypes(expr.InnerExpr, idContexts)

	return &ParenthesizedExpr{
		SourcePosition: expr.SourcePosition,
		InnerExpr:      inner,
		TypeIndex:      inner.GetTypeIndex(),
	}

}

//---------------------------------------------------------------------------------------------------------------------

func (t *typeChecker) typeCheckRecordExpr(expr *prior.RecordExpr, idContexts []uint64) IExpression {

	fields := make([]*RecordFieldExpr, 0)
	for _, field := range expr.Fields {
		fields = append(fields, t.typeCheckRecordFieldExpr(field, idContexts))
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
	typeIndex := t.TypePool.Put(recordType)

	return &RecordExpr{
		SourcePosition: expr.SourcePosition,
		Fields:         fields,
		TypeIndex:      typeIndex,
	}

}

//---------------------------------------------------------------------------------------------------------------------

func (t *typeChecker) typeCheckRecordFieldExpr(expr *prior.RecordFieldExpr, idContexts []uint64) *RecordFieldExpr {
	return &RecordFieldExpr{
		SourcePosition: expr.SourcePosition,
		FieldNameIndex: expr.FieldNameIndex,
		FieldValue:     t.checkTypes(expr.FieldValue, idContexts),
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (t *typeChecker) typeCheckStringLiteralExpr(expr *prior.StringLiteralExpr) IExpression {
	return &StringLiteralExpr{
		SourcePosition: expr.SourcePosition,
		ValueIndex:     expr.ValueIndex,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (t *typeChecker) typeCheckSubtractionExpr(expr *prior.SubtractionExpr, idContexts []uint64) IExpression {
	lhs := t.checkTypes(expr.Lhs, idContexts)
	rhs := t.checkTypes(expr.Rhs, idContexts)
	// TODO: ensure they're the same
	return &SubtractionExpr{
		SourcePosition: expr.SourcePosition,
		Lhs:            lhs,
		Rhs:            rhs,
		TypeIndex:      lhs.GetTypeIndex(),
	}
}

//---------------------------------------------------------------------------------------------------------------------
