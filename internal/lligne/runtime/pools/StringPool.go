//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package pools

//=====================================================================================================================

// StringPool holds a list of strings interned so that they can be retrieved by index.
type StringPool struct {
	strings []string
	indexes map[string]uint64
}

//---------------------------------------------------------------------------------------------------------------------

// NewStringPool creates a new empty string pool.
func NewStringPool() *StringPool {
	return &StringPool{
		strings: nil,
		indexes: make(map[string]uint64),
	}
}

//---------------------------------------------------------------------------------------------------------------------

// Freeze returns an immutable view of this string pool. The original mutable view should be abandoned afterward.
func (p *StringPool) Freeze() *StringConstantPool {
	return &StringConstantPool{
		strings: p.strings,
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

// StringConstantPool is an immutable view of a StringPool.
type StringConstantPool struct {
	strings []string
}

//---------------------------------------------------------------------------------------------------------------------

// Clone returns a mutable copy of this string pool.
func (p *StringConstantPool) Clone() *StringPool {
	result := NewStringPool()
	for _, str := range p.strings {
		result.Put(str)
	}
	return result
}

//---------------------------------------------------------------------------------------------------------------------

// Get returns the string at the given index.
func (p *StringConstantPool) Get(index uint64) string {
	return p.strings[index]
}

//=====================================================================================================================
