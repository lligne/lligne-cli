//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package records

import "lligne-cli/internal/lligne/runtime/types"

//=====================================================================================================================

type Record struct {
	TypeIndex   uint64
	FieldValues []RecordFieldValue
}

//=====================================================================================================================

type RecordFieldValue = uint64

//=====================================================================================================================

func AreRecordsEqual(p *types.TypePool, r *RecordPool, r1Index uint64, r2Index uint64) bool {

	r1 := r.Get(r1Index)
	r2 := r.Get(r2Index)

	if !areRecordTypesEquivalent(p, r, r1.TypeIndex, r2.TypeIndex) {
		return false
	}

	for i, f1 := range r1.FieldValues {
		r1Type := p.Get(r1.TypeIndex).(*types.RecordType)

		f1Type := p.Get(r1Type.FieldTypeIndexes[i])
		f2 := r2.FieldValues[i]

		if f1Type.Category() == types.TypeCategoryRecord {
			if !AreRecordsEqual(p, r, f1, f2) {
				return false
			}
		} else if f1 != f2 {
			return false
		}
	}

	return true
}

//=====================================================================================================================

func areRecordTypesEquivalent(p *types.TypePool, r *RecordPool, type1Index uint64, type2Index uint64) bool {

	type1 := p.Get(type1Index).(*types.RecordType)
	type2 := p.Get(type2Index).(*types.RecordType)

	if len(type1.FieldTypeIndexes) != len(type2.FieldTypeIndexes) {
		return false
	}

	for i, field1Name := range type1.FieldNameIndexes {
		field2Name := type2.FieldNameIndexes[i]
		if field1Name != field2Name {
			return false
		}

		field1TypeIndex := type1.FieldTypeIndexes[i]
		field2TypeIndex := type2.FieldTypeIndexes[i]

		field1Type := p.Get(field1TypeIndex)

		if field1Type.Category() == types.TypeCategoryRecord {
			if !areRecordTypesEquivalent(p, r, field1TypeIndex, field2TypeIndex) {
				return false
			}
		} else if field1TypeIndex != field2TypeIndex {
			return false
		}
	}

	return true
}

//=====================================================================================================================
