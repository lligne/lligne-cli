//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package records

//=====================================================================================================================

type Record struct {
	TypeIndex   uint64
	FieldValues []RecordFieldValue
}

//---------------------------------------------------------------------------------------------------------------------

func (r1 *Record) Equal(r2 Record) bool {

	// TODO: type synonyms - use canonical type
	if r1.TypeIndex != r2.TypeIndex {
		return false
	}

	for i, f1 := range r1.FieldValues {
		f2 := r2.FieldValues[i]

		// TODO: deep equality
		if f1 != f2 {
			return false
		}
	}

	return true
}

//=====================================================================================================================

type RecordFieldValue = uint64

//=====================================================================================================================
