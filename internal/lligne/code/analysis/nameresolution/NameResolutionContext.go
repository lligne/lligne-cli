//
// # Data types related to Lligne name resolution.
//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package nameresolution

import "lligne-cli/internal/lligne/runtime/pools"

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
	fieldReferenceNames           map[pools.NameIndex]NameUsage
	whereNames                    map[pools.NameIndex]NameUsage
	recordsUnderConstructionNames map[pools.NameIndex]NameUsage
	topLevelNames                 map[pools.NameIndex]NameUsage
}

//---------------------------------------------------------------------------------------------------------------------

func NewNameResolutionContext() *NameResolutionContext {
	return &NameResolutionContext{
		fieldReferenceNames:           make(map[pools.NameIndex]NameUsage),
		whereNames:                    make(map[pools.NameIndex]NameUsage),
		recordsUnderConstructionNames: make(map[pools.NameIndex]NameUsage),
		topLevelNames:                 make(map[pools.NameIndex]NameUsage),
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

func (c *NameResolutionContext) LookUpName(nameIndex pools.NameIndex) NameUsage {
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

func makeNameUsageMap(nameIndexes []pools.NameIndex, mechanism ResolutionMechanism) map[pools.NameIndex]NameUsage {
	result := make(map[pools.NameIndex]NameUsage)

	for index, nameIndex := range nameIndexes {
		result[nameIndex] = NameUsage{
			FieldIndex: uint64(index),
			Mechanism:  mechanism,
		}
	}

	return result
}

//=====================================================================================================================
