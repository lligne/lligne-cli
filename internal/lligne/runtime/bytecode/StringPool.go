//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package bytecode

//=====================================================================================================================

// StringPool holds a list of strings interned so that they can be retrieved by index.
type StringPool struct {
	strings []string
	indexes map[string]uint64
}

//---------------------------------------------------------------------------------------------------------------------

// NewStringConstantPool creates a new empty string pool.
func NewStringConstantPool() StringPool {
	return StringPool{
		strings: nil,
		indexes: make(map[string]uint64),
	}
}

//---------------------------------------------------------------------------------------------------------------------

// Get returns the string at the given index.
func (p *StringPool) Get(index uint64) string {
	return p.strings[index]
}

//---------------------------------------------------------------------------------------------------------------------

// Put looks for the string already in the pool. It adds it if not there.
// Returns the index of the new or existing entry.
func (p *StringPool) Put(value string) uint64 {
	result, found := p.indexes[value]

	if !found {
		result = uint64(len(p.strings))
		p.strings = append(p.strings, value)
		p.indexes[value] = result
	}

	return result
}

//=====================================================================================================================