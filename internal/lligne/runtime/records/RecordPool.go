//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package records

//=====================================================================================================================

// RecordPool holds a list of records stored so that they can be retrieved by index.
type RecordPool struct {
	records []Record
}

//---------------------------------------------------------------------------------------------------------------------

// NewRecordPool creates a new empty record pool.
func NewRecordPool() *RecordPool {
	return &RecordPool{
		records: nil,
	}
}

//---------------------------------------------------------------------------------------------------------------------

// Freeze returns an immutable view of this string pool. The original mutable view should be abandoned afterward.
func (p *RecordPool) Freeze() *RecordConstantPool {
	return &RecordConstantPool{
		records: p.records,
	}
}

//---------------------------------------------------------------------------------------------------------------------

// Get returns the string at the given index.
func (p *RecordPool) Get(index uint64) Record {
	return p.records[index]
}

//---------------------------------------------------------------------------------------------------------------------

// Put adds a record to the pool.
// Returns the index of the new or existing entry.
func (p *RecordPool) Put(value Record) uint64 {
	result := uint64(len(p.records))
	p.records = append(p.records, value)

	return result
}

//=====================================================================================================================

// RecordConstantPool is an immutable view of a RecordPool.
type RecordConstantPool struct {
	records []Record
}

//---------------------------------------------------------------------------------------------------------------------

// Clone returns a mutable copy of this string pool.
func (p *RecordConstantPool) Clone() *RecordPool {
	result := NewRecordPool()
	for _, str := range p.records {
		result.Put(str)
	}
	return result
}

//---------------------------------------------------------------------------------------------------------------------

// Get returns the string at the given index.
func (p *RecordConstantPool) Get(index uint64) Record {
	return p.records[index]
}

//=====================================================================================================================
