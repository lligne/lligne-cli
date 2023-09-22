//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package pools

//=====================================================================================================================

type NameIndex uint64

//=====================================================================================================================

type StringIndex uint64

//=====================================================================================================================

// Pool holds a list of strings interned so that they can be retrieved by index.
type Pool[Index NameIndex | StringIndex] struct {
	strings []string
	indexes map[string]Index
}

//---------------------------------------------------------------------------------------------------------------------

// NewPool creates a new empty string pool.
func newPool[Index NameIndex | StringIndex]() *Pool[Index] {
	return &Pool[Index]{
		strings: nil,
		indexes: make(map[string]Index),
	}
}

//---------------------------------------------------------------------------------------------------------------------

// Freeze returns an immutable view of this string pool. The original mutable view should be abandoned afterward.
func (p *Pool[Index]) Freeze() *ConstantPool[Index] {
	return &ConstantPool[Index]{
		strings: p.strings,
	}
}

//---------------------------------------------------------------------------------------------------------------------

// Get returns the string at the given index.
func (p *Pool[Index]) Get(index Index) string {
	return p.strings[index]
}

//---------------------------------------------------------------------------------------------------------------------

// Put looks for the string already in the pool. It adds it if not there.
// Returns the index of the new or existing entry.
func (p *Pool[Index]) Put(value string) Index {
	result, found := p.indexes[value]

	if !found {
		result = Index(len(p.strings))
		p.strings = append(p.strings, value)
		p.indexes[value] = result
	}

	return result
}

//=====================================================================================================================

// StringConstantPool is an immutable view of a StringPool.
type ConstantPool[Index NameIndex | StringIndex] struct {
	strings []string
}

//---------------------------------------------------------------------------------------------------------------------

// Clone returns a mutable copy of this string pool.
func (p *ConstantPool[Index]) Clone() *Pool[Index] {
	result := newPool[Index]()
	for _, str := range p.strings {
		result.Put(str)
	}
	return result
}

//---------------------------------------------------------------------------------------------------------------------

// Get returns the string at the given index.
func (p *ConstantPool[Index]) Get(index Index) string {
	return p.strings[index]
}

//=====================================================================================================================

type StringPool = Pool[StringIndex]

func NewStringPool() *StringPool {
	return newPool[StringIndex]()
}

type StringConstantPool = ConstantPool[StringIndex]

//=====================================================================================================================

type NamePool = Pool[NameIndex]

func NewNamePool() *NamePool {
	return newPool[NameIndex]()
}

type NameConstantPool = ConstantPool[NameIndex]

//=====================================================================================================================
