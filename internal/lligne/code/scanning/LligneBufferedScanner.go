//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package scanning

//---------------------------------------------------------------------------------------------------------------------

// ILligneBufferedScanner allows reading tokens with one token of lookahead.
type ILligneBufferedScanner interface {
	ILligneScanner
	AdvanceTokenIfType(tokenType LligneTokenType) bool
	PeekToken() LligneToken
	PeekTokenIsType(tokenType LligneTokenType) bool
}

//---------------------------------------------------------------------------------------------------------------------

// LligneBufferedScanner converts a string of Lligne source code into tokens with one token of lookahead.
type lligneBufferedScanner struct {
	scanner   ILligneScanner
	nextToken LligneToken
}

//---------------------------------------------------------------------------------------------------------------------

// NewLligneBufferedScanner allocates a new buffered scanner that wraps a given inner scanner.
func NewLligneBufferedScanner(scanner ILligneScanner) ILligneBufferedScanner {
	return &lligneBufferedScanner{
		scanner:   scanner,
		nextToken: scanner.ReadToken(),
	}
}

//---------------------------------------------------------------------------------------------------------------------

// AdvanceTokenIfType consumes one token if it has the given type. It ignores the token itself
// and returns a flag for whether the token was consumed.
func (s *lligneBufferedScanner) AdvanceTokenIfType(tokenType LligneTokenType) bool {
	if s.nextToken.TokenType == tokenType {
		s.nextToken = s.scanner.ReadToken()
		return true
	}
	return false
}

//---------------------------------------------------------------------------------------------------------------------

// GetOrigin determines a token origin from the inner scanner.
func (s *lligneBufferedScanner) GetOrigin(sourcePos int) LligneOrigin {
	return s.scanner.GetOrigin(sourcePos)
}

//---------------------------------------------------------------------------------------------------------------------

// PeekToken returns the lookahead token.
func (s *lligneBufferedScanner) PeekToken() LligneToken {
	return s.nextToken
}

//---------------------------------------------------------------------------------------------------------------------

// PeekTokenIsType determines whether the lookahead token has given type.
func (s *lligneBufferedScanner) PeekTokenIsType(tokenType LligneTokenType) bool {
	return s.nextToken.TokenType == tokenType
}

//---------------------------------------------------------------------------------------------------------------------

// ReadToken consumes and returns the current token then fills in the next lookahead token.
func (s *lligneBufferedScanner) ReadToken() LligneToken {
	result := s.nextToken
	s.nextToken = s.scanner.ReadToken()
	return result
}

//---------------------------------------------------------------------------------------------------------------------
