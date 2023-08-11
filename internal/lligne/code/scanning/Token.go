//
// # Data types related to Lligne token scanning.
//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package scanning

//=====================================================================================================================

// Token is an abstract token of type TokenType occurring at SourceOffset with length SourceLength in its source code.
type Token struct {
	SourceOffset uint32
	SourceLength uint16
	TokenType    TokenType
}

//=====================================================================================================================
