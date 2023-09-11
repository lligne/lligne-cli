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

func AreRecordsEqual(p *types.TypePool, r1 *Record, r2 *Record) bool {

	r1Type := p.Get(r1.TypeIndex).(*types.RecordType)
	r2Type := p.Get(r2.TypeIndex).(*types.RecordType)

	for i, field1Name := range r1Type.FieldNameIndexes {
		field2Name := r2Type.FieldNameIndexes[i]
		if field1Name != field2Name {
			return false
		}

		field1TypeIndex := r1Type.FieldTypeIndexes[i]
		field2TypeIndex := r2Type.FieldTypeIndexes[i]

		// TODO: nested record type equality
		if field1TypeIndex != field2TypeIndex {
			return false
		}
	}

	for i, f1 := range r1.FieldValues {
		f2 := r2.FieldValues[i]

		// TODO: deep equality for nested records
		if f1 != f2 {
			return false
		}
	}

	return true
}

//=====================================================================================================================
