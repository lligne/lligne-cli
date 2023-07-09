//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package scanning

//---------------------------------------------------------------------------------------------------------------------

type ILligneScannerBuffer interface {
	AdvanceTokenIfType(tokenType LligneTokenType) bool
	PeekToken() LligneToken
	PeekTokenIsType(tokenType LligneTokenType) bool
	ReadToken() LligneToken
}

//---------------------------------------------------------------------------------------------------------------------

// LligneScanner converts a string of Lligne source code into tokens.
type lligneScannerBuffer struct {
	scanner   ILligneScanner
	nextToken LligneToken
}

//---------------------------------------------------------------------------------------------------------------------

// NewLligneScanner allocates a new scanner for given sourceCode from the given fileName.
func NewLligneScannerBuffer(scanner ILligneScanner) ILligneScannerBuffer {
	return &lligneScannerBuffer{
		scanner:   scanner,
		nextToken: scanner.ReadToken(),
	}
}

//---------------------------------------------------------------------------------------------------------------------

func (s *lligneScannerBuffer) AdvanceTokenIfType(tokenType LligneTokenType) bool {
	if s.nextToken.TokenType == tokenType {
		s.nextToken = s.scanner.ReadToken()
		return true
	}
	return false
}

//---------------------------------------------------------------------------------------------------------------------

func (s *lligneScannerBuffer) PeekToken() LligneToken {
	return s.nextToken
}

//---------------------------------------------------------------------------------------------------------------------

func (s *lligneScannerBuffer) PeekTokenIsType(tokenType LligneTokenType) bool {
	return s.nextToken.TokenType == tokenType
}

//---------------------------------------------------------------------------------------------------------------------

func (s *lligneScannerBuffer) ReadToken() LligneToken {
	result := s.nextToken
	s.nextToken = s.scanner.ReadToken()
	return result
}

//---------------------------------------------------------------------------------------------------------------------
