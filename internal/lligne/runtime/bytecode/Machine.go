//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package bytecode

//=====================================================================================================================

// Machine is a stack of operands for bytecode operations.
type Machine struct {
	Stack     [1000]uint64
	Top       int
	IP        int
	IsRunning bool
}

//---------------------------------------------------------------------------------------------------------------------

func NewMachine() *Machine {
	return &Machine{Top: -1, IsRunning: true}
}

//=====================================================================================================================
