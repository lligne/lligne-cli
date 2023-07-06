//
// # Data types related to token origins.
//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package scanning

import (
	"fmt"
)

//---------------------------------------------------------------------------------------------------------------------

// LligneOrigin is a record of where a parsed item originated.
type LligneOrigin struct {

	// The name of the source file.
	FileName string

	// The line in the file (1-based).
	Line int

	// The column in the file (1-based).
	Column int
}

//---------------------------------------------------------------------------------------------------------------------

// String converts a LligneOrigin to text for the file name and start of the origin.
func (o LligneOrigin) String() string {
	return fmt.Sprintf("%s(%d,%d)", o.FileName, o.Line, o.Column)
}

//---------------------------------------------------------------------------------------------------------------------
