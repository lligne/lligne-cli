//
// # Data types related to Lligne name resolution.
//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package nameresolution

//=====================================================================================================================

// ResolutionMechanism is an enumeration of how identifiers link to their origin.
//
// Name resolution rules:
// 1. When inside the right hand side of a field reference expression, find the name inside the left hand side of the expression or fail.
//
// 2. When inside the left hand side of a where expression, find the name inside the right hand side of the expression or continue.
// 3. When inside a record, find the name as a sibling field in the record or continue.
// 4. When inside a nested record, recursively find the name as a field of the parent record or continue.
// 5. Find the name inside the top level.
type ResolutionMechanism uint16

const (
	ResolutionMechanismUndefined ResolutionMechanism = iota
	ResolutionMechanismFieldReference
	ResolutionMechanismWhereField
	ResolutionMechanismRecordField
	ResolutionMechanismTopLevel
)

//=====================================================================================================================

type NameUsage struct {
	FieldIndex uint64
	Mechanism  ResolutionMechanism
}

//=====================================================================================================================

// TODO: make these maps instead of arrays

type NameResolutionContext struct {
	fieldReferenceNames           map[uint64]NameUsage
	whereNames                    map[uint64]NameUsage
	recordsUnderConstructionNames map[uint64]NameUsage
	topLevelNames                 map[uint64]NameUsage
}

//---------------------------------------------------------------------------------------------------------------------

func NewNameResolutionContext() *NameResolutionContext {
	return &NameResolutionContext{
		fieldReferenceNames:           make(map[uint64]NameUsage),
		whereNames:                    make(map[uint64]NameUsage),
		recordsUnderConstructionNames: make(map[uint64]NameUsage),
		topLevelNames:                 make(map[uint64]NameUsage),
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (c *NameResolutionContext) WithFieldReferenceLhs(fieldReferenceLhs IExpression) *NameResolutionContext {
	return &NameResolutionContext{
		fieldReferenceNames:           makeNameUsageMap(fieldReferenceLhs.GetFieldNameIndexes(), ResolutionMechanismFieldReference),
		whereNames:                    c.whereNames,
		recordsUnderConstructionNames: c.recordsUnderConstructionNames,
		topLevelNames:                 c.topLevelNames,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (c *NameResolutionContext) WithWhereRhs(whereRhs IExpression) *NameResolutionContext {
	return &NameResolutionContext{
		fieldReferenceNames:           c.fieldReferenceNames,
		whereNames:                    makeNameUsageMap(whereRhs.GetFieldNameIndexes(), ResolutionMechanismWhereField),
		recordsUnderConstructionNames: c.recordsUnderConstructionNames,
		topLevelNames:                 c.topLevelNames,
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (c *NameResolutionContext) LookUpName(nameIndex uint64) NameUsage {
	// Name resolution rules:
	// 1. When inside the right hand side of a field reference expression, find the name inside the left hand side of the expression or fail.
	//
	// 2. When inside the left hand side of a where expression, find the name inside the right hand side of the expression or continue.
	// 3. When inside a record, find the name as a sibling field in the record or continue.
	// 4. When inside a nested record, recursively find the name as a field of the parent record or continue.
	// 5. Find the name inside the top level.

	if len(c.fieldReferenceNames) > 0 {
		return c.fieldReferenceNames[nameIndex]
	}

	result := c.whereNames[nameIndex]

	// TODO: names within same or parent records if result is nil

	return result
}

//=====================================================================================================================

func makeNameUsageMap(nameIndexes []uint64, mechanism ResolutionMechanism) map[uint64]NameUsage {
	result := make(map[uint64]NameUsage)

	for index, nameIndex := range nameIndexes {
		result[nameIndex] = NameUsage{
			FieldIndex: uint64(index),
			Mechanism:  mechanism,
		}
	}

	return result
}

//=====================================================================================================================
