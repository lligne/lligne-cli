//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package bytecode

//=====================================================================================================================

type StringConstantPool struct {
	strings []string
	indexes map[string]uint16
}

//---------------------------------------------------------------------------------------------------------------------

func NewStringConstantPool() StringConstantPool {
	return StringConstantPool{
		strings: nil,
		indexes: make(map[string]uint16),
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (p *StringConstantPool) Get(index uint16) string {
	return p.strings[index]
}

//---------------------------------------------------------------------------------------------------------------------

func (p *StringConstantPool) Put(value string) uint16 {
	result, found := p.indexes[value]

	if !found {
		result = uint16(len(p.strings))
		p.strings = append(p.strings, value)
		p.indexes[value] = result
	}

	return result
}

//=====================================================================================================================
