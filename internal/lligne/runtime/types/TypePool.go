//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package types

//=====================================================================================================================

// TypePool holds a list of types interned so that they can be retrieved by index.
type TypePool struct {
	types         []IType
	indexes       map[IType]uint64
	indexesByName map[string]uint64
}

//---------------------------------------------------------------------------------------------------------------------

// NewTypePool creates a new empty type pool.
func NewTypePool() *TypePool {
	result := &TypePool{
		types:         nil,
		indexes:       make(map[IType]uint64),
		indexesByName: make(map[string]uint64),
	}

	// NOTE: Keep these in sync with BuiltInTypeIndex just below
	result.Put(UnitTypeInstance)
	result.Put(BoolTypeInstance)
	result.Put(Float64TypeInstance)
	result.Put(Int64TypeInstance)
	result.Put(StringTypeInstance)
	result.Put(TypeTypeInstance)

	return result
}

//---------------------------------------------------------------------------------------------------------------------

// BuiltInTypeIndex is an enumeration of known pool indexes for built-in types.
const (
	// NOTE: Keep these in sync with type pool initialization just above
	BuiltInTypeIndexUnit uint64 = iota
	BuiltInTypeIndexBool
	BuiltInTypeIndexFloat64
	BuiltInTypeIndexInt64
	BuiltInTypeIndexString
	BuiltInTypeIndexType
)

//---------------------------------------------------------------------------------------------------------------------

// Freeze returns an immutable view of this type pool. The original mutable view should be abandoned afterward.
func (p *TypePool) Freeze() *TypeConstantPool {
	return &TypeConstantPool{
		ITypes: p.types,
	}
}

//---------------------------------------------------------------------------------------------------------------------

// Get returns the type at the given index.
func (p *TypePool) Get(index uint64) IType {
	return p.types[index]
}

//---------------------------------------------------------------------------------------------------------------------

// GetByName returns the type with given name.
func (p *TypePool) GetByName(name string) IType {
	return p.types[p.indexesByName[name]]
}

//---------------------------------------------------------------------------------------------------------------------

// GetIndexByName returns the index of the type with given name.
func (p *TypePool) GetIndexByName(name string) uint64 {
	return p.indexesByName[name]
}

//---------------------------------------------------------------------------------------------------------------------

// Put looks for the type already in the pool. It adds it if not there.
// Returns the index of the new or existing entry.
func (p *TypePool) Put(value IType) uint64 {
	result, found := p.indexes[value]

	if !found {
		result = uint64(len(p.types))
		p.types = append(p.types, value)
		p.indexes[value] = result
		p.indexesByName[value.Name()] = result
	}

	return result
}

//=====================================================================================================================

// TypeConstantPool is an immutable view of a TypePool.
type TypeConstantPool struct {
	ITypes []IType
}

//---------------------------------------------------------------------------------------------------------------------

// Clone returns a mutable copy of this type pool.
func (p *TypeConstantPool) Clone() *TypePool {
	result := NewTypePool()
	for _, str := range p.ITypes {
		result.Put(str)
	}
	return result
}

//---------------------------------------------------------------------------------------------------------------------

// Get returns the type at the given index.
func (p *TypeConstantPool) Get(index uint64) IType {
	return p.ITypes[index]
}

//=====================================================================================================================
