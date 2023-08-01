//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package bytecode

//=====================================================================================================================

// ValueStack holds the stack of values for the virtual bytecode machine.
type ValueStack struct {
	values [1000]uint64
	size   int
}

//---------------------------------------------------------------------------------------------------------------------

const True64 uint64 = 0xFFFFFFFFFFFFFFFF

//---------------------------------------------------------------------------------------------------------------------

// PeekBool returns the top value on the stack as a boolean. The stack remains unchanged.
func (s *ValueStack) PeekBool() bool {
	return s.values[s.size-1] != 0
}

//---------------------------------------------------------------------------------------------------------------------

// PopBool returns the top value on the stack as a boolean and shrinks the stack by one.
func (s *ValueStack) PopBool() bool {
	s.size -= 1
	return s.values[s.size] != 0
}

//---------------------------------------------------------------------------------------------------------------------

// PushBool places a boolean on the top of the stack.
func (s *ValueStack) PushBool(value bool) {
	if value {
		s.values[s.size] = True64
	} else {
		s.values[s.size] = 0
	}
	s.size += 1
}

//---------------------------------------------------------------------------------------------------------------------

// PeekInt64 returns the top value on the stack as a 64 bit integer. The stack remains unchanged.
func (s *ValueStack) PeekInt64() int64 {
	return int64(s.values[s.size-1])
}

//---------------------------------------------------------------------------------------------------------------------

// PopInt64 returns the top value on the stack as a 64 bit integer and shrinks the stack by one.
func (s *ValueStack) PopInt64() int64 {
	s.size -= 1
	return int64(s.values[s.size])
}

//---------------------------------------------------------------------------------------------------------------------

// PushInt64 places a 64 bit integer on the top of the stack.
func (s *ValueStack) PushInt64(value int64) {
	s.values[s.size] = uint64(value)
	s.size += 1
}

//=====================================================================================================================
