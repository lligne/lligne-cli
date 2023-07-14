//
// # Data types related to Lligne token scanning.
//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package scanning

//=====================================================================================================================

// LligneToken is an abstract token occurring at SourceStartPos in its source file.
type LligneToken struct {
	TokenType      LligneTokenType
	Text           string
	SourceStartPos int
}

//=====================================================================================================================
